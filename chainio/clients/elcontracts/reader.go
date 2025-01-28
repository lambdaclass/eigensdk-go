package elcontracts

import (
	"context"
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

var errLegacyAVSsNotSupported = Error{3, "Other errors", "Method not supported for legacy AVSs", nil}

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
		wrappedError := CreateForNestedError("NewBindingsFromConfig", err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("DelegationManager")
		return false, wrappedError
	}

	isRegistered, err := r.delegationManager.IsOperator(
		&bind.CallOpts{Context: ctx},
		gethcommon.HexToAddress(operator.Address),
	)
	if err != nil {
		wrappedError := CreateForBindingError("delegationManager.IsOperator", err)
		return false, wrappedError
	}

	return isRegistered, nil
}

// GetStakerShares returns the amount of shares that a staker has in all of the strategies in which they have nonzero
// shares
func (r *ChainReader) GetStakerShares(
	ctx context.Context,
	stakerAddress gethcommon.Address,
) ([]gethcommon.Address, []*big.Int, error) {
	if r.delegationManager == nil {
		wrappedError := CreateErrorForMissingContract("DelegationManager")
		return nil, nil, wrappedError
	}

	addresses, shares, err := r.delegationManager.GetDepositedShares(&bind.CallOpts{Context: ctx}, stakerAddress)
	if err != nil {
		wrappedError := CreateForBindingError("delegationManager.GetDepositedShares", err)
		return nil, nil, wrappedError
	}

	return addresses, shares, nil
}

// GetDelegatedOperator returns the operator that a staker has delegated to
func (r *ChainReader) GetDelegatedOperator(
	ctx context.Context,
	stakerAddress gethcommon.Address,
	blockNumber *big.Int,
) (gethcommon.Address, error) {
	if r.delegationManager == nil {
		wrappedError := CreateErrorForMissingContract("DelegationManager")
		return gethcommon.Address{}, wrappedError
	}

	delegatedOperator, err := r.delegationManager.DelegatedTo(&bind.CallOpts{Context: ctx}, stakerAddress)
	if err != nil {
		wrappedError := CreateForBindingError("delegationManager.DelegatedTo", err)
		return gethcommon.Address{}, wrappedError
	}

	return delegatedOperator, nil
}

func (r *ChainReader) GetOperatorDetails(
	ctx context.Context,
	operator types.Operator,
) (types.Operator, error) {
	if r.delegationManager == nil {
		wrappedError := CreateErrorForMissingContract("DelegationManager")
		return types.Operator{}, wrappedError
	}

	delegationManagerAddress, err := r.delegationManager.DelegationApprover(
		&bind.CallOpts{Context: ctx},
		gethcommon.HexToAddress(operator.Address),
	)
	// This call should not fail since it's a getter
	if err != nil {
		wrappedError := CreateForBindingError("delegationManager.DelegationApprover", err)
		return types.Operator{}, wrappedError
	}

	// Should we check if (allocationManager != nil)?
	isSet, delay, err := r.allocationManager.GetAllocationDelay(
		&bind.CallOpts{
			Context: ctx,
		},
		gethcommon.HexToAddress(operator.Address),
	)
	// This call should not fail
	if err != nil {
		wrappedError := CreateForBindingError("allocationManager.GetAllocationDelay", err)
		return types.Operator{}, wrappedError
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
		wrappedError := CreateForBindingError("strategy contract", err)
		return nil, gethcommon.Address{}, wrappedError
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: ctx})
	if err != nil {
		wrappedError := CreateForBindingError("token contract", err)
		return nil, gethcommon.Address{}, wrappedError
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
		wrappedError := CreateForBindingError("strategy contract", err)
		return nil, nil, gethcommon.Address{}, wrappedError
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: ctx})
	if err != nil {
		wrappedError := CreateForBindingError("token contract", err)
		return nil, nil, gethcommon.Address{}, wrappedError
	}
	contractUnderlyingToken, err := erc20.NewContractIERC20(underlyingTokenAddr, r.ethClient)
	// This call should not fail, if the strategy does not have an underlying token then it would enter the if above
	if err != nil {
		wrappedError := CreateForBindingError("erc20 token contract", err)
		return nil, nil, gethcommon.Address{}, wrappedError
	}
	return contractStrategy, contractUnderlyingToken, underlyingTokenAddr, nil
}

