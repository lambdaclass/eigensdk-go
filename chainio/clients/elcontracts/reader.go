package elcontracts

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	permissioncontroller "github.com/Layr-Labs/eigensdk-go/contracts/bindings/PermissionController"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type Config struct {
	DelegationManagerAddress     gethcommon.Address
	AvsDirectoryAddress          gethcommon.Address
	RewardsCoordinatorAddress    gethcommon.Address
	PermissionsControllerAddress gethcommon.Address
}

type ChainReader struct {
	logger               logging.Logger
	delegationManager    *delegationmanager.ContractDelegationManager
	strategyManager      *strategymanager.ContractStrategyManager
	avsDirectory         *avsdirectory.ContractAVSDirectory
	rewardsCoordinator   *rewardscoordinator.ContractIRewardsCoordinator
	allocationManager    *allocationmanager.ContractAllocationManager
	permissionController *permissioncontroller.ContractPermissionController
	ethClient            eth.HttpBackend
}

var errLegacyAVSsNotSupported = errors.New("method not supported for legacy AVSs")

func NewChainReader(
	delegationManager *delegationmanager.ContractDelegationManager,
	strategyManager *strategymanager.ContractStrategyManager,
	avsDirectory *avsdirectory.ContractAVSDirectory,
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator,
	allocationManager *allocationmanager.ContractAllocationManager,
	permissionController *permissioncontroller.ContractPermissionController,
	logger logging.Logger,
	ethClient eth.HttpBackend,
) *ChainReader {
	logger = logger.With(logging.ComponentKey, "elcontracts/reader")

	return &ChainReader{
		delegationManager:    delegationManager,
		strategyManager:      strategyManager,
		avsDirectory:         avsDirectory,
		rewardsCoordinator:   rewardsCoordinator,
		allocationManager:    allocationManager,
		permissionController: permissionController,
		logger:               logger,
		ethClient:            ethClient,
	}
}

