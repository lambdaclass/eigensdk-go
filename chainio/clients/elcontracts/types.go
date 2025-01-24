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

type StakerSharesRequest struct {
	blockNumber   *big.Int
	StakerAddress common.Address
}

type StakerSharesResponse struct {
	StrategiesAddresses []common.Address
	Shares              []*big.Int
}

type DelegatedOperatorRequest struct {
	blockNumber   *big.Int
	StakerAddress common.Address
}

type DelegatedOperatorResponse struct {
	OperatorAddress common.Address
}

type OperatorDetailsRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type OperatorDetailsResponse struct {
	OperatorAddress           common.Address
	DelegationApproverAddress common.Address
	AllocationDelay           uint32
}

type StrategyAndUnderlyingTokenRequest struct {
	blockNumber     *big.Int
	StrategyAddress common.Address
}

type StrategyAndUnderlyingTokenResponse struct {
	StrategyContract       *strategy.ContractIStrategy
	UnderlyingTokenAddress common.Address
}

type StrategyAndUnderlyingERC20TokenRequest struct {
	blockNumber     *big.Int
	StrategyAddress common.Address
}

type StrategyAndUnderlyingERC20TokenResponse struct {
	StrategyContract       *strategy.ContractIStrategy
	ERC20Bindings          erc20.ContractIERC20Methods
	UnderlyingTokenAddress common.Address
}

type OperatorSharesInStrategyRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
	StrategyAddress common.Address
}

type OperatorSharesInStrategyResponse struct {
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

type DistributionRootsLengthRequest struct {
	blockNumber *big.Int
}

type DistributionRootsLengthResponse struct {
	Length *big.Int
}

type CurrRewardsCalculationEndTimestampRequest struct {
	blockNumber *big.Int
}

type CurrRewardsCalculationEndTimestampResponse struct {
	Timestamp uint32
}

type CurrentClaimableDistributionRootRequest struct {
	blockNumber *big.Int
}

type CurrentClaimableDistributionRootResponse struct {
	DistributionRoot rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot
}

type RootIndexFromHashRequest struct {
	blockNumber *big.Int
	RootHash    [32]byte
}

type RootIndexFromHashResponse struct {
	RootIndex uint32
}

type CumulativeClaimedRequest struct {
	blockNumber    *big.Int
	ClaimerAddress common.Address
	TokenAddress   common.Address
}

type CumulativeClaimedResponse struct {
	CumulativeClaimed *big.Int
}

type CheckClaimRequest struct {
	blockNumber *big.Int
	Claim       rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim
}

type CheckClaimResponse struct {
	IsValid bool
}

type OperatorAVSSplitRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
	AvsAddress      common.Address
}

type OperatorAVSSplitResponse struct {
	Split uint16
}

type OperatorPISplitRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type OperatorPISplitResponse struct {
	Split uint16
}

type MaxMagnitudes0Request struct {
	blockNumber         *big.Int
	OperatorAddress     common.Address
	StrategiesAddresses []common.Address
}

type MaxMagnitudes0Response struct {
	MaxMagnitudes []uint64
}

type AllocationInfoRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
	StrategyAddress common.Address
}

type AllocationInfoResponse struct {
	AllocationInfo []AllocationInfo
}

type OperatorSharesRequest struct {
	blockNumber         *big.Int
	OperatorAddress     common.Address
	StrategiesAddresses []common.Address
}

type OperatorSharesResponse struct {
	Shares []*big.Int
}

type OperatorsSharesRequest struct {
	blockNumber         *big.Int
	OperatorsAddresses  []common.Address
	StrategiesAddresses []common.Address
}

type OperatorsSharesResponse struct {
	Shares [][]*big.Int
}

type NumOperatorSetsForOperatorRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type NumOperatorSetsForOperatorResponse struct {
	NumOperatorSets *big.Int
}

type OperatorSetsForOperatorRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type OperatorSetsForOperatorResponse struct {
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

type OperatorsForOperatorSetRequest struct {
	blockNumber *big.Int
	OperatorSet allocationmanager.OperatorSet
}

type OperatorsForOperatorSetResponse struct {
	Operators []common.Address
}

type NumOperatorsForOperatorSetRequest struct {
	blockNumber *big.Int
	OperatorSet allocationmanager.OperatorSet
}

type NumOperatorsForOperatorSetResponse struct {
	NumOperators *big.Int
}

type StrategiesForOperatorSetRequest struct {
	blockNumber *big.Int
	OperatorSet allocationmanager.OperatorSet
}

type StrategiesForOperatorSetResponse struct {
	StrategiesAddresses []common.Address
}

type SlashableSharesRequest struct {
	blockNumber         *big.Int
	OperatorAddress     common.Address
	OperatorSet         allocationmanager.OperatorSet
	StrategiesAddresses []common.Address
}

type SlashableSharesResponse struct {
	SlashableShares map[common.Address]*big.Int
}

type AllocatableMagnitudeRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
	StrategyAddress common.Address
}

type AllocatableMagnitudeResponse struct {
	AllocatableMagnitude uint64
}

type SlashableSharesForOperatorSetsRequest struct {
	blockNumber  *big.Int
	OperatorSets []allocationmanager.OperatorSet
}

type SlashableSharesForOperatorSetsResponse struct {
	OperatorSetStakes []OperatorSetStakes
}

type SlashableSharesForOperatorSetsBeforeRequest struct {
	blockNumber  *big.Int
	OperatorSets []allocationmanager.OperatorSet
	FutureBlock  uint32
}

type SlashableSharesForOperatorSetsBeforeResponse struct {
	OperatorSetStakes []OperatorSetStakes
}

type AllocationDelayRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type AllocationDelayResponse struct {
	AllocationDelay uint32
}

type RegisteredSetsRequest struct {
	blockNumber     *big.Int
	OperatorAddress common.Address
}

type RegisteredSetsResponse struct {
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
