package avsregistry_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestReaderMethods(t *testing.T) {
	clients, _ := testclients.BuildTestClients(t)
	chainReader := clients.ReadClients.AvsRegistryChainReader

	quorumNumbers := types.QuorumNums{0}

	t.Run("get quorum state", func(t *testing.T) {
		count, err := chainReader.GetQuorumCount(&bind.CallOpts{})
		require.NoError(t, err)
		require.NotNil(t, count)
	})

	t.Run("get operator stake in quorums at current block", func(t *testing.T) {
		stake, err := chainReader.GetOperatorsStakeInQuorumsAtCurrentBlock(&bind.CallOpts{}, quorumNumbers)
		require.NoError(t, err)
		require.NotNil(t, stake)
	})

	t.Run("get operator stake in quorums at block", func(t *testing.T) {
		stake, err := chainReader.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{}, quorumNumbers, 100)
		require.NoError(t, err)
		require.NotNil(t, stake)
	})

	t.Run("get operator address in quorums at current block", func(t *testing.T) {
		addresses, err := chainReader.GetOperatorAddrsInQuorumsAtCurrentBlock(&bind.CallOpts{}, quorumNumbers)
		require.NoError(t, err)
		require.NotNil(t, addresses)
	})

	t.Run(
		"get operators stake in quorums of operator at block returns error for non-registered operator",
		func(t *testing.T) {
			operatorAddress := common.Address{0x1}
			operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
			require.NoError(t, err)

			_, _, err = chainReader.GetOperatorsStakeInQuorumsOfOperatorAtBlock(&bind.CallOpts{}, operatorId, 100)
			require.Error(t, err)
			require.Contains(t, err.Error(), "Failed to get operators state")
		})

	t.Run(
		"get single operator stake in quorums of operator at current block returns error for non-registered operator",
		func(t *testing.T) {
			operatorAddress := common.Address{0x1}
			operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
			require.NoError(t, err)

			stakes, err := chainReader.GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(&bind.CallOpts{}, operatorId)
			require.NoError(t, err)
			require.Equal(t, 0, len(stakes))
		})

	t.Run("get check signatures indices returns error for non-registered operator", func(t *testing.T) {
		operatorAddress := common.Address{0x1}
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		_, err = chainReader.GetCheckSignaturesIndices(
			&bind.CallOpts{},
			100,
			quorumNumbers,
			[]types.OperatorId{operatorId},
		)
		require.Contains(t, err.Error(), "Failed to get check signatures indices")
	})

	t.Run("get operator id", func(t *testing.T) {
		operatorAddress := common.Address{0x1}
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		require.NotNil(t, operatorId)
	})

	t.Run("get operator from id returns zero address for non-registered operator", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		operatorId, err := chainReader.GetOperatorId(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)

		retrievedAddress, err := chainReader.GetOperatorFromId(&bind.CallOpts{}, operatorId)
		require.NoError(t, err)
		require.Equal(t, retrievedAddress, common.Address{0x0})
	})

	t.Run("query registration detail", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		quorums, err := chainReader.QueryRegistrationDetail(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		require.Equal(t, 1, len(quorums))
	})

	t.Run("is operator registered", func(t *testing.T) {
		operatorAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
		isRegistered, err := chainReader.IsOperatorRegistered(&bind.CallOpts{}, operatorAddress)
		require.NoError(t, err)
		require.False(t, isRegistered)
	})

	t.Run(
		"query existing registered operator pub keys", func(t *testing.T) {
			addresses, pubKeys, err := chainReader.QueryExistingRegisteredOperatorPubKeys(
				context.Background(),
				big.NewInt(0),
				nil,
				nil,
			)
			require.NoError(t, err)
			require.Equal(t, 0, len(pubKeys))
			require.Equal(t, 0, len(addresses))
		})

	t.Run(
		"query existing registered operator sockets", func(t *testing.T) {
			address_to_sockets, err := chainReader.QueryExistingRegisteredOperatorSockets(
				context.Background(),
				big.NewInt(0),
				nil,
				nil,
			)
			require.NoError(t, err)
			require.Equal(t, 0, len(address_to_sockets))
		})
}

// Test that the reader returns an error when the configuration is invalid.
func TestReaderWithInvalidConfiguration(t *testing.T) {
	_, anvilHttpEndpoint := testclients.BuildTestClients(t)

	config := avsregistry.Config{}
	chainReader, err := testclients.NewTestAVSChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	quorumNumbers := types.QuorumNums{0}

	tests := []struct {
		name        string
		runFunc     func() error
		expectError bool
	}{
		{
			name: "get quorum state",
			runFunc: func() error {
				_, err := chainReader.GetQuorumCount(&bind.CallOpts{})
				return err
			},
			expectError: true,
		},
		{
			name: "get operator stake in quorums at current block",
			runFunc: func() error {
				_, err := chainReader.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{}, quorumNumbers, 100)
				return err
			},
			expectError: true,
		},
		{
			name: "get operator address in quorums at current block",
			runFunc: func() error {
				_, err := chainReader.GetOperatorAddrsInQuorumsAtCurrentBlock(&bind.CallOpts{}, quorumNumbers)
				return err
			},
			expectError: true,
		},
		{
			name: "get operators stake in quorums of operator at block",
			runFunc: func() error {
				randomOperatorId := types.OperatorId{99}
				_, _, err := chainReader.GetOperatorsStakeInQuorumsOfOperatorAtBlock(
					&bind.CallOpts{},
					randomOperatorId,
					100,
				)
				return err
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s with invalid config", tc.name), func(t *testing.T) {
			err := tc.runFunc()
			if tc.expectError {
				require.Error(t, err, "Expected error for %s", tc.name)
			}
		})
	}
}
