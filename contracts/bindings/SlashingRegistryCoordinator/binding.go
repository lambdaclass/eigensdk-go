// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSlashingRegistryCoordinator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// ISlashingRegistryCoordinatorTypesOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesOperatorInfo struct {
	OperatorId [32]byte
	Status     uint8
}

// ISlashingRegistryCoordinatorTypesOperatorKickParam is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesOperatorKickParam struct {
	QuorumNumber uint8
	Operator     common.Address
}

// ISlashingRegistryCoordinatorTypesOperatorSetParam is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesOperatorSetParam struct {
	MaxOperatorCount        uint32
	KickBIPsOfOperatorStake uint16
	KickBIPsOfTotalStake    uint16
}

// ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate is an auto generated low-level Go binding around an user-defined struct.
type ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	QuorumBitmap          *big.Int
}

// IStakeRegistryTypesStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type IStakeRegistryTypesStrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ContractSlashingRegistryCoordinatorMetaData contains all meta data concerning the ContractSlashingRegistryCoordinator contract.
var ContractSlashingRegistryCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"},{\"name\":\"_indexRegistry\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"},{\"name\":\"_socketRegistry\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_CHURN_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorChurnApprovalDigestHash\",\"inputs\":[{\"name\":\"registeringOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registeringOperatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"churnApprover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createSlashableStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"lookAheadPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createTotalDelegatedStakeQuorum\",\"inputs\":[{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"minimumStake\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"strategyParams\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRegistryTypes.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectionCooldown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ejector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentQuorumBitmap\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorInfo\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISlashingRegistryCoordinatorTypes.OperatorStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromId\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumISlashingRegistryCoordinatorTypes.OperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapUpdateByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.QuorumBitmapUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBitmap\",\"type\":\"uint192\",\"internalType\":\"uint192\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_accountIdentifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChurnApproverSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isM2Quorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastEjectionTimestamp\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"m2QuorumsDisabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numRegistries\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetsEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumUpdateBlockNumber\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registries\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAccountIdentifier\",\"inputs\":[{\"name\":\"_accountIdentifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChurnApprover\",\"inputs\":[{\"name\":\"_churnApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjectionCooldown\",\"inputs\":[{\"name\":\"_ejectionCooldown\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEjector\",\"inputs\":[{\"name\":\"_ejector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"socketRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISocketRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChurnApproverUpdated\",\"inputs\":[{\"name\":\"prevChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectorUpdated\",\"inputs\":[{\"name\":\"prevEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetParamsUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlashingRegistryCoordinatorTypes.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumBlockNumberUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegisteredForQuorums\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapUpdateIsAfterBlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BitmapValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayLengthTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BytesArrayNotOrdered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotChurnSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotKickOperatorAboveThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotReregisterYet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChurnApproverSaltUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientStakeForChurn\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRegistrationType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxQuorumsReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NextBitmapUpdateIsBeforeBlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegisteredForQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEjector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumOperatorCountMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]}]",
	Bin: "0x610200604052348015610010575f5ffd5b506040516155d13803806155d183398101604081905261002f9161027e565b604080518082018252601681527f4156535265676973747279436f6f7264696e61746f720000000000000000000060208083019182528351808501909452600684526576302e302e3160d01b908401528151902060e08190527f6bda7e3f385e48841048390444cced5cc795af87758af67622e5f4f0882c4a996101008190524660a0528993899389938993899389939290917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6101318184846040805160208101859052908101839052606081018290524660808201523060a08201525f9060c0016040516020818303038152906040528051906020012090509392505050565b6080523060c052610120525050506001600160a01b0382169050610168576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b03908116610140529485166101a052928416610180529083166101c052821661016052166101e05261019f6101aa565b505050505050610301565b5f54610100900460ff16156102155760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610265575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461027b575f5ffd5b50565b5f5f5f5f5f5f60c08789031215610293575f5ffd5b865161029e81610267565b60208801519096506102af81610267565b60408801519095506102c081610267565b60608801519094506102d181610267565b60808801519093506102e281610267565b60a08801519092506102f381610267565b809150509295509295509295565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c0516101e05161519761043a5f395f8181610838015281816121d30152818161297b01526129ee01525f818161076201528181610ede015281816112cd0152818161238f0152818161283b0152612e7801525f81816106790152818161125a01528181611b8e015281816122730152818161230e015281816127c001528181612dd6015261360701525f818161063f01528181610c9401528181611298015281816124040152818161274701528181612b0c01528181612b830152612d5d01525f81816109090152611d2e01525f818161070101528181610b1d015281816113db015261196701525f6132ef01525f61333e01525f61331901525f61327201525f61329c01525f6132c601526151975ff3fe608060405234801561000f575f5ffd5b506004361061034c575f3560e01c80635df45946116101c95780639feab859116100fe578063ca8aa7c71161009e578063ea32afae11610079578063ea32afae14610904578063f2fde38b1461092b578063fabc1cbc1461093e578063fd39105a14610951575f5ffd5b8063ca8aa7c714610833578063d72d8dd61461085a578063e65797ad14610862575f5ffd5b8063adcf73f7116100d9578063adcf73f7146107c7578063b2d8678d146107da578063c391425e146107ec578063ca0de8821461080c575f5ffd5b80639feab85914610784578063a4d7871f146107ab578063a96f783e146107be575f5ffd5b806384ca5213116101695780638da5cb5b116101445780638da5cb5b146107235780639aa1653d1461072b5780639d8e0c231461074a5780639e9923c21461075d575f5ffd5b806384ca5213146106d6578063871ef049146106e9578063886f1195146106fc575f5ffd5b80636e3b17db116101a45780636e3b17db1461069b578063715018a6146106ae57806381f936d2146106b65780638281ab75146106c3575f5ffd5b80635df459461461063a5780636347c900146106615780636830483514610674575f5ffd5b8063249a0c421161029f5780635140a5481161023f578063595c6a671161021a578063595c6a67146105f85780635ac86ab7146106005780635b0b829f1461061f5780635c975abb14610632575f5ffd5b80635140a548146105b2578063530b97a4146105c55780635865c60c146105d8575f5ffd5b806329d1e0c31161027a57806329d1e0c3146105595780632cdd1e861461056c5780633c2a7f4c1461057f5780633eef3a511461059f575f5ffd5b8063249a0c421461051457806328f61b3114610533578063296bb06414610546575f5ffd5b80630d3f21341161030a578063136439dd116102e5578063136439dd14610473578063143e5915146104865780631478851f146104995780631eb812da146104cb575f5ffd5b80630d3f213414610419578063125e05841461042c57806313542a4e1461044b575f5ffd5b8062cf2ab51461035057806303fd34921461036557806304ec635114610397578063054310e6146103c25780630764cb93146103ed5780630cf4b76714610406575b5f5ffd5b61036361035e366004613ee9565b61098c565b005b610384610373366004613f1a565b5f9081526098602052604090205490565b6040519081526020015b60405180910390f35b6103aa6103a5366004613f42565b610a81565b6040516001600160c01b03909116815260200161038e565b609d546103d5906001600160a01b031681565b6040516001600160a01b03909116815260200161038e565b60a1546103d5906201000090046001600160a01b031681565b610363610414366004613fe8565b610a99565b610363610427366004613f1a565b610afb565b61038461043a366004614019565b609f6020525f908152604090205481565b610384610459366004614019565b6001600160a01b03165f9081526099602052604090205490565b610363610481366004613f1a565b610b08565b610363610494366004614019565b610bdd565b6104bb6104a7366004613f1a565b609a6020525f908152604090205460ff1681565b604051901515815260200161038e565b6104de6104d9366004614034565b610bee565b60408051825163ffffffff908116825260208085015190911690820152918101516001600160c01b03169082015260600161038e565b610384610522366004614069565b609b6020525f908152604090205481565b609e546103d5906001600160a01b031681565b6103d5610554366004613f1a565b610c7c565b610363610567366004614019565b610d05565b61036361057a366004614019565b610d16565b61059261058d366004614019565b610d27565b60405161038e9190614082565b6103636105ad3660046141a3565b610da5565b6103636105c036600461424f565b610de4565b6103636105d336600461432c565b61112e565b6105eb6105e6366004614019565b611354565b60405161038e91906143c4565b6103636113c6565b6104bb61060e366004614069565b6001805460ff9092161b9081161490565b61036361062d3660046143df565b611475565b600154610384565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b6103d561066f366004613f1a565b611491565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b6103636106a9366004614411565b6114b9565b61036361157b565b60a1546104bb9060ff1681565b6103636106d136600461445d565b61158c565b6103846106e436600461454d565b6115a1565b6103aa6106f7366004613f1a565b6115ea565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b6103d56115f4565b6096546107389060ff1681565b60405160ff909116815260200161038e565b610363610758366004614616565b61160c565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b6103847f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de681565b6104bb6107b9366004614069565b611674565b61038460a05481565b6103636107d5366004614658565b61167e565b60a1546104bb90610100900460ff1681565b6107ff6107fa3660046146cf565b6118e1565b60405161038e9190614774565b6103847f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a81565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b609c54610384565b6108d0610870366004614069565b60408051606080820183525f808352602080840182905292840181905260ff9490941684526097825292829020825193840183525463ffffffff8116845261ffff600160201b8204811692850192909252600160301b9004169082015290565b60408051825163ffffffff16815260208084015161ffff90811691830191909152928201519092169082015260600161038e565b6103d57f000000000000000000000000000000000000000000000000000000000000000081565b610363610939366004614019565b6118ef565b61036361094c366004613f1a565b611965565b61097f61095f366004614019565b6001600160a01b03165f9081526099602052604090206001015460ff1690565b60405161038e91906147bc565b6001546002906004908116036109b55760405163840a48d560e01b815260040160405180910390fd5b5f5b8251811015610a7c575f8382815181106109d3576109d36147ca565b6020908102919091018101516001600160a01b0381165f9081526099835260408082208151808301909252805482526001810154939550919390929083019060ff166002811115610a2657610a26614390565b6002811115610a3757610a37614390565b90525080519091505f610a4982611a7c565b90505f610a5e826001600160c01b0316611a88565b9050610a6b858583611b51565b5050600190930192506109b7915050565b505050565b5f610a8f6098858585611c33565b90505b9392505050565b6001335f9081526099602052604090206001015460ff166002811115610ac157610ac1614390565b14610adf5760405163aba4733960e01b815260040160405180910390fd5b335f90815260996020526040902054610af89082611d17565b50565b610b03611dc2565b60a055565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610b6a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b8e91906147de565b610bab57604051631d77d47760e21b815260040160405180910390fd5b6001548181168114610bd05760405163c61dca5d60e01b815260040160405180910390fd5b610bd982611e21565b5050565b610be5611dc2565b610af881611e5e565b604080516060810182525f80825260208201819052918101919091525f838152609860205260409020805483908110610c2957610c296147ca565b5f91825260209182902060408051606081018252919092015463ffffffff8082168352600160201b820416938201939093526001600160c01b03600160401b909304929092169082015290505b92915050565b6040516308f6629d60e31b8152600481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906347b314e890602401602060405180830381865afa158015610ce1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c7691906147fd565b610d0d611dc2565b610af881611e88565b610d1e611dc2565b610af881611ef1565b604080518082019091525f8082526020820152610c76610da07f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de684604051602001610d859291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120611f5a565b611fa6565b610dad611dc2565b60a15460ff16610dd057604051635b77901960e01b815260040160405180910390fd5b610dde848484600185612030565b50505050565b600154600290600490811603610e0d5760405163840a48d560e01b815260040160405180910390fd5b610e5283838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505060965460ff16915061246f9050565b5083518214610e745760405163aaad13f760e01b815260040160405180910390fd5b5f5b82811015611127575f848483818110610e9157610e916147ca565b885192013560f81c92505f9188915084908110610eb057610eb06147ca565b60209081029190910101516040516379a0849160e11b815260ff841660048201529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f341092290602401602060405180830381865afa158015610f23573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f479190614818565b63ffffffff16815114610f6d57604051638e5aeee760e01b815260040160405180910390fd5b5f805b82518110156110ce575f838281518110610f8c57610f8c6147ca565b6020908102919091018101516001600160a01b0381165f9081526099835260408082208151808301909252805482526001810154939550919390929083019060ff166002811115610fdf57610fdf614390565b6002811115610ff057610ff0614390565b90525080519091505f61100282611a7c565b905060016001600160c01b03821660ff8a161c8116146110355760405163d053aa2160e01b815260040160405180910390fd5b856001600160a01b0316846001600160a01b0316116110675760405163ba50f91160e01b815260040160405180910390fd5b506110c183838d8b8e61107b826001614847565b926110889392919061485a565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611b5192505050565b5090925050600101610f70565b5060ff83165f818152609b6020908152604091829020439081905591519182527f46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4910160405180910390a2505050806001019050610e76565b5050505050565b5f54610100900460ff161580801561114c57505f54600160ff909116105b806111655750303b15801561116557505f5460ff166001145b6111cd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff1916600117905580156111ee575f805461ff0019166101001790555b6111f7866124a3565b61120085611e88565b61120983611e21565b61121284611ef1565b61121b82611e5e565b609c8054600181810183555f8390527faf85b9071dfafeac1409d3f1d19bafc9bc7c37974cde8df0ee6168f0086e539c91820180546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166001600160a01b03199283161790925584548084018655840180547f000000000000000000000000000000000000000000000000000000000000000084169083161790558454928301909455910180547f00000000000000000000000000000000000000000000000000000000000000009092169190921617905560a1805461010161ffff19909116179055801561134c575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b604080518082019091525f80825260208201526001600160a01b0382165f908152609960209081526040918290208251808401909352805483526001810154909183019060ff1660028111156113ac576113ac614390565b60028111156113bd576113bd614390565b90525092915050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611428573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061144c91906147de565b61146957604051631d77d47760e21b815260040160405180910390fd5b6114735f19611e21565b565b61147d611dc2565b81611487816124f4565b610a7c838361251d565b609c81815481106114a0575f80fd5b5f918252602090912001546001600160a01b0316905081565b6114c16125c2565b6001600160a01b0382165f908152609f60209081526040808320429055609990915281208054609654919290916114fc90859060ff1661246f565b90505f61150883611a7c565b905060018085015460ff16600281111561152457611524614390565b14801561153957506001600160c01b03821615155b801561155757506115576001600160c01b0383811690831681161490565b1561134c5761156686866125ed565b60a15460ff161561134c5761134c8686612874565b611583611dc2565b6114735f6124a3565b611594611dc2565b610a7c8383835f5f612030565b5f6115e07f4d404e3276e7ac2163d8ee476afa6a41d1f68fb71f2d8b6546b24e55ce01b72a8787878787604051602001610d8596959493929190614881565b9695505050505050565b5f610c7682611a7c565b5f6116076064546001600160a01b031690565b905090565b6116146129e3565b6001805460029081160361163b5760405163840a48d560e01b815260040160405180910390fd5b60a15460ff1661165e57604051635b77901960e01b815260040160405180910390fd5b5f61166883612a2c565b9050610dde84826125ed565b5f610c7682612ad4565b6116866129e3565b600180545f91908116036116ad5760405163840a48d560e01b815260040160405180910390fd5b60a15460ff166116d057604051635b77901960e01b815260040160405180910390fd5b5f6116da85612a2c565b90505f80806116eb86880188614a05565b9250925092505f6116fc8a83612aeb565b90505f84600181111561171157611711614390565b036117b8575f6117238b838887612c19565b5190505f5b86518110156117b1575f878281518110611744576117446147ca565b0160209081015160f81c5f8181526097909252604090912054845191925063ffffffff169084908490811061177b5761177b6147ca565b602002602001015163ffffffff1611156117a857604051633cb89c9760e01b815260040160405180910390fd5b50600101611728565b505061180d565b60018460018111156117cc576117cc614390565b036117f4575f806117df898b018b614a60565b945094505050506117b18c8489888686612efd565b60405163354bb8ab60e01b815260040160405180910390fd5b60016001600160a01b038b165f9081526099602052604090206001015460ff16600281111561183e5761183e614390565b146118d557604080518082018252828152600160208083018281526001600160a01b038f165f908152609990925293902082518155925183820180549394939192909160ff19169083600281111561189857611898614390565b0217905550506040518291506001600160a01b038c16907fe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe905f90a35b50505050505050505050565b6060610a9260988484613123565b6118f7611dc2565b6001600160a01b03811661195c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016111c4565b610af8816124a3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119c1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119e591906147fd565b6001600160a01b0316336001600160a01b031614611a165760405163794821ff60e01b815260040160405180910390fd5b60015480198219811614611a3d5760405163c61dca5d60e01b815260040160405180910390fd5b600182905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020015b60405180910390a25050565b5f610c766098836131d2565b60605f5f611a958461323c565b61ffff166001600160401b03811115611ab057611ab0613db2565b6040519080825280601f01601f191660200182016040528015611ada576020820181803683370190505b5090505f805b825182108015611af1575061010081105b15611b47576001811b935085841615611b37578060f81b838381518110611b1a57611b1a6147ca565b60200101906001600160f81b03191690815f1a9053508160010191505b611b4081614b59565b9050611ae0565b5090949350505050565b600182602001516002811115611b6957611b69614390565b14611b7357505050565b81516040516333567f7f60e11b81525f906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906366acfefe90611bc790889086908890600401614b9f565b6020604051808303815f875af1158015611be3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c079190614bce565b90506001600160c01b038116156111275761112785611c2e836001600160c01b0316611a88565b6125ed565b5f838152602085905260408120805482919084908110611c5557611c556147ca565b5f91825260209182902060408051606081018252929091015463ffffffff808216808552600160201b8304821695850195909552600160401b9091046001600160c01b03169183019190915290925085161015611cc557604051636cb19aff60e01b815260040160405180910390fd5b602081015163ffffffff161580611ceb5750806020015163ffffffff168463ffffffff16105b611d085760405163bbba60cb60e01b815260040160405180910390fd5b6040015190505b949350505050565b6040516378219b3f60e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f043367e90611d659085908590600401614bf4565b5f604051808303815f87803b158015611d7c575f5ffd5b505af1158015611d8e573d5f5f3e3d5ffd5b50505050817fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa82604051611a709190614c0c565b33611dcb6115f4565b6001600160a01b0316146114735760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016111c4565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b60a180546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b609d54604080516001600160a01b03928316815291831660208301527f315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c910160405180910390a1609d80546001600160a01b0319166001600160a01b0392909216919091179055565b609e54604080516001600160a01b03928316815291831660208301527f8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9910160405180910390a1609e80546001600160a01b0319166001600160a01b0392909216919091179055565b5f610c76611f66613266565b8360405161190160f01b602082015260228101839052604281018290525f9060620160405160208183030381529060405280519060200120905092915050565b604080518082019091525f80825260208201525f8080611fd35f5160206151425f395f51905f5286614c32565b90505b611fdf8161338c565b90935091505f5160206151425f395f51905f528283098303612017576040805180820190915290815260208101919091529392505050565b5f5160206151425f395f51905f52600182089050611fd6565b60965460ff1660c0811061205757604051633cb89c9760e01b815260040160405180910390fd5b612062816001614c45565b6096805460ff191660ff9290921691909117905580612081818861251d565b60a15460ff168015612099575061209781612ad4565b155b15612244576040805160018082528183019092525f91816020015b604080518082019091525f8152606060208201528152602001906001900390816120b45790505090505f86516001600160401b038111156120f7576120f7613db2565b604051908082528060200260200182016040528015612120578160200160208202803683370190505b5090505f5b875181101561217d57878181518110612140576121406147ca565b60200260200101515f015182828151811061215d5761215d6147ca565b6001600160a01b0390921660209283029190910190910152600101612125565b5060405180604001604052808460ff1663ffffffff16815260200182815250825f815181106121ae576121ae6147ca565b602090810291909101015260a154604051630130fc2760e51b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169263261f84e0926122149262010000909204909116908690600401614c5e565b5f604051808303815f87803b15801561222b575f5ffd5b505af115801561223d573d5f5f3e3d5ffd5b5050505050505b5f84600181111561225757612257614390565b036122de57604051633aea0b9d60e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906375d4173a906122ac9084908a908a90600401614d78565b5f604051808303815f87803b1580156122c3575f5ffd5b505af11580156122d5573d5f5f3e3d5ffd5b50505050612377565b60018460018111156122f2576122f2614390565b0361237757604051630662d3e160e51b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063cc5a7c20906123499084908a9088908b90600401614da2565b5f604051808303815f87803b158015612360575f5ffd5b505af1158015612372573d5f5f3e3d5ffd5b505050505b60405163136ca0f960e11b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906326d941f2906024015f604051808303815f87803b1580156123d8575f5ffd5b505af11580156123ea573d5f5f3e3d5ffd5b505060405163136ca0f960e11b815260ff841660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692506326d941f291506024015b5f604051808303815f87803b158015612450575f5ffd5b505af1158015612462573d5f5f3e3d5ffd5b5050505050505050505050565b5f5f61247a84613408565b9050808360ff166001901b11610a925760405163ca95733360e01b815260040160405180910390fd5b606480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60965460ff90811690821610610af857604051637310cff560e11b815260040160405180910390fd5b60ff82165f81815260976020908152604091829020845181548684018051888701805163ffffffff90951665ffffffffffff199094168417600160201b61ffff938416021767ffff0000000000001916600160301b95831695909502949094179094558551918252518316938101939093525116918101919091527f3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac90606001611a70565b609e546001600160a01b03163314611473576040516376d8ab1760e11b815260040160405180910390fd5b6001600160a01b0382165f9081526099602052604081208054909161261182611a7c565b905060018084015460ff16600281111561262d5761262d614390565b1461264b5760405163aba4733960e01b815260040160405180910390fd5b6096545f9061265e90869060ff1661246f565b90506001600160c01b038116612687576040516368b6a87560e11b815260040160405180910390fd5b61269e6001600160c01b0382811690841681161490565b6126bb5760405163d053aa2160e01b815260040160405180910390fd5b6001600160c01b03818116198316166126d484826134c3565b6001600160c01b038116612730576001600160a01b0387165f81815260996020526040808220600101805460ff19166002179055518692917f396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e491a35b60405163f4e24fe560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f4e24fe59061277e908a908a90600401614dd8565b5f604051808303815f87803b158015612795575f5ffd5b505af11580156127a7573d5f5f3e3d5ffd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd91506127f99087908a90600401614bf4565b5f604051808303815f87803b158015612810575f5ffd5b505af1158015612822573d5f5f3e3d5ffd5b505060405163bd29b8cd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063bd29b8cd91506124399087908a90600401614bf4565b5f81516001600160401b0381111561288e5761288e613db2565b6040519080825280602002602001820160405280156128b7578160200160208202803683370190505b5090505f805b8351811015612936575f8482815181106128d9576128d96147ca565b016020015160f81c90506128ec81612ad4565b1561292d5760ff8116848461290081614b59565b955081518110612912576129126147ca565b602002602001019063ffffffff16908163ffffffff16815250505b506001016128bd565b508015610dde57808252604080516060810182526001600160a01b03868116825260a154620100009004811660208301528183018590529151636e3492b560e01b81527f000000000000000000000000000000000000000000000000000000000000000090921691636e3492b5916129b091600401614dfb565b5f604051808303815f87803b1580156129c7575f5ffd5b505af11580156129d9573d5f5f3e3d5ffd5b5050505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611473576040516323d871a560e01b815260040160405180910390fd5b60605f82516001600160401b03811115612a4857612a48613db2565b6040519080825280601f01601f191660200182016040528015612a72576020820181803683370190505b5090505f5b8351811015612acd57838181518110612a9257612a926147ca565b602002602001015160f81b828281518110612aaf57612aaf6147ca565b60200101906001600160f81b03191690815f1a905350600101612a77565b5092915050565b60a2545f90610c76908360ff161c60019081161490565b6040516309aa152760e11b81526001600160a01b0383811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906313542a4e90602401602060405180830381865afa158015612b53573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b779190614e69565b90505f819003610c76577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bf79ce588484612bbb87610d27565b6040518463ffffffff1660e01b8152600401612bd993929190614ea2565b6020604051808303815f875af1158015612bf5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a929190614e69565b612c3d60405180606001604052806060815260200160608152602001606081525090565b6096545f90612c5090859060ff1661246f565b90505f612c5c86611a7c565b90506001600160c01b038216612c85576040516313ca465760e01b815260040160405180910390fd5b8082166001600160c01b031615612caf57604051630c6816cd60e01b815260040160405180910390fd5b60a0546001600160a01b0388165f908152609f60205260409020546001600160c01b0383811690851617914291612ce69190614847565b10612d0457604051631968677d60e11b815260040160405180910390fd5b612d0e87826134c3565b867fec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa86604051612d3e9190614c0c565b60405180910390a2604051631fd93ca960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633fb2795290612d94908b908a90600401614dd8565b5f604051808303815f87803b158015612dab575f5ffd5b505af1158015612dbd573d5f5f3e3d5ffd5b5050604051632550477760e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063255047779150612e11908b908b908b90600401614b9f565b5f604051808303815f875af1158015612e2c573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612e539190810190614f7f565b60408087019190915260208601919091525162bff04d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169062bff04d90612eae908a908a90600401614bf4565b5f604051808303815f875af1158015612ec9573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612ef09190810190614fd8565b8452505050949350505050565b8351825114612f1f5760405163aaad13f760e01b815260040160405180910390fd5b612f2b868684846134cf565b5f612f3887878787612c19565b90505f5b85518110156129d9575f60975f888481518110612f5b57612f5b6147ca565b0160209081015160f81c82528181019290925260409081015f208151606081018352905463ffffffff811680835261ffff600160201b8304811695840195909552600160301b90910490931691810191909152845180519193509084908110612fc657612fc66147ca565b602002602001015163ffffffff16111561311a5761305a878381518110612fef57612fef6147ca565b602001015160f81c60f81b60f81c84604001518481518110613013576130136147ca565b60200260200101518b86602001518681518110613032576130326147ca565b602002602001015189878151811061304c5761304c6147ca565b60200260200101518661357a565b6040805160018082528183019092525f9160208201818036833701905050905087838151811061308c5761308c6147ca565b602001015160f81c60f81b815f815181106130a9576130a96147ca565b60200101906001600160f81b03191690815f1a9053506130e68684815181106130d4576130d46147ca565b602002602001015160200151826125ed565b60a15460ff161561311857613118868481518110613106576131066147ca565b60200260200101516020015182612874565b505b50600101612f3c565b60605f82516001600160401b0381111561313f5761313f613db2565b604051908082528060200260200182016040528015613168578160200160208202803683370190505b5090505f5b83518110156131c95761319a868686848151811061318d5761318d6147ca565b60200260200101516136fb565b8282815181106131ac576131ac6147ca565b63ffffffff9092166020928302919091019091015260010161316d565b50949350505050565b5f818152602083905260408120548082036131f0575f915050610c76565b5f838152602085905260409020613208600183615067565b81548110613218576132186147ca565b5f91825260209091200154600160401b90046001600160c01b03169150610c769050565b5f805b8215610c7657613250600184615067565b909216918061325e8161507a565b91505061323f565b5f306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156132be57507f000000000000000000000000000000000000000000000000000000000000000046145b156132e857507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b5f80805f5160206151425f395f51905f5260035f5160206151425f395f51905f52865f5160206151425f395f51905f52888909090890505f6133fc827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f525f5160206151425f395f51905f52613813565b91959194509092505050565b5f6101008251111561342d57604051637da54e4760e11b815260040160405180910390fd5b81515f0361343c57505f919050565b5f5f835f81518110613450576134506147ca565b0160200151600160f89190911c81901b92505b84518110156134ba5784818151811061347e5761347e6147ca565b0160200151600160f89190911c1b91508282116134ae57604051631019106960e31b815260040160405180910390fd5b91811791600101613463565b50909392505050565b610bd96098838361388c565b6020808201515f908152609a909152604090205460ff161561350457604051636fbefec360e11b815260040160405180910390fd5b428160400151101561352957604051630819bdcd60e01b815260040160405180910390fd5b602080820180515f908152609a909252604091829020805460ff19166001179055609d54905191830151610dde926001600160a01b039092169161357391889188918891906115a1565b8351613a45565b6020808301516001600160a01b038082165f81815260999094526040909320549192908716036135bd576040516356168b4160e11b815260040160405180910390fd5b8760ff16845f015160ff16146135e657604051638e5aeee760e01b815260040160405180910390fd5b604051635401ed2760e01b81526004810182905260ff891660248201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635401ed2790604401602060405180830381865afa158015613654573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613678919061509a565b90506136848185613a6d565b6001600160601b0316866001600160601b0316116136b557604051634c44995d60e01b815260040160405180910390fd5b6136bf8885613a90565b6001600160601b0316816001600160601b0316106136f05760405163b187e86960e01b815260040160405180910390fd5b505050505050505050565b5f81815260208490526040812054815b8181101561377e57600161371f8284615067565b6137299190615067565b92508463ffffffff16865f8681526020019081526020015f208463ffffffff1681548110613759576137596147ca565b5f9182526020909120015463ffffffff1611613776575050610a92565b60010161370b565b5060405162461bcd60e51b815260206004820152605c60248201527f5265676973747279436f6f7264696e61746f722e67657451756f72756d42697460448201527f6d6170496e6465784174426c6f636b4e756d6265723a206e6f206269746d617060648201527f2075706461746520666f756e6420666f72206f70657261746f72496400000000608482015260a4016111c4565b5f5f61381d613d76565b613825613d94565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828061386257fe5b50826138815760405163d51edae360e01b815260040160405180910390fd5b505195945050505050565b5f8281526020849052604081205490819003613930575f83815260208581526040808320815160608101835263ffffffff43811682528185018681526001600160c01b03808a16958401958652845460018101865594885295909620915191909201805495519351909416600160401b026001600160401b03938316600160201b0267ffffffffffffffff1990961691909216179390931716919091179055610dde565b5f838152602085905260408120613948600184615067565b81548110613958576139586147ca565b5f918252602090912001805490915063ffffffff4381169116036139995780546001600160401b0316600160401b6001600160c01b03851602178155611127565b805463ffffffff438116600160201b81810267ffffffff00000000199094169390931784555f8781526020898152604080832081516060810183529485528483018481526001600160c01b03808c1693870193845282546001810184559286529390942094519401805493519151909216600160401b026001600160401b0391861690960267ffffffffffffffff19909316939094169290921717919091169190911790555050505050565b613a50838383613aa9565b610a7c57604051638baa579f60e01b815260040160405180910390fd5b60208101515f9061271090613a869061ffff16856150b5565b610a9291906150d7565b60408101515f9061271090613a869061ffff16856150b5565b5f5f5f613ab68585613bee565b90925090505f816004811115613ace57613ace614390565b148015613aec5750856001600160a01b0316826001600160a01b0316145b15613afc57600192505050610a92565b5f5f876001600160a01b0316631626ba7e60e01b8888604051602401613b23929190614bf4565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051613b619190615104565b5f60405180830381855afa9150503d805f8114613b99576040519150601f19603f3d011682016040523d82523d5f602084013e613b9e565b606091505b5091509150818015613bb1575080516020145b8015613be257508051630b135d3f60e11b90613bd6908301602090810190840161511a565b6001600160e01b031916145b98975050505050505050565b5f5f8251604103613c22576020830151604084015160608501515f1a613c1687828585613c59565b94509450505050613c52565b8251604003613c4b5760208301516040840151613c40868383613d3e565b935093505050613c52565b505f905060025b9250929050565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115613c8e57505f90506003613d35565b8460ff16601b14158015613ca657508460ff16601c14155b15613cb657505f90506004613d35565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613d07573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116613d2f575f60019250925050613d35565b91505f90505b94509492505050565b5f806001600160ff1b03831681613d5a60ff86901c601b614847565b9050613d6887828885613c59565b935093505050935093915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715613de857613de8613db2565b60405290565b604080519081016001600160401b0381118282101715613de857613de8613db2565b604051601f8201601f191681016001600160401b0381118282101715613e3857613e38613db2565b604052919050565b5f6001600160401b03821115613e5857613e58613db2565b5060051b60200190565b6001600160a01b0381168114610af8575f5ffd5b5f82601f830112613e85575f5ffd5b8135613e98613e9382613e40565b613e10565b8082825260208201915060208360051b860101925085831115613eb9575f5ffd5b602085015b83811015613edf578035613ed181613e62565b835260209283019201613ebe565b5095945050505050565b5f60208284031215613ef9575f5ffd5b81356001600160401b03811115613f0e575f5ffd5b611d0f84828501613e76565b5f60208284031215613f2a575f5ffd5b5035919050565b63ffffffff81168114610af8575f5ffd5b5f5f5f60608486031215613f54575f5ffd5b833592506020840135613f6681613f31565b929592945050506040919091013590565b5f82601f830112613f86575f5ffd5b8135602083015f5f6001600160401b03841115613fa557613fa5613db2565b50601f8301601f1916602001613fba81613e10565b915050828152858383011115613fce575f5ffd5b828260208301375f92810160200192909252509392505050565b5f60208284031215613ff8575f5ffd5b81356001600160401b0381111561400d575f5ffd5b611d0f84828501613f77565b5f60208284031215614029575f5ffd5b8135610a9281613e62565b5f5f60408385031215614045575f5ffd5b50508035926020909101359150565b803560ff81168114614064575f5ffd5b919050565b5f60208284031215614079575f5ffd5b610a9282614054565b815181526020808301519082015260408101610c76565b803561ffff81168114614064575f5ffd5b5f606082840312156140ba575f5ffd5b6140c2613dc6565b905081356140cf81613f31565b81526140dd60208301614099565b60208201526140ee60408301614099565b604082015292915050565b6001600160601b0381168114610af8575f5ffd5b5f82601f83011261411c575f5ffd5b813561412a613e9382613e40565b8082825260208201915060208360061b86010192508583111561414b575f5ffd5b602085015b83811015613edf5760408188031215614167575f5ffd5b61416f613dee565b813561417a81613e62565b8152602082013561418a816140f9565b6020828101919091529084529290920191604001614150565b5f5f5f5f60c085870312156141b6575f5ffd5b6141c086866140aa565b935060608501356141d0816140f9565b925060808501356001600160401b038111156141ea575f5ffd5b6141f68782880161410d565b92505060a085013561420781613f31565b939692955090935050565b5f5f83601f840112614222575f5ffd5b5081356001600160401b03811115614238575f5ffd5b602083019150836020828501011115613c52575f5ffd5b5f5f5f60408486031215614261575f5ffd5b83356001600160401b03811115614276575f5ffd5b8401601f81018613614286575f5ffd5b8035614294613e9382613e40565b8082825260208201915060208360051b8501019250888311156142b5575f5ffd5b602084015b838110156142f55780356001600160401b038111156142d7575f5ffd5b6142e68b602083890101613e76565b845250602092830192016142ba565b50955050505060208401356001600160401b03811115614313575f5ffd5b61431f86828701614212565b9497909650939450505050565b5f5f5f5f5f60a08688031215614340575f5ffd5b853561434b81613e62565b9450602086013561435b81613e62565b9350604086013561436b81613e62565b925060608601359150608086013561438281613e62565b809150509295509295909350565b634e487b7160e01b5f52602160045260245ffd5b600381106143c057634e487b7160e01b5f52602160045260245ffd5b9052565b815181526020808301516040830191612acd908401826143a4565b5f5f608083850312156143f0575f5ffd5b6143f983614054565b915061440884602085016140aa565b90509250929050565b5f5f60408385031215614422575f5ffd5b823561442d81613e62565b915060208301356001600160401b03811115614447575f5ffd5b61445385828601613f77565b9150509250929050565b5f5f5f60a0848603121561446f575f5ffd5b61447985856140aa565b92506060840135614489816140f9565b915060808401356001600160401b038111156144a3575f5ffd5b6144af8682870161410d565b9150509250925092565b5f82601f8301126144c8575f5ffd5b81356144d6613e9382613e40565b8082825260208201915060208360061b8601019250858311156144f7575f5ffd5b602085015b83811015613edf5760408188031215614513575f5ffd5b61451b613dee565b61452482614054565b8152602082013561453481613e62565b60208281019190915290845292909201916040016144fc565b5f5f5f5f5f60a08688031215614561575f5ffd5b853561456c81613e62565b94506020860135935060408601356001600160401b0381111561458d575f5ffd5b614599888289016144b9565b9598949750949560608101359550608001359392505050565b5f82601f8301126145c1575f5ffd5b81356145cf613e9382613e40565b8082825260208201915060208360051b8601019250858311156145f0575f5ffd5b602085015b83811015613edf57803561460881613f31565b8352602092830192016145f5565b5f5f60408385031215614627575f5ffd5b823561463281613e62565b915060208301356001600160401b0381111561464c575f5ffd5b614453858286016145b2565b5f5f5f5f6060858703121561466b575f5ffd5b843561467681613e62565b935060208501356001600160401b03811115614690575f5ffd5b61469c878288016145b2565b93505060408501356001600160401b038111156146b7575f5ffd5b6146c387828801614212565b95989497509550505050565b5f5f604083850312156146e0575f5ffd5b82356146eb81613f31565b915060208301356001600160401b03811115614705575f5ffd5b8301601f81018513614715575f5ffd5b8035614723613e9382613e40565b8082825260208201915060208360051b850101925087831115614744575f5ffd5b6020840193505b8284101561476657833582526020938401939091019061474b565b809450505050509250929050565b602080825282518282018190525f918401906040840190835b818110156147b157835163ffffffff1683526020938401939092019160010161478d565b509095945050505050565b60208101610c7682846143a4565b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156147ee575f5ffd5b81518015158114610a92575f5ffd5b5f6020828403121561480d575f5ffd5b8151610a9281613e62565b5f60208284031215614828575f5ffd5b8151610a9281613f31565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610c7657610c76614833565b5f5f85851115614868575f5ffd5b83861115614874575f5ffd5b5050820193919092039150565b5f60c0820188835260018060a01b038816602084015286604084015260c0606084015280865180835260e0850191506020880192505f5b818110156148ee578351805160ff1684526020908101516001600160a01b031681850152909301926040909201916001016148b8565b50506080840195909552505060a00152949350505050565b803560028110614064575f5ffd5b5f60408284031215614924575f5ffd5b61492c613dee565b823581526020928301359281019290925250919050565b5f82601f830112614952575f5ffd5b61495a613dee565b80604084018581111561496b575f5ffd5b845b818110156147b157803584526020938401930161496d565b5f818303610100811215614997575f5ffd5b61499f613dc6565b91506149ab8484614914565b82526149ba8460408501614914565b60208301526080607f19820112156149d0575f5ffd5b506149d9613dee565b6149e68460808501614943565b81526149f58460c08501614943565b6020820152604082015292915050565b5f5f5f6101408486031215614a18575f5ffd5b614a2184614906565b925060208401356001600160401b03811115614a3b575f5ffd5b614a4786828701613f77565b925050614a578560408601614985565b90509250925092565b5f5f5f5f5f6101808688031215614a75575f5ffd5b614a7e86614906565b945060208601356001600160401b03811115614a98575f5ffd5b614aa488828901613f77565b945050614ab48760408801614985565b92506101408601356001600160401b03811115614acf575f5ffd5b614adb888289016144b9565b9250506101608601356001600160401b03811115614af7575f5ffd5b860160608189031215614b08575f5ffd5b614b10613dc6565b81356001600160401b03811115614b25575f5ffd5b614b318a828501613f77565b8252506020828101359082015260409182013591810191909152949793965091945092919050565b5f60018201614b6a57614b6a614833565b5060010190565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60018060a01b0384168152826020820152606060408201525f614bc56060830184614b71565b95945050505050565b5f60208284031215614bde575f5ffd5b81516001600160c01b0381168114610a92575f5ffd5b828152604060208201525f610a8f6040830184614b71565b602081525f610a926020830184614b71565b634e487b7160e01b5f52601260045260245ffd5b5f82614c4057614c40614c1e565b500690565b60ff8181168382160190811115610c7657610c76614833565b5f6040820160018060a01b03851683526040602084015280845180835260608501915060608160051b8601019250602086015f5b82811015614d1357868503605f190184528151805163ffffffff168652602090810151604082880181905281519088018190529101905f9060608801905b80831015614cfb5783516001600160a01b031682526020938401936001939093019290910190614cd0565b50965050506020938401939190910190600101614c92565b5092979650505050505050565b5f8151808452602084019350602083015f5b82811015614d6e57815180516001600160a01b031687526020908101516001600160601b03168188015260409096019590910190600101614d32565b5093949350505050565b60ff841681526001600160601b0383166020820152606060408201525f614bc56060830184614d20565b60ff851681526001600160601b038416602082015263ffffffff83166040820152608060608201525f6115e06080830184614d20565b6001600160a01b03831681526040602082018190525f90610a8f90830184614b71565b602080825282516001600160a01b039081168383015283820151166040808401919091528301516060808401528051608084018190525f929190910190829060a08501905b80831015613edf5763ffffffff8451168252602082019150602084019350600183019250614e40565b5f60208284031215614e79575f5ffd5b5051919050565b805f5b6002811015610dde578151845260209384019390910190600101614e83565b6001600160a01b03841681528251805160208084019190915201516040820152610160810160208481015180516060850152908101516080840152506040840151614ef160a084018251614e80565b60200151614f0260e0840182614e80565b5082516101208301526020830151610140830152611d0f565b5f82601f830112614f2a575f5ffd5b8151614f38613e9382613e40565b8082825260208201915060208360051b860101925085831115614f59575f5ffd5b602085015b83811015613edf578051614f71816140f9565b835260209283019201614f5e565b5f5f60408385031215614f90575f5ffd5b82516001600160401b03811115614fa5575f5ffd5b614fb185828601614f1b565b92505060208301516001600160401b03811115614fcc575f5ffd5b61445385828601614f1b565b5f60208284031215614fe8575f5ffd5b81516001600160401b03811115614ffd575f5ffd5b8201601f8101841361500d575f5ffd5b805161501b613e9382613e40565b8082825260208201915060208360051b85010192508683111561503c575f5ffd5b6020840193505b828410156115e057835161505681613f31565b825260209384019390910190615043565b81810381811115610c7657610c76614833565b5f61ffff821661ffff810361509157615091614833565b60010192915050565b5f602082840312156150aa575f5ffd5b8151610a92816140f9565b6001600160601b038181168382160290811690818114612acd57612acd614833565b5f6001600160601b038316806150ef576150ef614c1e565b806001600160601b0384160491505092915050565b5f82518060208501845e5f920191825250919050565b5f6020828403121561512a575f5ffd5b81516001600160e01b031981168114610a92575f5ffdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a26469706673582212206a2ad208f238c76f31569e46ed8871c3a213f57f7495fc7089df2a687f22d7a464736f6c634300081b0033",
}

// ContractSlashingRegistryCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSlashingRegistryCoordinatorMetaData.ABI instead.
var ContractSlashingRegistryCoordinatorABI = ContractSlashingRegistryCoordinatorMetaData.ABI

// ContractSlashingRegistryCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSlashingRegistryCoordinatorMetaData.Bin instead.
var ContractSlashingRegistryCoordinatorBin = ContractSlashingRegistryCoordinatorMetaData.Bin

// DeployContractSlashingRegistryCoordinator deploys a new Ethereum contract, binding an instance of ContractSlashingRegistryCoordinator to it.
func DeployContractSlashingRegistryCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _stakeRegistry common.Address, _blsApkRegistry common.Address, _indexRegistry common.Address, _socketRegistry common.Address, _allocationManager common.Address, _pauserRegistry common.Address) (common.Address, *types.Transaction, *ContractSlashingRegistryCoordinator, error) {
	parsed, err := ContractSlashingRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSlashingRegistryCoordinatorBin), backend, _stakeRegistry, _blsApkRegistry, _indexRegistry, _socketRegistry, _allocationManager, _pauserRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSlashingRegistryCoordinator{ContractSlashingRegistryCoordinatorCaller: ContractSlashingRegistryCoordinatorCaller{contract: contract}, ContractSlashingRegistryCoordinatorTransactor: ContractSlashingRegistryCoordinatorTransactor{contract: contract}, ContractSlashingRegistryCoordinatorFilterer: ContractSlashingRegistryCoordinatorFilterer{contract: contract}}, nil
}

