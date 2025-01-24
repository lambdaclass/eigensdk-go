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
	blockNumber *big.Int,
	request GetOperatorAVSSplitRequest,
) (GetOperatorAVSSplitResponse, error) {
	if r.rewardsCoordinator == nil {
		return GetOperatorAVSSplitResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	split, err := r.rewardsCoordinator.GetOperatorAVSSplit(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
		request.AvsAddress,
	)
	if err != nil {
		return GetOperatorAVSSplitResponse{}, utils.WrapError("failed to get operator AVS split", err)
	}

	return GetOperatorAVSSplitResponse{Split: split}, nil
}

func (r *ChainReader) GetOperatorPISplit(
	ctx context.Context,
	blockNumber *big.Int,
	request GetOperatorPISplitRequest,
) (GetOperatorPISplitResponse, error) {
	if r.rewardsCoordinator == nil {
		return GetOperatorPISplitResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	split, err := r.rewardsCoordinator.GetOperatorPISplit(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return GetOperatorPISplitResponse{}, utils.WrapError("failed to get operator PI split", err)
	}

	return GetOperatorPISplitResponse{Split: split}, nil
}

func (r *ChainReader) GetAllocatableMagnitude(
	ctx context.Context,
	blockNumber *big.Int,
	request GetAllocatableMagnitudeRequest,
) (GetAllocatableMagnitudeResponse, error) {
	if r.allocationManager == nil {
		return GetAllocatableMagnitudeResponse{}, errors.New("AllocationManager contract not provided")
	}

	magnitude, err := r.allocationManager.GetAllocatableMagnitude(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
		request.StrategyAddress,
	)
	if err != nil {
		return GetAllocatableMagnitudeResponse{}, utils.WrapError("failed to get allocatable magnitude", err)
	}

	return GetAllocatableMagnitudeResponse{AllocatableMagnitude: magnitude}, nil
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
	blockNumber *big.Int,
	request GetOperatorsSharesRequest,
) (GetOperatorsSharesResponse, error) {
	if r.delegationManager == nil {
		return GetOperatorsSharesResponse{}, errors.New("DelegationManager contract not provided")
	}

	shares, err := r.delegationManager.GetOperatorsShares(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorsAddresses,
		request.StrategiesAddresses,
	)
	if err != nil {
		return GetOperatorsSharesResponse{}, utils.WrapError("failed to get operators shares", err)
	}

	return GetOperatorsSharesResponse{Shares: shares}, nil
}

// GetNumOperatorSetsForOperator returns the number of operator sets that an operator is part of
// Doesn't include M2 AVSs
func (r *ChainReader) GetNumOperatorSetsForOperator(
	ctx context.Context,
	blockNumber *big.Int,
	request GetNumOperatorSetsForOperatorRequest,
) (GetNumOperatorSetsForOperatorResponse, error) {
	if r.allocationManager == nil {
		return GetNumOperatorSetsForOperatorResponse{}, errors.New("AllocationManager contract not provided")
	}
	opSets, err := r.allocationManager.GetAllocatedSets(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return GetNumOperatorSetsForOperatorResponse{}, err
	}
	return GetNumOperatorSetsForOperatorResponse{NumOperatorSets: big.NewInt(int64(len(opSets)))}, nil
}

// GetOperatorSetsForOperator returns the list of operator sets that an operator is part of
// Doesn't include M2 AVSs
func (r *ChainReader) GetOperatorSetsForOperator(
	ctx context.Context,
	blockNumber *big.Int,
	request GetOperatorSetsForOperatorRequest,
) (GetOperatorSetsForOperatorResponse, error) {
	if r.allocationManager == nil {
		return GetOperatorSetsForOperatorResponse{}, errors.New("AllocationManager contract not provided")
	}
	// TODO: we're fetching max int64 operatorSets here. What's the practical limit for timeout by RPC? do we need to
	// paginate?
	opSets, err := r.allocationManager.GetAllocatedSets(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return GetOperatorSetsForOperatorResponse{}, err
	}

	return GetOperatorSetsForOperatorResponse{OperatorSets: opSets}, nil
}

// IsOperatorRegisteredWithOperatorSet returns if an operator is registered with a specific operator set
func (r *ChainReader) IsOperatorRegisteredWithOperatorSet(
	ctx context.Context,
	blockNumber *big.Int,
	request IsOperatorRegisteredWithOperatorSetRequest,
) (IsOperatorRegisteredResponse, error) {
	if request.OperatorSet.Id == 0 {
		// this is an M2 AVS
		if r.avsDirectory == nil {
			return IsOperatorRegisteredResponse{}, errors.New("AVSDirectory contract not provided")
		}

		status, err := r.avsDirectory.AvsOperatorStatus(
			&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
			request.OperatorSet.Avs,
			request.OperatorAddress,
		)
		// This call should not fail since it's a getter
		if err != nil {
			return IsOperatorRegisteredResponse{
					IsRegistered: false,
				}, utils.WrapError(
					"failed to check the operator status",
					err,
				)
		}

		return IsOperatorRegisteredResponse{IsRegistered: status == 1}, nil
	} else {
		if r.allocationManager == nil {
			return IsOperatorRegisteredResponse{IsRegistered: false}, errors.New("AllocationManager contract not provided")
		}
		registeredOperatorSets, err := r.allocationManager.GetRegisteredSets(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, request.OperatorAddress)
		// This call should not fail since it's a getter
		if err != nil {
			return IsOperatorRegisteredResponse{IsRegistered: false}, utils.WrapError("failed to get registered operator sets", err)
		}
		for _, registeredOperatorSet := range registeredOperatorSets {
			if registeredOperatorSet.Id == request.OperatorSet.Id && registeredOperatorSet.Avs == request.OperatorSet.Avs {
				return IsOperatorRegisteredResponse{IsRegistered: true}, nil
			}
		}

		return IsOperatorRegisteredResponse{IsRegistered: false}, nil
	}
}

// GetOperatorsForOperatorSet returns the list of operators in a specific operator set
// Not supported for M2 AVSs
func (r *ChainReader) GetOperatorsForOperatorSet(
	ctx context.Context,
	blockNumber *big.Int,
	request GetOperatorsForOperatorSetRequest,
) (GetOperatorsForOperatorSetResponse, error) {
	if request.OperatorSet.Id == 0 {
		return GetOperatorsForOperatorSetResponse{}, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return GetOperatorsForOperatorSetResponse{}, errors.New("AllocationManager contract not provided")
		}
		members, err := r.allocationManager.GetMembers(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, request.OperatorSet)
		if err != nil {
			return GetOperatorsForOperatorSetResponse{}, utils.WrapError("failed to get members", err)
		}

		return GetOperatorsForOperatorSetResponse{Operators: members}, nil
	}
}

// GetNumOperatorsForOperatorSet returns the number of operators in a specific operator set
func (r *ChainReader) GetNumOperatorsForOperatorSet(
	ctx context.Context,
	blockNumber *big.Int,
	request GetNumOperatorsForOperatorSetRequest,
) (GetNumOperatorsForOperatorSetResponse, error) {
	if request.OperatorSet.Id == 0 {
		return GetNumOperatorsForOperatorSetResponse{}, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return GetNumOperatorsForOperatorSetResponse{}, errors.New("AllocationManager contract not provided")
		}

		memberCount, err := r.allocationManager.GetMemberCount(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, request.OperatorSet)
		if err != nil {
			return GetNumOperatorsForOperatorSetResponse{}, utils.WrapError("failed to get member count", err)
		}

		return GetNumOperatorsForOperatorSetResponse{NumOperators: memberCount}, nil
	}
}