func (r *ChainReader) GetOperatorSharesInStrategy(
	ctx context.Context,
	operatorAddr gethcommon.Address,
	strategyAddr gethcommon.Address,
) (*big.Int, error) {
	if r.delegationManager == nil {
		wrappedError := CreateErrorForMissingContract("DelegationManager")
		return &big.Int{}, wrappedError
	}

	shares, err := r.delegationManager.OperatorShares(
		&bind.CallOpts{Context: ctx},
		operatorAddr,
		strategyAddr,
	)
	if err != nil {
		wrappedError := CreateForBindingError("delegationManager.OperatorShares", err)
		return &big.Int{}, wrappedError
	}

	return shares, nil
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
		wrappedError := CreateErrorForMissingContract("DelegationManager")
		return [32]byte{}, wrappedError
	}

	digestHash, err := r.delegationManager.CalculateDelegationApprovalDigestHash(
		&bind.CallOpts{Context: ctx},
		staker,
		operator,
		delegationApprover,
		approverSalt,
		expiry,
	)
	if err != nil {
		wrappedError := CreateForBindingError("delegationManager.CalculateDelegationApprovalDigestHash", err)
		return [32]byte{}, wrappedError
	}

	return digestHash, nil
}

func (r *ChainReader) CalculateOperatorAVSRegistrationDigestHash(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
	salt [32]byte,
	expiry *big.Int,
) ([32]byte, error) {
	if r.avsDirectory == nil {
		wrappedError := CreateErrorForMissingContract("AVSDirectory")
		return [32]byte{}, wrappedError
	}

	digestHash, err := r.avsDirectory.CalculateOperatorAVSRegistrationDigestHash(
		&bind.CallOpts{Context: ctx},
		operator,
		avs,
		salt,
		expiry,
	)
	if err != nil {
		wrappedError := CreateForBindingError("avsDirectory.CalculateOperatorAVSRegistrationDigestHash", err)
		return [32]byte{}, wrappedError
	}

	return digestHash, nil
}

func (r *ChainReader) GetDistributionRootsLength(ctx context.Context) (*big.Int, error) {
	if r.rewardsCoordinator == nil {
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return &big.Int{}, wrappedError
	}

	distributionRootsLength, err := r.rewardsCoordinator.GetDistributionRootsLength(&bind.CallOpts{Context: ctx})
	if err != nil {
		wrappedError := CreateForBindingError("rewardsCoordinator.GetDistributionRootsLength", err)
		return &big.Int{}, wrappedError
	}

	return distributionRootsLength, nil
}

func (r *ChainReader) CurrRewardsCalculationEndTimestamp(ctx context.Context) (uint32, error) {
	if r.rewardsCoordinator == nil {
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return 0, wrappedError
	}

	endTimestamp, err := r.rewardsCoordinator.CurrRewardsCalculationEndTimestamp(&bind.CallOpts{Context: ctx})
	if err != nil {
		wrappedError := CreateForBindingError("rewardsCoordinator.CurrRewardsCalculationEndTimestamp", err)
		return 0, wrappedError
	}

	return endTimestamp, nil
}

func (r *ChainReader) GetCurrentClaimableDistributionRoot(
	ctx context.Context,
) (rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot, error) {
	if r.rewardsCoordinator == nil {
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot{}, wrappedError
	}

	distributionRoot, err := r.rewardsCoordinator.GetCurrentClaimableDistributionRoot(&bind.CallOpts{Context: ctx})
	if err != nil {
		wrappedError := CreateForBindingError("rewardsCoordinator.GetCurrentClaimableDistributionRoot", err)
		return rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot{}, wrappedError
	}

	return distributionRoot, nil
}

func (r *ChainReader) GetRootIndexFromHash(
	ctx context.Context,
	rootHash [32]byte,
) (uint32, error) {
	if r.rewardsCoordinator == nil {
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return 0, wrappedError
	}

	rootIndex, err := r.rewardsCoordinator.GetRootIndexFromHash(&bind.CallOpts{Context: ctx}, rootHash)
	if err != nil {
		wrappedError := CreateForBindingError("rewardsCoordinator.GetRootIndexFromHash", err)
		return 0, wrappedError
	}

	return rootIndex, nil
}

func (r *ChainReader) GetCumulativeClaimed(
	ctx context.Context,
	earner gethcommon.Address,
	token gethcommon.Address,
) (*big.Int, error) {
	if r.rewardsCoordinator == nil {
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return nil, wrappedError
	}

	cumulativeClaimed, err := r.rewardsCoordinator.CumulativeClaimed(&bind.CallOpts{Context: ctx}, earner, token)
	if err != nil {
		wrappedError := CreateForBindingError("rewardsCoordinator.CumulativeClaimed", err)
		return nil, wrappedError
	}

	return cumulativeClaimed, nil
}