// ContractSlashingRegistryCoordinatorMethods is an auto generated interface around an Ethereum contract.
type ContractSlashingRegistryCoordinatorMethods interface {
	ContractSlashingRegistryCoordinatorCalls
	ContractSlashingRegistryCoordinatorTransacts
	ContractSlashingRegistryCoordinatorFilters
}

// ContractSlashingRegistryCoordinatorCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractSlashingRegistryCoordinatorCalls interface {
	OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	AccountIdentifier(opts *bind.CallOpts) (common.Address, error)

	AllocationManager(opts *bind.CallOpts) (common.Address, error)

	BlsApkRegistry(opts *bind.CallOpts) (common.Address, error)

	CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error)

	ChurnApprover(opts *bind.CallOpts) (common.Address, error)

	EjectionCooldown(opts *bind.CallOpts) (*big.Int, error)

	Ejector(opts *bind.CallOpts) (common.Address, error)

	GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetOperator(opts *bind.CallOpts, operator common.Address) (ISlashingRegistryCoordinatorTypesOperatorInfo, error)

	GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error)

	GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (ISlashingRegistryCoordinatorTypesOperatorSetParam, error)

	GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error)

	GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error)

	GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error)

	GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate, error)

	IndexRegistry(opts *bind.CallOpts) (common.Address, error)

	IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error)

	IsM2Quorum(opts *bind.CallOpts, quorumNumber uint8) (bool, error)

	LastEjectionTimestamp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	M2QuorumsDisabled(opts *bind.CallOpts) (bool, error)

	NumRegistries(opts *bind.CallOpts) (*big.Int, error)

	OperatorSetsEnabled(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error)

	QuorumCount(opts *bind.CallOpts) (uint8, error)

	QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error)

	Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error)

	SocketRegistry(opts *bind.CallOpts) (common.Address, error)

	StakeRegistry(opts *bind.CallOpts) (common.Address, error)
}

// ContractSlashingRegistryCoordinatorTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractSlashingRegistryCoordinatorTransacts interface {
	CreateSlashableStakeQuorum(opts *bind.TransactOpts, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error)

	CreateTotalDelegatedStakeQuorum(opts *bind.TransactOpts, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error)

	DeregisterOperator(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error)

	EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _accountIdentifier common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAccountIdentifier(opts *bind.TransactOpts, _accountIdentifier common.Address) (*types.Transaction, error)

	SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error)

	SetEjectionCooldown(opts *bind.TransactOpts, _ejectionCooldown *big.Int) (*types.Transaction, error)

	SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error)

	SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error)

	UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error)

	UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error)
}

// ContractSlashingRegistryCoordinatorFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractSlashingRegistryCoordinatorFilters interface {
	FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator, error)
	WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error)
	ParseChurnApproverUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorChurnApproverUpdated, error)

	FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorEjectorUpdatedIterator, error)
	WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorEjectorUpdated) (event.Subscription, error)
	ParseEjectorUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorEjectorUpdated, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractSlashingRegistryCoordinatorInitialized, error)

	FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator, error)
	WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorDeregistered(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorDeregistered, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorRegistered, error)

	FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator, error)
	WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseOperatorSetParamsUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated, error)

	FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator, error)
	WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorSocketUpdate(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorSocketUpdate, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractSlashingRegistryCoordinatorOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractSlashingRegistryCoordinatorOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractSlashingRegistryCoordinatorPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractSlashingRegistryCoordinatorPaused, error)

	FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error)
	WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseQuorumBlockNumberUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractSlashingRegistryCoordinatorUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractSlashingRegistryCoordinatorUnpaused, error)
}

