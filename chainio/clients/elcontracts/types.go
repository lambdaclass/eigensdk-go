package elcontracts

import (
	"math/big"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"

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

// READER TYPES

type OperatorRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
}

type OperatorResponse struct {
	Operator types.Operator
}

type StrategyRequest struct {
	BlockNumber     *big.Int
	StrategyAddress common.Address
}

type StrategyTokenResponse struct {
	StrategyContract strategy.ContractIStrategy
	TokenAddress     common.Address
}

type StrategyERC20TokenResponse struct {
	StrategyContract        strategy.ContractIStrategy
	UnderlyingERC20Contract erc20.ContractIERC20Methods
	UnderlyingTokenAddress  common.Address
}

type OperatorRegisterResponse struct {
	IsRegistered bool
}

type StakerRequest struct {
	BlockNumber   *big.Int
	StakerAddress common.Address
}

type StakerSharesResponse struct {
	StrategiesAddresses []common.Address
	Shares              []*big.Int
}

type DelegateOperatorResponse struct {
	OperatorAddress common.Address
}

type SharesInStrategyRequest struct {
	BlockNumber     *big.Int
	OperatorAddress common.Address
	StrategyAddress common.Address
}

type SharesResponse struct {
	Shares *big.Int
}