func NewReaderFromConfig(
	cfg Config,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) (*ChainReader, error) {
	elContractBindings, err := NewBindingsFromConfig(
		cfg,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewChainReader(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AllocationManager,
		elContractBindings.PermissionController,
		logger,
		ethClient,
	), nil
}

func (r *ChainReader) IsOperatorRegistered(
	ctx context.Context,
	operator types.Operator,
) (bool, error) {
	if r.delegationManager == nil {
		return false, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.IsOperator(&bind.CallOpts{Context: ctx}, gethcommon.HexToAddress(operator.Address))
}

// GetStakerShares returns the amount of shares that a staker has in all of the strategies in which they have nonzero
// shares
func (r *ChainReader) GetStakerShares(
	ctx context.Context,
	stakerAddress gethcommon.Address,
) ([]gethcommon.Address, []*big.Int, error) {
	if r.delegationManager == nil {
		return nil, nil, errors.New("DelegationManager contract not provided")
	}
	return r.delegationManager.GetDepositedShares(&bind.CallOpts{Context: ctx}, stakerAddress)
}

// GetDelegatedOperator returns the operator that a staker has delegated to
func (r *ChainReader) GetDelegatedOperator(
	ctx context.Context,
	stakerAddress gethcommon.Address,
	blockNumber *big.Int,
) (gethcommon.Address, error) {
	if r.delegationManager == nil {
		return gethcommon.Address{}, errors.New("DelegationManager contract not provided")
	}
	return r.delegationManager.DelegatedTo(&bind.CallOpts{Context: ctx}, stakerAddress)
}

func (r *ChainReader) GetOperatorDetails(
	ctx context.Context,
	operator types.Operator,
) (types.Operator, error) {
	if r.delegationManager == nil {
		return types.Operator{}, errors.New("DelegationManager contract not provided")
	}

	delegationManagerAddress, err := r.delegationManager.DelegationApprover(
		&bind.CallOpts{Context: ctx},
		gethcommon.HexToAddress(operator.Address),
	)
	// This call should not fail since it's a getter
	if err != nil {
		return types.Operator{}, err
	}

	isSet, delay, err := r.allocationManager.GetAllocationDelay(
		&bind.CallOpts{
			Context: ctx,
		},
		gethcommon.HexToAddress(operator.Address),
	)
	// This call should not fail
	if err != nil {
		return types.Operator{}, err
	}

	var allocationDelay uint32
	if isSet {
		allocationDelay = delay
	} else {
		allocationDelay = 0
	}

	return types.Operator{
		Address:                   operator.Address,
		DelegationApproverAddress: delegationManagerAddress.Hex(),
		AllocationDelay:           allocationDelay,
	}, nil
}

// GetStrategyAndUnderlyingToken returns the strategy contract and the underlying token address
func (r *ChainReader) GetStrategyAndUnderlyingToken(
	ctx context.Context,
	strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	// This call should not fail since it's an init
	if err != nil {
		return nil, gethcommon.Address{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return contractStrategy, underlyingTokenAddr, nil
}

// GetStrategyAndUnderlyingERC20Token returns the strategy contract, the erc20 bindings for the underlying token
// and the underlying token address
func (r *ChainReader) GetStrategyAndUnderlyingERC20Token(
	ctx context.Context,
	strategyAddr gethcommon.Address,
) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error) {
	contractStrategy, err := strategy.NewContractIStrategy(strategyAddr, r.ethClient)
	// This call should not fail since it's an init
	if err != nil {
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	contractUnderlyingToken, err := erc20.NewContractIERC20(underlyingTokenAddr, r.ethClient)
	// This call should not fail, if the strategy does not have an underlying token then it would enter the if above
	if err != nil {
		return nil, nil, gethcommon.Address{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return contractStrategy, contractUnderlyingToken, underlyingTokenAddr, nil
}

func (r *ChainReader) GetOperatorSharesInStrategy(
	ctx context.Context,
	operatorAddr gethcommon.Address,
	strategyAddr gethcommon.Address,
) (*big.Int, error) {
	if r.delegationManager == nil {
		return &big.Int{}, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.OperatorShares(
		&bind.CallOpts{Context: ctx},
		operatorAddr,
		strategyAddr,
	)
}

func (r *ChainReader) CalculateDelegationApprovalDigestHash(
	ctx context.Context,
	staker gethcommon.Address,
	operator gethcommon.Address,
	delegationApprover gethcommon.Address,
	approverSalt [32]byte,
	expiry *big.Int,
) ([32]byte, error) {
	if r.delegationManager == nil {
		return [32]byte{}, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.CalculateDelegationApprovalDigestHash(
		&bind.CallOpts{Context: ctx},
		staker,
		operator,
		delegationApprover,
		approverSalt,
		expiry,
	)
}

func (r *ChainReader) CalculateOperatorAVSRegistrationDigestHash(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
	salt [32]byte,
	expiry *big.Int,
) ([32]byte, error) {
	if r.avsDirectory == nil {
		return [32]byte{}, errors.New("AVSDirectory contract not provided")
	}

	return r.avsDirectory.CalculateOperatorAVSRegistrationDigestHash(
		&bind.CallOpts{Context: ctx},
		operator,
		avs,
		salt,
		expiry,
	)
}

func (r *ChainReader) GetDistributionRootsLength(ctx context.Context) (*big.Int, error) {
	if r.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetDistributionRootsLength(&bind.CallOpts{Context: ctx})
}

func (r *ChainReader) CurrRewardsCalculationEndTimestamp(ctx context.Context) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.CurrRewardsCalculationEndTimestamp(&bind.CallOpts{Context: ctx})
}

func (r *ChainReader) GetCurrentClaimableDistributionRoot(
	ctx context.Context,
) (rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot, error) {
	if r.rewardsCoordinator == nil {
		return rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot{}, errors.New(
			"RewardsCoordinator contract not provided",
		)
	}

	return r.rewardsCoordinator.GetCurrentClaimableDistributionRoot(&bind.CallOpts{Context: ctx})
}

func (r *ChainReader) GetRootIndexFromHash(
	ctx context.Context,
	rootHash [32]byte,
) (uint32, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetRootIndexFromHash(&bind.CallOpts{Context: ctx}, rootHash)
}

func (r *ChainReader) GetCumulativeClaimed(
	ctx context.Context,
	earner gethcommon.Address,
	token gethcommon.Address,
) (*big.Int, error) {
	if r.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.CumulativeClaimed(&bind.CallOpts{Context: ctx}, earner, token)
}

func (r *ChainReader) CheckClaim(
	ctx context.Context,
	claim rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
) (bool, error) {
	if r.rewardsCoordinator == nil {
		return false, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.CheckClaim(&bind.CallOpts{Context: ctx}, claim)
}

func (r *ChainReader) GetOperatorAVSSplit(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
) (uint16, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetOperatorAVSSplit(&bind.CallOpts{Context: ctx}, operator, avs)
}

func (r *ChainReader) GetOperatorPISplit(
	ctx context.Context,
	operator gethcommon.Address,
) (uint16, error) {
	if r.rewardsCoordinator == nil {
		return 0, errors.New("RewardsCoordinator contract not provided")
	}

	return r.rewardsCoordinator.GetOperatorPISplit(&bind.CallOpts{Context: ctx}, operator)
}

func (r *ChainReader) GetAllocatableMagnitude(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddress gethcommon.Address,
) (uint64, error) {
	if r.allocationManager == nil {
		return 0, errors.New("AllocationManager contract not provided")
	}

	return r.allocationManager.GetAllocatableMagnitude(&bind.CallOpts{Context: ctx}, operatorAddress, strategyAddress)
}

func (r *ChainReader) GetMaxMagnitudes(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([]uint64, error) {
	if r.allocationManager == nil {
		return []uint64{}, errors.New("AllocationManager contract not provided")
	}

	return r.allocationManager.GetMaxMagnitudes0(&bind.CallOpts{Context: ctx}, operatorAddress, strategyAddresses)
}

func (r *ChainReader) GetAllocationInfo(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddress gethcommon.Address,
) ([]AllocationInfo, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	opSets, allocationInfo, err := r.allocationManager.GetStrategyAllocations(
		&bind.CallOpts{Context: ctx},
		operatorAddress,
		strategyAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, err
	}

	allocationsInfo := make([]AllocationInfo, len(opSets))
	for i, opSet := range opSets {
		allocationsInfo[i] = AllocationInfo{
			OperatorSetId:    opSet.Id,
			AvsAddress:       opSet.Avs,
			CurrentMagnitude: big.NewInt(int64(allocationInfo[i].CurrentMagnitude)),
			PendingDiff:      allocationInfo[i].PendingDiff,
			EffectBlock:      allocationInfo[i].EffectBlock,
		}
	}

	return allocationsInfo, nil
}

func (r *ChainReader) GetOperatorShares(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([]*big.Int, error) {
	if r.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	return r.delegationManager.GetOperatorShares(&bind.CallOpts{
		Context: ctx,
	}, operatorAddress, strategyAddresses)
}

func (r *ChainReader) GetOperatorsShares(
	ctx context.Context,
	operatorAddresses []gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([][]*big.Int, error) {
	if r.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}
	return r.delegationManager.GetOperatorsShares(&bind.CallOpts{Context: ctx}, operatorAddresses, strategyAddresses)
}

// GetNumOperatorSetsForOperator returns the number of operator sets that an operator is part of
// Doesn't include M2 AVSs
func (r *ChainReader) GetNumOperatorSetsForOperator(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) (*big.Int, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}
	opSets, err := r.allocationManager.GetAllocatedSets(&bind.CallOpts{Context: ctx}, operatorAddress)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(len(opSets))), nil
}

// GetOperatorSetsForOperator returns the list of operator sets that an operator is part of
// Doesn't include M2 AVSs
func (r *ChainReader) GetOperatorSetsForOperator(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) ([]allocationmanager.OperatorSet, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}
	// TODO: we're fetching max int64 operatorSets here. What's the practical limit for timeout by RPC? do we need to
	// paginate?
	return r.allocationManager.GetAllocatedSets(&bind.CallOpts{Context: ctx}, operatorAddress)
}

// IsOperatorRegisteredWithOperatorSet returns if an operator is registered with a specific operator set
func (r *ChainReader) IsOperatorRegisteredWithOperatorSet(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	operatorSet allocationmanager.OperatorSet,
) (bool, error) {
	if operatorSet.Id == 0 {
		// this is an M2 AVS
		if r.avsDirectory == nil {
			return false, errors.New("AVSDirectory contract not provided")
		}

		status, err := r.avsDirectory.AvsOperatorStatus(&bind.CallOpts{Context: ctx}, operatorSet.Avs, operatorAddress)
		// This call should not fail since it's a getter
		if err != nil {
			return false, err
		}

		return status == 1, nil
	} else {
		if r.allocationManager == nil {
			return false, errors.New("AllocationManager contract not provided")
		}
		registeredOperatorSets, err := r.allocationManager.GetRegisteredSets(&bind.CallOpts{Context: ctx}, operatorAddress)
		// This call should not fail since it's a getter
		if err != nil {
			return false, err
		}
		for _, registeredOperatorSet := range registeredOperatorSets {
			if registeredOperatorSet.Id == operatorSet.Id && registeredOperatorSet.Avs == operatorSet.Avs {
				return true, nil
			}
		}

		return false, nil
	}
}

// GetOperatorsForOperatorSet returns the list of operators in a specific operator set
// Not supported for M2 AVSs
func (r *ChainReader) GetOperatorsForOperatorSet(
	ctx context.Context,
	operatorSet allocationmanager.OperatorSet,
) ([]gethcommon.Address, error) {
	if operatorSet.Id == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return nil, errors.New("AllocationManager contract not provided")
		}

		return r.allocationManager.GetMembers(&bind.CallOpts{Context: ctx}, operatorSet)
	}
}

// GetNumOperatorsForOperatorSet returns the number of operators in a specific operator set
func (r *ChainReader) GetNumOperatorsForOperatorSet(
	ctx context.Context,
	operatorSet allocationmanager.OperatorSet,
) (*big.Int, error) {
	if operatorSet.Id == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return nil, errors.New("AllocationManager contract not provided")
		}

		return r.allocationManager.GetMemberCount(&bind.CallOpts{Context: ctx}, operatorSet)
	}
}

// GetStrategiesForOperatorSet returns the list of strategies that an operator set takes into account
// Not supported for M2 AVSs
func (r *ChainReader) GetStrategiesForOperatorSet(
	ctx context.Context,
	operatorSet allocationmanager.OperatorSet,
) ([]gethcommon.Address, error) {
	if operatorSet.Id == 0 {
		return nil, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return nil, errors.New("AllocationManager contract not provided")
		}

		return r.allocationManager.GetStrategiesInOperatorSet(&bind.CallOpts{Context: ctx}, operatorSet)
	}
}

func (r *ChainReader) GetSlashableShares(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	operatorSet allocationmanager.OperatorSet,
	strategies []gethcommon.Address,
) (map[gethcommon.Address]*big.Int, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	currentBlock, err := r.ethClient.BlockNumber(ctx)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, err
	}

	slashableShares, err := r.allocationManager.GetMinimumSlashableStake(
		&bind.CallOpts{Context: ctx},
		operatorSet,
		[]gethcommon.Address{operatorAddress},
		strategies,
		uint32(currentBlock),
	)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, err
	}
	if len(slashableShares) == 0 {
		return nil, errors.New("no slashable shares found for operator")
	}

	slashableShareStrategyMap := make(map[gethcommon.Address]*big.Int)
	for i, strat := range strategies {
		// The reason we use 0 here is because we only have one operator in the list
		slashableShareStrategyMap[strat] = slashableShares[0][i]
	}

	return slashableShareStrategyMap, nil
}

