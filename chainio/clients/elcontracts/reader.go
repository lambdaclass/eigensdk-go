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
	request IsOperatorRegisteredRequest,
) (IsOperatorRegisteredResponse, error) {
	if r.delegationManager == nil {
		return IsOperatorRegisteredResponse{}, errors.New("DelegationManager contract not provided")
	}

	isOperatorRegistered, err := r.delegationManager.IsOperator(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
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
	request StakerSharesRequest,
) (StakerSharesResponse, error) {
	if r.delegationManager == nil {
		return StakerSharesResponse{}, errors.New("DelegationManager contract not provided")
	}

	strategies, shares, err := r.delegationManager.GetDepositedShares(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.StakerAddress,
	)
	if err != nil {
		return StakerSharesResponse{}, utils.WrapError("failed to get staker shares", err)
	}

	return StakerSharesResponse{StrategiesAddresses: strategies, Shares: shares}, nil
}

// GetDelegatedOperator returns the operator that a staker has delegated to
func (r *ChainReader) GetDelegatedOperator(
	ctx context.Context,
	request DelegatedOperatorRequest,
) (DelegatedOperatorResponse, error) {
	if r.delegationManager == nil {
		return DelegatedOperatorResponse{}, errors.New("DelegationManager contract not provided")
	}

	operator, err := r.delegationManager.DelegatedTo(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.StakerAddress,
	)
	if err != nil {
		return DelegatedOperatorResponse{}, utils.WrapError("failed to get delegated operator", err)
	}

	return DelegatedOperatorResponse{OperatorAddress: operator}, nil
}

func (r *ChainReader) GetOperatorDetails(
	ctx context.Context,
	request OperatorDetailsRequest,
) (OperatorDetailsResponse, error) {
	if r.delegationManager == nil {
		return OperatorDetailsResponse{}, errors.New("DelegationManager contract not provided")
	}

	delegationManagerAddress, err := r.delegationManager.DelegationApprover(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return OperatorDetailsResponse{}, utils.WrapError("failed to get delegation approver", err)
	}

	isSet, delay, err := r.allocationManager.GetAllocationDelay(
		&bind.CallOpts{
			Context:     ctx,
			BlockNumber: request.BlockNumber,
		},
		request.OperatorAddress,
	)
	// This call should not fail
	if err != nil {
		return OperatorDetailsResponse{}, utils.WrapError("failed to get allocation delay", err)
	}

	var allocationDelay uint32
	if isSet {
		allocationDelay = delay
	} else {
		allocationDelay = 0
	}

	return OperatorDetailsResponse{
		OperatorAddress:           request.OperatorAddress,
		DelegationApproverAddress: delegationManagerAddress,
		AllocationDelay:           allocationDelay,
	}, nil
}

// GetStrategyAndUnderlyingToken returns the strategy contract and the underlying token address
func (r *ChainReader) GetStrategyAndUnderlyingToken(
	ctx context.Context,
	request StrategyAndUnderlyingTokenRequest,
) (StrategyAndUnderlyingTokenResponse, error) {
	contractStrategy, err := strategy.NewContractIStrategy(request.StrategyAddress, r.ethClient)
	// This call should not fail since it's an init
	if err != nil {
		return StrategyAndUnderlyingTokenResponse{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
	)
	if err != nil {
		return StrategyAndUnderlyingTokenResponse{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return StrategyAndUnderlyingTokenResponse{
		StrategyContract:       contractStrategy,
		UnderlyingTokenAddress: underlyingTokenAddr,
	}, nil
}

// GetStrategyAndUnderlyingERC20Token returns the strategy contract, the erc20 bindings for the underlying token
// and the underlying token address
func (r *ChainReader) GetStrategyAndUnderlyingERC20Token(
	ctx context.Context,
	request StrategyAndUnderlyingERC20TokenRequest,
) (StrategyAndUnderlyingERC20TokenResponse, error) {
	contractStrategy, err := strategy.NewContractIStrategy(request.StrategyAddress, r.ethClient)
	// This call should not fail since it's an init
	if err != nil {
		return StrategyAndUnderlyingERC20TokenResponse{}, utils.WrapError("Failed to fetch strategy contract", err)
	}
	underlyingTokenAddr, err := contractStrategy.UnderlyingToken(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
	)
	if err != nil {
		return StrategyAndUnderlyingERC20TokenResponse{}, utils.WrapError("Failed to fetch token contract", err)
	}
	contractUnderlyingToken, err := erc20.NewContractIERC20(underlyingTokenAddr, r.ethClient)
	// This call should not fail, if the strategy does not have an underlying token then it would enter the if above
	if err != nil {
		return StrategyAndUnderlyingERC20TokenResponse{}, utils.WrapError("Failed to fetch token contract", err)
	}
	return StrategyAndUnderlyingERC20TokenResponse{
		StrategyContract:       contractStrategy,
		ERC20Bindings:          contractUnderlyingToken,
		UnderlyingTokenAddress: underlyingTokenAddr,
	}, nil
}

func (r *ChainReader) GetOperatorSharesInStrategy(
	ctx context.Context,
	request OperatorSharesInStrategyRequest,
) (OperatorSharesInStrategyResponse, error) {
	if r.delegationManager == nil {
		return OperatorSharesInStrategyResponse{}, errors.New("DelegationManager contract not provided")
	}

	shares, err := r.delegationManager.OperatorShares(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
		request.StrategyAddress,
	)
	if err != nil {
		return OperatorSharesInStrategyResponse{}, utils.WrapError("failed to get operator shares in strategy", err)
	}

	return OperatorSharesInStrategyResponse{Shares: shares}, nil
}

func (r *ChainReader) CalculateDelegationApprovalDigestHash(
	ctx context.Context,
	request CalculateDelegationApprovalDigestHashRequest,
) (CalculateDelegationApprovalDigestHashResponse, error) {
	if r.delegationManager == nil {
		return CalculateDelegationApprovalDigestHashResponse{}, errors.New("DelegationManager contract not provided")
	}

	digestHash, err := r.delegationManager.CalculateDelegationApprovalDigestHash(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
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
	request CalculateOperatorAVSRegistrationDigestHashRequest,
) (CalculateOperatorAVSRegistrationDigestHashResponse, error) {
	if r.avsDirectory == nil {
		return CalculateOperatorAVSRegistrationDigestHashResponse{}, errors.New("AVSDirectory contract not provided")
	}

	digestHash, err := r.avsDirectory.CalculateOperatorAVSRegistrationDigestHash(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
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
	request DistributionRootsLengthRequest,
) (DistributionRootsLengthResponse, error) {
	if r.rewardsCoordinator == nil {
		return DistributionRootsLengthResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	rootLength, err := r.rewardsCoordinator.GetDistributionRootsLength(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
	)
	if err != nil {
		return DistributionRootsLengthResponse{}, utils.WrapError("failed to get distribution roots length", err)
	}

	return DistributionRootsLengthResponse{Length: rootLength}, nil
}

func (r *ChainReader) CurrRewardsCalculationEndTimestamp(
	ctx context.Context,
	request CurrRewardsCalculationEndTimestampRequest,
) (CurrRewardsCalculationEndTimestampResponse, error) {
	if r.rewardsCoordinator == nil {
		return CurrRewardsCalculationEndTimestampResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	timestamp, err := r.rewardsCoordinator.CurrRewardsCalculationEndTimestamp(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
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
	request CurrentClaimableDistributionRootRequest,
) (CurrentClaimableDistributionRootResponse, error) {
	if r.rewardsCoordinator == nil {
		return CurrentClaimableDistributionRootResponse{}, errors.New(
			"RewardsCoordinator contract not provided",
		)
	}

	root, err := r.rewardsCoordinator.GetCurrentClaimableDistributionRoot(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
	)
	if err != nil {
		return CurrentClaimableDistributionRootResponse{}, utils.WrapError(
			"failed to get current claimable distribution root",
			err,
		)
	}

	return CurrentClaimableDistributionRootResponse{DistributionRoot: root}, nil
}

func (r *ChainReader) GetRootIndexFromHash(
	ctx context.Context,
	request RootIndexFromHashRequest,
) (RootIndexFromHashResponse, error) {
	if r.rewardsCoordinator == nil {
		return RootIndexFromHashResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	rootIndex, err := r.rewardsCoordinator.GetRootIndexFromHash(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.RootHash,
	)
	if err != nil {
		return RootIndexFromHashResponse{}, utils.WrapError("failed to get root index from hash", err)
	}

	return RootIndexFromHashResponse{RootIndex: rootIndex}, nil
}

func (r *ChainReader) GetCumulativeClaimed(
	ctx context.Context,
	request CumulativeClaimedRequest,
) (CumulativeClaimedResponse, error) {
	if r.rewardsCoordinator == nil {
		return CumulativeClaimedResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	cumulativeClaimed, err := r.rewardsCoordinator.CumulativeClaimed(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.ClaimerAddress,
		request.TokenAddress,
	)
	if err != nil {
		return CumulativeClaimedResponse{}, utils.WrapError("failed to get cumulative claimed", err)
	}

	return CumulativeClaimedResponse{CumulativeClaimed: cumulativeClaimed}, nil
}

func (r *ChainReader) CheckClaim(
	ctx context.Context,
	request CheckClaimRequest,
) (CheckClaimResponse, error) {
	if r.rewardsCoordinator == nil {
		return CheckClaimResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	isClaimed, err := r.rewardsCoordinator.CheckClaim(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.Claim,
	)
	if err != nil {
		return CheckClaimResponse{}, utils.WrapError("failed to check claim", err)
	}

	return CheckClaimResponse{IsValid: isClaimed}, nil
}

func (r *ChainReader) GetOperatorAVSSplit(
	ctx context.Context,
	request OperatorAVSSplitRequest,
) (OperatorAVSSplitResponse, error) {
	if r.rewardsCoordinator == nil {
		return OperatorAVSSplitResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	split, err := r.rewardsCoordinator.GetOperatorAVSSplit(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
		request.AvsAddress,
	)
	if err != nil {
		return OperatorAVSSplitResponse{}, utils.WrapError("failed to get operator AVS split", err)
	}

	return OperatorAVSSplitResponse{Split: split}, nil
}

func (r *ChainReader) GetOperatorPISplit(
	ctx context.Context,
	request OperatorPISplitRequest,
) (OperatorPISplitResponse, error) {
	if r.rewardsCoordinator == nil {
		return OperatorPISplitResponse{}, errors.New("RewardsCoordinator contract not provided")
	}

	split, err := r.rewardsCoordinator.GetOperatorPISplit(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return OperatorPISplitResponse{}, utils.WrapError("failed to get operator PI split", err)
	}

	return OperatorPISplitResponse{Split: split}, nil
}

func (r *ChainReader) GetAllocatableMagnitude(
	ctx context.Context,
	request AllocatableMagnitudeRequest,
) (AllocatableMagnitudeResponse, error) {
	if r.allocationManager == nil {
		return AllocatableMagnitudeResponse{}, errors.New("AllocationManager contract not provided")
	}

	magnitude, err := r.allocationManager.GetAllocatableMagnitude(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
		request.StrategyAddress,
	)
	if err != nil {
		return AllocatableMagnitudeResponse{}, utils.WrapError("failed to get allocatable magnitude", err)
	}

	return AllocatableMagnitudeResponse{AllocatableMagnitude: magnitude}, nil
}

func (r *ChainReader) GetMaxMagnitudes(
	ctx context.Context,
	request MaxMagnitudes0Request,
) (MaxMagnitudes0Response, error) {
	if r.allocationManager == nil {
		return MaxMagnitudes0Response{}, errors.New("AllocationManager contract not provided")
	}

	maxMagnitudes, err := r.allocationManager.GetMaxMagnitudes0(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
		request.StrategiesAddresses,
	)
	if err != nil {
		return MaxMagnitudes0Response{}, utils.WrapError("failed to get max magnitudes", err)
	}

	return MaxMagnitudes0Response{MaxMagnitudes: maxMagnitudes}, nil
}

func (r *ChainReader) GetAllocationInfo(
	ctx context.Context,
	request AllocationInfoRequest,
) (AllocationInfoResponse, error) {
	if r.allocationManager == nil {
		return AllocationInfoResponse{}, errors.New("AllocationManager contract not provided")
	}

	opSets, allocationInfo, err := r.allocationManager.GetStrategyAllocations(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
		request.StrategyAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return AllocationInfoResponse{}, err
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

	return AllocationInfoResponse{AllocationInfo: allocationsInfo}, nil
}

func (r *ChainReader) GetOperatorShares(
	ctx context.Context,
	request OperatorSharesRequest,
) (OperatorSharesResponse, error) {
	if r.delegationManager == nil {
		return OperatorSharesResponse{}, errors.New("DelegationManager contract not provided")
	}

	shares, err := r.delegationManager.GetOperatorShares(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
		request.StrategiesAddresses)
	if err != nil {
		return OperatorSharesResponse{}, utils.WrapError("failed to get operator shares", err)
	}

	return OperatorSharesResponse{Shares: shares}, nil
}

func (r *ChainReader) GetOperatorsShares(
	ctx context.Context,
	request OperatorsSharesRequest,
) (OperatorsSharesResponse, error) {
	if r.delegationManager == nil {
		return OperatorsSharesResponse{}, errors.New("DelegationManager contract not provided")
	}

	shares, err := r.delegationManager.GetOperatorsShares(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorsAddresses,
		request.StrategiesAddresses,
	)
	if err != nil {
		return OperatorsSharesResponse{}, utils.WrapError("failed to get operators shares", err)
	}

	return OperatorsSharesResponse{Shares: shares}, nil
}

// GetNumOperatorSetsForOperator returns the number of operator sets that an operator is part of
// Doesn't include M2 AVSs
func (r *ChainReader) GetNumOperatorSetsForOperator(
	ctx context.Context,
	request NumOperatorSetsForOperatorRequest,
) (NumOperatorSetsForOperatorResponse, error) {
	if r.allocationManager == nil {
		return NumOperatorSetsForOperatorResponse{}, errors.New("AllocationManager contract not provided")
	}
	opSets, err := r.allocationManager.GetAllocatedSets(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return NumOperatorSetsForOperatorResponse{}, err
	}
	return NumOperatorSetsForOperatorResponse{NumOperatorSets: big.NewInt(int64(len(opSets)))}, nil
}

// GetOperatorSetsForOperator returns the list of operator sets that an operator is part of
// Doesn't include M2 AVSs
func (r *ChainReader) GetOperatorSetsForOperator(
	ctx context.Context,
	request OperatorSetsForOperatorRequest,
) (OperatorSetsForOperatorResponse, error) {
	if r.allocationManager == nil {
		return OperatorSetsForOperatorResponse{}, errors.New("AllocationManager contract not provided")
	}
	// TODO: we're fetching max int64 operatorSets here. What's the practical limit for timeout by RPC? do we need to
	// paginate?
	opSets, err := r.allocationManager.GetAllocatedSets(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return OperatorSetsForOperatorResponse{}, err
	}

	return OperatorSetsForOperatorResponse{OperatorSets: opSets}, nil
}

// IsOperatorRegisteredWithOperatorSet returns if an operator is registered with a specific operator set
func (r *ChainReader) IsOperatorRegisteredWithOperatorSet(
	ctx context.Context,
	request IsOperatorRegisteredWithOperatorSetRequest,
) (IsOperatorRegisteredResponse, error) {
	if request.OperatorSet.Id == 0 {
		// this is an M2 AVS
		if r.avsDirectory == nil {
			return IsOperatorRegisteredResponse{}, errors.New("AVSDirectory contract not provided")
		}

		status, err := r.avsDirectory.AvsOperatorStatus(
			&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
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
		registeredOperatorSets, err := r.allocationManager.GetRegisteredSets(&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber}, request.OperatorAddress)
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
	request OperatorsForOperatorSetRequest,
) (OperatorsForOperatorSetResponse, error) {
	if request.OperatorSet.Id == 0 {
		return OperatorsForOperatorSetResponse{}, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return OperatorsForOperatorSetResponse{}, errors.New("AllocationManager contract not provided")
		}
		members, err := r.allocationManager.GetMembers(&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber}, request.OperatorSet)
		if err != nil {
			return OperatorsForOperatorSetResponse{}, utils.WrapError("failed to get members", err)
		}

		return OperatorsForOperatorSetResponse{Operators: members}, nil
	}
}

// GetNumOperatorsForOperatorSet returns the number of operators in a specific operator set
func (r *ChainReader) GetNumOperatorsForOperatorSet(
	ctx context.Context,
	request NumOperatorsForOperatorSetRequest,
) (NumOperatorsForOperatorSetResponse, error) {
	if request.OperatorSet.Id == 0 {
		return NumOperatorsForOperatorSetResponse{}, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return NumOperatorsForOperatorSetResponse{}, errors.New("AllocationManager contract not provided")
		}

		memberCount, err := r.allocationManager.GetMemberCount(&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber}, request.OperatorSet)
		if err != nil {
			return NumOperatorsForOperatorSetResponse{}, utils.WrapError("failed to get member count", err)
		}

		return NumOperatorsForOperatorSetResponse{NumOperators: memberCount}, nil
	}
}

// GetStrategiesForOperatorSet returns the list of strategies that an operator set takes into account
// Not supported for M2 AVSs
func (r *ChainReader) GetStrategiesForOperatorSet(
	ctx context.Context,
	request StrategiesForOperatorSetRequest,
) (StrategiesForOperatorSetResponse, error) {
	if request.OperatorSet.Id == 0 {
		return StrategiesForOperatorSetResponse{}, errLegacyAVSsNotSupported
	} else {
		if r.allocationManager == nil {
			return StrategiesForOperatorSetResponse{}, errors.New("AllocationManager contract not provided")
		}

		strategies, err := r.allocationManager.GetStrategiesInOperatorSet(
			&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
			request.OperatorSet,
		)
		if err != nil {
			return StrategiesForOperatorSetResponse{}, utils.WrapError("failed to get strategies", err)
		}

		return StrategiesForOperatorSetResponse{StrategiesAddresses: strategies}, nil
	}
}

func (r *ChainReader) GetSlashableShares(
	ctx context.Context,
	request SlashableSharesRequest,
) (SlashableSharesResponse, error) {
	if r.allocationManager == nil {
		return SlashableSharesResponse{}, errors.New("AllocationManager contract not provided")
	}

	// TODO: Is necessary to get the block number here? Or should we use the one passed as argument?
	currentBlock, err := r.ethClient.BlockNumber(ctx)
	// This call should not fail since it's a getter
	if err != nil {
		return SlashableSharesResponse{}, err
	}

	slashableShares, err := r.allocationManager.GetMinimumSlashableStake(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorSet,
		[]gethcommon.Address{request.OperatorAddress},
		request.StrategiesAddresses,
		uint32(currentBlock),
	)
	// This call should not fail since it's a getter
	if err != nil {
		return SlashableSharesResponse{}, err
	}
	if len(slashableShares) == 0 {
		return SlashableSharesResponse{}, errors.New("no slashable shares found for operator")
	}

	slashableShareStrategyMap := make(map[gethcommon.Address]*big.Int)
	for i, strat := range request.StrategiesAddresses {
		// The reason we use 0 here is because we only have one operator in the list
		slashableShareStrategyMap[strat] = slashableShares[0][i]
	}

	return SlashableSharesResponse{SlashableShares: slashableShareStrategyMap}, nil
}

// GetSlashableSharesForOperatorSets returns the strategies the operatorSets take into account, their
// operators, and the minimum amount of shares that are slashable by the operatorSets.
// Not supported for M2 AVSs
func (r *ChainReader) GetSlashableSharesForOperatorSets(
	ctx context.Context,
	request SlashableSharesForOperatorSetsRequest,
) (SlashableSharesForOperatorSetsResponse, error) {
	currentBlock, err := r.ethClient.BlockNumber(ctx)
	// This call should not fail since it's a getter
	if err != nil {
		return SlashableSharesForOperatorSetsResponse{}, err
	}

	requestBefore := SlashableSharesForOperatorSetsBeforeRequest{
		BlockNumber:  request.BlockNumber,
		OperatorSets: request.OperatorSets,
		FutureBlock:  uint32(currentBlock),
	}

	resp, err := r.GetSlashableSharesForOperatorSetsBefore(ctx, requestBefore)
	if err != nil {
		return SlashableSharesForOperatorSetsResponse{}, utils.WrapError(
			"failed to get slashable shares for operator sets",
			err,
		)
	}

	return SlashableSharesForOperatorSetsResponse(resp), nil
}

// GetSlashableSharesForOperatorSetsBefore returns the strategies the operatorSets take into account, their
// operators, and the minimum amount of shares slashable by the
// operatorSets before a given timestamp.
// Timestamp must be in the future. Used to underestimate future slashable stake.
// Not supported for M2 AVSs
func (r *ChainReader) GetSlashableSharesForOperatorSetsBefore(
	ctx context.Context,
	request SlashableSharesForOperatorSetsBeforeRequest,
) (SlashableSharesForOperatorSetsBeforeResponse, error) {
	operatorSetStakes := make([]OperatorSetStakes, len(request.OperatorSets))
	for i, operatorSet := range request.OperatorSets {
		requestOperator := OperatorsForOperatorSetRequest{
			BlockNumber: request.BlockNumber,
			OperatorSet: operatorSet,
		}
		responseOperators, err := r.GetOperatorsForOperatorSet(ctx, requestOperator)
		if err != nil {
			return SlashableSharesForOperatorSetsBeforeResponse{}, utils.WrapError(
				"failed to get operators for operator set",
				err,
			)
		}

		requestStrategies := StrategiesForOperatorSetRequest{
			BlockNumber: request.BlockNumber,
			OperatorSet: operatorSet,
		}
		responseStrategies, err := r.GetStrategiesForOperatorSet(ctx, requestStrategies)
		// If operator setId is 0 will fail on if above
		if err != nil {
			return SlashableSharesForOperatorSetsBeforeResponse{}, utils.WrapError(
				"failed to get strategies for operator set",
				err,
			)
		}

		slashableShares, err := r.allocationManager.GetMinimumSlashableStake(
			&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
			allocationmanager.OperatorSet{
				Id:  operatorSet.Id,
				Avs: operatorSet.Avs,
			},
			responseOperators.Operators,
			responseStrategies.StrategiesAddresses,
			request.FutureBlock,
		)
		// This call should not fail since it's a getter
		if err != nil {
			return SlashableSharesForOperatorSetsBeforeResponse{}, utils.WrapError(
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

	return SlashableSharesForOperatorSetsBeforeResponse{OperatorSetStakes: operatorSetStakes}, nil
}

func (r *ChainReader) GetAllocationDelay(
	ctx context.Context,
	request AllocationDelayRequest,
) (AllocationDelayResponse, error) {
	if r.allocationManager == nil {
		return AllocationDelayResponse{}, errors.New("AllocationManager contract not provided")
	}
	isSet, delay, err := r.allocationManager.GetAllocationDelay(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
	)
	// This call should not fail since it's a getter
	if err != nil {
		return AllocationDelayResponse{}, utils.WrapError("failed to get allocation delay", err)
	}
	if !isSet {
		return AllocationDelayResponse{}, errors.New("allocation delay not set")
	}
	return AllocationDelayResponse{AllocationDelay: delay}, nil
}

func (r *ChainReader) GetRegisteredSets(
	ctx context.Context,
	request RegisteredSetsRequest,
) (RegisteredSetsResponse, error) {
	if r.allocationManager == nil {
		return RegisteredSetsResponse{}, errors.New("AllocationManager contract not provided")
	}
	reigsteredSets, err := r.allocationManager.GetRegisteredSets(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return RegisteredSetsResponse{}, utils.WrapError("failed to get registered sets", err)
	}

	return RegisteredSetsResponse{OperatorSets: reigsteredSets}, nil
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

func (r *ChainReader) ListAppointees(
	ctx context.Context,
	request ListAppointeesRequest,
) (ListAppointeesResponse, error) {
	if r.permissionController == nil {
		return ListAppointeesResponse{}, errors.New("PermissionController contract not provided")
	}

	appointees, err := r.permissionController.GetAppointees(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
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
	request ListAppointeePermissionsRequest,
) (ListAppointeePermissionsResponse, error) {
	if r.permissionController == nil {
		return ListAppointeePermissionsResponse{}, errors.New("PermissionController contract not provided")
	}

	targets, selectors, err := r.permissionController.GetAppointeePermissions(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
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
	request ListPendingAdminsRequest,
) (ListPendingAdminsResponse, error) {
	if r.permissionController == nil {
		return ListPendingAdminsResponse{}, errors.New("PermissionController contract not provided")
	}

	pendingAdmins, err := r.permissionController.GetPendingAdmins(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
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
	request ListAdminsRequest,
) (ListAdminsResponse, error) {
	if r.permissionController == nil {
		return ListAdminsResponse{}, errors.New("PermissionController contract not provided")
	}

	admins, err := r.permissionController.GetAdmins(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
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
	request IsPendingAdminRequest,
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

func (r *ChainReader) IsAdmin(
	ctx context.Context,
	request IsAdminRequest,
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
