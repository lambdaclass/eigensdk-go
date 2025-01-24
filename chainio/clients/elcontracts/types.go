package elcontracts

import (
	"math/big"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	"github.com/ethereum/go-ethereum/common"
)

type OperatorSetStakes struct {
	OperatorSet     allocationmanager.OperatorSet
	Strategies      []common.Address
	Operators       []common.Address
	SlashableStakes [][]*big.Int
}

type PendingDeallocation struct {
	MagnitudeDiff        uint64
	CompletableTimestamp uint32
}

type AllocationInfo struct {
	CurrentMagnitude *big.Int
	PendingDiff      *big.Int
	EffectBlock      uint32
	OperatorSetId    uint32
	AvsAddress       common.Address
}

type DeregistrationRequest struct {
	AVSAddress     common.Address
	OperatorSetIds []uint32
	WaitForReceipt bool
}

type RegistrationRequest struct {
	OperatorAddress common.Address
	AVSAddress      common.Address
	OperatorSetIds  []uint32
	WaitForReceipt  bool
	BlsKeyPair      *bls.KeyPair
	Socket          string
}
type RemovePermissionRequest struct {
	AccountAddress   common.Address
	AppointeeAddress common.Address
	Target           common.Address
	Selector         [4]byte
	WaitForReceipt   bool
}

type SetPermissionRequest struct {
	AccountAddress   common.Address
	AppointeeAddress common.Address
	Target           common.Address
	Selector         [4]byte
	WaitForReceipt   bool
}

type AcceptAdminRequest struct {
	AccountAddress common.Address
	WaitForReceipt bool
}

type AddPendingAdminRequest struct {
	AccountAddress common.Address
	AdminAddress   common.Address
	WaitForReceipt bool
}

type RemoveAdminRequest struct {
	AccountAddress common.Address
	AdminAddress   common.Address
	WaitForReceipt bool
}

type RemovePendingAdminRequest struct {
	AccountAddress common.Address
	AdminAddress   common.Address
	WaitForReceipt bool
}

// Reader structs
type IsOperatorRegisteredRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type IsOperatorRegisteredResponse struct {
	IsRegistered bool
}

type GetStakerSharesRequest struct {
	blockNumber   *big.Int
	StakerAddress common.Address
}

type GetStakerSharesResponse struct {
	StrategiesAddresses []common.Address
	Shares              []*big.Int
}

type GetDelegatedOperatorRequest struct {
	blockNumber   *big.Int
	StakerAddress common.Address
}

type GetDelegatedOperatorResponse struct {
	OperatorAddress common.Address
}

type GetOperatorDetailsRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type GetOperatorDetailsResponse struct {
	OperatorAddress           common.Address
	DelegationApproverAddress common.Address
	AllocationDelay           uint32
}

type GetStrategyAndUnderlyingTokenRequest struct {
	blockNumber     *big.Int
	StrategyAddress common.Address
}

type GetStrategyAndUnderlyingTokenResponse struct {
	StrategyContract       *strategy.ContractIStrategy
	UnderlyingTokenAddress common.Address
}

type GetStrategyAndUnderlyingERC20TokenRequest struct {
	blockNumber     *big.Int
	StrategyAddress common.Address
}

type GetStrategyAndUnderlyingERC20TokenResponse struct {
	StrategyContract       *strategy.ContractIStrategy
	ERC20Bindings          erc20.ContractIERC20Methods
	UnderlyingTokenAddress common.Address
}

type GetOperatorSharesInStrategyRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
	StrategyAddress common.Address
}

type GetOperatorSharesInStrategyResponse struct {
	Shares *big.Int
}

type CalculateDelegationApprovalDigestHashRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
	StakerAddress   common.Address
	ApproverAddress common.Address
	ApproverSalt    [32]byte
	Expiry          *big.Int
}

type CalculateDelegationApprovalDigestHashResponse struct {
	DigestHash [32]byte
}

type CalculateOperatorAVSRegistrationDigestHashRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
	AVSAddress      common.Address
	Salt            [32]byte
	Expiry          *big.Int
}

type CalculateOperatorAVSRegistrationDigestHashResponse struct {
	DigestHash [32]byte
}

type GetDistributionRootsLengthRequest struct {
	blockNumber *big.Int
}

type GetDistributionRootsLengthResponse struct {
	Length *big.Int
}

type CurrRewardsCalculationEndTimestampRequest struct {
	blockNumber *big.Int
}

type CurrRewardsCalculationEndTimestampResponse struct {
	Timestamp uint32
}

type GetCurrentClaimableDistributionRootRequest struct {
	blockNumber *big.Int
}

type GetCurrentClaimableDistributionRootResponse struct {
	DistributionRoot rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot
}

type GetRootIndexFromHashRequest struct {
	blockNumber *big.Int
	RootHash    [32]byte
}

type GetRootIndexFromHashResponse struct {
	RootIndex uint32
}

type GetCumulativeClaimedRequest struct {
	blockNumber    *big.Int
	ClaimerAddress common.Address
	TokenAddress   common.Address
}

type GetCumulativeClaimedResponse struct {
	CumulativeClaimed *big.Int
}

type CheckClaimRequest struct {
	blockNumber *big.Int
	Claim       rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim
}

type CheckClaimResponse struct {
	IsValid bool
}

type GetOperatorAVSSplitRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
	AvsAddress      common.Address
}

type GetOperatorAVSSplitResponse struct {
	Split uint16
}

type GetOperatorPISplitRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type GetOperatorPISplitResponse struct {
	Split uint16
}

type GetMaxMagnitudes0Request struct {
	blockNumber         *big.Int
	OperatorAddress     common.Address
	StrategiesAddresses []common.Address
}

type GetMaxMagnitudes0Response struct {
	MaxMagnitudes []uint64
}

type GetAllocationInfoRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
	StrategyAddress common.Address
}

type GetAllocationInfoResponse struct {
	AllocationInfo []AllocationInfo
}

type GetOperatorSharesRequest struct {
	blockNumber         *big.Int
	OperatorAddress     common.Address
	StrategiesAddresses []common.Address
}

type GetOperatorSharesResponse struct {
	Shares []*big.Int
}

type GetOperatorsSharesRequest struct {
	blockNumber         *big.Int
	OperatorsAddresses  []common.Address
	StrategiesAddresses []common.Address
}

type GetOperatorsSharesResponse struct {
	Shares [][]*big.Int
}

type GetNumOperatorSetsForOperatorRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type GetNumOperatorSetsForOperatorResponse struct {
	NumOperatorSets *big.Int
}

type GetOperatorSetsForOperatorRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type GetOperatorSetsForOperatorResponse struct {
	OperatorSets []allocationmanager.OperatorSet
}

type IsOperatorRegisteredWithOperatorSetRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
	OperatorSet     allocationmanager.OperatorSet
}

type IsOperatorRegisteredWithOperatorSetResponse struct {
	IsRegistered bool
}

type GetOperatorsForOperatorSetRequest struct {
	blockNumber *big.Int
	OperatorSet allocationmanager.OperatorSet
}

type GetOperatorsForOperatorSetResponse struct {
	Operators []common.Address
}

type GetNumOperatorsForOperatorSetRequest struct {
	blockNumber *big.Int
	OperatorSet allocationmanager.OperatorSet
}

type GetNumOperatorsForOperatorSetResponse struct {
	NumOperators *big.Int
}

type GetStrategiesForOperatorSetRequest struct {
	blockNumber *big.Int
	OperatorSet allocationmanager.OperatorSet
}

type GetStrategiesForOperatorSetResponse struct {
	StrategiesAddresses []common.Address
}

type GetSlashableSharesRequest struct {
	blockNumber         *big.Int
	OperatorAddress     common.Address
	OperatorSet         allocationmanager.OperatorSet
	StrategiesAddresses []common.Address
}

type GetSlashableSharesResponse struct {
	SlashableShares map[common.Address]*big.Int
}

type GetAllocatableMagnitudeRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
	StrategyAddress common.Address
}

type GetAllocatableMagnitudeResponse struct {
	AllocatableMagnitude uint64
}

type GetSlashableSharesForOperatorSetsRequest struct {
	blockNumber  *big.Int
	OperatorSets []allocationmanager.OperatorSet
}

// Original struct was OperatorSetStakes
type GetSlashableSharesForOperatorSetsResponse struct {
	OperatorSetStakes []OperatorSetStakes
}

type GetSlashableSharesForOperatorSetsBeforeRequest struct {
	blockNumber  *big.Int
	OperatorSets []allocationmanager.OperatorSet
	FutureBlock  uint32
}

type GetSlashableSharesForOperatorSetsBeforeResponse struct {
	OperatorSetStakes []OperatorSetStakes
}

type GetAllocationDelayRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type GetAllocationDelayResponse struct {
	AllocationDelay uint32
}

type GetRegisteredSetsRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type GetRegisteredSetsResponse struct {
	OperatorSets []allocationmanager.OperatorSet
}

type CanCallRequest struct {
	blockNumber      *big.Int
	AccountAddress   common.Address
	AppointeeAddress common.Address
	Target           common.Address
	Selector         [4]byte
}

type CanCallResponse struct {
	CanCall bool
}

type ListAppointeesRequest struct {
	blockNumber    *big.Int
	AccountAddress common.Address
	Target         common.Address
	Select         [4]byte
}

type ListAppointeesResponse struct {
	Appointees []common.Address
}

type ListAppointeePermissionsRequest struct {
	blockNumber      *big.Int
	AccountAddress   common.Address
	AppointeeAddress common.Address
}

type ListAppointeePermissionsResponse struct {
	AppointeeAddress []common.Address
	Selector         [][4]byte
}

type ListPendingAdminsRequest struct {
	blockNumber    *big.Int
	AccountAddress common.Address
}

type ListPendingAdminsResponse struct {
	PendingAdmins []common.Address
}

type ListAdminsRequest struct {
	blockNumber    *big.Int
	AccountAddress common.Address
}

type ListAdminsResponse struct {
	Admins []common.Address
}

type IsPendingAdminRequest struct {
	blockNumber         *big.Int
	AccountAddress      common.Address
	PendingAdminAddress common.Address
}

type IsPendingAdminResponse struct {
	IsPendingAdmin bool
}

type IsAdminRequest struct {
	blockNumber    *big.Int
	AccountAddress common.Address
	AdminAddress   common.Address
}

type IsAdminResponse struct {
	IsAdmin bool
}