func (r *ChainReader) CheckClaim(
	ctx context.Context,
	claim rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
) (bool, error) {
	if r.rewardsCoordinator == nil {
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return false, wrappedError
	}

	claimChecked, err := r.rewardsCoordinator.CheckClaim(&bind.CallOpts{Context: ctx}, claim)
	if err != nil {
		wrappedError := CreateForBindingError("rewardsCoordinator.CheckClaim", err)
		return false, wrappedError
	}

	return claimChecked, nil
}

func (r *ChainReader) GetOperatorAVSSplit(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
) (uint16, error) {
	if r.rewardsCoordinator == nil {
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return 0, wrappedError
	}

	operatorSplit, err := r.rewardsCoordinator.GetOperatorAVSSplit(&bind.CallOpts{Context: ctx}, operator, avs)
	if err != nil {
		wrappedError := CreateForBindingError("rewardsCoordinator.GetOperatorAVSSplit", err)
		return 0, wrappedError
	}

	return operatorSplit, nil
}

func (r *ChainReader) GetOperatorPISplit(
	ctx context.Context,
	operator gethcommon.Address,
) (uint16, error) {
	if r.rewardsCoordinator == nil {
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return 0, wrappedError
	}

	operatorSplit, err := r.rewardsCoordinator.GetOperatorPISplit(&bind.CallOpts{Context: ctx}, operator)
	if err != nil {
		wrappedError := CreateForBindingError("rewardsCoordinator.GetOperatorPISplit", err)
		return 0, wrappedError
	}

	return operatorSplit, nil
}

func (r *ChainReader) GetAllocatableMagnitude(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddress gethcommon.Address,
) (uint64, error) {
	if r.allocationManager == nil {
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return 0, wrappedError
	}

	allocatableMagnitude, err := r.allocationManager.GetAllocatableMagnitude(
		&bind.CallOpts{Context: ctx},
		operatorAddress,
		strategyAddress,
	)
	if err != nil {
		wrappedError := CreateForBindingError("allocationManager.GetAllocatableMagnitude", err)
		return 0, wrappedError
	}

	return allocatableMagnitude, nil
}

