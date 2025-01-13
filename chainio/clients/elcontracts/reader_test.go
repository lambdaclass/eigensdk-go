package elcontracts_test

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChainReader(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	operator := types.Operator{
		Address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
	}

	t.Run("is operator registered", func(t *testing.T) {
		isOperator, err := clients.ElChainReader.IsOperatorRegistered(ctx, operator)
		assert.NoError(t, err)
		assert.Equal(t, isOperator, true)
	})

	t.Run("get operator details", func(t *testing.T) {
		operatorDetails, err := clients.ElChainReader.GetOperatorDetails(ctx, operator)
		assert.NoError(t, err)
		assert.NotNil(t, operatorDetails)
		assert.Equal(t, operator.Address, operatorDetails.Address)
	})

	t.Run("get strategy and underlying token", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategy, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingToken(
			ctx,
			strategyAddr,
		)
		assert.NoError(t, err)
		assert.NotNil(t, strategy)
		assert.NotEqual(t, common.Address{}, underlyingTokenAddr)

		erc20Token, err := erc20.NewContractIERC20(underlyingTokenAddr, clients.EthHttpClient)
		assert.NoError(t, err)

		tokenName, err := erc20Token.Name(&bind.CallOpts{})
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenName)
	})

	t.Run("get strategy and underlying ERC20 token", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategy, contractUnderlyingToken, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingERC20Token(
			ctx,
			strategyAddr,
		)
		assert.NoError(t, err)
		assert.NotNil(t, strategy)
		assert.NotEqual(t, common.Address{}, underlyingTokenAddr)
		assert.NotNil(t, contractUnderlyingToken)

		tokenName, err := contractUnderlyingToken.Name(&bind.CallOpts{})
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenName)
	})

	t.Run("get operator shares in strategy", func(t *testing.T) {
		shares, err := clients.ElChainReader.GetOperatorSharesInStrategy(
			ctx,
			common.HexToAddress(operator.Address),
			contractAddrs.Erc20MockStrategy,
		)
		assert.NoError(t, err)
		assert.NotZero(t, shares)
	})

	t.Run("calculate delegation approval digest hash", func(t *testing.T) {
		staker := common.Address{0x0}
		delegationApprover := common.Address{0x0}
		approverSalt := [32]byte{}
		expiry := big.NewInt(0)
		digest, err := clients.ElChainReader.CalculateDelegationApprovalDigestHash(
			ctx,
			staker,
			common.HexToAddress(operator.Address),
			delegationApprover,
			approverSalt,
			expiry,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, digest)
	})

	t.Run("calculate operator AVS registration digest hash", func(t *testing.T) {
		avs := common.Address{0x0}
		salt := [32]byte{}
		expiry := big.NewInt(0)
		digest, err := clients.ElChainReader.CalculateOperatorAVSRegistrationDigestHash(
			ctx,
			common.HexToAddress(operator.Address),
			avs,
			salt,
			expiry,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, digest)
	})

	t.Run("staker shares test (GetStakerShares)", func(t *testing.T) {
		strategies, shares, err := clients.ElChainReader.GetStakerShares(
			ctx,
			common.HexToAddress(operator.Address),
		)
		assert.NotZero(t, len(strategies))            // Strategies has at least one element
		assert.NotZero(t, len(shares))                // Shares has at least one element
		assert.Equal(t, len(strategies), len(shares)) // Strategies has the same ammount of elements as
		assert.NoError(t, err)
	})

	t.Run("get delegated operator", func(t *testing.T) {
		// The delegated operator of an operator is the operator itself
		val := big.NewInt(0)
		address, err := clients.ElChainReader.GetDelegatedOperator(
			ctx,
			common.HexToAddress(operator.Address),
			val,
		)

		assert.NoError(t, err)
		assert.Equal(t, address.String(), operator.Address)
	})

	t.Run("get current claimable distribution root without submitted roots is zero", func(t *testing.T) {
		contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

		rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
		config := elcontracts.Config{
			DelegationManagerAddress:  contractAddrs.DelegationManager,
			RewardsCoordinatorAddress: rewardsCoordinatorAddr,
		}

		chainReader, err := NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
		require.NoError(t, err)

		root, err := chainReader.GetCurrentClaimableDistributionRoot(
			ctx,
		)

		assert.NoError(t, err)
		assert.Zero(t, root)
	})

	t.Run("get current claimable distribution root with and without submitted roots", func(t *testing.T) {
		// Verifies GetCurrentClaimableDistributionRoot returns 0 if no root and the root if there's one
		root := [32]byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		}

		rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
		config := elcontracts.Config{
			DelegationManagerAddress:  contractAddrs.DelegationManager,
			RewardsCoordinatorAddress: rewardsCoordinatorAddr,
		}

		chainReader, err := NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
		require.NoError(t, err)

		// Create and configure rewards coordinator
		ethClient, err := ethclient.Dial(anvilHttpEndpoint)
		require.NoError(t, err)
		rewardsCoordinator, err := rewardscoordinator.NewContractIRewardsCoordinator(rewardsCoordinatorAddr, ethClient)
		require.NoError(t, err)

		// Set delay to zero to inmediatly operate with coordinator
		receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", uint32(0))
		require.NoError(t, err)
		require.Equal(t, receipt.Status, uint64(1))

		// Create txManager to send transactions to the Ethereum node
		txManager, err := NewTestTxManager(anvilHttpEndpoint, "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
		require.NoError(t, err)
		noSendTxOpts, err := txManager.GetNoSendTxOpts()
		require.NoError(t, err)

		rewardsUpdater := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

		// Change the rewards updater to be able to submit the new root
		tx, err := rewardsCoordinator.SetRewardsUpdater(noSendTxOpts, rewardsUpdater)
		require.NoError(t, err)

		waitForReceipt := true
		_, err = txManager.Send(context.Background(), tx, waitForReceipt)
		require.NoError(t, err)

		// Check that if there is no root submitted the result is zero
		distr_root, err := chainReader.GetCurrentClaimableDistributionRoot(
			ctx,
		)
		assert.NoError(t, err)
		assert.Zero(t, distr_root.Root)

		currRewardsCalculationEndTimestamp, err := chainReader.CurrRewardsCalculationEndTimestamp(context.Background())
		require.NoError(t, err)

		tx, err = rewardsCoordinator.SubmitRoot(noSendTxOpts, root, currRewardsCalculationEndTimestamp+1)
		require.NoError(t, err)

		_, err = txManager.Send(context.Background(), tx, waitForReceipt)
		require.NoError(t, err)

		// Check that if there is a root submitted the result is that root
		distr_root, err = chainReader.GetCurrentClaimableDistributionRoot(
			ctx,
		)
		assert.NoError(t, err)
		assert.Equal(t, distr_root.Root, root)
	})

	t.Run("get root index from root hash with and without submitted roots", func(t *testing.T) {
		rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
		config := elcontracts.Config{
			DelegationManagerAddress:  contractAddrs.DelegationManager,
			RewardsCoordinatorAddress: rewardsCoordinatorAddr,
		}

		chainReader, err := NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
		require.NoError(t, err)

		// Create and configure rewards coordinator
		ethClient, err := ethclient.Dial(anvilHttpEndpoint)
		require.NoError(t, err)
		rewardsCoordinator, err := rewardscoordinator.NewContractIRewardsCoordinator(rewardsCoordinatorAddr, ethClient)
		require.NoError(t, err)

		// Set delay to zero to inmediatly operate with coordinator
		receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", uint32(0))
		require.NoError(t, err)
		require.Equal(t, receipt.Status, uint64(1))

		// Create txManager to send transactions to the Ethereum node
		txManager, err := NewTestTxManager(anvilHttpEndpoint, "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
		require.NoError(t, err)
		noSendTxOpts, err := txManager.GetNoSendTxOpts()
		require.NoError(t, err)

		rewardsUpdater := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

		// Change the rewards updater to be able to submit the new root
		tx, err := rewardsCoordinator.SetRewardsUpdater(noSendTxOpts, rewardsUpdater)
		require.NoError(t, err)

		waitForReceipt := true
		_, err = txManager.Send(context.Background(), tx, waitForReceipt)
		require.NoError(t, err)

		root := [32]byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		}

		// Check that if there is no root submitted the result is an InvalidRoot error
		root_index, err := chainReader.GetRootIndexFromHash(
			ctx,
			root,
		)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "execution reverted: custom error 0x504570e3")
		assert.Zero(t, root_index)

		currRewardsCalculationEndTimestamp, err := chainReader.CurrRewardsCalculationEndTimestamp(context.Background())
		require.NoError(t, err)

		tx, err = rewardsCoordinator.SubmitRoot(noSendTxOpts, root, currRewardsCalculationEndTimestamp+1)
		require.NoError(t, err)

		_, err = txManager.Send(context.Background(), tx, waitForReceipt)
		require.NoError(t, err)

		root2 := [32]byte{
			0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		}

		currRewardsCalculationEndTimestamp2, err := chainReader.CurrRewardsCalculationEndTimestamp(context.Background())
		require.NoError(t, err)

		tx, err = rewardsCoordinator.SubmitRoot(noSendTxOpts, root2, currRewardsCalculationEndTimestamp2+1)
		require.NoError(t, err)

		_, err = txManager.Send(context.Background(), tx, waitForReceipt)
		require.NoError(t, err)

		// Check that the first root inserted is the first indexed (zero)
		root_index, err = chainReader.GetRootIndexFromHash(
			ctx,
			root,
		)
		assert.NoError(t, err)
		assert.Equal(t, root_index, uint32(0))

		// Check that the second root inserted is the second indexed (zero)
		root_index, err = chainReader.GetRootIndexFromHash(
			ctx,
			root2,
		)
		assert.NoError(t, err)
		assert.Equal(t, root_index, uint32(1))
	})

	t.Run("get cumulative claimed rewards", func(t *testing.T) {
		rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
		config := elcontracts.Config{
			DelegationManagerAddress:  contractAddrs.DelegationManager,
			RewardsCoordinatorAddress: rewardsCoordinatorAddr,
		}

		//chainWriter, err := NewTestChainWriterFromConfig(anvilHttpEndpoint, "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", config)
		//require.NoError(t, err)
		chainReader, err := NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
		require.NoError(t, err)

		// Create and configure rewards coordinator
		ethClient, err := ethclient.Dial(anvilHttpEndpoint)
		require.NoError(t, err)
		rewardsCoordinator, err := rewardscoordinator.NewContractIRewardsCoordinator(rewardsCoordinatorAddr, ethClient)
		require.NoError(t, err)

		// Set delay to zero to inmediatly operate with coordinator
		receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", uint32(0))
		require.NoError(t, err)
		require.Equal(t, receipt.Status, uint64(1))

		// Create txManager to send transactions to the Ethereum node
		txManager, err := NewTestTxManager(anvilHttpEndpoint, "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
		require.NoError(t, err)
		noSendTxOpts, err := txManager.GetNoSendTxOpts()
		require.NoError(t, err)

		rewardsUpdater := common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

		// Change the rewards updater to be able to submit the new root
		tx, err := rewardsCoordinator.SetRewardsUpdater(noSendTxOpts, rewardsUpdater)
		require.NoError(t, err)

		waitForReceipt := true
		_, err = txManager.Send(context.Background(), tx, waitForReceipt)
		require.NoError(t, err)

		strategyAddr := contractAddrs.Erc20MockStrategy
		strategy, contractUnderlyingToken, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingERC20Token(
			ctx,
			strategyAddr,
		)
		assert.NoError(t, err)
		assert.NotNil(t, strategy)
		assert.NotEqual(t, common.Address{}, underlyingTokenAddr)
		assert.NotNil(t, contractUnderlyingToken)

		claimed, err := chainReader.GetCumulativeClaimed(ctx, common.HexToAddress("f39Fd6e51aad88F6F4ce6aB8827279cffFb92266"), underlyingTokenAddr)
		assert.True(t, claimed.Cmp(big.NewInt(0)) == 0)
		assert.NoError(t, err)

		// This tests that without claims result is zero, should add claims and test that the cumulative increases
	})

}

