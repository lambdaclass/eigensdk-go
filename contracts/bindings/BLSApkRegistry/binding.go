// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractBLSApkRegistry

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

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSApkRegistryTypesApkUpdate is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryTypesApkUpdate struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}

// IBLSApkRegistryTypesPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryTypesPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// ContractBLSApkRegistryMetaData contains all meta data concerning the ContractBLSApkRegistry contract.
var ContractBLSApkRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_slashingRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"apkHistory\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"apkHash\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentApk\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApk\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkHashAtBlockNumberAndIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkHistoryLength\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistryTypes.ApkUpdate\",\"components\":[{\"name\":\"apkHash\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromPubkeyHash\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredPubkey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorToPubkey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorToPubkeyHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyHashToOperator\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerBLSPublicKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistryTypes.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"pubkeyRegistrationMessageHash\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewPubkeyRegistration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BLSPubkeyAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockNumberNotLatest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockNumberTooRecent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECMulFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECPairingFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSSignatureOrPrivateKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRegistryCoordinatorOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroPubKey\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051611b00380380611b0083398101604081905261002e91610108565b6001600160a01b0381166080528061004461004b565b5050610135565b5f54610100900460ff16156100b65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610106575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610118575f5ffd5b81516001600160a01b038116811461012e575f5ffd5b9392505050565b6080516119ac6101545f395f81816103030152610cdc01526119ac5ff3fe608060405234801561000f575f5ffd5b5060043610610110575f3560e01c80636d14a9871161009e578063bf79ce581161006e578063bf79ce58146103bf578063d5254a8c146103d2578063de29fac0146103f2578063e8bb9ae614610411578063f4e24fe514610439575f5ffd5b80636d14a987146102fe5780637916cea6146103255780637ff81a8714610366578063a3db80e214610399575f5ffd5b80633fb27952116100e45780633fb27952146101d657806347b314e8146101e95780635f61a88414610229578063605747d51461028357806368bccaac146102d1575f5ffd5b8062a1f4cb1461011457806313542a4e1461015457806326d941f21461018a578063377ed99d1461019f575b5f5ffd5b61013a61012236600461146c565b60036020525f90815260409020805460019091015482565b604080519283526020830191909152015b60405180910390f35b61017c61016236600461146c565b6001600160a01b03165f9081526001602052604090205490565b60405190815260200161014b565b61019d61019836600461149c565b61044c565b005b6101c16101ad36600461149c565b60ff165f9081526004602052604090205490565b60405163ffffffff909116815260200161014b565b61019d6101e4366004611523565b61050c565b6102116101f73660046115c8565b5f908152600260205260409020546001600160a01b031690565b6040516001600160a01b03909116815260200161014b565b61027661023736600461149c565b604080518082019091525f80825260208201525060ff165f90815260056020908152604091829020825180840190935280548352600101549082015290565b60405161014b91906115df565b6102966102913660046115f6565b610588565b60408051825167ffffffffffffffff1916815260208084015163ffffffff90811691830191909152928201519092169082015260600161014b565b6102e46102df36600461161e565b610619565b60405167ffffffffffffffff19909116815260200161014b565b6102117f000000000000000000000000000000000000000000000000000000000000000081565b6103386103333660046115f6565b6106fc565b6040805167ffffffffffffffff19909416845263ffffffff928316602085015291169082015260600161014b565b61037961037436600461146c565b610743565b60408051835181526020938401519381019390935282015260600161014b565b61013a6103a736600461149c565b60056020525f90815260409020805460019091015482565b61017c6103cd366004611662565b6107b9565b6103e56103e03660046116ba565b610a5f565b60405161014b919061172c565b61017c61040036600461146c565b60016020525f908152604090205481565b61021161041f3660046115c8565b60026020525f90815260409020546001600160a01b031681565b61019d610447366004611523565b610c6a565b610454610cd1565b60ff81165f9081526004602052604090205415610484576040516310cda51760e21b815260040160405180910390fd5b60ff165f908152600460209081526040808320815160608101835284815263ffffffff4381168286019081528285018781528454600181018655948852959096209151919092018054955194518316600160e01b026001600160e01b0395909316600160c01b026001600160e01b03199096169190931c179390931791909116919091179055565b610514610cd1565b5f61051e83610743565b50905061052b8282610d1c565b7f73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e8361056b856001600160a01b03165f9081526001602052604090205490565b8460405161057b93929190611774565b60405180910390a1505050565b604080516060810182525f808252602080830182905282840182905260ff8616825260049052919091208054839081106105c4576105c46117bf565b5f91825260209182902060408051606081018252919092015467ffffffffffffffff1981841b16825263ffffffff600160c01b8204811694830194909452600160e01b90049092169082015290505b92915050565b60ff83165f90815260046020526040812080548291908490811061063f5761063f6117bf565b5f91825260209182902060408051606081018252919092015467ffffffffffffffff1981841b16825263ffffffff600160c01b82048116948301859052600160e01b9091048116928201929092529250851610156106b057604051633d22884160e01b815260040160405180910390fd5b604081015163ffffffff1615806106d65750806040015163ffffffff168463ffffffff16105b6106f357604051636fe02d4b60e01b815260040160405180910390fd5b51949350505050565b6004602052815f5260405f208181548110610715575f80fd5b5f91825260209091200154604081901b925063ffffffff600160c01b820481169250600160e01b9091041683565b604080518082019091525f80825260208201526001600160a01b0382165f818152600360209081526040808320815180830183528154815260019182015481850152948452909152812054909190806107af576040516325ec6c1f60e01b815260040160405180910390fd5b9094909350915050565b5f6107c2610cd1565b5f6107ee6107d8368690038601604087016117d3565b80515f9081526020918201519091526040902090565b90507fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5810361083057604051630cc7509160e01b815260040160405180910390fd5b6001600160a01b0385165f9081526001602052604081205414610866576040516342ee68b560e01b815260040160405180910390fd5b5f818152600260205260409020546001600160a01b03161561089b57604051634c334c9760e11b815260040160405180910390fd5b604080515f917f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001916108f3918835916020808b0135928b01359160608c01359160808d019160c08e01918d35918e8201359101611804565b604051602081830303815290604052805190602001205f1c6109159190611846565b90506109ae61094e61093983610933368a90038a0160408b016117d3565b90610f03565b610948368990038901896117d3565b90610f73565b610956610fe7565b610997610988856109336040805180820182525f80825260209182015281518083019092526001825260029082015290565b610948368a90038a018a6117d3565b6109a9368a90038a0160808b016118a7565b6110a7565b6109cb5760405163a72d026360e01b815260040160405180910390fd5b6001600160a01b0386165f8181526003602090815260408083208982018035825560608b013560019283015590835281842087905586845260029092529182902080546001600160a01b0319168417905590517fe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba382804191610a4e9160808a01906118e6565b60405180910390a250949350505050565b60605f8367ffffffffffffffff811115610a7b57610a7b6114b5565b604051908082528060200260200182016040528015610aa4578160200160208202803683370190505b5090505f5b84811015610c61575f868683818110610ac457610ac46117bf565b919091013560f81c5f818152600460205260409020549092509050801580610b24575060ff82165f9081526004602052604081208054909190610b0957610b096117bf565b5f91825260209091200154600160c01b900463ffffffff1686105b15610bb55760405162461bcd60e51b815260206004820152605160248201527f424c5341706b52656769737472792e67657441706b496e64696365734174426c60448201527f6f636b4e756d6265723a20626c6f636b4e756d626572206973206265666f7265606482015270207468652066697273742075706461746560781b608482015260a40160405180910390fd5b805b8015610c565760ff83165f9081526004602052604090208790610bdb600184611924565b81548110610beb57610beb6117bf565b5f91825260209091200154600160c01b900463ffffffff1611610c4457610c13600182611924565b858581518110610c2557610c256117bf565b602002602001019063ffffffff16908163ffffffff1681525050610c56565b80610c4e81611937565b915050610bb7565b505050600101610aa9565b50949350505050565b610c72610cd1565b5f610c7c83610743565b509050610c9182610c8c836112de565b610d1c565b7ff843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e8361056b856001600160a01b03165f9081526001602052604090205490565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610d1a57604051637070f3b160e11b815260040160405180910390fd5b565b604080518082019091525f80825260208201525f5b8351811015610efd575f848281518110610d4d57610d4d6117bf565b0160209081015160f81c5f8181526004909252604082205490925090819003610d8957604051637310cff560e11b815260040160405180910390fd5b60ff82165f908152600560209081526040918290208251808401909352805483526001015490820152610dbc9086610f73565b60ff83165f818152600560209081526040808320855180825586840180516001938401559085525183528184209484526004909252822093975091929091610e049085611924565b81548110610e1457610e146117bf565b5f918252602090912001805490915063ffffffff438116600160c01b9092041603610e525780546001600160c01b031916604083901c178155610eed565b805463ffffffff438116600160e01b8181026001600160e01b0394851617855560ff88165f908152600460209081526040808320815160608101835267ffffffffffffffff198b16815280840196875280830185815282546001810184559286529390942093519301805495519251871690940291909516600160c01b026001600160e01b0319949094169190941c17919091179092161790555b505060019092019150610d319050565b50505050565b604080518082019091525f8082526020820152610f1e61139a565b835181526020808501519082015260408082018490525f908360608460076107d05a03fa90508080610f4c57fe5b5080610f6b57604051632319df1960e11b815260040160405180910390fd5b505092915050565b604080518082019091525f8082526020820152610f8e6113b8565b835181526020808501518183015283516040808401919091529084015160608301525f908360808460066107d05a03fa90508080610fc857fe5b5080610f6b5760405163d4b68fd760e01b815260040160405180910390fd5b610fef6113d6565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820182528581526020808201859052825180840190935285835282018390525f916110d56113fb565b5f5b600281101561128c575f6110ec82600661194c565b9050848260028110611100576111006117bf565b60200201515183611111835f611963565b600c8110611121576111216117bf565b6020020152848260028110611138576111386117bf565b6020020151602001518382600161114f9190611963565b600c811061115f5761115f6117bf565b6020020152838260028110611176576111766117bf565b6020020151515183611189836002611963565b600c8110611199576111996117bf565b60200201528382600281106111b0576111b06117bf565b60200201515160016020020151836111c9836003611963565b600c81106111d9576111d96117bf565b60200201528382600281106111f0576111f06117bf565b6020020151602001515f6002811061120a5761120a6117bf565b60200201518361121b836004611963565b600c811061122b5761122b6117bf565b6020020152838260028110611242576112426117bf565b60200201516020015160016002811061125d5761125d6117bf565b60200201518361126e836005611963565b600c811061127e5761127e6117bf565b6020020152506001016110d7565b5061129561141a565b5f6020826101808560086107d05a03fa905080806112af57fe5b50806112ce576040516324ccc79360e21b815260040160405180910390fd5b5051151598975050505050505050565b604080518082019091525f8082526020820152815115801561130257506020820151155b1561131f575050604080518082019091525f808252602082015290565b6040518060400160405280835f015181526020017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4784602001516113639190611846565b61138d907f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47611924565b905292915050565b919050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806113e9611438565b81526020016113f6611438565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b80356001600160a01b0381168114611395575f5ffd5b5f6020828403121561147c575f5ffd5b61148582611456565b9392505050565b803560ff81168114611395575f5ffd5b5f602082840312156114ac575f5ffd5b6114858261148c565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff811182821017156114ec576114ec6114b5565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561151b5761151b6114b5565b604052919050565b5f5f60408385031215611534575f5ffd5b61153d83611456565b9150602083013567ffffffffffffffff811115611558575f5ffd5b8301601f81018513611568575f5ffd5b803567ffffffffffffffff811115611582576115826114b5565b611595601f8201601f19166020016114f2565b8181528660208385010111156115a9575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f602082840312156115d8575f5ffd5b5035919050565b815181526020808301519082015260408101610613565b5f5f60408385031215611607575f5ffd5b6116108361148c565b946020939093013593505050565b5f5f5f60608486031215611630575f5ffd5b6116398461148c565b9250602084013563ffffffff81168114611651575f5ffd5b929592945050506040919091013590565b5f5f5f838503610160811215611676575f5ffd5b61167f85611456565b9350610100601f1982011215611693575f5ffd5b602085019250604061011f19820112156116ab575f5ffd5b50610120840190509250925092565b5f5f5f604084860312156116cc575f5ffd5b833567ffffffffffffffff8111156116e2575f5ffd5b8401601f810186136116f2575f5ffd5b803567ffffffffffffffff811115611708575f5ffd5b866020828401011115611719575f5ffd5b6020918201979096509401359392505050565b602080825282518282018190525f918401906040840190835b8181101561176957835163ffffffff16835260209384019390920191600101611745565b509095945050505050565b60018060a01b0384168152826020820152606060408201525f82518060608401528060208501608085015e5f608082850101526080601f19601f830116840101915050949350505050565b634e487b7160e01b5f52603260045260245ffd5b5f60408284031280156117e4575f5ffd5b506117ed6114c9565b823581526020928301359281019290925250919050565b888152876020820152866040820152856060820152604085608083013760408460c0830137610100810192909252610120820152610140019695505050505050565b5f8261186057634e487b7160e01b5f52601260045260245ffd5b500690565b5f82601f830112611874575f5ffd5b61187c6114c9565b80604084018581111561188d575f5ffd5b845b8181101561176957803584526020938401930161188f565b5f60808284031280156118b8575f5ffd5b506118c16114c9565b6118cb8484611865565b81526118da8460408501611865565b60208201529392505050565b823581526020808401359082015260c0810160408381840137604080840160808401379392505050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561061357610613611910565b5f8161194557611945611910565b505f190190565b808202811582820484141761061357610613611910565b808201808211156106135761061361191056fea26469706673582212201a108f0f99b10263412bf9af9892644cec0ec35c9b4928809983016b0a83858764736f6c634300081b0033",
}

// ContractBLSApkRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractBLSApkRegistryMetaData.ABI instead.
var ContractBLSApkRegistryABI = ContractBLSApkRegistryMetaData.ABI

// ContractBLSApkRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractBLSApkRegistryMetaData.Bin instead.
var ContractBLSApkRegistryBin = ContractBLSApkRegistryMetaData.Bin

// DeployContractBLSApkRegistry deploys a new Ethereum contract, binding an instance of ContractBLSApkRegistry to it.
func DeployContractBLSApkRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _slashingRegistryCoordinator common.Address) (common.Address, *types.Transaction, *ContractBLSApkRegistry, error) {
	parsed, err := ContractBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBLSApkRegistryBin), backend, _slashingRegistryCoordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractBLSApkRegistry{ContractBLSApkRegistryCaller: ContractBLSApkRegistryCaller{contract: contract}, ContractBLSApkRegistryTransactor: ContractBLSApkRegistryTransactor{contract: contract}, ContractBLSApkRegistryFilterer: ContractBLSApkRegistryFilterer{contract: contract}}, nil
}

// ContractBLSApkRegistryMethods is an auto generated interface around an Ethereum contract.
type ContractBLSApkRegistryMethods interface {
	ContractBLSApkRegistryCalls
	ContractBLSApkRegistryTransacts
	ContractBLSApkRegistryFilters
}

// ContractBLSApkRegistryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractBLSApkRegistryCalls interface {
	ApkHistory(opts *bind.CallOpts, quorumNumber uint8, arg1 *big.Int) (struct {
		ApkHash               [24]byte
		UpdateBlockNumber     uint32
		NextUpdateBlockNumber uint32
	}, error)

	CurrentApk(opts *bind.CallOpts, quorumNumber uint8) (struct {
		X *big.Int
		Y *big.Int
	}, error)

	GetApk(opts *bind.CallOpts, quorumNumber uint8) (BN254G1Point, error)

	GetApkHashAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error)

	GetApkHistoryLength(opts *bind.CallOpts, quorumNumber uint8) (uint32, error)

	GetApkIndicesAtBlockNumber(opts *bind.CallOpts, quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error)

	GetApkUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IBLSApkRegistryTypesApkUpdate, error)

	GetOperatorFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error)

	GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address) (BN254G1Point, [32]byte, error)

	OperatorToPubkey(opts *bind.CallOpts, operator common.Address) (struct {
		X *big.Int
		Y *big.Int
	}, error)

	OperatorToPubkeyHash(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	PubkeyHashToOperator(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error)

	RegistryCoordinator(opts *bind.CallOpts) (common.Address, error)
}

// ContractBLSApkRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractBLSApkRegistryTransacts interface {
	DeregisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)

	InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error)

	RegisterBLSPublicKey(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)
}

// ContractBLSApkRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractBLSApkRegistryFilters interface {
	FilterInitialized(opts *bind.FilterOpts) (*ContractBLSApkRegistryInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractBLSApkRegistryInitialized, error)

	FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractBLSApkRegistryNewPubkeyRegistrationIterator, error)
	WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryNewPubkeyRegistration, operator []common.Address) (event.Subscription, error)
	ParseNewPubkeyRegistration(log types.Log) (*ContractBLSApkRegistryNewPubkeyRegistration, error)

	FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorAddedToQuorumsIterator, error)
	WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorAddedToQuorums) (event.Subscription, error)
	ParseOperatorAddedToQuorums(log types.Log) (*ContractBLSApkRegistryOperatorAddedToQuorums, error)

	FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator, error)
	WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorRemovedFromQuorums) (event.Subscription, error)
	ParseOperatorRemovedFromQuorums(log types.Log) (*ContractBLSApkRegistryOperatorRemovedFromQuorums, error)
}

// ContractBLSApkRegistry is an auto generated Go binding around an Ethereum contract.
type ContractBLSApkRegistry struct {
	ContractBLSApkRegistryCaller     // Read-only binding to the contract
	ContractBLSApkRegistryTransactor // Write-only binding to the contract
	ContractBLSApkRegistryFilterer   // Log filterer for contract events
}

