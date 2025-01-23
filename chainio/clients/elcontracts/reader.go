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
	blockNumber *big.Int,
	request IsOperatorRegisteredRequest,
) (IsOperatorRegisteredResponse, error) {
	if r.delegationManager == nil {
		return IsOperatorRegisteredResponse{}, errors.New("DelegationManager contract not provided")
	}

	isOperatorRegistered, err := r.delegationManager.IsOperator(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return IsOperatorRegisteredResponse{}, utils.WrapError("failed to check if operator is registered", err)
	}

	return IsOperatorRegisteredResponse{IsRegistered: isOperatorRegistered}, nil
}

// GetStakerShares returns the amount of shares that a staker has in all of the strategies in which they have nonzero
// shares
func (r *ChainReader) GetStakerShares(
	ctx context.Context,
	blockNumer *big.Int,
	request GetStakerSharesRequest,
) (GetStakerSharesResponse, error) {
	if r.delegationManager == nil {
		return GetStakerSharesResponse{}, errors.New("DelegationManager contract not provided")
	}

	strategies, shares, err := r.delegationManager.GetDepositedShares(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumer},
		request.StakerAddress,
	)
	if err != nil {
		return GetStakerSharesResponse{}, utils.WrapError("failed to get staker shares", err)
	}

	return GetStakerSharesResponse{StrategiesAddresses: strategies, Shares: shares}, nil
}

// GetDelegatedOperator returns the operator that a staker has delegated to
func (r *ChainReader) GetDelegatedOperator(
	ctx context.Context,
	blockNumber *big.Int,
	request GetDelegatedOperatorRequest,
) (GetDelegatedOperatorResponse, error) {
	if r.delegationManager == nil {
		return GetDelegatedOperatorResponse{}, errors.New("DelegationManager contract not provided")
	}

	operator, err := r.delegationManager.DelegatedTo(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.StakerAddress,
	)
	if err != nil {
		return GetDelegatedOperatorResponse{}, utils.WrapError("failed to get delegated operator", err)
	}

	return GetDelegatedOperatorResponse{OperatorAddress: operator}, nil
}

// TODO: This return type should be types.Operator or GetOperatorDetailsResponse?
func (r *ChainReader) GetOperatorDetails(
	ctx context.Context,
	blockNumber *big.Int,
	request GetOperatorDetailsRequest,
) (GetOperatorDetailsResponse, error) {
	if r.delegationManager == nil {
		return GetOperatorDetailsResponse{}, errors.New("DelegationManager contract not provided")
	}

	delegationManagerAddress, err := r.delegationManager.DelegationApprover(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return GetOperatorDetailsResponse{}, utils.WrapError("failed to get delegation approver", err)
	}

	isSet, delay, err := r.allocationManager.GetAllocationDelay(
		&bind.CallOpts{
			Context:     ctx,
			BlockNumber: blockNumber,
		},
		request.OperatorAddress,
	)
	// This call should not fail
	if err != nil {
		return GetOperatorDetailsResponse{}, utils.WrapError("failed to get allocation delay", err)
	}

	var allocationDelay uint32
	if isSet {
		allocationDelay = delay
	} else {
		allocationDelay = 0
	}

	return GetOperatorDetailsResponse{
		OperatorAddress:           request.OperatorAddress,
		DelegationApproverAddress: delegationManagerAddress,
		AllocationDelay:           allocationDelay,
	}, nil
}