// ContractSlashingRegistryCoordinator is an auto generated Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinator struct {
	ContractSlashingRegistryCoordinatorCaller     // Read-only binding to the contract
	ContractSlashingRegistryCoordinatorTransactor // Write-only binding to the contract
	ContractSlashingRegistryCoordinatorFilterer   // Log filterer for contract events
}

// ContractSlashingRegistryCoordinator implements the ContractSlashingRegistryCoordinatorMethods interface.
var _ ContractSlashingRegistryCoordinatorMethods = (*ContractSlashingRegistryCoordinator)(nil)

// ContractSlashingRegistryCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSlashingRegistryCoordinatorCaller implements the ContractSlashingRegistryCoordinatorCalls interface.
var _ ContractSlashingRegistryCoordinatorCalls = (*ContractSlashingRegistryCoordinatorCaller)(nil)

// ContractSlashingRegistryCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSlashingRegistryCoordinatorTransactor implements the ContractSlashingRegistryCoordinatorTransacts interface.
var _ ContractSlashingRegistryCoordinatorTransacts = (*ContractSlashingRegistryCoordinatorTransactor)(nil)

// ContractSlashingRegistryCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSlashingRegistryCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSlashingRegistryCoordinatorFilterer implements the ContractSlashingRegistryCoordinatorFilters interface.
var _ ContractSlashingRegistryCoordinatorFilters = (*ContractSlashingRegistryCoordinatorFilterer)(nil)

// ContractSlashingRegistryCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSlashingRegistryCoordinatorSession struct {
	Contract     *ContractSlashingRegistryCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                        // Call options to use throughout this session
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractSlashingRegistryCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSlashingRegistryCoordinatorCallerSession struct {
	Contract *ContractSlashingRegistryCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                              // Call options to use throughout this session
}

// ContractSlashingRegistryCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSlashingRegistryCoordinatorTransactorSession struct {
	Contract     *ContractSlashingRegistryCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                              // Transaction auth options to use throughout this session
}

// ContractSlashingRegistryCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinatorRaw struct {
	Contract *ContractSlashingRegistryCoordinator // Generic contract binding to access the raw methods on
}

// ContractSlashingRegistryCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinatorCallerRaw struct {
	Contract *ContractSlashingRegistryCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSlashingRegistryCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSlashingRegistryCoordinatorTransactorRaw struct {
	Contract *ContractSlashingRegistryCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSlashingRegistryCoordinator creates a new instance of ContractSlashingRegistryCoordinator, bound to a specific deployed contract.
func NewContractSlashingRegistryCoordinator(address common.Address, backend bind.ContractBackend) (*ContractSlashingRegistryCoordinator, error) {
	contract, err := bindContractSlashingRegistryCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinator{ContractSlashingRegistryCoordinatorCaller: ContractSlashingRegistryCoordinatorCaller{contract: contract}, ContractSlashingRegistryCoordinatorTransactor: ContractSlashingRegistryCoordinatorTransactor{contract: contract}, ContractSlashingRegistryCoordinatorFilterer: ContractSlashingRegistryCoordinatorFilterer{contract: contract}}, nil
}

// NewContractSlashingRegistryCoordinatorCaller creates a new read-only instance of ContractSlashingRegistryCoordinator, bound to a specific deployed contract.
func NewContractSlashingRegistryCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*ContractSlashingRegistryCoordinatorCaller, error) {
	contract, err := bindContractSlashingRegistryCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorCaller{contract: contract}, nil
}

// NewContractSlashingRegistryCoordinatorTransactor creates a new write-only instance of ContractSlashingRegistryCoordinator, bound to a specific deployed contract.
func NewContractSlashingRegistryCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSlashingRegistryCoordinatorTransactor, error) {
	contract, err := bindContractSlashingRegistryCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorTransactor{contract: contract}, nil
}

// NewContractSlashingRegistryCoordinatorFilterer creates a new log filterer instance of ContractSlashingRegistryCoordinator, bound to a specific deployed contract.
func NewContractSlashingRegistryCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSlashingRegistryCoordinatorFilterer, error) {
	contract, err := bindContractSlashingRegistryCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorFilterer{contract: contract}, nil
}

// bindContractSlashingRegistryCoordinator binds a generic wrapper to an already deployed contract.
func bindContractSlashingRegistryCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSlashingRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSlashingRegistryCoordinator.Contract.ContractSlashingRegistryCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.ContractSlashingRegistryCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.ContractSlashingRegistryCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSlashingRegistryCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.contract.Transact(opts, method, params...)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) OPERATORCHURNAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "OPERATOR_CHURN_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// OPERATORCHURNAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0xca0de882.
//
// Solidity: function OPERATOR_CHURN_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) OPERATORCHURNAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.OPERATORCHURNAPPROVALTYPEHASH(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "PUBKEY_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PUBKEYREGISTRATIONTYPEHASH(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// AccountIdentifier is a free data retrieval call binding the contract method 0x0764cb93.
//
// Solidity: function accountIdentifier() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) AccountIdentifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "accountIdentifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountIdentifier is a free data retrieval call binding the contract method 0x0764cb93.
//
// Solidity: function accountIdentifier() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) AccountIdentifier() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.AccountIdentifier(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// AccountIdentifier is a free data retrieval call binding the contract method 0x0764cb93.
//
// Solidity: function accountIdentifier() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) AccountIdentifier() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.AccountIdentifier(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) AllocationManager() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.AllocationManager(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) AllocationManager() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.AllocationManager(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.BlsApkRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.BlsApkRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) CalculateOperatorChurnApprovalDigestHash(opts *bind.CallOpts, registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "calculateOperatorChurnApprovalDigestHash", registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractSlashingRegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// CalculateOperatorChurnApprovalDigestHash is a free data retrieval call binding the contract method 0x84ca5213.
//
// Solidity: function calculateOperatorChurnApprovalDigestHash(address registeringOperator, bytes32 registeringOperatorId, (uint8,address)[] operatorKickParams, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) CalculateOperatorChurnApprovalDigestHash(registeringOperator common.Address, registeringOperatorId [32]byte, operatorKickParams []ISlashingRegistryCoordinatorTypesOperatorKickParam, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CalculateOperatorChurnApprovalDigestHash(&_ContractSlashingRegistryCoordinator.CallOpts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) ChurnApprover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "churnApprover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) ChurnApprover() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.ChurnApprover(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// ChurnApprover is a free data retrieval call binding the contract method 0x054310e6.
//
// Solidity: function churnApprover() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) ChurnApprover() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.ChurnApprover(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) EjectionCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "ejectionCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) EjectionCooldown() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.EjectionCooldown(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// EjectionCooldown is a free data retrieval call binding the contract method 0xa96f783e.
//
// Solidity: function ejectionCooldown() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) EjectionCooldown() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.EjectionCooldown(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Ejector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "ejector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Ejector() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Ejector(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Ejector is a free data retrieval call binding the contract method 0x28f61b31.
//
// Solidity: function ejector() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Ejector() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Ejector(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getCurrentQuorumBitmap", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetOperator(opts *bind.CallOpts, operator common.Address) (ISlashingRegistryCoordinatorTypesOperatorInfo, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getOperator", operator)

	if err != nil {
		return *new(ISlashingRegistryCoordinatorTypesOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlashingRegistryCoordinatorTypesOperatorInfo)).(*ISlashingRegistryCoordinatorTypesOperatorInfo)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetOperator(operator common.Address) (ISlashingRegistryCoordinatorTypesOperatorInfo, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperator(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetOperator(operator common.Address) (ISlashingRegistryCoordinatorTypesOperatorInfo, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperator(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getOperatorFromId", operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorFromId(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorFromId(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorId(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorId(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (ISlashingRegistryCoordinatorTypesOperatorSetParam, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getOperatorSetParams", quorumNumber)

	if err != nil {
		return *new(ISlashingRegistryCoordinatorTypesOperatorSetParam), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlashingRegistryCoordinatorTypesOperatorSetParam)).(*ISlashingRegistryCoordinatorTypesOperatorSetParam)

	return out0, err

}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetOperatorSetParams(quorumNumber uint8) (ISlashingRegistryCoordinatorTypesOperatorSetParam, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractSlashingRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetOperatorSetParams(quorumNumber uint8) (ISlashingRegistryCoordinatorTypesOperatorSetParam, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractSlashingRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getOperatorStatus", operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorStatus(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetOperatorStatus(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapAtBlockNumberByIndex", operatorId, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapHistoryLength", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapIndicesAtBlockNumber", blockNumber, operatorIds)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractSlashingRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractSlashingRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapUpdateByIndex", operatorId, index)

	if err != nil {
		return *new(ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate)).(*ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate)

	return out0, err

}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId, index)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (ISlashingRegistryCoordinatorTypesQuorumBitmapUpdate, error) {
	return _ContractSlashingRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractSlashingRegistryCoordinator.CallOpts, operatorId, index)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) IndexRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "indexRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) IndexRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.IndexRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) IndexRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.IndexRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) IsChurnApproverSaltUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "isChurnApproverSaltUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// IsChurnApproverSaltUsed is a free data retrieval call binding the contract method 0x1478851f.
