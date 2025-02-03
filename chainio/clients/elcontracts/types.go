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

// TxOption represents a Ethereum transaction option.
type TxOption struct {
	Opts *bind.TransactOpts
}

// OperatorRequest represents a request that requires an operator's address.
type OperatorRequest struct {
	BlockNumber    *big.Int
	Operator       types.Operator
	WaitForReceipt bool
}

// OperatorMetadataRequest represents a request that updates operator metadata.
type OperatorMetadataRequest struct {
	OperatorAddress common.Address
	Uri             string
	WaitForReceipt  bool
}

// DepositRequest represents a request to deposit funds into a strategy.
type DepositRequest struct {
	StrategyAddress common.Address
	Amount          *big.Int
	WaitForReceipt  bool
}

// ClaimerRequest represents a request to set a claimer
type ClaimerRequest struct {
	BlockNumber    *big.Int
	ClaimerAddress common.Address
	WaitForReceipt bool
}

// ClaimProcessRequest represents a request to process a claim for rewards.
type ClaimProcessRequest struct {
	BlockNumber      *big.Int
	Claim            rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim
	RecipientAddress common.Address
	WaitForReceipt   bool
}

type ClaimsProcessRequest struct {
	BlockNumber      *big.Int
	Claims           []rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim
	RecipientAddress common.Address
	WaitForReceipt   bool
}

// OperatorAVSSplitRequest represents a request to set an operator's AVS split.
type OperatorAVSSplitRequest struct {
	OperatorAddress common.Address
	AVSAddress      common.Address
	Split           uint16
	WaitForReceipt  bool
}

// OperatorPISplitRequest represents a request to set an operator's Programmatic Incentive split.
type OperatorPISplitRequest struct {
	OperatorAddress common.Address
	Split           uint16
	WaitForReceipt  bool
}