// ContractBLSApkRegistry implements the ContractBLSApkRegistryMethods interface.
var _ ContractBLSApkRegistryMethods = (*ContractBLSApkRegistry)(nil)

// ContractBLSApkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSApkRegistryCaller implements the ContractBLSApkRegistryCalls interface.
var _ ContractBLSApkRegistryCalls = (*ContractBLSApkRegistryCaller)(nil)

// ContractBLSApkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSApkRegistryTransactor implements the ContractBLSApkRegistryTransacts interface.
var _ ContractBLSApkRegistryTransacts = (*ContractBLSApkRegistryTransactor)(nil)

// ContractBLSApkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractBLSApkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSApkRegistryFilterer implements the ContractBLSApkRegistryFilters interface.
var _ ContractBLSApkRegistryFilters = (*ContractBLSApkRegistryFilterer)(nil)

// ContractBLSApkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractBLSApkRegistrySession struct {
	Contract     *ContractBLSApkRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractBLSApkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractBLSApkRegistryCallerSession struct {
	Contract *ContractBLSApkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ContractBLSApkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractBLSApkRegistryTransactorSession struct {
	Contract     *ContractBLSApkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ContractBLSApkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractBLSApkRegistryRaw struct {
	Contract *ContractBLSApkRegistry // Generic contract binding to access the raw methods on
}

// ContractBLSApkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryCallerRaw struct {
	Contract *ContractBLSApkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractBLSApkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryTransactorRaw struct {
	Contract *ContractBLSApkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractBLSApkRegistry creates a new instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistry(address common.Address, backend bind.ContractBackend) (*ContractBLSApkRegistry, error) {
	contract, err := bindContractBLSApkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistry{ContractBLSApkRegistryCaller: ContractBLSApkRegistryCaller{contract: contract}, ContractBLSApkRegistryTransactor: ContractBLSApkRegistryTransactor{contract: contract}, ContractBLSApkRegistryFilterer: ContractBLSApkRegistryFilterer{contract: contract}}, nil
}

// NewContractBLSApkRegistryCaller creates a new read-only instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractBLSApkRegistryCaller, error) {
	contract, err := bindContractBLSApkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryCaller{contract: contract}, nil
}

// NewContractBLSApkRegistryTransactor creates a new write-only instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractBLSApkRegistryTransactor, error) {
	contract, err := bindContractBLSApkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryTransactor{contract: contract}, nil
}

// NewContractBLSApkRegistryFilterer creates a new log filterer instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractBLSApkRegistryFilterer, error) {
	contract, err := bindContractBLSApkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryFilterer{contract: contract}, nil
}

