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
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

const ANVIL_FIRST_ADDRESS = "f39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
const ANVIL_FIRST_PRIVATE_KEY = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

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
}

func TestAdminFunctions(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	assert.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	assert.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	permissionControllerAddr := common.HexToAddress("0x610178dA211FEF7D417bC0e6FeD39F05609AD788")
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}

	operatorAddr := common.HexToAddress(ANVIL_FIRST_ADDRESS)
	privateKeyHex := ANVIL_FIRST_PRIVATE_KEY
	accountChainWriter, err := NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	assert.NoError(t, err)

	pendingAdminAddr := common.HexToAddress("14dC79964da2C08b23698B3D3cc7Ca32193d9955")
	pendingAdminPrivateKeyHex := "4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"
	adminChainWriter, err := NewTestChainWriterFromConfig(anvilHttpEndpoint, pendingAdminPrivateKeyHex, config)
	assert.NoError(t, err)

	chainReader, err := NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	assert.NoError(t, err)

	t.Run("non-existent pending admin", func(t *testing.T) {
		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdminAddr)
		assert.NoError(t, err)
		assert.False(t, isPendingAdmin)
	})

	t.Run("list pending admins when empty", func(t *testing.T) {
		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdminAddr)
		assert.NoError(t, err)
		assert.False(t, isPendingAdmin)

		listPendingAdmins, err := chainReader.ListPendingAdmins(context.Background(), operatorAddr)
		assert.NoError(t, err)
		assert.Empty(t, listPendingAdmins)
	})

	t.Run("add pending admin and list", func(t *testing.T) {
		request := elcontracts.AddPendingAdminRequest{
			AccountAddress: operatorAddr,
			AdminAddress:   pendingAdminAddr,
			WaitForReceipt: true,
		}

		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdminAddr)
		assert.NoError(t, err)
		assert.False(t, isPendingAdmin)
		
		receipt, err := accountChainWriter.AddPendingAdmin(context.Background(), request)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)

		isPendingAdmin, err = chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdminAddr)
		assert.NoError(t, err)
		assert.True(t, isPendingAdmin)

		listPendingAdmins, err := chainReader.ListPendingAdmins(context.Background(), operatorAddr)
		assert.NoError(t, err)
		assert.NotEmpty(t, listPendingAdmins)
		assert.Len(t, listPendingAdmins, 1)
	})

	t.Run("non-existent admin", func(t *testing.T) {
		isAdmin, err := chainReader.IsAdmin(context.Background(), operatorAddr, pendingAdminAddr)
		assert.NoError(t, err)
		assert.False(t, isAdmin)
	})

	t.Run("existing admin", func(t *testing.T) {
		acceptAdminRequest := elcontracts.AcceptAdminRequest{
			AccountAddress: operatorAddr,
			WaitForReceipt: true,
		}

		receipt, err := adminChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)

		isAdmin, err := chainReader.IsAdmin(context.Background(), operatorAddr, pendingAdminAddr)
		assert.NoError(t, err)
		assert.True(t, isAdmin)
	})

	t.Run("list admins", func(t *testing.T) {
		listAdmins, err := chainReader.ListAdmins(context.Background(), operatorAddr)
		assert.NoError(t, err)
		assert.Len(t, listAdmins, 1)

		admin := listAdmins[0]
		isAdmin, err := chainReader.IsAdmin(context.Background(), operatorAddr, admin)
		assert.NoError(t, err)
		assert.True(t, isAdmin)
	})
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