//
// Solidity: function isChurnApproverSaltUsed(bytes32 ) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) IsChurnApproverSaltUsed(arg0 [32]byte) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.IsChurnApproverSaltUsed(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 quorumNumber) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) IsM2Quorum(opts *bind.CallOpts, quorumNumber uint8) (bool, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "isM2Quorum", quorumNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 quorumNumber) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) IsM2Quorum(quorumNumber uint8) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.IsM2Quorum(&_ContractSlashingRegistryCoordinator.CallOpts, quorumNumber)
}

// IsM2Quorum is a free data retrieval call binding the contract method 0xa4d7871f.
//
// Solidity: function isM2Quorum(uint8 quorumNumber) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) IsM2Quorum(quorumNumber uint8) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.IsM2Quorum(&_ContractSlashingRegistryCoordinator.CallOpts, quorumNumber)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) LastEjectionTimestamp(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "lastEjectionTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.LastEjectionTimestamp(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// LastEjectionTimestamp is a free data retrieval call binding the contract method 0x125e0584.
//
// Solidity: function lastEjectionTimestamp(address ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) LastEjectionTimestamp(arg0 common.Address) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.LastEjectionTimestamp(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// M2QuorumsDisabled is a free data retrieval call binding the contract method 0xb2d8678d.
//
// Solidity: function m2QuorumsDisabled() view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) M2QuorumsDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "m2QuorumsDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// M2QuorumsDisabled is a free data retrieval call binding the contract method 0xb2d8678d.
//
// Solidity: function m2QuorumsDisabled() view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) M2QuorumsDisabled() (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.M2QuorumsDisabled(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// M2QuorumsDisabled is a free data retrieval call binding the contract method 0xb2d8678d.
//
// Solidity: function m2QuorumsDisabled() view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) M2QuorumsDisabled() (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.M2QuorumsDisabled(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) NumRegistries(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "numRegistries")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) NumRegistries() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.NumRegistries(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) NumRegistries() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.NumRegistries(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// OperatorSetsEnabled is a free data retrieval call binding the contract method 0x81f936d2.
//
// Solidity: function operatorSetsEnabled() view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) OperatorSetsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "operatorSetsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSetsEnabled is a free data retrieval call binding the contract method 0x81f936d2.
//
// Solidity: function operatorSetsEnabled() view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) OperatorSetsEnabled() (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.OperatorSetsEnabled(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// OperatorSetsEnabled is a free data retrieval call binding the contract method 0x81f936d2.
//
// Solidity: function operatorSetsEnabled() view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) OperatorSetsEnabled() (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.OperatorSetsEnabled(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Owner() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Owner(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Owner() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Owner(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Paused(index uint8) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Paused(&_ContractSlashingRegistryCoordinator.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Paused(index uint8) (bool, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Paused(&_ContractSlashingRegistryCoordinator.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Paused0() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Paused0(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Paused0() (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Paused0(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) PauserRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PauserRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PauserRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "pubkeyRegistrationMessageHash", operator)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractSlashingRegistryCoordinator.CallOpts, operator)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) QuorumCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "quorumCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) QuorumCount() (uint8, error) {
	return _ContractSlashingRegistryCoordinator.Contract.QuorumCount(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) QuorumCount() (uint8, error) {
	return _ContractSlashingRegistryCoordinator.Contract.QuorumCount(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) QuorumUpdateBlockNumber(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "quorumUpdateBlockNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 ) view returns(uint256)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) QuorumUpdateBlockNumber(arg0 uint8) (*big.Int, error) {
	return _ContractSlashingRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "registries", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Registries(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Registries(&_ContractSlashingRegistryCoordinator.CallOpts, arg0)
}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) SocketRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "socketRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SocketRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SocketRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// SocketRegistry is a free data retrieval call binding the contract method 0xea32afae.
//
// Solidity: function socketRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) SocketRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SocketRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSlashingRegistryCoordinator.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) StakeRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.StakeRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractSlashingRegistryCoordinator.Contract.StakeRegistry(&_ContractSlashingRegistryCoordinator.CallOpts)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) CreateSlashableStakeQuorum(opts *bind.TransactOpts, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "createSlashableStakeQuorum", operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) CreateSlashableStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CreateSlashableStakeQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateSlashableStakeQuorum is a paid mutator transaction binding the contract method 0x3eef3a51.
//
// Solidity: function createSlashableStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams, uint32 lookAheadPeriod) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) CreateSlashableStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams, lookAheadPeriod uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CreateSlashableStakeQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams, lookAheadPeriod)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) CreateTotalDelegatedStakeQuorum(opts *bind.TransactOpts, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "createTotalDelegatedStakeQuorum", operatorSetParams, minimumStake, strategyParams)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) CreateTotalDelegatedStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CreateTotalDelegatedStakeQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// CreateTotalDelegatedStakeQuorum is a paid mutator transaction binding the contract method 0x8281ab75.
//
// Solidity: function createTotalDelegatedStakeQuorum((uint32,uint16,uint16) operatorSetParams, uint96 minimumStake, (address,uint96)[] strategyParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) CreateTotalDelegatedStakeQuorum(operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam, minimumStake *big.Int, strategyParams []IStakeRegistryTypesStrategyParams) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.CreateTotalDelegatedStakeQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorSetParams, minimumStake, strategyParams)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "deregisterOperator", operator, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) DeregisterOperator(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.DeregisterOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x9d8e0c23.
//
// Solidity: function deregisterOperator(address operator, uint32[] operatorSetIds) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) DeregisterOperator(operator common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.DeregisterOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, operatorSetIds)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "ejectOperator", operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.EjectOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.EjectOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, address _accountIdentifier) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _accountIdentifier common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "initialize", _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _accountIdentifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, address _accountIdentifier) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _accountIdentifier common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Initialize(&_ContractSlashingRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _accountIdentifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address _initialOwner, address _churnApprover, address _ejector, uint256 _initialPausedStatus, address _accountIdentifier) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) Initialize(_initialOwner common.Address, _churnApprover common.Address, _ejector common.Address, _initialPausedStatus *big.Int, _accountIdentifier common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Initialize(&_ContractSlashingRegistryCoordinator.TransactOpts, _initialOwner, _churnApprover, _ejector, _initialPausedStatus, _accountIdentifier)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Pause(&_ContractSlashingRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Pause(&_ContractSlashingRegistryCoordinator.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) PauseAll() (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PauseAll(&_ContractSlashingRegistryCoordinator.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.PauseAll(&_ContractSlashingRegistryCoordinator.TransactOpts)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "registerOperator", operator, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) RegisterOperator(operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.RegisterOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xadcf73f7.
//
// Solidity: function registerOperator(address operator, uint32[] operatorSetIds, bytes data) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) RegisterOperator(operator common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.RegisterOperator(&_ContractSlashingRegistryCoordinator.TransactOpts, operator, operatorSetIds, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.RenounceOwnership(&_ContractSlashingRegistryCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.RenounceOwnership(&_ContractSlashingRegistryCoordinator.TransactOpts)
}

// SetAccountIdentifier is a paid mutator transaction binding the contract method 0x143e5915.
//
// Solidity: function setAccountIdentifier(address _accountIdentifier) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) SetAccountIdentifier(opts *bind.TransactOpts, _accountIdentifier common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "setAccountIdentifier", _accountIdentifier)
}

// SetAccountIdentifier is a paid mutator transaction binding the contract method 0x143e5915.
//
// Solidity: function setAccountIdentifier(address _accountIdentifier) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SetAccountIdentifier(_accountIdentifier common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetAccountIdentifier(&_ContractSlashingRegistryCoordinator.TransactOpts, _accountIdentifier)
}

// SetAccountIdentifier is a paid mutator transaction binding the contract method 0x143e5915.
//
// Solidity: function setAccountIdentifier(address _accountIdentifier) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) SetAccountIdentifier(_accountIdentifier common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetAccountIdentifier(&_ContractSlashingRegistryCoordinator.TransactOpts, _accountIdentifier)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) SetChurnApprover(opts *bind.TransactOpts, _churnApprover common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "setChurnApprover", _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetChurnApprover(&_ContractSlashingRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetChurnApprover is a paid mutator transaction binding the contract method 0x29d1e0c3.
//
// Solidity: function setChurnApprover(address _churnApprover) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) SetChurnApprover(_churnApprover common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetChurnApprover(&_ContractSlashingRegistryCoordinator.TransactOpts, _churnApprover)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) SetEjectionCooldown(opts *bind.TransactOpts, _ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "setEjectionCooldown", _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetEjectionCooldown(&_ContractSlashingRegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjectionCooldown is a paid mutator transaction binding the contract method 0x0d3f2134.
//
// Solidity: function setEjectionCooldown(uint256 _ejectionCooldown) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) SetEjectionCooldown(_ejectionCooldown *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetEjectionCooldown(&_ContractSlashingRegistryCoordinator.TransactOpts, _ejectionCooldown)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) SetEjector(opts *bind.TransactOpts, _ejector common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "setEjector", _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetEjector(&_ContractSlashingRegistryCoordinator.TransactOpts, _ejector)
}