func (r *ChainReader) GetMaxMagnitudes(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([]uint64, error) {
	if r.allocationManager == nil {
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return []uint64{}, wrappedError
	}

	maxMagnitudes, err := r.allocationManager.GetMaxMagnitudes0(
		&bind.CallOpts{Context: ctx},
		operatorAddress,
		strategyAddresses,
	)
	if err != nil {
		wrappedError := CreateForBindingError("allocationManager.GetMaxMagnitudes0", err)
		return []uint64{}, wrappedError
	}

	return maxMagnitudes, nil
}

func (r *ChainReader) GetAllocationInfo(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	strategyAddress gethcommon.Address,
) ([]AllocationInfo, error) {
	if r.allocationManager == nil {
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return nil, wrappedError
	}

	opSets, allocationInfo, err := r.allocationManager.GetStrategyAllocations(
		&bind.CallOpts{Context: ctx},
		operatorAddress,
		strategyAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		wrappedError := CreateForBindingError("allocationManager.GetStrategyAllocations", err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("DelegationManager")
		return nil, wrappedError
	}

	operatorShares, err := r.delegationManager.GetOperatorShares(&bind.CallOpts{
		Context: ctx,
	}, operatorAddress, strategyAddresses)
	if err != nil {
		wrappedError := CreateForBindingError("delegationManager.GetOperatorShares", err)
		return nil, wrappedError
	}

	return operatorShares, nil
}

func (r *ChainReader) GetOperatorsShares(
	ctx context.Context,
	operatorAddresses []gethcommon.Address,
	strategyAddresses []gethcommon.Address,
) ([][]*big.Int, error) {
	if r.delegationManager == nil {
		wrappedError := CreateErrorForMissingContract("DelegationManager")
		return nil, wrappedError
	}

	operatorsShares, err := r.delegationManager.GetOperatorsShares(
		&bind.CallOpts{Context: ctx},
		operatorAddresses,
		strategyAddresses,
	)
	if err != nil {
		wrappedError := CreateForBindingError("delegationManager.GetOperatorsShares", err)
		return nil, wrappedError
	}

	return operatorsShares, nil
}

// GetNumOperatorSetsForOperator returns the number of operator sets that an operator is part of
// Doesn't include M2 AVSs
func (r *ChainReader) GetNumOperatorSetsForOperator(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) (*big.Int, error) {
	if r.allocationManager == nil {
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return nil, wrappedError
	}
	opSets, err := r.allocationManager.GetAllocatedSets(&bind.CallOpts{Context: ctx}, operatorAddress)
	if err != nil {
		wrappedError := CreateForBindingError("allocationManager.GetAllocatedSets", err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return nil, wrappedError
	}
	// TODO: we're fetching max int64 operatorSets here. What's the practical limit for timeout by RPC? do we need to
	// paginate?
	allocatedSets, err := r.allocationManager.GetAllocatedSets(&bind.CallOpts{Context: ctx}, operatorAddress)
	if err != nil {
		wrappedError := CreateForBindingError("allocationManager.GetAllocatedSets", err)
		return nil, wrappedError
	}

	return allocatedSets, nil
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
			wrappedError := CreateErrorForMissingContract("AVSDirectory")
			return false, wrappedError
		}

		status, err := r.avsDirectory.AvsOperatorStatus(&bind.CallOpts{Context: ctx}, operatorSet.Avs, operatorAddress)
		// This call should not fail since it's a getter
		if err != nil {
			wrappedError := CreateForBindingError("avsDirectory.AvsOperatorStatus", err)
			return false, wrappedError
		}

		return status == 1, nil
	} else {
		if r.allocationManager == nil {
			wrappedError := CreateErrorForMissingContract("AllocationManager")
			return false, wrappedError
		}
		registeredOperatorSets, err := r.allocationManager.GetRegisteredSets(&bind.CallOpts{Context: ctx}, operatorAddress)
		// This call should not fail since it's a getter
		if err != nil {
			wrappedError := CreateForBindingError("allocationManager.GetRegisteredSets", err)
			return false, wrappedError
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
			wrappedError := CreateErrorForMissingContract("AllocationManager")
			return nil, wrappedError
		}

		members, err := r.allocationManager.GetMembers(&bind.CallOpts{Context: ctx}, operatorSet)
		if err != nil {
			wrappedError := CreateForBindingError("allocationManager.GetMembers", err)
			return nil, wrappedError
		}

		return members, nil
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
			wrappedError := CreateErrorForMissingContract("AllocationManager")
			return nil, wrappedError
		}

		memberCount, err := r.allocationManager.GetMemberCount(&bind.CallOpts{Context: ctx}, operatorSet)
		if err != nil {
			wrappedError := CreateForBindingError("allocationManager.GetMemberCount", err)
			return nil, wrappedError
		}

		return memberCount, nil
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
			wrappedError := CreateErrorForMissingContract("AllocationManager")
			return nil, wrappedError
		}

		strategiesInSet, err := r.allocationManager.GetStrategiesInOperatorSet(&bind.CallOpts{Context: ctx}, operatorSet)
		if err != nil {
			wrappedError := CreateForBindingError("allocationManager.GetStrategiesInOperatorSet", err)
			return nil, wrappedError
		}

		return strategiesInSet, nil
	}
}

func (r *ChainReader) GetSlashableShares(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	operatorSet allocationmanager.OperatorSet,
	strategies []gethcommon.Address,
) (map[gethcommon.Address]*big.Int, error) {
	if r.allocationManager == nil {
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return nil, wrappedError
	}

	currentBlock, err := r.ethClient.BlockNumber(ctx)
	// This call should not fail since it's a getter
	if err != nil {
		wrappedError := CreateForBindingError("ethClient.BlockNumber", err)
		return nil, wrappedError
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
		wrappedError := CreateForBindingError("allocationManager.GetMinimumSlashableStake", err)
		return nil, wrappedError
	}
	if len(slashableShares) == 0 {
		wrappedError := Error{3, "Other errors", "No slashable shares found for operator", err}
		return nil, wrappedError
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
		wrappedError := CreateForBindingError("ethClient.BlockNumber", err)
		return nil, wrappedError
	}

	operatorSetStakes, err := r.GetSlashableSharesForOperatorSetsBefore(ctx, operatorSets, uint32(currentBlock))
	if err != nil {
		wrappedError := CreateForNestedError("GetSlashableSharesForOperatorSetsBefore", err)
		return nil, wrappedError
	}

	return operatorSetStakes, nil
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
			wrappedError := CreateForNestedError("GetOperatorsForOperatorSet", err)
			return nil, wrappedError
		}

		strategies, err := r.GetStrategiesForOperatorSet(ctx, operatorSet)
		// If operator setId is 0 will fail on if above
		if err != nil {
			wrappedError := CreateForNestedError("GetStrategiesForOperatorSet", err)
			return nil, wrappedError
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
			wrappedError := CreateForBindingError("allocationManager.GetMinimumSlashableStake", err)
			return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return 0, wrappedError
	}
	isSet, delay, err := r.allocationManager.GetAllocationDelay(&bind.CallOpts{Context: ctx}, operatorAddress)
	// This call should not fail since it's a getter
	if err != nil {
		wrappedError := CreateForBindingError("allocationManager.GetAllocationDelay", err)
		return 0, wrappedError
	}
	if !isSet {
		wrappedError := Error{3, "Other errors", "Allocation delay not set", err}
		return 0, wrappedError
	}
	return delay, nil
}

