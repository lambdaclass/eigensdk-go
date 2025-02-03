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

// OperatorRequest represents a request that requires an operator's address
// If `BlockNumber` is nil, the latest block will be used
type OperatorRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
}

// OperatorSetCountResponse contains the number of operator sets
type OperatorSetCountResponse struct {
	Count *big.Int // Could this be a uint?
}

// OperatorSetsResponse contains a list of operator sets
type OperatorSetsResponse struct {
	OperatorSets []allocationmanager.OperatorSet
}

// RegisteredOperatorSetRequest represents a request to check if an operator is registered with a specific operator set
// If `BlockNumber` is nil, the latest block will be used
type RegisteredOperatorSetRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
	OperatorSet     allocationmanager.OperatorSet
}

// OperatorRegistrationResponse represents whether an operator is registered with a given set
type OperatorRegistrationResponse struct {
	IsRegistered bool
}

// OperatorsResponse contains a list of operator addresses
type OperatorsResponse struct {
	OperatorAddresses []common.Address
}

// OperatorSetRequest represents a request that requires an operator set
// If `BlockNumber` is nil, the latest block will be used
type OperatorSetRequest struct {
	BlockNumber *big.Int
	OperatorSet allocationmanager.OperatorSet
}

// OperatorSetsRequest represents a request that requires a list of operator sets
// If `BlockNumber` is nil, the latest block will be used
type OperatorSetsRequest struct {
	BlockNumber  *big.Int
	OperatorSets []allocationmanager.OperatorSet
}

// OperatorCountResponse contains the number of operators for a given operator sets
type OperatorCountResponse struct {
	Count *big.Int
}

// StrategiesResponse contains a list of strategy addresses
type StrategiesResponse struct {
	StrategyAddresses []common.Address
}

type AllocationDelayResponse struct {
	Delay uint32
}

// OperatorsStrategiesRequest represents a request that requires an operator and multiple strategies.
type OperatorsStrategiesRequest struct {
	BlockNumber       *big.Int
	OperatorAddress   common.Address
	StrategyAddresses []common.Address
	OperatorSet       allocationmanager.OperatorSet
}

type StrategySlashableSharesResponse struct {
	StrategyShares map[common.Address]*big.Int
}

// OperatorSetsBeforeRequest extends OperatorSetsRequest with a future block number.
type OperatorSetsBeforeRequest struct {
	FutureBlock  uint32
	OperatorSets []allocationmanager.OperatorSet
}

// OperatorSetStakesResponse represents the slashable stakes for a list of operator sets.
type OperatorSetStakesResponse struct {
	OperatorSetStakes []OperatorSetStakes
}
