package avsregistry_test

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestReaderMethods(t *testing.T) {
	clients, _ := testclients.BuildTestClients(t)
	chainReader := clients.ReadClients.AvsRegistryChainReader

	quorumNumbers := types.QuorumNums{0}

	t.Run("get quorum state", func(t *testing.T) {
		request := avsregistry.QuorumCountRequest{
			BlockNumber: nil,
		}
		response, err := chainReader.GetQuorumCount(context.Background(), request)
		require.NoError(t, err)
		require.NotNil(t, response.QuorumCount)
	})

	t.Run("get operator stake in quorums at current block", func(t *testing.T) {
		request := avsregistry.OperatorsStakeInQuorumAtCurrentBlockRequest{
			BlockNumber:   nil,
			QuorumNumbers: quorumNumbers,
		}
		response, err := chainReader.GetOperatorsStakeInQuorumsAtCurrentBlock(context.Background(), request)
		require.NoError(t, err)
		require.NotNil(t, response.OperatorsStakeInQuorum)
	})

	t.Run("get operator stake in quorums at block", func(t *testing.T) {
		request := avsregistry.OperatorsStakeInQuorumAtBlockRequest{
			BlockNumber:           nil,
			QuorumNumbers:         quorumNumbers,
			HistoricalBlockNumber: 100,
		}
		response, err := chainReader.GetOperatorsStakeInQuorumsAtBlock(context.Background(), request)
		require.NoError(t, err)
		require.NotNil(t, response.OperatorsStakeInQuorum)
	})

	t.Run("get operator address in quorums at current block", func(t *testing.T) {
		request := avsregistry.OperatorAddrsInQuorumsAtCurrentBlockRequest{
			BlockNumber:   nil,
			QuorumNumbers: quorumNumbers,
		}
		response, err := chainReader.GetOperatorAddrsInQuorumsAtCurrentBlock(context.Background(), request)
		require.NoError(t, err)
		require.NotNil(t, response.OperatorAddrsInQuorums)
	})

	t.Run(
		"get operators stake in quorums of operator at block returns error for non-registered operator",
		func(t *testing.T) {
			operatorAddress := common.Address{0x1}
			request := avsregistry.OperatorIdRequest{
				BlockNumber:     nil,
				OperatorAddress: operatorAddress,
			}
			response, err := chainReader.GetOperatorId(context.Background(), request)
			require.NoError(t, err)

			operatorStakeRequest := avsregistry.OperatorsStakeInQuorumsOfOperatorAtBlockRequest{
				BlockNumber:           nil,
				OperatorId:            response.OperatorId,
				HistoricalBlockNumber: 100,
			}

			_, err = chainReader.GetOperatorsStakeInQuorumsOfOperatorAtBlock(context.Background(), operatorStakeRequest)
			require.Error(t, err)
			require.Contains(t, err.Error(), "Failed to get operators state")
		})

	t.Run(
		"get single operator stake in quorums of operator at current block returns error for non-registered operator",
		func(t *testing.T) {
			operatorAddress := common.Address{0x1}
			request := avsregistry.OperatorIdRequest{
				BlockNumber:     nil,
				OperatorAddress: operatorAddress,
			}
			response, err := chainReader.GetOperatorId(context.Background(), request)
			require.NoError(t, err)

			stakesRequest := avsregistry.OperatorStakeInQuorumsOfOperatorAtCurrentBlockRequest{
				BlockNumber: nil,
				OperatorId:  response.OperatorId,
			}
			responseOperators, err := chainReader.GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(context.Background(), stakesRequest)
			require.NoError(t, err)
			require.Equal(t, 0, len(responseOperators.QuorumStakes))
		})

	t.Run("get check signatures indices returns error for non-registered operator", func(t *testing.T) {
		operatorAddress := common.Address{0x1}
		request := avsregistry.OperatorIdRequest{
			BlockNumber:     nil,
			OperatorAddress: operatorAddress,
		}
		response, err := chainReader.GetOperatorId(context.Background(), request)
		require.NoError(t, err)

		_, err = chainReader.GetCheckSignaturesIndices(
			context.Background(),
			avsregistry.SignaturesIndicesRequest{
				BlockNumber:          nil,
				ReferenceBlockNumber: 100,
				QuorumNumbers:        quorumNumbers,
				NonSignerOperatorIds: []types.OperatorId{response.OperatorId},
			},
		)
		require.Contains(t, err.Error(), "Failed to get check signatures indices")
	})

	t.Run("get operator id", func(t *testing.T) {
		operatorAddress := common.Address{0x1}
		request := avsregistry.OperatorIdRequest{
			BlockNumber:     nil,
			OperatorAddress: operatorAddress,
		}
		response, err := chainReader.GetOperatorId(context.Background(), request)
		require.NoError(t, err)
		require.NotNil(t, response.OperatorId)
	})

	t.Run("get operator from id returns zero address for non-registered operator", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		request := avsregistry.OperatorIdRequest{
			BlockNumber:     nil,
			OperatorAddress: operatorAddress,
		}
		response, err := chainReader.GetOperatorId(context.Background(), request)
		require.NoError(t, err)

		responseOperator, err := chainReader.GetOperatorFromId(context.Background(), avsregistry.OperatorFromIdRequest{
			BlockNumber: nil,
			OperatorId:  response.OperatorId,
		})
		require.NoError(t, err)
		require.Equal(t, responseOperator.OperatorAddress, common.Address{0x0})
	})

	t.Run("query registration detail", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		request := avsregistry.RegistrationDetailRequest{
			BlockNumber:     nil,
			OperatorAddress: operatorAddress,
		}
		response, err := chainReader.QueryRegistrationDetail(context.Background(), request)
		require.NoError(t, err)
		require.Equal(t, 1, len(response.Quorums))
	})

	t.Run("is operator registered", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		request := avsregistry.OperatorRegisteredRequest{
			BlockNumber:     nil,
			OperatorAddress: operatorAddress,
		}
		response, err := chainReader.IsOperatorRegistered(context.Background(), request)
		require.NoError(t, err)
		require.False(t, response.IsRegistered)
	})

	t.Run(
		"query existing registered operator pub keys", func(t *testing.T) {
			request := avsregistry.OperatorQueryRequest{
				BlockNumber: nil,
				StartBlock:  0,
				StopBlock:   0,
				BlockRange:  0,
			}
			response, err := chainReader.QueryExistingRegisteredOperatorPubKeys(
				context.Background(),
				request,
			)
			require.NoError(t, err)
			require.Equal(t, 0, len(response.OperatorAddresses))
			require.Equal(t, 0, len(response.OperatorsPubkeys))
		})

	t.Run(
		"query existing registered operator sockets", func(t *testing.T) {
			request := avsregistry.OperatorQueryRequest{
				BlockNumber: nil,
				StartBlock:  0,
				StopBlock:   0,
				BlockRange:  0,
			}
			response, err := chainReader.QueryExistingRegisteredOperatorSockets(
				context.Background(),
				request,
			)
			require.NoError(t, err)
			require.Equal(t, 0, len(response.Sockets))
		})
}