// GetStrategiesForOperatorSet returns the list of strategies that an operator set takes into account
// Not supported for M2 AVSs
func (r *ChainReader) GetStrategiesForOperatorSet(
	ctx context.Context,
	blockNumber *big.Int,
	request GetStrategiesForOperatorSetRequest,
) (GetStrategiesForOperatorSetResponse, error) {
	if request.OperatorSet.Id == 0 {
		return GetStrategiesForOperatorSetResponse{}, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return GetStrategiesForOperatorSetResponse{}, errors.New("AllocationManager contract not provided")
		}

		strategies, err := r.allocationManager.GetStrategiesInOperatorSet(
			&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
			request.OperatorSet,
		)
		if err != nil {
			return GetStrategiesForOperatorSetResponse{}, utils.WrapError("failed to get strategies", err)
		}

		return GetStrategiesForOperatorSetResponse{StrategiesAddresses: strategies}, nil
	}
}

func (r *ChainReader) GetSlashableShares(
	ctx context.Context,
	blockNumber *big.Int,
	request GetSlashableSharesRequest,
) (GetSlashableSharesResponse, error) {
	if r.allocationManager == nil {
		return GetSlashableSharesResponse{}, errors.New("AllocationManager contract not provided")
	}

	// TODO: Is necessary to get the block number here? Or should we use the one passed as argument?
	// currentBlock, err := r.ethClient.BlockNumber(ctx)
	// This call should not fail since it's a getter
	// if err != nil {
	// 	return GetSlashableSharesResponse{}, err
	// }

	slashableShares, err := r.allocationManager.GetMinimumSlashableStake(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorSet,
		[]gethcommon.Address{request.OperatorAddress},
		request.StrategiesAddresses,
		uint32(blockNumber.Uint64()),
	)
	// This call should not fail since it's a getter
	if err != nil {
		return GetSlashableSharesResponse{}, err
	}
	if len(slashableShares) == 0 {
		return GetSlashableSharesResponse{}, errors.New("no slashable shares found for operator")
	}

	slashableShareStrategyMap := make(map[gethcommon.Address]*big.Int)
	for i, strat := range request.StrategiesAddresses {
		// The reason we use 0 here is because we only have one operator in the list
		slashableShareStrategyMap[strat] = slashableShares[0][i]
	}

	return GetSlashableSharesResponse{SlashableShares: slashableShareStrategyMap}, nil
}

