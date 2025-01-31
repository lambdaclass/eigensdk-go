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

type ApprovalDigestHashRequest struct {
	BlockNumber       *big.Int
	StakerAddress     common.Address
	OperatorAddress   common.Address
	DelegationAddress common.Address
	ApproverSalt      [32]byte
	Expiry            *big.Int
}

type AVSRegistrationDigestHashRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
	AVSAddress      common.Address
	Salt            [32]byte
	Expiry          *big.Int
}

type DigestHashResponse struct {
	DigestHash [32]byte
}

// This should be an specific struct like BlockNumberRequest?
type RootRequest struct {
	BlockNumer *big.Int
}

type RootLengthResponse struct {
	Length *big.Int
}

// This should be an specific struct like BlockNumberRequest?
type RewardsEndTimestampRequest struct {
	BlockNumber *big.Int
}

type EndTimestampResponse struct {
	EndTimestamp uint32
}

type ClaimableDistributionRootResponse struct {
	DistributionRoot rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot
}

type RootHashRequest struct {
	BlockNumber *big.Int
	RootHash    [32]byte
}

type RootIndexResponse struct {
	Index uint32
}

type CumulativeClaimedRequest struct {
	BlockNumber   *big.Int
	EarnerAddress common.Address
	TokenAddress  common.Address
}

type CumulativeClaimedResponse struct {
	CumulativeClaimed *big.Int
}

type ClaimRequest struct {
	BlockNumber *big.Int
	Claim       rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim
}

type ClaimResponse struct {
	CheckClaim bool
}
