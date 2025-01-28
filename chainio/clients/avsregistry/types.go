package avsregistry

import (
	"crypto/ecdsa"
	"math/big"

	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type TxOptions struct {
	Options *bind.TransactOpts
}

type QuorumCountRequest struct {
	BlockNumber *big.Int
}

type QuorumCountResponse struct {
	QuorumCount uint8
}

type SignaturesIndicesRequest struct {
	BlockNumber          *big.Int
	ReferenceBlockNumber uint32
	QuorumNumbers        types.QuorumNums
	NonSignerOperatorIds []types.OperatorId
}

type SignaturesIndicesResponse struct {
	SignaturesIndices opstateretriever.OperatorStateRetrieverCheckSignaturesIndices
}

type OperatorIdRequest struct {
	BlockNumber     *big.Int
	OperatorAddress gethcommon.Address
}

type OperatorIdResponse struct {
	OperatorId [32]byte
}

type OperatorFromIdRequest struct {
	BlockNumber *big.Int
	OperatorId  types.OperatorId
}

type OperatorFromIdResponse struct {
	OperatorAddress gethcommon.Address
}

type RegistrationDetailRequest struct {
	BlockNumber     *big.Int
	OperatorAddress gethcommon.Address
}

type RegistrationDetailResponse struct {
	Quorums []bool
}

type OperatorRegistrationRequest struct {
	BlockNumber     *big.Int
	OperatorAddress gethcommon.Address
}

type OperatorRegistrationResponse struct {
	IsRegistered bool
}

type OperatorQueryRequest struct {
	BlockNumber *big.Int
	StartBlock  uint64
	StopBlock   uint64
	BlockRange  uint64
}

type OperatorPubKeysResponse struct {
	OperatorAddresses []gethcommon.Address
	OperatorsPubkeys  []types.OperatorPubkeys
}

type OperatorSocketsResponse struct {
	Sockets map[types.OperatorId]types.Socket
}

type OperatorsStakeInQuorumAtCurrentBlockRequest struct {
	BlockNumber   *big.Int
	QuorumNumbers types.QuorumNums
}

type OperatorsStakeInQuorumAtBlockRequest struct {
	BlockNumber           *big.Int
	HistoricalBlockNumber uint32
	QuorumNumbers         types.QuorumNums
}

type OperatorsStakeInQuorumResponse struct {
	OperatorsStakeInQuorum [][]opstateretriever.OperatorStateRetrieverOperator
}

type OperatorAddrsInQuorumsAtCurrentBlockRequest struct {
	BlockNumber   *big.Int
	QuorumNumbers types.QuorumNums
}

type OperatorAddrsInQuorumsAtCurrentBlockResponse struct {
	OperatorAddrsInQuorums [][]gethcommon.Address
}

type OperatorsStakeInQuorumsByOperatorAtBlockRequest struct {
	BlockNumber           *big.Int
	HistoricalBlockNumber uint32
	OperatorId            types.OperatorId
}

type OperatorsStakeInQuorumsByOperatorAtCurrentBlockRequest struct {
	BlockNumber *big.Int
	OperatorId  types.OperatorId
}

type OperatorsStakeInQuorumsByOperatorResponse struct {
	QuorumNumbers           types.QuorumNums
	OperatorsStakesInQuorum [][]opstateretriever.OperatorStateRetrieverOperator
}

type OperatorQuorumStakeAtCurrentBlockRequest struct {
	BlockNumber *big.Int
	OperatorId  types.OperatorId
}

type OperatorStakeInQuorumsOfOperatorResponse struct {
	QuorumStakes map[types.QuorumNum]types.StakeAmount
}

type OperatorRegisterRequest struct {
	OperatorEcdsaPrivateKey *ecdsa.PrivateKey
	BlsKeyPair              *bls.KeyPair
	QuorumNumbers           types.QuorumNums
	Socket                  string
	WaitForReceipt          bool
}

type StakesOfEntireOperatorSetForQuorumsRequest struct {
	OperatorsPerQuorum [][]gethcommon.Address
	QuorumNumbers      types.QuorumNums
	WaitForReceipt     bool
}

type StakesOfOperatorSubsetForAllQuorumsRequest struct {
	OperatorsAddresses []gethcommon.Address
	WaitForReceipt     bool
}

type OperatorDeregisterRequest struct {
	QuorumNumbers  types.QuorumNums
	Pubkey         regcoord.BN254G1Point
	WaitForReceipt bool
}

type OperatorDeregisterOperatorSetsRequest struct {
	OperatorSetIds types.OperatorSetIds
	Operator       types.Operator
	Pubkey         regcoord.BN254G1Point
	WaitForReceipt bool
}

type SocketUpdateRequest struct {
	Socket         types.Socket
	WaitForReceipt bool
}

type OperatorRegisterInQuorumWithAVSRequest struct {
	OperatorEcdsaPrivateKey            *ecdsa.PrivateKey
	OperatorToAvsRegistrationSigSalt   [32]byte
	OperatorToAvsRegistrationSigExpiry *big.Int
	BlsKeyPair                         *bls.KeyPair
	QuorumNumbers                      types.QuorumNums
	Socket                             string
	WaitForReceipt                     bool
}
