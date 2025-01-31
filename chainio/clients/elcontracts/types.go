package elcontracts

import (
	"math/big"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
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

// ApprovalDigestHashRequest represents the request parameters required to calculate the delegation approval digest
// hash. If `BlockNumber` is nil, the latest block will be used.
type ApprovalDigestHashRequest struct {
	BlockNumber       *big.Int
	StakerAddress     common.Address
	OperatorAddress   common.Address
	DelegationAddress common.Address
	ApproverSalt      [32]byte
	Expiry            *big.Int
}

// AVSRegistrationDigestHashRequest represents the request required to calculate the operator AVS registration digest
// hash. If `BlockNumber` is nil, the latest block will be used
type AVSRegistrationDigestHashRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
	AVSAddress      common.Address
	Salt            [32]byte
	Expiry          *big.Int
}

// DigestHashResponse contains the calculated digest hash.
type DigestHashResponse struct {
	DigestHash [32]byte
}

// This should be an specific struct like BlockNumberRequest?
// If `BlockNumber` is nil, the latest block will be used
type RootRequest struct {
	BlockNumer *big.Int
}

// RootLengthResponse contains the length of the distribution roots.
type RootLengthResponse struct {
	Length *big.Int
}

// This should be an specific struct like BlockNumberRequest?
// If `BlockNumber` is nil, the latest block will be used
type RewardsEndTimestampRequest struct {
	BlockNumber *big.Int
}

// EndTimestampResponse contains the rewards calculation end timestamp
type EndTimestampResponse struct {
	EndTimestamp uint32
}

// ClaimableDistributionRootResponse contains the current claimable distribution root
type ClaimableDistributionRootResponse struct {
	DistributionRoot rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot
}

// RootHashRequest represents a request to retrieve the index of a root hash. If `BlockNumber` is nil, the latest block
// will be used
type RootHashRequest struct {
	BlockNumber *big.Int
	RootHash    [32]byte
}

// RootIndexResponse contains the index corresponding to a given root hash
type RootIndexResponse struct {
	Index uint32
}

// CumulativeClaimedRequest represents a request to fetch the cumulative claimed rewards
// for a specific earner address and token address. If `BlockNumber` is nil, the latest block will be used
type CumulativeClaimedRequest struct {
	BlockNumber   *big.Int
	EarnerAddress common.Address
	TokenAddress  common.Address
}

// CumulativeClaimedResponse contains the cumulative claimed amount
type CumulativeClaimedResponse struct {
	CumulativeClaimed *big.Int
}

// ClaimRequest represents a request to verify a claim
// If `BlockNumber` is nil, the latest block will be used
type ClaimRequest struct {
	BlockNumber *big.Int
	Claim       rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim
}

// ClaimResponse contains the verification result of a claim
type ClaimResponse struct {
	CheckClaim bool
}
