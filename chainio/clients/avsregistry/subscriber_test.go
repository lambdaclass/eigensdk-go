package avsregistry_test

import (
	"context"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubscriberAvsRegistry(t *testing.T) {
	client, _ := testclients.BuildTestClients(t)
	chainSubscriber := client.AvsRegistryChainSubscriber
	chainWriter := client.AvsRegistryChainWriter

	pubKeyRegistrations, event, err := chainSubscriber.SubscribeToNewPubkeyRegistrations()
	defer event.Unsubscribe()
	require.NoError(t, err)

	// Emit a NewPubkeyRegistration event creating a new operator
	keypair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err, "Failed to create BLS key pair")

	ecdsaPrivateKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err, "Failed to parse ECDSA private key")

	quorumNumbers := types.QuorumNums{0}

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

	select {
	case newPubkeyRegistration := <-pubKeyRegistrations:
		expectedOperator := crypto.PubkeyToAddress(ecdsaPrivateKey.PublicKey)
		assert.Equal(t, expectedOperator, newPubkeyRegistration.Operator)
	case <-time.After(10 * time.Second):
		// Throw an error if the event is not received within 10 seconds, making the test fail
		t.Fatal("Timed out waiting for NewPubkeyRegistration event")
	}
}