// GetSlashableSharesForOperatorSets returns the strategies the operatorSets take into account, their
// operators, and the minimum amount of shares that are slashable by the operatorSets.
// Not supported for M2 AVSs
func (r *ChainReader) GetSlashableSharesForOperatorSets(
	ctx context.Context,
	operatorSets []allocationmanager.OperatorSet,
) ([]OperatorSetStakes, error) {
	currentBlock, err := r.ethClient.BlockNumber(ctx)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, err
	}
	return r.GetSlashableSharesForOperatorSetsBefore(ctx, operatorSets, uint32(currentBlock))
}

// GetSlashableSharesForOperatorSetsBefore returns the strategies the operatorSets take into account, their
// operators, and the minimum amount of shares slashable by the
// operatorSets before a given timestamp.
// Timestamp must be in the future. Used to underestimate future slashable stake.
// Not supported for M2 AVSs
func (r *ChainReader) GetSlashableSharesForOperatorSetsBefore(
	ctx context.Context,
	operatorSets []allocationmanager.OperatorSet,
	futureBlock uint32,
) ([]OperatorSetStakes, error) {
	operatorSetStakes := make([]OperatorSetStakes, len(operatorSets))
	for i, operatorSet := range operatorSets {
		operators, err := r.GetOperatorsForOperatorSet(ctx, operatorSet)
		if err != nil {
			return nil, err
		}

		strategies, err := r.GetStrategiesForOperatorSet(ctx, operatorSet)
		// If operator setId is 0 will fail on if above
		if err != nil {
			return nil, err
		}

		slashableShares, err := r.allocationManager.GetMinimumSlashableStake(
			&bind.CallOpts{Context: ctx},
			allocationmanager.OperatorSet{
				Id:  operatorSet.Id,
				Avs: operatorSet.Avs,
			},
			operators,
			strategies,
			futureBlock,
		)
		// This call should not fail since it's a getter
		if err != nil {
			return nil, err
		}

		operatorSetStakes[i] = OperatorSetStakes{
			OperatorSet:     operatorSet,
			Strategies:      strategies,
			Operators:       operators,
			SlashableStakes: slashableShares,
		}
	}

	return operatorSetStakes, nil
}

