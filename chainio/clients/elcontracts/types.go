package elcontracts

import (
	"math/big"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
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

// IsAdminRequest represents a request to check if the caller is an admin of an account
type AdminCheckRequest struct {
	BlockNumber    *big.Int
	AccountAddress common.Address
	AdminAddress   common.Address
}

// IsAdminResponse shows whether an address is an admin.
type IsAdminResponse struct {
	IsAdmin bool
}

// IsPendingAdminRequest represents a request to check if an address is a pending admin of the account
type PendingAdminCheckRequest struct {
	BlockNumber         *big.Int
	AccountAddress      common.Address
	PendingAdminAddress common.Address
}

// IsPendingAdminResponse shows whether an address is a pending admin.
type IsPendingAdminResponse struct {
	IsPendingAdmin bool
}

// AccountRequest represents the address of an specific account
type AccountRequest struct {
	BlockNumber    *big.Int
	AccountAddress common.Address
}

// AdminsResponse shows the list of admins of an account
type AdminsResponse struct {
	Admins []common.Address
}

// PendingAdminsResponse shows the list of pending admins of an account
type PendingAdminsResponse struct {
	PendingAdmins []common.Address
}

// AppointeesListRequest represents a request to get the list of appointees of an account
type AppointeesListRequest struct {
	BlockNumber    *big.Int
	AccountAddress common.Address
	TargetAddress  common.Address
	Selector       [4]byte
}

// AppointeesResponse shows the list of appointees of an account
type AppointeesResponse struct {
	Appointees []common.Address
}

type AppointeePermissionsListRequest struct {
	BlockNumber      *big.Int
	AccountAddress   common.Address
	AppointeeAddress common.Address
}

type AppointeePermissionsListResponse struct {
	TargetAddresses []common.Address
	Selectors       [][4]byte
}

// CanCallRequest represents a request to check if an account can call a target
type CanCallRequest struct {
	BlockNumber      *big.Int
	AccountAddress   common.Address
	AppointeeAddress common.Address
	Target           common.Address
	Selector         [4]byte
}

// CanCallResponse shows whether an account can call a target
type CanCallResponse struct {
	CanCall bool
}
