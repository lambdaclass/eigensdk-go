package avsregistry

import (
	"math/big"

	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

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

type OperatorRegisteredRequest struct {
	BlockNumber     *big.Int
	OperatorAddress gethcommon.Address
}

type OperatorRegisteredResponse struct {
	IsRegistered bool
}

// Uint64 or Uint32? There was a comment in the original code that
// wants to use uint instead of big.Int
type OperatorQueryRequest struct {
	BlockNumber *big.Int
	StartBlock  uint64
	StopBlock   uint64
	BlockRange  uint64
}

type OperatorPubKeysRequestResponse struct {
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

type OperatorsStakeInQuorumsOfOperatorAtBlockRequest struct {
	BlockNumber           *big.Int
	HistoricalBlockNumber uint32
	OperatorId            types.OperatorId
}

type OperatorsStakeInQuorumsOfOperatorAtCurrentBlockRequest struct {
	BlockNumber *big.Int
	OperatorId  types.OperatorId
}

type OperatorsStakeInQuorumsOfOperatorResponse struct {
	QuorumNumbers           types.QuorumNums
	OperatorsStakesInQuorum [][]opstateretriever.OperatorStateRetrieverOperator
}

type OperatorStakeInQuorumsOfOperatorAtCurrentBlockRequest struct {
	BlockNumber *big.Int
	OperatorId  types.OperatorId
}

type OperatorStakeInQuorumsOfOperatorResponse struct {
	QuorumStakes map[types.QuorumNum]types.StakeAmount
}
