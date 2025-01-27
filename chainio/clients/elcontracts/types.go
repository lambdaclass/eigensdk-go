package elcontracts

import (
	"math/big"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"

	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

type TxOptions struct {
	Options *bind.TransactOpts
}

type RegisterOperatorRequest struct {
	Operator       types.Operator
	WaitForReceipt bool
}

type RegisterOperatorSetsRequest struct {
	RegistryCoordinatorAddress common.Address
	OperatorAddress            common.Address
	AVSAddress                 common.Address
	OperatorSetIds             []uint32
	BlsKeyPair                 *bls.KeyPair
	Socket                     string
	WaitForReceipt             bool
}

type OperatorDetailsRequest struct {
	Operator       types.Operator
	WaitForReceipt bool
}

type MetadataURIRequest struct {
	OperatorAddress common.Address
	Uri             string
	WaitForReceipt  bool
}

type ERC20IntoStrategyRequest struct {
	StrategyAddress common.Address
	Amount          *big.Int
	WaitForReceipt  bool
}

type ClaimForRequest struct {
	Claimer        common.Address
	WaitForReceipt bool
}

type ClaimProcessRequest struct {
	Claim            rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim
	RecipientAddress common.Address
	WaitForReceipt   bool
}

type ClaimsProcessRequest struct {
	Claims           []rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim
	RecipientAddress common.Address
	WaitForReceipt   bool
}

type OperatorAVSSplitRequest struct {
	OperatorAddress common.Address
	AVSAddress      common.Address
	Split           uint16
	WaitForReceipt  bool
}

type OperatorPISplitRequest struct {
	OperatorAddress common.Address
	Split           uint16
	WaitForReceipt  bool
}

type OperatorSetDeregisterRequest struct {
	OperatorAddress common.Address
	AVSAddress      common.Address
	OperatorSetIds  []uint32
	WaitForReceipt  bool
}

type AllocationModifyRequest struct {
	OperatorAddress common.Address
	Allocations     []allocationmanager.IAllocationManagerTypesAllocateParams
	WaitForReceipt  bool
}

type AllocationDelayRequest struct {
	OperatorAddress common.Address
	Delay           uint32
	WaitForReceipt  bool
}

type PermissionRemoveRequest struct {
	AccountAddress   common.Address
	AppointeeAddress common.Address
	Target           common.Address
	Selector         [4]byte
	WaitForReceipt   bool
}

type PermissionSetRequest struct {
	AccountAddress   common.Address
	AppointeeAddress common.Address
	Target           common.Address
	Selector         [4]byte
	WaitForReceipt   bool
}

type AdminAcceptRequest struct {
	AccountAddress common.Address
	WaitForReceipt bool
}

type PendingAdminAcceptRequest struct {
	AccountAddress common.Address
	AdminAddress   common.Address
	WaitForReceipt bool
}

type AdminRemoveRequest struct {
	AccountAddress common.Address
	AdminAddress   common.Address
	WaitForReceipt bool
}

type PendingAdminRemoveRequest struct {
	AccountAddress common.Address
	AdminAddress   common.Address
	WaitForReceipt bool
}
