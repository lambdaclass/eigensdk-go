package avsregistry_test

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	regcoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	servicemanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
		ServiceManagerAddress:         contractAddrs.ServiceManager,
	}

	chainWriter, err := testclients.NewTestAvsRegistryWriterFromConfig(anvilHttpEndpoint, operatorPrivateKeyHex, config)
	require.NoError(t, err)

	keypair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)

	addr := gethcommon.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	quorumNumbers := types.QuorumNums{0}

	subCtx, cancelFn := context.WithCancel(context.Background())
	cancelFn()

	t.Run("update socket without being registered", func(t *testing.T) {
		receipt, err := chainWriter.UpdateSocket(
			context.Background(),
			types.Socket("102901920192019201902910291209"),
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("register operator", func(t *testing.T) {
		receipt, err := chainWriter.RegisterOperator(
			context.Background(),
			ecdsaPrivateKey,
			keypair,
			quorumNumbers,
			"",
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update stake of operator subset", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfOperatorSubsetForAllQuorums(
			context.Background(),
			[]gethcommon.Address{addr},
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update stake of entire operator set", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfEntireOperatorSetForQuorums(
			context.Background(),
			[][]gethcommon.Address{{addr}},
			quorumNumbers,
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("deregister operator", func(t *testing.T) {
		receipt, err := chainWriter.DeregisterOperator(
			context.Background(),
			quorumNumbers,
			chainioutils.ConvertToBN254G1Point(keypair.PubKey),
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("update socket", func(t *testing.T) {
		receipt, err := chainWriter.RegisterOperator(
			context.Background(),
			ecdsaPrivateKey,
			keypair,
			quorumNumbers,
			"",
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)

		receipt, err = chainWriter.UpdateSocket(
			context.Background(),
			types.Socket(""),
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)
	})

	t.Run("set rewards initiator", func(t *testing.T) {
		// Set up to create a ServiceManager binding
		ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
		require.NoError(t, err)

		contractBlsRegistryCoordinator, err := regcoordinator.NewContractRegistryCoordinator(
			contractAddrs.RegistryCoordinator,
			ethHttpClient,
		)
		require.NoError(t, err)

		serviceManagerAddr, err := contractBlsRegistryCoordinator.ServiceManager(&bind.CallOpts{})
		require.NoError(t, err)

		serviceManager, err := servicemanager.NewContractServiceManagerBase(
			serviceManagerAddr,
			ethHttpClient,
		)
		require.NoError(t, err)

		// Check that, first, initiator address is anvil first address
		initiator_initial_addr, err := serviceManager.RewardsInitiator(&bind.CallOpts{})
		require.NoError(t, err)
		assert.Equal(t, initiator_initial_addr.String(), testutils.ANVIL_FIRST_ADDRESS)

		// Modify initiator address, set it as anvil second address
		new_initiator_addr := gethcommon.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
		receipt, err := chainWriter.SetRewardsInitiator(
			context.Background(),
			new_initiator_addr,
			true,
		)
		require.NoError(t, err)
		require.NotNil(t, receipt)

		// Check that, after modifying it, initiator address is now anvil second address
		initiator_modified_addr, err := serviceManager.RewardsInitiator(&bind.CallOpts{})
		require.NoError(t, err)
		assert.Equal(t, initiator_modified_addr, new_initiator_addr)
	})

	// Error cases
	t.Run("fail register operator cancelling context", func(t *testing.T) {
		receipt, err := chainWriter.RegisterOperator(
			subCtx,
			ecdsaPrivateKey,
			keypair,
			quorumNumbers,
			"",
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update stake of operator subset cancelling context", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfOperatorSubsetForAllQuorums(
			subCtx,
			[]gethcommon.Address{addr},
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update stake of entire operator set cancelling context", func(t *testing.T) {
		receipt, err := chainWriter.UpdateStakesOfEntireOperatorSetForQuorums(
			subCtx,
			[][]gethcommon.Address{{addr}},
			quorumNumbers,
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update stake of entire operator set because of quorum length", func(t *testing.T) {
		// Fails because operators per quorum length is distinct from quorum numbers
		receipt, err := chainWriter.UpdateStakesOfEntireOperatorSetForQuorums(
			context.Background(),
			[][]gethcommon.Address{{addr, addr}},
			quorumNumbers,
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail deregister operator cancelling context", func(t *testing.T) {
		receipt, err := chainWriter.DeregisterOperator(
			subCtx,
			quorumNumbers,
			chainioutils.ConvertToBN254G1Point(keypair.PubKey),
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail deregister operator because of operator not registered", func(t *testing.T) {
		quorumNumbers := types.QuorumNums{}
		receipt, err := chainWriter.DeregisterOperator(
			context.Background(),
			quorumNumbers,
			chainioutils.ConvertToBN254G1Point(keypair.PubKey),
			true,
		)
		assert.Error(t, err)
		assert.Nil(t, receipt)
	})

	t.Run("fail update socket cancelling context", func(t *testing.T) {
		receipt, err := chainWriter.UpdateSocket(
			subCtx,
			types.Socket(""),
			true,
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
