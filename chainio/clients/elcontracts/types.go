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

type OperatorAVSSplitRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
	AVSAddress      common.Address
}

type SplitResponse struct {
	Split uint16
}

type OperatorRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
}

type OperatorStrategyRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
	StrategyAddress common.Address
}

type AllocatableResponse struct {
	Allocatable uint64
}

type OperatorStrategiesRequest struct {
	BlockNumber       *big.Int
	OperatorAddress   common.Address
	StrategyAddresses []common.Address
}

type MaxMagnitudesResponse struct {
	MaxMagnitudes []uint64
}

type AllocationResponse struct {
	AllocationInfo []AllocationInfo
}

type OperatorSharesResponse struct {
	Shares []*big.Int
}

type OperatorsStrategiesRequest struct {
	BlockNumber       *big.Int
	OperatorAddresses []common.Address
	StrategyAddresses []common.Address
}

type OperatorsSharesResponse struct {
	Shares [][]*big.Int
}