// GetSlashableSharesForOperatorSets returns the strategies the operatorSets take into account, their
// operators, and the minimum amount of shares that are slashable by the operatorSets.
// Not supported for M2 AVSs
// VOY X ACAAAA
func (r *ChainReader) GetSlashableSharesForOperatorSets(
	ctx context.Context,
	blockNumber *big.Int,
	request GetSlashableSharesForOperatorSetsRequest,
) (GetSlashableSharesForOperatorSetsResponse, error) {
	// TODO: Is necessary to get the block number here? Or should we use the one passed as argument?
	// currentBlock, err := r.ethClient.BlockNumber(ctx)
	// // This call should not fail since it's a getter
	// if err != nil {
	// 	return nil, err
	// }
	return r.GetSlashableSharesForOperatorSetsBefore(ctx, uint32(blockNumber.Uint64()), request)
}

// GetSlashableSharesForOperatorSetsBefore returns the strategies the operatorSets take into account, their
// operators, and the minimum amount of shares slashable by the
// operatorSets before a given timestamp.
// Timestamp must be in the future. Used to underestimate future slashable stake.
// Not supported for M2 AVSs

// TODO: Should we use the block number instead of the futureBlock?
func (r *ChainReader) GetSlashableSharesForOperatorSetsBefore(
	ctx context.Context,
	futureBlock uint32,
	request GetSlashableSharesForOperatorSetsRequest,
) (GetSlashableSharesForOperatorSetsResponse, error) {
	operatorSetStakes := make([]OperatorSetStakes, len(request.OperatorSets))
	for i, operatorSet := range request.OperatorSets {
		requestOperator := GetOperatorsForOperatorSetRequest{
			OperatorSet: operatorSet,
		}
		responseOperators, err := r.GetOperatorsForOperatorSet(ctx, nil, requestOperator)
		if err != nil {
			return GetSlashableSharesForOperatorSetsResponse{}, utils.WrapError(
				"failed to get operators for operator set",
				err,
			)
		}

		requestStrategies := GetStrategiesForOperatorSetRequest{
			OperatorSet: operatorSet,
		}
		// blockNumber should be nil or futureBlock?
		responseStrategies, err := r.GetStrategiesForOperatorSet(ctx, nil, requestStrategies)
		// If operator setId is 0 will fail on if above
		if err != nil {
			return GetSlashableSharesForOperatorSetsResponse{}, utils.WrapError(
				"failed to get strategies for operator set",
				err,
			)
		}

		slashableShares, err := r.allocationManager.GetMinimumSlashableStake(
			&bind.CallOpts{Context: ctx},
			allocationmanager.OperatorSet{
				Id:  operatorSet.Id,
				Avs: operatorSet.Avs,
			},
			responseOperators.Operators,
			responseStrategies.StrategiesAddresses,
			futureBlock,
		)
		// This call should not fail since it's a getter
		if err != nil {
			return GetSlashableSharesForOperatorSetsResponse{}, utils.WrapError(
				"failed to get minimum slashable stake",
				err,
			)
		}

		operatorSetStakes[i] = OperatorSetStakes{
			OperatorSet:     operatorSet,
			Strategies:      responseStrategies.StrategiesAddresses,
			Operators:       responseOperators.Operators,
			SlashableStakes: slashableShares,
		}
	}

	return GetSlashableSharesForOperatorSetsResponse{OperatorSetStakes: operatorSetStakes}, nil
}

func (r *ChainReader) GetAllocationDelay(
	ctx context.Context,
	blockNumber *big.Int,
	request GetAllocationDelayRequest,
) (GetAllocationDelayResponse, error) {
	if r.allocationManager == nil {
		return GetAllocationDelayResponse{}, errors.New("AllocationManager contract not provided")
	}
	isSet, delay, err := r.allocationManager.GetAllocationDelay(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return GetAllocationDelayResponse{}, utils.WrapError("failed to get allocation delay", err)
	}
	if !isSet {
		return GetAllocationDelayResponse{}, errors.New("allocation delay not set")
	}
	return GetAllocationDelayResponse{AllocationDelay: delay}, nil
}