func (r *ChainReader) GetAllocationDelay(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) (uint32, error) {
	if r.allocationManager == nil {
		return 0, errors.New("AllocationManager contract not provided")
	}
	isSet, delay, err := r.allocationManager.GetAllocationDelay(&bind.CallOpts{Context: ctx}, operatorAddress)
	// This call should not fail since it's a getter
	if err != nil {
		return 0, err
	}
	if !isSet {
		return 0, errors.New("allocation delay not set")
	}
	return delay, nil
}

func (r *ChainReader) GetRegisteredSets(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) ([]allocationmanager.OperatorSet, error) {
	if r.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}
	return r.allocationManager.GetRegisteredSets(&bind.CallOpts{Context: ctx}, operatorAddress)
}

func (r *ChainReader) CanCall(
	ctx context.Context,
	request CanCallRequest,
) (CanCallResponse, error) {
	if r.permissionController == nil {
		return CanCallResponse{}, errors.New("PermissionController contract not provided")
	}

	canCall, err := r.permissionController.CanCall(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.AccountAddress,
		request.AppointeeAddress,
		request.Target,
		request.Selector,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return CanCallResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return CanCallResponse{CanCall: canCall}, nil
}

// ListAppointees returns the list of appointees of an account
func (r *ChainReader) ListAppointees(
	ctx context.Context,
	request AppointeesListRequest,
) (AppointeesResponse, error) {
	if r.permissionController == nil {
		return AppointeesResponse{}, errors.New("PermissionController contract not provided")
	}

	appointees, err := r.permissionController.GetAppointees(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.AccountAddress,
		request.TargetAddress,
		request.Selector,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return AppointeesResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return AppointeesResponse{Appointees: appointees}, nil
}

// ListAppointeePermissions returns the list of permissions of an appointee
func (r *ChainReader) ListAppointeePermissions(
	ctx context.Context,
	request AppointeePermissionsListRequest,
) (AppointeePermissionsListResponse, error) {
	if r.permissionController == nil {
		return AppointeePermissionsListResponse{}, errors.New("PermissionController contract not provided")
	}

	targets, selectors, err := r.permissionController.GetAppointeePermissions(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.AccountAddress,
		request.AppointeeAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return AppointeePermissionsListResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return AppointeePermissionsListResponse{TargetAddresses: targets, Selectors: selectors}, nil
}

// ListPendingAdmins returns the list of pending admins of an account
func (r *ChainReader) ListPendingAdmins(
	ctx context.Context,
	request AccountRequest,
) (PendingAdminsResponse, error) {
	if r.permissionController == nil {
		return PendingAdminsResponse{}, errors.New("PermissionController contract not provided")
	}

	pendingAdmins, err := r.permissionController.GetPendingAdmins(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.AccountAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return PendingAdminsResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return PendingAdminsResponse{PendingAdmins: pendingAdmins}, nil
}

// ListAdmins returns the list of admins of an account
func (r *ChainReader) ListAdmins(
	ctx context.Context,
	request AccountRequest,
) (AdminsResponse, error) {
	if r.permissionController == nil {
		return AdminsResponse{}, errors.New("PermissionController contract not provided")
	}

	admins, err := r.permissionController.GetAdmins(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.AccountAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return AdminsResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return AdminsResponse{Admins: admins}, nil
}

// IsPendingAdmin returns if an address is a pending admin of an account
func (r *ChainReader) IsPendingAdmin(
	ctx context.Context,
	request PendingAdminCheckRequest,
) (IsPendingAdminResponse, error) {
	if r.permissionController == nil {
		return IsPendingAdminResponse{}, errors.New("PermissionController contract not provided")
	}

	isPendingAdmin, err := r.permissionController.IsPendingAdmin(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.AccountAddress,
		request.PendingAdminAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return IsPendingAdminResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return IsPendingAdminResponse{IsPendingAdmin: isPendingAdmin}, nil
}

// IsAdmin returns if an address is an admin of an account
func (r *ChainReader) IsAdmin(
	ctx context.Context,
	request AdminCheckRequest,
) (IsAdminResponse, error) {
	if r.permissionController == nil {
		return IsAdminResponse{}, errors.New("PermissionController contract not provided")
	}

	isAdmin, err := r.permissionController.IsAdmin(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.AccountAddress,
		request.AdminAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return IsAdminResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return IsAdminResponse{IsAdmin: isAdmin}, nil
}
