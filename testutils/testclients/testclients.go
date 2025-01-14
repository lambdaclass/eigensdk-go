package testclients

import (
	"context"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

// Starts an anvil container and builds the ChainIO Clients for testing.
func BuildTestClients(t *testing.T) (*clients.Clients, string) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	anvilWsEndpoint, err := anvilC.Endpoint(context.Background(), "ws")
	require.NoError(t, err)
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKeyHex)
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	require.NoError(t, err)

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 anvilHttpEndpoint,
		EthWsUrl:                   anvilWsEndpoint,
		RegistryCoordinatorAddr:    contractAddrs.RegistryCoordinator.String(),
		OperatorStateRetrieverAddr: contractAddrs.OperatorStateRetriever.String(),
		AvsName:                    "exampleAvs",
		PromMetricsIpPortAddress:   ":9090",
	}

	clients, err := clients.BuildAll(
		chainioConfig,
		ecdsaPrivateKey,
		logger,
	)
	require.NoError(t, err)
	return clients, anvilHttpEndpoint
}

// Creates a testing ChainWriter from an httpEndpoint, private key and config.
// This is needed because the existing testclients.BuildTestClients returns a
// ChainReader with a null rewardsCoordinator, which is required for some of the tests.
func NewTestChainReaderFromConfig(
	httpEndpoint string,
	config elcontracts.Config,
) (*elcontracts.ChainReader, error) {
	testConfig := testutils.GetDefaultTestConfig()
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})
	ethHttpClient, err := ethclient.Dial(httpEndpoint)
	if err != nil {
		return nil, utils.WrapError("Failed to create eth client", err)
	}

	testReader, err := elcontracts.NewReaderFromConfig(
		config,
		ethHttpClient,
		logger,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to create chain reader from config", err)
	}
	return testReader, nil
}

// Creates a testing ChainWriter from an httpEndpoint, private key and config.
// This is needed because the existing testclients.BuildTestClients returns a
// ChainWriter with a null rewardsCoordinator, which is required for some of the tests.
func NewTestChainWriterFromConfig(
	httpEndpoint string,
	privateKeyHex string,
	config elcontracts.Config,
) (*elcontracts.ChainWriter, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, utils.WrapError("Failed convert hex string to ecdsa private key", err)
	}
	testConfig := testutils.GetDefaultTestConfig()
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})
	ethHttpClient, err := ethclient.Dial(httpEndpoint)
	if err != nil {
		return nil, utils.WrapError("Failed to create eth client", err)
	}
	chainid, err := ethHttpClient.ChainID(context.Background())
	if err != nil {
		return nil, utils.WrapError("Failed to get chain id", err)
	}
	promReg := prometheus.NewRegistry()
	eigenMetrics := metrics.NewEigenMetrics("", "", promReg, logger)
	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: privateKey}, chainid)
	if err != nil {
		return nil, utils.WrapError("Failed to create the signer from the given config", err)
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, addr, logger)
	if err != nil {
		return nil, utils.WrapError("Failed to create wallet", err)
	}
	txManager := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, addr)
	testWriter, err := elcontracts.NewWriterFromConfig(
		config,
		ethHttpClient,
		logger,
		eigenMetrics,
		txManager,
	)
	if err != nil {
		return nil, err
	}
	return testWriter, nil
}

func NewTestTxManager(httpEndpoint string, privateKeyHex string) (*txmgr.SimpleTxManager, error) {
	testConfig := testutils.GetDefaultTestConfig()
	ethHttpClient, err := ethclient.Dial(httpEndpoint)
	if err != nil {
		return nil, utils.WrapError("Failed to create eth client", err)
	}

	chainid, err := ethHttpClient.ChainID(context.Background())
	if err != nil {
		return nil, utils.WrapError("Failed to retrieve chain id", err)
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, utils.WrapError("Failed to convert hex string to private key", err)
	}
	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: privateKey}, chainid)
	if err != nil {
		return nil, utils.WrapError("Failed to create signer", err)
	}

	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, addr, logger)
	if err != nil {
		return nil, utils.WrapError("Failed to create wallet", err)
	}

	txManager := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, addr)
	return txManager, nil
}