func (r *ChainReader) GetRegisteredSets(
	ctx context.Context,
	blockNumber *big.Int,
	request GetRegisteredSetsRequest,
) (GetRegisteredSetsResponse, error) {
	if r.allocationManager == nil {
		return GetRegisteredSetsResponse{}, errors.New("AllocationManager contract not provided")
	}
	reigsteredSets, err := r.allocationManager.GetRegisteredSets(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return GetRegisteredSetsResponse{}, utils.WrapError("failed to get registered sets", err)
	}

	return GetRegisteredSetsResponse{OperatorSets: reigsteredSets}, nil
}

func (r *ChainReader) CanCall(
	ctx context.Context,
	blockNumber *big.Int,
	request CanCallRequest,
) (CanCallResponse, error) {
	if r.permissionController == nil {
		return CanCallResponse{}, errors.New("PermissionController contract not provided")
	}

	canCall, err := r.permissionController.CanCall(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
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

func (r *ChainReader) ListAppointees(
	ctx context.Context,
	blockNumber *big.Int,
	request ListAppointeesRequest,
) (ListAppointeesResponse, error) {
	if r.permissionController == nil {
		return ListAppointeesResponse{}, errors.New("PermissionController contract not provided")
	}

	appointees, err := r.permissionController.GetAppointees(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.AccountAddress,
		request.Target,
		request.Select,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return ListAppointeesResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return ListAppointeesResponse{Appointees: appointees}, nil
}

func (r *ChainReader) ListAppointeePermissions(
	ctx context.Context,
	blockNumber *big.Int,
	request ListAppointeePermissionsRequest,
) (ListAppointeePermissionsResponse, error) {
	if r.permissionController == nil {
		return ListAppointeePermissionsResponse{}, errors.New("PermissionController contract not provided")
	}

	targets, selectors, err := r.permissionController.GetAppointeePermissions(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.AccountAddress,
		request.AppointeeAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return ListAppointeePermissionsResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return ListAppointeePermissionsResponse{AppointeeAddress: targets, Selector: selectors}, nil
}

func (r *ChainReader) ListPendingAdmins(
	ctx context.Context,
	blockNumber *big.Int,
	request ListPendingAdminsRequest,
) (ListPendingAdminsResponse, error) {
	if r.permissionController == nil {
		return ListPendingAdminsResponse{}, errors.New("PermissionController contract not provided")
	}

	pendingAdmins, err := r.permissionController.GetPendingAdmins(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.AccountAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return ListPendingAdminsResponse{}, utils.WrapError("call to permission controller failed", err)
	}

	return ListPendingAdminsResponse{PendingAdmins: pendingAdmins}, nil
}

func (r *ChainReader) ListAdmins(
	ctx context.Context,
	blockNumber *big.Int,
	request ListAdminsRequest,
) (ListAdminsResponse, error) {
	if r.permissionController == nil {
		return ListAdminsResponse{}, errors.New("PermissionController contract not provided")
	}

	admins, err := r.permissionController.GetAdmins(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.AccountAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return ListAdminsResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return ListAdminsResponse{Admins: admins}, nil
}

func (r *ChainReader) IsPendingAdmin(
	ctx context.Context,
	blockNumber *big.Int,
	request IsPendingAdminRequest,
) (IsPendingAdminResponse, error) {
	if r.permissionController == nil {
		return IsPendingAdminResponse{}, errors.New("PermissionController contract not provided")
	}

	isPendingAdmin, err := r.permissionController.IsPendingAdmin(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.AccountAddress,
		request.PendingAdminAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return IsPendingAdminResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return IsPendingAdminResponse{IsPendingAdmin: isPendingAdmin}, nil
}

func (r *ChainReader) IsAdmin(
	ctx context.Context,
	blockNumber *big.Int,
	request IsAdminRequest,
) (IsAdminResponse, error) {
	if r.permissionController == nil {
		return IsAdminResponse{}, errors.New("PermissionController contract not provided")
	}

	isAdmin, err := r.permissionController.IsAdmin(
		&bind.CallOpts{Context: ctx, BlockNumber: blockNumber},
		request.AccountAddress,
		request.AdminAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return IsAdminResponse{}, utils.WrapError("call to permission controller failed", err)
	}
	return IsAdminResponse{IsAdmin: isAdmin}, nil
}
