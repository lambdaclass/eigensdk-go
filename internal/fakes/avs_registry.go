package fakes

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	apkregistrybindings "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/types"

	"github.com/ethereum/go-ethereum/common"
)

type TestOperator struct {
	OperatorAddr     common.Address
	OperatorInfo     types.OperatorInfo
	ContractG1Pubkey apkregistrybindings.BN254G1Point
	ContractG2Pubkey apkregistrybindings.BN254G2Point
	OperatorId       types.OperatorId
}

type FakeAVSRegistryReader struct {
	opAddress  []types.OperatorAddr
	opPubKeys  []types.OperatorPubkeys
	operatorId types.OperatorId
	socket     types.Socket
	err        error
}

func NewFakeAVSRegistryReader(
	opr *TestOperator,
	err error,
) *FakeAVSRegistryReader {
	if opr == nil {
		return &FakeAVSRegistryReader{}
	}
	return &FakeAVSRegistryReader{
		opAddress:  []common.Address{opr.OperatorAddr},
		opPubKeys:  []types.OperatorPubkeys{opr.OperatorInfo.Pubkeys},
		socket:     opr.OperatorInfo.Socket,
		operatorId: opr.OperatorId,
		err:        err,
	}
}

func (f *FakeAVSRegistryReader) QueryExistingRegisteredOperatorPubKeys(
	ctx context.Context,
	request avsregistry.OperatorQueryRequest,
) (avsregistry.OperatorPubKeysResponse, error) {
	return avsregistry.OperatorPubKeysResponse{OperatorAddresses: f.opAddress, OperatorsPubkeys: f.opPubKeys}, f.err
}

func (f *FakeAVSRegistryReader) QueryExistingRegisteredOperatorSockets(
	ctx context.Context,
	request avsregistry.OperatorQueryRequest,
) (avsregistry.OperatorSocketsResponse, error) {
	if len(f.opPubKeys) == 0 {
		return avsregistry.OperatorSocketsResponse{}, nil
	}

	return avsregistry.OperatorSocketsResponse{Sockets: map[types.OperatorId]types.Socket{
		types.OperatorIdFromG1Pubkey(f.opPubKeys[0].G1Pubkey): f.socket,
	}}, nil
}

func (f *FakeAVSRegistryReader) GetOperatorFromId(
	ctx context.Context,
	request avsregistry.OperatorFromIdRequest,
) (avsregistry.OperatorFromIdResponse, error) {
	return avsregistry.OperatorFromIdResponse{OperatorAddress: f.opAddress[0]}, f.err
}

func (f *FakeAVSRegistryReader) GetOperatorsStakeInQuorumsAtBlock(
	ctx context.Context,
	request avsregistry.OperatorsStakeInQuorumAtBlockRequest,
) (avsregistry.OperatorsStakeInQuorumResponse, error) {
	return avsregistry.OperatorsStakeInQuorumResponse{
		OperatorsStakeInQuorum: [][]opstateretriever.OperatorStateRetrieverOperator{
			{
				{
					OperatorId: f.operatorId,
					Stake:      big.NewInt(123),
				},
			},
		},
	}, nil
}

func (f *FakeAVSRegistryReader) GetCheckSignaturesIndices(
	ctx context.Context,
	request avsregistry.SignaturesIndicesRequest,
) (avsregistry.SignaturesIndicesResponse, error) {
	return avsregistry.SignaturesIndicesResponse{}, nil
}
