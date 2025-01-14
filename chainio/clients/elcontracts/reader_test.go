package elcontracts_test

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	ANVIL_FIRST_ADDRESS           = "f39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	ANVIL_SECOND_ADDRESS          = "70997970C51812dc3A010C7d01b50e0d17dc79C8"
	ANVIL_SECOND_PRIVATE_KEY      = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
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
}

func TestSlashableSharesFunctions(t *testing.T) {
	eigenClients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	avsAddress := common.HexToAddress(ANVIL_FIRST_ADDRESS)
	operatorAddress := common.HexToAddress(ANVIL_SECOND_ADDRESS)

	operatorPrivateKeyHex := ANVIL_SECOND_PRIVATE_KEY
	operatorClients, err := newTestClients(anvilHttpEndpoint, operatorPrivateKeyHex)
	require.NoError(t, err)

	erc20MockStrategyAddr := contractAddrs.Erc20MockStrategy
	operatorSetId := uint32(1)

	t.Run("create First OperatorSet", func(t *testing.T) {
		err := createOperatorSet(eigenClients, avsAddress, operatorSetId, erc20MockStrategyAddr)
		require.NoError(t, err)
	})

	t.Run("Register Operator to OperatorSets", func(t *testing.T) {
		keypair, err := bls.NewKeyPairFromString("0x01")
		require.NoError(t, err)

		request := elcontracts.RegistrationRequest{
			OperatorAddress: operatorAddress,
			AVSAddress:      avsAddress,
			OperatorSetIds:  []uint32{operatorSetId},
			WaitForReceipt:  true,
			Socket:          "socket",
			BlsKeyPair:      keypair,
		}

		registryCoordinatorAddress := contractAddrs.RegistryCoordinator
		receipt, err := operatorClients.ElChainWriter.RegisterForOperatorSets(
			context.Background(),
			registryCoordinatorAddress,
			request,
		)
		require.NoError(t, err)
		require.Equal(t, uint64(1), receipt.Status)
	})

	t.Run("get Slashable Shares for operatorSet1", func(t *testing.T) {
		operatorSet := allocationmanager.OperatorSet{
			Avs: avsAddress,
			Id:  operatorSetId,
		}
		strategies := []common.Address{erc20MockStrategyAddr}

		shares, err := eigenClients.ElChainReader.GetSlashableShares(
			context.Background(),
			operatorAddress,
			operatorSet,
			strategies,
		)
		require.NoError(t, err)
		require.NotEmpty(t, shares)
	})

	t.Run("get Slashable Shares for Multiple OperatorSets", func(t *testing.T) {
		operatorSets := []allocationmanager.OperatorSet{
			{Avs: avsAddress, Id: operatorSetId},
		}

		shares, err := eigenClients.ElChainReader.GetSlashableSharesForOperatorSets(
			context.Background(),
			operatorSets,
		)
		require.NoError(t, err)
		require.NotEmpty(t, shares)
		require.Len(t, shares, 1)
	})
}

func createOperatorSet(
	client *clients.Clients,
	avsAddress common.Address,
	operatorSetId uint32,
	erc20MockStrategyAddr common.Address,
) error {
	allocationManagerAddress := client.EigenlayerContractBindings.AllocationManagerAddr
	allocationManager, err := allocationmanager.NewContractAllocationManager(
		allocationManagerAddress,
		client.EthHttpClient,
	)
	if err != nil {
		return err
	}
	registryCoordinatorAddress := client.AvsRegistryContractBindings.RegistryCoordinatorAddr
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(
		registryCoordinatorAddress,
		client.EthHttpClient,
	)
	if err != nil {
		return err
	}

	noSendTxOpts, err := client.TxManager.GetNoSendTxOpts()
	if err != nil {
		return err
	}

	tx, err := allocationManager.SetAVSRegistrar(noSendTxOpts, avsAddress, registryCoordinatorAddress)
	if err != nil {
		return err
	}

	waitForReceipt := true

	_, err = client.TxManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return err
	}

	tx, err = registryCoordinator.EnableOperatorSets(noSendTxOpts)
	if err != nil {
		return err
	}

	_, err = client.TxManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return err
	}

	operatorSetParam := regcoord.IRegistryCoordinatorOperatorSetParam{
		MaxOperatorCount:        10,
		KickBIPsOfOperatorStake: 100,
		KickBIPsOfTotalStake:    1000,
	}
	minimumStake := big.NewInt(0)

	strategyParams := regcoord.IStakeRegistryStrategyParams{
		Strategy:   erc20MockStrategyAddr,
		Multiplier: big.NewInt(1),
	}
	strategyParamsArray := []regcoord.IStakeRegistryStrategyParams{strategyParams}
	lookAheadPeriod := uint32(0)
	tx, err = registryCoordinator.CreateSlashableStakeQuorum(
		noSendTxOpts,
		operatorSetParam,
		minimumStake,
		strategyParamsArray,
		lookAheadPeriod,
	)
	if err != nil {
		return err
	}

	_, err = client.TxManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return err
	}

	strategies := []common.Address{erc20MockStrategyAddr}
	operatorSetParams := allocationmanager.IAllocationManagerTypesCreateSetParams{
		OperatorSetId: operatorSetId,
		Strategies:    strategies,
	}
	operatorSetParamsArray := []allocationmanager.IAllocationManagerTypesCreateSetParams{operatorSetParams}
	tx, err = allocationManager.CreateOperatorSets(noSendTxOpts, avsAddress, operatorSetParamsArray)
	if err != nil {
		return err
	}

	_, err = client.TxManager.Send(context.Background(), tx, waitForReceipt)
	return err
}

func newTestClients(httpEndpoint string, privateKeyHex string) (*clients.Clients, error) {
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(httpEndpoint)
	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 httpEndpoint,
		EthWsUrl:                   httpEndpoint,
		RegistryCoordinatorAddr:    contractAddrs.RegistryCoordinator.String(),
		OperatorStateRetrieverAddr: contractAddrs.OperatorStateRetriever.String(),
		AvsName:                    "exampleAvs",
		PromMetricsIpPortAddress:   ":9090",
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	testConfig := testutils.GetDefaultTestConfig()
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

	testClients, err := clients.BuildAll(
		chainioConfig,
		privateKey,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return testClients, nil
}
