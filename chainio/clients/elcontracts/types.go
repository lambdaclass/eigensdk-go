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

// OperatorAVSSplitRequest is used to request the split of an operator for a specific AVS.
// If `BlockNumber` is nil, the latest block will be used
type OperatorAVSSplitRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
	AVSAddress      common.Address
}

// SplitResponse represents the split of an operator
type SplitResponse struct {
	Split uint16
}

// OperatorRequest is used to represent the address of an operator.
// If `BlockNumber` is nil, the latest block will be used
type OperatorRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
}

// OperatorStrategyRequest is used to represent the address of an operator and a strategy.
// If `BlockNumber` is nil, the latest block will be used
type OperatorStrategyRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
	StrategyAddress common.Address
}

// AllocatableResponse shows the amount of allocatable magnitude for an operator
type AllocatableResponse struct {
	Allocatable uint64
}

// OperatorStrategiesRequest is used to represent the addresses of an operator and strategies.
// If `BlockNumber` is nil, the latest block will be used
type OperatorStrategiesRequest struct {
	BlockNumber       *big.Int
	OperatorAddress   common.Address
	StrategyAddresses []common.Address
}

// MaxMagnitudesResponse is used to represent the max magnitudes for an operator for each strategy
type MaxMagnitudesResponse struct {
	MaxMagnitudes []uint64
}

// AllocationResponse is used to represent the allocation info for specific operator and strategy
type AllocationResponse struct {
	AllocationInfo []AllocationInfo
}

// OperatorSharesResponse is used to represent the shares of an operator for each strategy
type OperatorSharesResponse struct {
	Shares []*big.Int
}

// OperatorsStrategiesRequest represents the addresses of operators and strategies.
// If `BlockNumber` is nil, the latest block will be used
type OperatorsStrategiesRequest struct {
	BlockNumber       *big.Int
	OperatorAddresses []common.Address
	StrategyAddresses []common.Address
}

// OperatorsSharesResponse shows the shares of operators for each strategy
type OperatorsSharesResponse struct {
	Shares [][]*big.Int
}