// The functions below will be replaced for those placed in testutils/testclients/testclients.go
func setTestRewardsCoordinatorActivationDelay(
	httpEndpoint string,
	privateKeyHex string,
	activationDelay uint32,
) (*gethtypes.Receipt, error) {
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(httpEndpoint)
	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	ethHttpClient, err := ethclient.Dial(httpEndpoint)
	if err != nil {
		return nil, utils.WrapError("Failed to create eth client", err)
	}

	rewardsCoordinator, err := rewardscoordinator.NewContractIRewardsCoordinator(rewardsCoordinatorAddr, ethHttpClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create rewards coordinator", err)
	}

	txManager, err := NewTestTxManager(httpEndpoint, privateKeyHex)
	if err != nil {
		return nil, utils.WrapError("Failed to create tx manager", err)
	}

	noSendOpts, err := txManager.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("Failed to get NoSend tx opts", err)
	}

	tx, err := rewardsCoordinator.SetActivationDelay(noSendOpts, activationDelay)
	if err != nil {
		return nil, utils.WrapError("Failed to create SetActivationDelay tx", err)
	}

	receipt, err := txManager.Send(context.Background(), tx, true)
	if err != nil {
		return nil, utils.WrapError("Failed to send SetActivationDelay tx", err)
	}
	return receipt, err
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