// bindContractBLSApkRegistry binds a generic wrapper to an already deployed contract.
func bindContractBLSApkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBLSApkRegistry.Contract.ContractBLSApkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.ContractBLSApkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.ContractBLSApkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBLSApkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.contract.Transact(opts, method, params...)
}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 quorumNumber, uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) ApkHistory(opts *bind.CallOpts, quorumNumber uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "apkHistory", quorumNumber, arg1)

	outstruct := new(struct {
		ApkHash               [24]byte
		UpdateBlockNumber     uint32
		NextUpdateBlockNumber uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ApkHash = *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)
	outstruct.UpdateBlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.NextUpdateBlockNumber = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 quorumNumber, uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) ApkHistory(quorumNumber uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _ContractBLSApkRegistry.Contract.ApkHistory(&_ContractBLSApkRegistry.CallOpts, quorumNumber, arg1)
}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 quorumNumber, uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) ApkHistory(quorumNumber uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _ContractBLSApkRegistry.Contract.ApkHistory(&_ContractBLSApkRegistry.CallOpts, quorumNumber, arg1)
}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 quorumNumber) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) CurrentApk(opts *bind.CallOpts, quorumNumber uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "currentApk", quorumNumber)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 quorumNumber) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) CurrentApk(quorumNumber uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.CurrentApk(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 quorumNumber) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) CurrentApk(quorumNumber uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.CurrentApk(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApk(opts *bind.CallOpts, quorumNumber uint8) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApk", quorumNumber)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _ContractBLSApkRegistry.Contract.GetApk(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _ContractBLSApkRegistry.Contract.GetApk(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkHashAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkHashAtBlockNumberAndIndex", quorumNumber, blockNumber, index)

	if err != nil {
		return *new([24]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)

	return out0, err

}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkHistoryLength(opts *bind.CallOpts, quorumNumber uint8) (uint32, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkHistoryLength", quorumNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHistoryLength(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHistoryLength(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkIndicesAtBlockNumber(opts *bind.CallOpts, quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkIndicesAtBlockNumber", quorumNumbers, blockNumber)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_ContractBLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_ContractBLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IBLSApkRegistryTypesApkUpdate, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkUpdateAtIndex", quorumNumber, index)

	if err != nil {
		return *new(IBLSApkRegistryTypesApkUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSApkRegistryTypesApkUpdate)).(*IBLSApkRegistryTypesApkUpdate)

	return out0, err

}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryTypesApkUpdate, error) {
	return _ContractBLSApkRegistry.Contract.GetApkUpdateAtIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryTypesApkUpdate, error) {
	return _ContractBLSApkRegistry.Contract.GetApkUpdateAtIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetOperatorFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getOperatorFromPubkeyHash", pubkeyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_ContractBLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_ContractBLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorId(&_ContractBLSApkRegistry.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorId(&_ContractBLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address) (BN254G1Point, [32]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getRegisteredPubkey", operator)

	if err != nil {
		return *new(BN254G1Point), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetRegisteredPubkey(&_ContractBLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetRegisteredPubkey(&_ContractBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address operator) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) OperatorToPubkey(opts *bind.CallOpts, operator common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "operatorToPubkey", operator)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address operator) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) OperatorToPubkey(operator common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkey(&_ContractBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address operator) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) OperatorToPubkey(operator common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkey(&_ContractBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) OperatorToPubkeyHash(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "operatorToPubkeyHash", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) OperatorToPubkeyHash(operator common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkeyHash(&_ContractBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) OperatorToPubkeyHash(operator common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkeyHash(&_ContractBLSApkRegistry.CallOpts, operator)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address operator)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) PubkeyHashToOperator(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "pubkeyHashToOperator", pubkeyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address operator)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) PubkeyHashToOperator(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.PubkeyHashToOperator(&_ContractBLSApkRegistry.CallOpts, pubkeyHash)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address operator)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) PubkeyHashToOperator(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.PubkeyHashToOperator(&_ContractBLSApkRegistry.CallOpts, pubkeyHash)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.RegistryCoordinator(&_ContractBLSApkRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.RegistryCoordinator(&_ContractBLSApkRegistry.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "deregisterOperator", operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.DeregisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.DeregisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "initializeQuorum", quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.InitializeQuorum(&_ContractBLSApkRegistry.TransactOpts, quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.InitializeQuorum(&_ContractBLSApkRegistry.TransactOpts, quorumNumber)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "registerBLSPublicKey", operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterBLSPublicKey(&_ContractBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterBLSPublicKey(&_ContractBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "registerOperator", operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// ContractBLSApkRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryInitializedIterator struct {
	Event *ContractBLSApkRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryInitialized)
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
		it.Event = new(ContractBLSApkRegistryInitialized)
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
func (it *ContractBLSApkRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryInitialized represents a Initialized event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractBLSApkRegistryInitializedIterator, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryInitializedIterator{contract: _ContractBLSApkRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryInitialized)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseInitialized(log types.Log) (*ContractBLSApkRegistryInitialized, error) {
	event := new(ContractBLSApkRegistryInitialized)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBLSApkRegistryNewPubkeyRegistrationIterator is returned from FilterNewPubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewPubkeyRegistration events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryNewPubkeyRegistrationIterator struct {
	Event *ContractBLSApkRegistryNewPubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryNewPubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryNewPubkeyRegistration)
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
		it.Event = new(ContractBLSApkRegistryNewPubkeyRegistration)
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
func (it *ContractBLSApkRegistryNewPubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryNewPubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryNewPubkeyRegistration represents a NewPubkeyRegistration event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryNewPubkeyRegistration struct {
	Operator common.Address
	PubkeyG1 BN254G1Point
	PubkeyG2 BN254G2Point
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPubkeyRegistration is a free log retrieval operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractBLSApkRegistryNewPubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryNewPubkeyRegistrationIterator{contract: _ContractBLSApkRegistry.contract, event: "NewPubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewPubkeyRegistration is a free log subscription operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryNewPubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryNewPubkeyRegistration)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
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

// ParseNewPubkeyRegistration is a log parse operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseNewPubkeyRegistration(log types.Log) (*ContractBLSApkRegistryNewPubkeyRegistration, error) {
	event := new(ContractBLSApkRegistryNewPubkeyRegistration)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBLSApkRegistryOperatorAddedToQuorumsIterator is returned from FilterOperatorAddedToQuorums and is used to iterate over the raw logs and unpacked data for OperatorAddedToQuorums events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorAddedToQuorumsIterator struct {
	Event *ContractBLSApkRegistryOperatorAddedToQuorums // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryOperatorAddedToQuorums)
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
		it.Event = new(ContractBLSApkRegistryOperatorAddedToQuorums)
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
func (it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryOperatorAddedToQuorums represents a OperatorAddedToQuorums event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorAddedToQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToQuorums is a free log retrieval operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorAddedToQuorumsIterator, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryOperatorAddedToQuorumsIterator{contract: _ContractBLSApkRegistry.contract, event: "OperatorAddedToQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToQuorums is a free log subscription operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorAddedToQuorums) (event.Subscription, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryOperatorAddedToQuorums)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
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

// ParseOperatorAddedToQuorums is a log parse operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseOperatorAddedToQuorums(log types.Log) (*ContractBLSApkRegistryOperatorAddedToQuorums, error) {
	event := new(ContractBLSApkRegistryOperatorAddedToQuorums)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator is returned from FilterOperatorRemovedFromQuorums and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromQuorums events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator struct {
	Event *ContractBLSApkRegistryOperatorRemovedFromQuorums // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
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
		it.Event = new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
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
func (it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryOperatorRemovedFromQuorums represents a OperatorRemovedFromQuorums event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorRemovedFromQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromQuorums is a free log retrieval operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator{contract: _ContractBLSApkRegistry.contract, event: "OperatorRemovedFromQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromQuorums is a free log subscription operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorRemovedFromQuorums) (event.Subscription, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
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

// ParseOperatorRemovedFromQuorums is a log parse operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseOperatorRemovedFromQuorums(log types.Log) (*ContractBLSApkRegistryOperatorRemovedFromQuorums, error) {
	event := new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