// SetEjector is a paid mutator transaction binding the contract method 0x2cdd1e86.
//
// Solidity: function setEjector(address _ejector) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) SetEjector(_ejector common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetEjector(&_ContractSlashingRegistryCoordinator.TransactOpts, _ejector)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "setOperatorSetParams", quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractSlashingRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParams) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.SetOperatorSetParams(&_ContractSlashingRegistryCoordinator.TransactOpts, quorumNumber, operatorSetParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.TransferOwnership(&_ContractSlashingRegistryCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.TransferOwnership(&_ContractSlashingRegistryCoordinator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Unpause(&_ContractSlashingRegistryCoordinator.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.Unpause(&_ContractSlashingRegistryCoordinator.TransactOpts, newPausedStatus)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) UpdateOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "updateOperators", operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateOperators(&_ContractSlashingRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] operators) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) UpdateOperators(operators []common.Address) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateOperators(&_ContractSlashingRegistryCoordinator.TransactOpts, operators)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateOperatorsForQuorum(&_ContractSlashingRegistryCoordinator.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactor) UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.contract.Transact(opts, "updateSocket", socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateSocket(&_ContractSlashingRegistryCoordinator.TransactOpts, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorTransactorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _ContractSlashingRegistryCoordinator.Contract.UpdateSocket(&_ContractSlashingRegistryCoordinator.TransactOpts, socket)
}

// ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator is returned from FilterChurnApproverUpdated and is used to iterate over the raw logs and unpacked data for ChurnApproverUpdated events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator struct {
	Event *ContractSlashingRegistryCoordinatorChurnApproverUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorChurnApproverUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSlashingRegistryCoordinatorChurnApproverUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorChurnApproverUpdated represents a ChurnApproverUpdated event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorChurnApproverUpdated struct {
	PrevChurnApprover common.Address
	NewChurnApprover  common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChurnApproverUpdated is a free log retrieval operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorChurnApproverUpdatedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "ChurnApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchChurnApproverUpdated is a free log subscription operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorChurnApproverUpdated)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChurnApproverUpdated is a log parse operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseChurnApproverUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorChurnApproverUpdated, error) {
	event := new(ContractSlashingRegistryCoordinatorChurnApproverUpdated)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorEjectorUpdatedIterator is returned from FilterEjectorUpdated and is used to iterate over the raw logs and unpacked data for EjectorUpdated events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorEjectorUpdatedIterator struct {
	Event *ContractSlashingRegistryCoordinatorEjectorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSlashingRegistryCoordinatorEjectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorEjectorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSlashingRegistryCoordinatorEjectorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSlashingRegistryCoordinatorEjectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorEjectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorEjectorUpdated represents a EjectorUpdated event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorEjectorUpdated struct {
	PrevEjector common.Address
	NewEjector  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEjectorUpdated is a free log retrieval operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorEjectorUpdatedIterator, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorEjectorUpdatedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "EjectorUpdated", logs: logs, sub: sub}, nil
}

// WatchEjectorUpdated is a free log subscription operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorEjectorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorEjectorUpdated)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEjectorUpdated is a log parse operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseEjectorUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorEjectorUpdated, error) {
	event := new(ContractSlashingRegistryCoordinatorEjectorUpdated)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorInitializedIterator struct {
	Event *ContractSlashingRegistryCoordinatorInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSlashingRegistryCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSlashingRegistryCoordinatorInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSlashingRegistryCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorInitialized represents a Initialized event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractSlashingRegistryCoordinatorInitializedIterator, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorInitializedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorInitialized)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseInitialized(log types.Log) (*ContractSlashingRegistryCoordinatorInitialized, error) {
	event := new(ContractSlashingRegistryCoordinatorInitialized)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator struct {
	Event *ContractSlashingRegistryCoordinatorOperatorDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorOperatorDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSlashingRegistryCoordinatorOperatorDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorOperatorDeregistered represents a OperatorDeregistered event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorDeregistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorOperatorDeregisteredIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorOperatorDeregistered)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseOperatorDeregistered(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorDeregistered, error) {
	event := new(ContractSlashingRegistryCoordinatorOperatorDeregistered)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorRegisteredIterator struct {
	Event *ContractSlashingRegistryCoordinatorOperatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSlashingRegistryCoordinatorOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorOperatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSlashingRegistryCoordinatorOperatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSlashingRegistryCoordinatorOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorOperatorRegistered represents a OperatorRegistered event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorRegistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorOperatorRegisteredIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorOperatorRegistered)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRegistered is a log parse operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseOperatorRegistered(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorRegistered, error) {
	event := new(ContractSlashingRegistryCoordinatorOperatorRegistered)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator is returned from FilterOperatorSetParamsUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetParamsUpdated events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator struct {
	Event *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated represents a OperatorSetParamsUpdated event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated struct {
	QuorumNumber      uint8
	OperatorSetParams ISlashingRegistryCoordinatorTypesOperatorSetParam
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetParamsUpdated is a free log retrieval operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorOperatorSetParamsUpdatedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "OperatorSetParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetParamsUpdated is a free log subscription operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSetParamsUpdated is a log parse operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseOperatorSetParamsUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated, error) {
	event := new(ContractSlashingRegistryCoordinatorOperatorSetParamsUpdated)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator is returned from FilterOperatorSocketUpdate and is used to iterate over the raw logs and unpacked data for OperatorSocketUpdate events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator struct {
	Event *ContractSlashingRegistryCoordinatorOperatorSocketUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorOperatorSocketUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSlashingRegistryCoordinatorOperatorSocketUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorOperatorSocketUpdate represents a OperatorSocketUpdate event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOperatorSocketUpdate struct {
	OperatorId [32]byte
	Socket     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketUpdate is a free log retrieval operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterOperatorSocketUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorOperatorSocketUpdateIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "OperatorSocketUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketUpdate is a free log subscription operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOperatorSocketUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "OperatorSocketUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorOperatorSocketUpdate)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSocketUpdate is a log parse operation binding the contract event 0xec2963ab21c1e50e1e582aa542af2e4bf7bf38e6e1403c27b42e1c5d6e621eaa.
//
// Solidity: event OperatorSocketUpdate(bytes32 indexed operatorId, string socket)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseOperatorSocketUpdate(log types.Log) (*ContractSlashingRegistryCoordinatorOperatorSocketUpdate, error) {
	event := new(ContractSlashingRegistryCoordinatorOperatorSocketUpdate)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOwnershipTransferredIterator struct {
	Event *ContractSlashingRegistryCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSlashingRegistryCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSlashingRegistryCoordinatorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSlashingRegistryCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractSlashingRegistryCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorOwnershipTransferredIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorOwnershipTransferred)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*ContractSlashingRegistryCoordinatorOwnershipTransferred, error) {
	event := new(ContractSlashingRegistryCoordinatorOwnershipTransferred)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorPausedIterator struct {
	Event *ContractSlashingRegistryCoordinatorPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSlashingRegistryCoordinatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSlashingRegistryCoordinatorPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSlashingRegistryCoordinatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorPaused represents a Paused event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractSlashingRegistryCoordinatorPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorPausedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorPaused)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParsePaused(log types.Log) (*ContractSlashingRegistryCoordinatorPaused, error) {
	event := new(ContractSlashingRegistryCoordinatorPaused)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator is returned from FilterQuorumBlockNumberUpdated and is used to iterate over the raw logs and unpacked data for QuorumBlockNumberUpdated events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator struct {
	Event *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated represents a QuorumBlockNumberUpdated event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated struct {
	QuorumNumber uint8
	Blocknumber  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumBlockNumberUpdated is a free log retrieval operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdatedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "QuorumBlockNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumBlockNumberUpdated is a free log subscription operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseQuorumBlockNumberUpdated is a log parse operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseQuorumBlockNumberUpdated(log types.Log) (*ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated, error) {
	event := new(ContractSlashingRegistryCoordinatorQuorumBlockNumberUpdated)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSlashingRegistryCoordinatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorUnpausedIterator struct {
	Event *ContractSlashingRegistryCoordinatorUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSlashingRegistryCoordinatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSlashingRegistryCoordinatorUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSlashingRegistryCoordinatorUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSlashingRegistryCoordinatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSlashingRegistryCoordinatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSlashingRegistryCoordinatorUnpaused represents a Unpaused event raised by the ContractSlashingRegistryCoordinator contract.
type ContractSlashingRegistryCoordinatorUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractSlashingRegistryCoordinatorUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSlashingRegistryCoordinatorUnpausedIterator{contract: _ContractSlashingRegistryCoordinator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractSlashingRegistryCoordinatorUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSlashingRegistryCoordinator.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSlashingRegistryCoordinatorUnpaused)
				if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSlashingRegistryCoordinator *ContractSlashingRegistryCoordinatorFilterer) ParseUnpaused(log types.Log) (*ContractSlashingRegistryCoordinatorUnpaused, error) {
	event := new(ContractSlashingRegistryCoordinatorUnpaused)
	if err := _ContractSlashingRegistryCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
