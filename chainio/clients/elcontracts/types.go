package elcontracts

import (
	"math/big"

	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
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

// Reader structs
type IsOperatorRegisteredRequest struct {
	OperatorAddress common.Address
}

type IsOperatorRegisteredResponse struct {
	IsRegistered bool
}

type GetStakerSharesRequest struct {
	StakerAddress common.Address
}

type GetStakerSharesResponse struct {
	StrategiesAddresses []common.Address
	Shares              []*big.Int
}

type GetDelegatedOperatorRequest struct {
	StakerAddress common.Address
}

type GetDelegatedOperatorResponse struct {
	OperatorAddress common.Address
}

type GetOperatorDetailsRequest struct {
	OperatorAddress common.Address
}

type GetOperatorDetailsResponse struct {
	OperatorAddress           common.Address
	DelegationApproverAddress common.Address
	AllocationDelay           uint32
}

type GetStrategyAndUnderlyingTokenRequest struct {
	StrategyAddress common.Address
}

type GetStrategyAndUnderlyingTokenResponse struct {
	StrategyContract       *strategy.ContractIStrategy
	UnderlyingTokenAddress common.Address
}

type GetStrategyAndUnderlyingERC20TokenRequest struct {
	StrategyAddress common.Address
}

type GetStrategyAndUnderlyingERC20TokenResponse struct {
	StrategyContract       *strategy.ContractIStrategy
	ERC20Bindings          erc20.ContractIERC20Methods
	UnderlyingTokenAddress common.Address
}

type GetOperatorSharesInStrategyRequest struct {
	OperatorAddress common.Address
	StrategyAddress common.Address
}

type GetOperatorSharesInStrategyResponse struct {
	Shares *big.Int
}

type CalculateDelegationApprovalDigestHashRequest struct {
	OperatorAddress common.Address
	StakerAddress   common.Address
	ApproverAddress common.Address
	ApproverSalt    [32]byte
	Expiry          *big.Int
}

type CalculateDelegationApprovalDigestHashResponse struct {
	DigestHash [32]byte
}

type CalculateOperatorAVSRegistrationDigestHashRequest struct {
	OperatorAddress common.Address
	AVSAddress      common.Address
	Salt            [32]byte
	Expiry          *big.Int
}

type CalculateOperatorAVSRegistrationDigestHashResponse struct {
	DigestHash [32]byte
}

type GetDistributionRootsLengthResponse struct {
	Length *big.Int
}

type CurrRewardsCalculationEndTimestampResponse struct {
	Timestamp uint32
}

type GetCurrentClaimableDistributionRootResponse struct {
	DistributionRoot rewardscoordinator.IRewardsCoordinatorTypesDistributionRoot
}