// GetStrategyAndUnderlyingToken returns the strategy contract and the underlying token address
func (r *ChainReader) GetStrategyAndUnderlyingToken(
	ctx context.Context,
	blockNumber *big.Int,
	request GetStrategyAndUnderlyingTokenRequest,
) (GetStrategyAndUnderlyingTokenResponse, error) {
	contractStrategy, err := strategy.NewContractIStrategy(request.StrategyAddress, r.ethClient)
	// This call should not fail since it's an init
	if err != nil {
		return GetStrategyAndUnderlyingTokenResponse{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return GetStrategyAndUnderlyingTokenResponse{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return GetStrategyAndUnderlyingTokenResponse{
		StrategyContract:       contractStrategy,
		UnderlyingTokenAddress: underlyingTokenAddr,
	}, nil
}

// GetStrategyAndUnderlyingERC20Token returns the strategy contract, the erc20 bindings for the underlying token
// and the underlying token address
func (r *ChainReader) GetStrategyAndUnderlyingERC20Token(
	ctx context.Context,
	blockNumber *big.Int,
	request GetStrategyAndUnderlyingERC20TokenRequest,
) (GetStrategyAndUnderlyingERC20TokenResponse, error) {
	contractStrategy, err := strategy.NewContractIStrategy(request.StrategyAddress, r.ethClient)
	// This call should not fail since it's an init
	if err != nil {
		return GetStrategyAndUnderlyingERC20TokenResponse{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return GetStrategyAndUnderlyingERC20TokenResponse{}, utils.WrapError("Failed to fetch token contract", err)
	}
	contractUnderlyingToken, err := erc20.NewContractIERC20(underlyingTokenAddr, r.ethClient)
	// This call should not fail, if the strategy does not have an underlying token then it would enter the if above
	if err != nil {
		return GetStrategyAndUnderlyingERC20TokenResponse{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return GetStrategyAndUnderlyingERC20TokenResponse{
		StrategyContract:       contractStrategy,
		ERC20Bindings:          contractUnderlyingToken,
		UnderlyingTokenAddress: underlyingTokenAddr,
	}, nil
}

func (r *ChainReader) GetOperatorSharesInStrategy(
	ctx context.Context,
	blockNumber *big.Int,
	request GetOperatorSharesInStrategyRequest,
) (GetOperatorSharesInStrategyResponse, error) {
	if r.delegationManager == nil {
		return GetOperatorSharesInStrategyResponse{}, errors.New("DelegationManager contract not provided")
	}

	shares, err := r.delegationManager.OperatorShares(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
		request.StrategyAddress,
	)
	if err != nil {
		return GetOperatorSharesInStrategyResponse{}, utils.WrapError("failed to get operator shares in strategy", err)
	}

	return GetOperatorSharesInStrategyResponse{Shares: shares}, nil
}

func (r *ChainReader) CalculateDelegationApprovalDigestHash(
	ctx context.Context,
	blockNumber *big.Int,
	request CalculateDelegationApprovalDigestHashRequest,
) (CalculateDelegationApprovalDigestHashResponse, error) {
	if r.delegationManager == nil {
		return CalculateDelegationApprovalDigestHashResponse{}, errors.New("DelegationManager contract not provided")
	}

	digestHash, err := r.delegationManager.CalculateDelegationApprovalDigestHash(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.StakerAddress,
		request.OperatorAddress,
		request.ApproverAddress,
		request.ApproverSalt,
		request.Expiry,
	)
	if err != nil {
		return CalculateDelegationApprovalDigestHashResponse{}, utils.WrapError(
			"failed to calculate delegation approval digest hash",
			err,
		)
	}

	return CalculateDelegationApprovalDigestHashResponse{DigestHash: digestHash}, nil
}

func (r *ChainReader) CalculateOperatorAVSRegistrationDigestHash(
	ctx context.Context,
	blockNumber *big.Int,
	request CalculateOperatorAVSRegistrationDigestHashRequest,
) (CalculateOperatorAVSRegistrationDigestHashResponse, error) {
	if r.avsDirectory == nil {
		return CalculateOperatorAVSRegistrationDigestHashResponse{}, errors.New("AVSDirectory contract not provided")
	}

	digestHash, err := r.avsDirectory.CalculateOperatorAVSRegistrationDigestHash(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
		request.AVSAddress,
		request.Salt,
		request.Expiry,
	)
	if err != nil {
		return CalculateOperatorAVSRegistrationDigestHashResponse{}, utils.WrapError(
			"failed to calculate operator AVS registration digest hash",
			err,
		)
	}

	return CalculateOperatorAVSRegistrationDigestHashResponse{DigestHash: digestHash}, nil
}

func (r *ChainReader) GetDistributionRootsLength(
	ctx context.Context,
	blockNumber *big.Int,
) (GetDistributionRootsLengthResponse, error) {
	if r.rewardsCoordinator == nil {
		return GetDistributionRootsLengthResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	rootLength, err := r.rewardsCoordinator.GetDistributionRootsLength(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
	)
	if err != nil {
		return GetDistributionRootsLengthResponse{}, utils.WrapError("failed to get distribution roots length", err)
	}

	return GetDistributionRootsLengthResponse{Length: rootLength}, nil
}

func (r *ChainReader) CurrRewardsCalculationEndTimestamp(
	ctx context.Context,
	blockNumber *big.Int,
) (CurrRewardsCalculationEndTimestampResponse, error) {
	if r.rewardsCoordinator == nil {
		return CurrRewardsCalculationEndTimestampResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	timestamp, err := r.rewardsCoordinator.CurrRewardsCalculationEndTimestamp(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
	)
	if err != nil {
		return CurrRewardsCalculationEndTimestampResponse{}, utils.WrapError(
			"failed to get current rewards calculation end timestamp",
			err,
		)
	}

	return CurrRewardsCalculationEndTimestampResponse{Timestamp: timestamp}, nil
}

func (r *ChainReader) GetCurrentClaimableDistributionRoot(
	ctx context.Context,
	blockNumber *big.Int,
) (GetCurrentClaimableDistributionRootResponse, error) {
	if r.rewardsCoordinator == nil {
		return GetCurrentClaimableDistributionRootResponse{}, errors.New(
			"RewardsCoordinator contract not provided",
		)
	}

	root, err := r.rewardsCoordinator.GetCurrentClaimableDistributionRoot(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
	)
	if err != nil {
		return GetCurrentClaimableDistributionRootResponse{}, utils.WrapError(
			"failed to get current claimable distribution root",
			err,
		)
	}

	return GetCurrentClaimableDistributionRootResponse{DistributionRoot: root}, nil
}

func (r *ChainReader) GetRootIndexFromHash(
	ctx context.Context,
	blockNumber *big.Int,
	request GetRootIndexFromHashRequest,
) (GetRootIndexFromHashResponse, error) {
	if r.rewardsCoordinator == nil {
		return GetRootIndexFromHashResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	rootIndex, err := r.rewardsCoordinator.GetRootIndexFromHash(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.RootHash,
	)
	if err != nil {
		return GetRootIndexFromHashResponse{}, utils.WrapError("failed to get root index from hash", err)
	}

	return GetRootIndexFromHashResponse{RootIndex: rootIndex}, nil
}

func (r *ChainReader) GetCumulativeClaimed(
	ctx context.Context,
	blockNumber *big.Int,
	request GetCumulativeClaimedRequest,
) (GetCumulativeClaimedResponse, error) {
	if r.rewardsCoordinator == nil {
		return GetCumulativeClaimedResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	cumulativeClaimed, err := r.rewardsCoordinator.CumulativeClaimed(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.ClaimerAddress,
		request.TokenAddress,
	)
	if err != nil {
		return GetCumulativeClaimedResponse{}, utils.WrapError("failed to get cumulative claimed", err)
	}

	return GetCumulativeClaimedResponse{CumulativeClaimed: cumulativeClaimed}, nil
}

func (r *ChainReader) CheckClaim(
	ctx context.Context,
	blockNumber *big.Int,
	request CheckClaimRequest,
) (CheckClaimResponse, error) {
	if r.rewardsCoordinator == nil {
		return CheckClaimResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	isClaimed, err := r.rewardsCoordinator.CheckClaim(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.Claim,
	)
	if err != nil {
		return CheckClaimResponse{}, utils.WrapError("failed to check claim", err)
	}

	return CheckClaimResponse{IsValid: isClaimed}, nil
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
	blockNumber *big.Int,
	request GetOperatorAVSSplitRequest,
) (GetOperatorAVSSplitResponse, error) {
	if r.rewardsCoordinator == nil {
		return GetOperatorAVSSplitResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	split, err := r.rewardsCoordinator.GetOperatorPISplit(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return GetOperatorAVSSplitResponse{}, utils.WrapError("failed to get operator PI split", err)
	}

	return GetOperatorAVSSplitResponse{Split: split}, nil
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
	blockNumber *big.Int,
	request GetMaxMagnitudes0Request,
) (GetMaxMagnitudes0Response, error) {
	if r.allocationManager == nil {
		return GetMaxMagnitudes0Response{}, errors.New("AllocationManager contract not provided")
	}

	maxMagnitudes, err := r.allocationManager.GetMaxMagnitudes0(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
		request.StrategiesAddresses,
	)
	if err != nil {
		return GetMaxMagnitudes0Response{}, utils.WrapError("failed to get max magnitudes", err)
	}

	return GetMaxMagnitudes0Response{MaxMagnitudes: maxMagnitudes}, nil
}

func (r *ChainReader) GetAllocationInfo(
	ctx context.Context,
	blockNumber *big.Int,
	request GetAllocationInfoRequest,
) (GetAllocationInfoResponse, error) {
	if r.allocationManager == nil {
		return GetAllocationInfoResponse{}, errors.New("AllocationManager contract not provided")
	}

	opSets, allocationInfo, err := r.allocationManager.GetStrategyAllocations(
		&bind.CallOpts{Context: ctx},
		request.OperatorAddress,
		request.StrategyAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return GetAllocationInfoResponse{}, err
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

	return GetAllocationInfoResponse{AllocationInfo: allocationsInfo}, nil
}

func (r *ChainReader) GetOperatorShares(
	ctx context.Context,
	blockNumber *big.Int,
	request GetOperatorSharesRequest,
) (GetOperatorSharesResponse, error) {
	if r.delegationManager == nil {
		return GetOperatorSharesResponse{}, errors.New("DelegationManager contract not provided")
	}

	shares, err := r.delegationManager.GetOperatorShares(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
		request.StrategiesAddresses)
	if err != nil {
		return GetOperatorSharesResponse{}, utils.WrapError("failed to get operator shares", err)
	}

	return GetOperatorSharesResponse{Shares: shares}, nil
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
	accountAddress gethcommon.Address,
	appointeeAddress gethcommon.Address,
	target gethcommon.Address,
	selector [4]byte,
) (bool, error) {
	canCall, err := r.permissionController.CanCall(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		appointeeAddress,
		target,
		selector,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return false, utils.WrapError("call to permission controller failed", err)
	}
	return canCall, nil
}

func (r *ChainReader) ListAppointees(
	ctx context.Context,
	accountAddress gethcommon.Address,
	target gethcommon.Address,
	selector [4]byte,
) ([]gethcommon.Address, error) {
	appointees, err := r.permissionController.GetAppointees(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		target,
		selector,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, utils.WrapError("call to permission controller failed", err)
	}
	return appointees, nil
}

func (r *ChainReader) ListAppointeePermissions(
	ctx context.Context,
	accountAddress gethcommon.Address,
	appointeeAddress gethcommon.Address,
) ([]gethcommon.Address, [][4]byte, error) {
	targets, selectors, err := r.permissionController.GetAppointeePermissions(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		appointeeAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, nil, utils.WrapError("call to permission controller failed", err)
	}
	return targets, selectors, nil
}

func (r *ChainReader) ListPendingAdmins(
	ctx context.Context,
	accountAddress gethcommon.Address,
) ([]gethcommon.Address, error) {
	pendingAdmins, err := r.permissionController.GetPendingAdmins(&bind.CallOpts{Context: ctx}, accountAddress)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, utils.WrapError("call to permission controller failed", err)
	}
	return pendingAdmins, nil
}

func (r *ChainReader) ListAdmins(
	ctx context.Context,
	accountAddress gethcommon.Address,
) ([]gethcommon.Address, error) {
	pendingAdmins, err := r.permissionController.GetAdmins(&bind.CallOpts{Context: ctx}, accountAddress)
	// This call should not fail since it's a getter
	if err != nil {
		return nil, utils.WrapError("call to permission controller failed", err)
	}
	return pendingAdmins, nil
}

func (r *ChainReader) IsPendingAdmin(
	ctx context.Context,
	accountAddress gethcommon.Address,
	pendingAdminAddress gethcommon.Address,
) (bool, error) {
	isPendingAdmin, err := r.permissionController.IsPendingAdmin(
		&bind.CallOpts{Context: ctx},
		accountAddress,
		pendingAdminAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return false, utils.WrapError("call to permission controller failed", err)
	}
	return isPendingAdmin, nil
}

func (r *ChainReader) IsAdmin(
	ctx context.Context,
	accountAddress gethcommon.Address,
	adminAddress gethcommon.Address,
) (bool, error) {
	isAdmin, err := r.permissionController.IsAdmin(&bind.CallOpts{Context: ctx}, accountAddress, adminAddress)
	// This call should not fail since it's a getter
	if err != nil {
		return false, utils.WrapError("call to permission controller failed", err)
	}
	return isAdmin, nil
}
