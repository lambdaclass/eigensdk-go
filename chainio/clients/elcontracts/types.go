package elcontracts

import (
	"math/big"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

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

type RegisterForOperatorSetRequest struct {
	OperatorAddress            common.Address
	AVSAddress                 common.Address
	RegistryCoordinatorAddress common.Address
	OperatorSetIds             []uint32
	WaitForReceipt             bool
	BlsKeyPair                 *bls.KeyPair
	Socket                     string
}

type PermissionRequest struct {
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

type PendingAdminRequest struct {
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

// TxOption represents a Ethereum transaction option.
type TxOption struct {
	Opts *bind.TransactOpts
}

// DeregisterFromOperatorSet is a request to deregister an operator from an operator set
type DeregisterFromOperatorSet struct {
	OperatorAddress common.Address
	AVSAddress      common.Address
	OperatorSetIds  []uint32
	WaitForReceipt  bool
}

// ModifyAllocationRequest is a request to modify allocated stake
type ModifyAllocationRequest struct {
	OperatorAddress common.Address
	Allocations     []allocationmanager.IAllocationManagerTypesAllocateParams
	WaitForReceipt  bool
}

// AllocationDelayRequest is a request to delay allocation
type AllocationDelayRequest struct {
	OperatorAddress common.Address
	Delay           uint32
	WaitForReceipt  bool
}
