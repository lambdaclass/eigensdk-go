package avsregistry_test

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriterMethods(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	operatorPrivateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	config := avsregistry.Config{
		RegistryCoordinatorAddress:    contractAddrs.RegistryCoordinator,
		OperatorStateRetrieverAddress: contractAddrs.OperatorStateRetriever,
	}

	chainWriter, err := testclients.NewTestAvsRegistryWriterFromConfig(anvilHttpEndpoint, operatorPrivateKeyHex, config)
	require.NoError(t, err)

	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, operatorPrivateKeyHex)
	require.NoError(t, err)
	opts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)
	txOptions := &avsregistry.TxOptions{
		Options: opts,
	}

	keypair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)

	addr := gethcommon.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	quorumNumbers := types.QuorumNums{0}

	t.Run("update socket without being registered", func(t *testing.T) {
		receipt, err := chainWriter.UpdateSocket(
			context.Background(),
			avsregistry.SocketUpdateRequest{
				Socket:         types.Socket("102901920192019201902910291209"),
				WaitForReceipt: true,
			},
			txOptions,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("register operator", func(t *testing.T) {
		request := avsregistry.OperatorRegisterRequest{
			OperatorEcdsaPrivateKey: ecdsaPrivateKey,
			BlsKeyPair:              keypair,
			QuorumNumbers:           quorumNumbers,
			Socket:                  "",
			WaitForReceipt:          true,
		}
		receipt, err := chainWriter.RegisterOperator(
			context.Background(),
			request,
			txOptions,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update stake of operator subset", func(t *testing.T) {
		request := avsregistry.StakesOfOperatorSubsetForAllQuorumsRequest{
			OperatorsAddresses: []gethcommon.Address{addr},
			WaitForReceipt:     true,
		}
		receipt, err := chainWriter.UpdateStakesOfOperatorSubsetForAllQuorums(
			context.Background(),
			request,
			txOptions,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update stake of entire operator set", func(t *testing.T) {
		request := avsregistry.StakesOfEntireOperatorSetForQuorumsRequest{
			OperatorsPerQuorum: [][]gethcommon.Address{{addr}},
			QuorumNumbers:      quorumNumbers,
			WaitForReceipt:     true,
		}
		receipt, err := chainWriter.UpdateStakesOfEntireOperatorSetForQuorums(
			context.Background(),
			request,
			txOptions,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("deregister operator", func(t *testing.T) {
		request := avsregistry.OperatorDeregisterRequest{
			QuorumNumbers:  quorumNumbers,
			Pubkey:         chainioutils.ConvertToBN254G1Point(keypair.PubKey),
			WaitForReceipt: true,
		}
		receipt, err := chainWriter.DeregisterOperator(
			context.Background(),
			request,
			txOptions,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update socket", func(t *testing.T) {
		request := avsregistry.OperatorRegisterRequest{
			OperatorEcdsaPrivateKey: ecdsaPrivateKey,
			BlsKeyPair:              keypair,
			QuorumNumbers:           quorumNumbers,
			Socket:                  "",
			WaitForReceipt:          true,
		}
		receipt, err := chainWriter.RegisterOperator(
			context.Background(),
			request,
			txOptions,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)

		requestSocket := avsregistry.SocketUpdateRequest{
			Socket:         types.Socket(""),
			WaitForReceipt: true,
		}
		receipt, err = chainWriter.UpdateSocket(
			context.Background(),
			requestSocket,
			txOptions,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	// Error cases
	t.Run("fail register operator cancelling context", func(t *testing.T) {
		subCtx, cancelFn := context.WithCancel(context.Background())
		cancelFn()
		receipt, err := chainWriter.RegisterOperator(
			subCtx,
			avsregistry.OperatorRegisterRequest{
				OperatorEcdsaPrivateKey: ecdsaPrivateKey,
				BlsKeyPair:              keypair,
				QuorumNumbers:           quorumNumbers,
				Socket:                  "",
				WaitForReceipt:          true,
			},
			txOptions,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update stake of operator subset cancelling context", func(t *testing.T) {
		subCtx, cancelFn := context.WithCancel(context.Background())
		cancelFn()
		receipt, err := chainWriter.UpdateStakesOfOperatorSubsetForAllQuorums(
			subCtx,
			avsregistry.StakesOfOperatorSubsetForAllQuorumsRequest{
				OperatorsAddresses: []gethcommon.Address{addr},
				WaitForReceipt:     true,
			},
			txOptions,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update stake of entire operator set cancelling context", func(t *testing.T) {
		subCtx, cancelFn := context.WithCancel(context.Background())
		cancelFn()
		receipt, err := chainWriter.UpdateStakesOfEntireOperatorSetForQuorums(
			subCtx,
			avsregistry.StakesOfEntireOperatorSetForQuorumsRequest{
				OperatorsPerQuorum: [][]gethcommon.Address{{addr}},
				QuorumNumbers:      quorumNumbers,
				WaitForReceipt:     true,
			},
			txOptions,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update stake of entire operator set because of quorum length", func(t *testing.T) {
		// Fails because operators per quorum length is distinct from quorum numbers
		receipt, err := chainWriter.UpdateStakesOfEntireOperatorSetForQuorums(
			context.Background(),
			avsregistry.StakesOfEntireOperatorSetForQuorumsRequest{
				OperatorsPerQuorum: [][]gethcommon.Address{{addr}},
				QuorumNumbers:      types.QuorumNums{0, 1},
				WaitForReceipt:     true,
			},
			txOptions,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail deregister operator cancelling context", func(t *testing.T) {
		subCtx, cancelFn := context.WithCancel(context.Background())
		cancelFn()
		receipt, err := chainWriter.DeregisterOperator(
			subCtx,
			avsregistry.OperatorDeregisterRequest{
				QuorumNumbers:  quorumNumbers,
				Pubkey:         chainioutils.ConvertToBN254G1Point(keypair.PubKey),
				WaitForReceipt: true,
			},
			txOptions,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail deregister operator because of operator not registered", func(t *testing.T) {
		request := avsregistry.OperatorDeregisterRequest{
			QuorumNumbers:  types.QuorumNums{},
			Pubkey:         chainioutils.ConvertToBN254G1Point(keypair.PubKey),
			WaitForReceipt: true,
		}
		receipt, err := chainWriter.DeregisterOperator(
			context.Background(),
			request,
			txOptions,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update socket cancelling context", func(t *testing.T) {
		subCtx, cancelFn := context.WithCancel(context.Background())

		cancelFn()
		receipt, err := chainWriter.UpdateSocket(
			subCtx,
			avsregistry.SocketUpdateRequest{
				Socket:         types.Socket(""),
				WaitForReceipt: true,
			},
			txOptions,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})
}

// Compliance test for BLS signature
func TestBlsSignature(t *testing.T) {
	// read input from JSON if available, otherwise use default values
	// Data taken from
	// https://github.com/Layr-Labs/eigensdk-compliance/blob/429459572302d9fd42c1184b7257703460ba09ca/data_files/bls_signature.json
	var defaultInput = struct {
		Message    string `json:"message"`
		BlsPrivKey string `json:"bls_priv_key"`
	}{
		Message:    "Hello, world!Hello, world!123456",
		BlsPrivKey: "12248929636257230549931416853095037629726205319386239410403476017439825112537",
	}

	testData := testutils.NewTestData(defaultInput)
	// The message to sign
	messageArray := []byte(testData.Input.Message)

	var messageArray32 [32]byte
	copy(messageArray32[:], messageArray)

	// The private key as a string
	privKey, _ := bls.NewPrivateKey(testData.Input.BlsPrivKey)
	keyPair := bls.NewKeyPair(privKey)

	sig := keyPair.SignMessage(messageArray32)

	x := sig.G1Affine.X.String()
	y := sig.G1Affine.Y.String()

	// Values taken from previous run of this test
	assert.Equal(t, x, "15790168376429033610067099039091292283117017641532256477437243974517959682102")
	assert.Equal(t, y, "4960450323239587206117776989095741074887370703941588742100855592356200866613")
}