func (r *ChainReader) GetRegisteredSets(
	ctx context.Context,
	operatorAddress gethcommon.Address,
) ([]allocationmanager.OperatorSet, error) {
	if r.allocationManager == nil {
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return nil, wrappedError
	}

	registeredSets, err := r.allocationManager.GetRegisteredSets(&bind.CallOpts{Context: ctx}, operatorAddress)
	if err != nil {
		wrappedError := CreateForBindingError("allocationManager.GetRegisteredSets", err)
		return nil, wrappedError
	}

	return registeredSets, nil
}

func (r *ChainReader) CanCall(
	ctx context.Context,
	accountAddress gethcommon.Address,
	appointeeAddress gethcommon.Address,
	target gethcommon.Address,
	selector [4]byte,
) (bool, error) {
	if r.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return false, wrappedError
	}

	canCall, err := r.permissionController.CanCall(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		appointeeAddress,
		target,
		selector,
	)
	// This call should not fail since it's a getter
	if err != nil {
		wrappedError := CreateForBindingError("permissionController.CanCall", err)
		return false, wrappedError
	}
	return canCall, nil
}

func (r *ChainReader) ListAppointees(
	ctx context.Context,
	accountAddress gethcommon.Address,
	target gethcommon.Address,
	selector [4]byte,
) ([]gethcommon.Address, error) {
	if r.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return nil, wrappedError
	}

	appointees, err := r.permissionController.GetAppointees(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		target,
		selector,
	)
	// This call should not fail since it's a getter
	if err != nil {
		wrappedError := CreateForBindingError("permissionController.GetAppointees", err)
		return nil, wrappedError
	}
	return appointees, nil
}

func (r *ChainReader) ListAppointeePermissions(
	ctx context.Context,
	accountAddress gethcommon.Address,
	appointeeAddress gethcommon.Address,
) ([]gethcommon.Address, [][4]byte, error) {
	if r.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return nil, nil, wrappedError
	}

	targets, selectors, err := r.permissionController.GetAppointeePermissions(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		appointeeAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		wrappedError := CreateForBindingError("permissionController.GetAppointeePermissions", err)
		return nil, nil, wrappedError
	}
	return targets, selectors, nil
}

func (r *ChainReader) ListPendingAdmins(
	ctx context.Context,
	accountAddress gethcommon.Address,
) ([]gethcommon.Address, error) {
	if r.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return nil, wrappedError
	}

	pendingAdmins, err := r.permissionController.GetPendingAdmins(&bind.CallOpts{Context: ctx}, accountAddress)
	// This call should not fail since it's a getter
	if err != nil {
		wrappedError := CreateForBindingError("permissionController.GetPendingAdmins", err)
		return nil, wrappedError
	}
	return pendingAdmins, nil
}

func (r *ChainReader) ListAdmins(
	ctx context.Context,
	accountAddress gethcommon.Address,
) ([]gethcommon.Address, error) {
	if r.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return nil, wrappedError
	}

	pendingAdmins, err := r.permissionController.GetAdmins(&bind.CallOpts{Context: ctx}, accountAddress)
	// This call should not fail since it's a getter
	if err != nil {
		wrappedError := CreateForBindingError("permissionController.GetAdmins", err)
		return nil, wrappedError
	}
	return pendingAdmins, nil
}

func (r *ChainReader) IsPendingAdmin(
	ctx context.Context,
	accountAddress gethcommon.Address,
	pendingAdminAddress gethcommon.Address,
) (bool, error) {
	if r.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return false, wrappedError
	}

	isPendingAdmin, err := r.permissionController.IsPendingAdmin(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		pendingAdminAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		wrappedError := CreateForBindingError("permissionController.IsPendingAdmin", err)
		return false, wrappedError
	}
	return isPendingAdmin, nil
}

func (r *ChainReader) IsAdmin(
	ctx context.Context,
	accountAddress gethcommon.Address,
	adminAddress gethcommon.Address,
) (bool, error) {
	if r.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return false, wrappedError
	}

	isAdmin, err := r.permissionController.IsAdmin(&bind.CallOpts{Context: ctx}, accountAddress, adminAddress)
	// This call should not fail since it's a getter
	if err != nil {
		wrappedError := CreateForBindingError("permissionController.IsAdmin", err)
		return false, wrappedError
	}
	return isAdmin, nil
}
