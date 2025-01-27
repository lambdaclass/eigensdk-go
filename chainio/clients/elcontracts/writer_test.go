package elcontracts_test

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	mockerc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockERC20"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterOperator(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	anvilWsEndpoint, err := anvilC.Endpoint(context.Background(), "ws")
	require.NoError(t, err)
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

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

	t.Run("register as an operator", func(t *testing.T) {
		// Fund the new address with 5 ether
		fundedAccount := "0x408EfD9C90d59298A9b32F4441aC9Df6A2d8C3E1"
		fundedPrivateKeyHex := "3339854a8622364bcd5650fa92eac82d5dccf04089f5575a761c9b7d3c405b1c"
		richPrivateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
		code, _, err := anvilC.Exec(
			context.Background(),
			[]string{"cast",
				"send",
				fundedAccount,
				"--value",
				"5ether",
				"--private-key",
				richPrivateKeyHex,
			},
		)
		assert.NoError(t, err)
		assert.Equal(t, 0, code)
		time.Sleep(500 * time.Millisecond) // wait for the account to be funded

		ecdsaPrivateKey, err := crypto.HexToECDSA(fundedPrivateKeyHex)
		require.NoError(t, err)

		clients, err := clients.BuildAll(
			chainioConfig,
			ecdsaPrivateKey,
			logger,
		)
		require.NoError(t, err)

		opts, err := clients.TxManager.GetNoSendTxOpts()
		require.NoError(t, err)

		txOptions := &elcontracts.TxOptions{
			WaitForReceipt: true,
			Options:        opts,
		}

		operator :=
			types.Operator{
				Address:                   fundedAccount,
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				MetadataUrl:               "https://madhur-test-public.s3.us-east-2.amazonaws.com/metadata.json",
			}

		request := elcontracts.RegisterOperatorRequest{
			Operator: operator,
		}

		receipt, err := clients.ElChainWriter.RegisterAsOperator(context.Background(), request, txOptions)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})

	t.Run("register as an operator already registered", func(t *testing.T) {
		operatorAddress := "0x408EfD9C90d59298A9b32F4441aC9Df6A2d8C3E1"
		operatorPrivateKeyHex := "3339854a8622364bcd5650fa92eac82d5dccf04089f5575a761c9b7d3c405b1c"

		ecdsaPrivateKey, err := crypto.HexToECDSA(operatorPrivateKeyHex)
		require.NoError(t, err)

		clients, err := clients.BuildAll(
			chainioConfig,
			ecdsaPrivateKey,
			logger,
		)
		require.NoError(t, err)

		opts, err := clients.TxManager.GetNoSendTxOpts()
		require.NoError(t, err)

		txOptions := &elcontracts.TxOptions{
			WaitForReceipt: true,
			Options:        opts,
		}

		operator :=
			types.Operator{
				Address:                   operatorAddress,
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				MetadataUrl:               "https://madhur-test-public.s3.us-east-2.amazonaws.com/metadata.json",
			}

		request := elcontracts.RegisterOperatorRequest{
			Operator: operator,
		}

		_, err = clients.ElChainWriter.RegisterAsOperator(context.Background(), request, txOptions)
		assert.Error(t, err)
	})
}

func TestRegisterAndDeregisterFromOperatorSets(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	operatorAddressHex := testutils.ANVIL_SECOND_ADDRESS
	operatorPrivateKeyHex := testutils.ANVIL_SECOND_PRIVATE_KEY

	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: contractAddrs.RewardsCoordinator,
	}

	// Create operator clients
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, operatorPrivateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, testutils.ANVIL_SECOND_PRIVATE_KEY)
	require.NoError(t, err)

	opts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	avsAddress := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	operatorSetId := uint32(1)
	erc20MockStrategyAddr := contractAddrs.Erc20MockStrategy

	// Create an operator set to register an operator on it
	err = createOperatorSet(
		anvilHttpEndpoint,
		testutils.ANVIL_FIRST_PRIVATE_KEY,
		avsAddress,
		operatorSetId,
		erc20MockStrategyAddr,
	)
	require.NoError(t, err)

	operatorAddress := common.HexToAddress(operatorAddressHex)
	keypair, err := bls.NewKeyPairFromString("0x01")
	require.NoError(t, err)

	request := elcontracts.RegisterOperatorSetsRequest{
		OperatorAddress:            operatorAddress,
		RegistryCoordinatorAddress: contractAddrs.RegistryCoordinator,
		AVSAddress:                 avsAddress,
		OperatorSetIds:             []uint32{operatorSetId},
		Socket:                     "socket",
		BlsKeyPair:                 keypair,
	}

	txOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        opts,
	}

	operatorSet := allocationmanager.OperatorSet{
		Avs: avsAddress,
		Id:  uint32(operatorSetId),
	}
	t.Run("register operator for operator set", func(t *testing.T) {
		receipt, err := chainWriter.RegisterForOperatorSets(
			context.Background(),
			request,
			txOptions,
		)

		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		isRegistered, err := chainReader.IsOperatorRegisteredWithOperatorSet(
			context.Background(),
			operatorAddress,
			operatorSet,
		)
		require.NoError(t, err)
		require.Equal(t, true, isRegistered)
	})

	t.Run("register operator for same operator set", func(t *testing.T) {
		_, err = chainWriter.RegisterForOperatorSets(
			context.Background(),
			request,
			txOptions,
		)
		require.Error(t, err, "cannot register an operator to an operator set that is already registered")
	})

	deregistrationRequest := elcontracts.OperatorSetDeregisterRequest{
		OperatorAddress: operatorAddress,
		AVSAddress:      avsAddress,
		OperatorSetIds:  []uint32{operatorSetId},
	}

	t.Run("deregister operator from operator set", func(t *testing.T) {
		receipt, err := chainWriter.DeregisterFromOperatorSets(
			context.Background(),
			deregistrationRequest,
			txOptions,
		)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		isRegistered, err := chainReader.IsOperatorRegisteredWithOperatorSet(
			context.Background(),
			operatorAddress,
			operatorSet,
		)
		require.NoError(t, err)
		require.False(t, isRegistered)
	})

	t.Run("deregister operator from operator set when not registered", func(t *testing.T) {
		_, err = chainWriter.DeregisterFromOperatorSets(
			context.Background(),
			deregistrationRequest,
			txOptions,
		)
		require.Error(t, err, "cannot deregister an operator that is not registered")
	})
}

func TestChainWriter(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	opts, err := clients.TxManager.GetNoSendTxOpts()
	require.NoError(t, err)

	txOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        opts,
	}

	t.Run("update operator details", func(t *testing.T) {
		walletModified, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
		assert.NoError(t, err)
		walletModifiedAddress := crypto.PubkeyToAddress(walletModified.PublicKey)

		operatorModified := types.Operator{
			Address:                   walletModifiedAddress.Hex(),
			DelegationApproverAddress: walletModifiedAddress.Hex(),
			MetadataUrl:               "eigensdk-go",
		}

		request := elcontracts.OperatorDetailsRequest{
			Operator: operatorModified,
		}

		receipt, err := clients.ElChainWriter.UpdateOperatorDetails(context.Background(), request, txOptions)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})

	t.Run("update operator details when address is not an operator", func(t *testing.T) {
		wrongOperatorModified := types.Operator{
			Address:                   testutils.ANVIL_THIRD_ADDRESS,
			DelegationApproverAddress: testutils.ANVIL_FIRST_ADDRESS,
			MetadataUrl:               "eigensdk-go",
		}

		request := elcontracts.OperatorDetailsRequest{
			Operator: wrongOperatorModified,
		}

		_, err := clients.ElChainWriter.UpdateOperatorDetails(
			context.Background(),
			request,
			txOptions,
		)
		assert.Error(t, err, "cannot update operator details for an address that is not an operator")
	})

	t.Run("update metadata URI", func(t *testing.T) {
		walletModified, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
		assert.NoError(t, err)
		walletModifiedAddress := crypto.PubkeyToAddress(walletModified.PublicKey)
		request := elcontracts.MetadataURIRequest{
			OperatorAddress: walletModifiedAddress,
			Uri:             "https://0.0.0.0",
		}

		receipt, err := clients.ElChainWriter.UpdateMetadataURI(
			context.Background(),
			request,
			txOptions,
		)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})

	t.Run("update metadata URI when address is not an operator", func(t *testing.T) {
		requestWronOperator := elcontracts.MetadataURIRequest{
			OperatorAddress: common.HexToAddress(testutils.ANVIL_THIRD_ADDRESS),
			Uri:             "https://0.0.0.0",
		}
		_, err := clients.ElChainWriter.UpdateMetadataURI(
			context.Background(),
			requestWronOperator,
			txOptions,
		)
		assert.Error(t, err, "cannot update metadata URI for an address that is not an operator")
	})

	t.Run("deposit ERC20 into strategy", func(t *testing.T) {
		request := elcontracts.ERC20IntoStrategyRequest{
			StrategyAddress: contractAddrs.Erc20MockStrategy,
			Amount:          big.NewInt(1),
		}
		receipt, err := clients.ElChainWriter.DepositERC20IntoStrategy(
			context.Background(),
			request,
			txOptions,
		)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})
}

func TestSetClaimerFor(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	opts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	txOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        opts,
	}

	request := elcontracts.ClaimForRequest{
		Claimer: common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS),
	}

	// call SetClaimerFor
	receipt, err := chainWriter.SetClaimerFor(context.Background(), request, txOptions)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)
}

func TestSetOperatorPISplit(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	activationDelay := uint32(0)
	// Set activation delay to zero so that the new PI split can be retrieved immediately after setting it
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	opts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	txOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        opts,
	}

	expectedInitialSplit := uint16(1000)
	initialSplit, err := chainReader.GetOperatorPISplit(context.Background(), operatorAddr)
	require.NoError(t, err)
	require.Equal(t, expectedInitialSplit, initialSplit)

	newSplit := initialSplit + 1
	// Set a new operator PI split

	request := elcontracts.OperatorPISplitRequest{
		OperatorAddress: operatorAddr,
		Split:           newSplit,
	}

	receipt, err = chainWriter.SetOperatorPISplit(context.Background(), request, txOptions)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Retrieve the operator PI split to check it has been set
	updatedSplit, err := chainReader.GetOperatorPISplit(context.Background(), operatorAddr)
	require.NoError(t, err)
	require.Equal(t, newSplit, updatedSplit)

	// Set a invalid operator PI split
	request.Split = uint16(10001)
	_, err = chainWriter.SetOperatorPISplit(context.Background(), request, txOptions)
	require.Error(t, err, "split must be less than 10000")
}

func TestSetOperatorAVSSplit(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	avsAddr := contractAddrs.ServiceManager
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	activationDelay := uint32(0)

	// Set activation delay to zero so that the new AVS split can be retrieved immediately after setting it
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	opts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	txOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        opts,
	}

	expectedInitialSplit := uint16(1000)
	initialSplit, err := chainReader.GetOperatorAVSSplit(context.Background(), operatorAddr, avsAddr)
	require.NoError(t, err)
	require.Equal(t, expectedInitialSplit, initialSplit)

	newSplit := initialSplit + 1
	// Set a new operator AVS split
	request := elcontracts.OperatorAVSSplitRequest{
		OperatorAddress: operatorAddr,
		AVSAddress:      avsAddr,
		Split:           newSplit,
	}

	receipt, err = chainWriter.SetOperatorAVSSplit(
		context.Background(),
		request,
		txOptions,
	)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Retrieve the operator AVS split to check it has been set
	updatedSplit, err := chainReader.GetOperatorAVSSplit(context.Background(), operatorAddr, avsAddr)
	require.NoError(t, err)
	require.Equal(t, newSplit, updatedSplit)

	// Set a invalid operator AVS split
	request.Split = uint16(10001)
	_, err = chainWriter.SetOperatorAVSSplit(
		context.Background(),
		request,
		txOptions,
	)
	require.Error(t, err, "split must be less than 10000")
}

func TestSetAllocationDelay(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	opts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	txOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        opts,
	}

	t.Run("set allocation delay", func(t *testing.T) {
		delay := uint32(10)
		request := elcontracts.AllocationDelayRequest{
			OperatorAddress: operatorAddr,
			Delay:           delay,
		}
		receipt, err := chainWriter.SetAllocationDelay(context.Background(), request, txOptions)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)
	})

	t.Run("set allocation delay with invalid caller", func(t *testing.T) {
		invalidCaller := common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
		delay := uint32(20)
		request := elcontracts.AllocationDelayRequest{
			OperatorAddress: invalidCaller,
			Delay:           delay,
		}
		_, err = chainWriter.SetAllocationDelay(
			context.Background(),
			request,
			txOptions,
		)
		require.Error(t, err, "cannot set allocation delay with an invalid caller")
	})
}

func TestSetAndRemovePermission(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	permissionControllerAddr := common.HexToAddress(testutils.PERMISSION_CONTROLLER_ADDRESS)

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)
	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	accountAddress := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	appointeeAddress := common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
	target := common.HexToAddress(testutils.ANVIL_THIRD_ADDRESS)
	selector := [4]byte{0, 1, 2, 3}

	setPermissionRequest := elcontracts.PermissionSetRequest{
		AccountAddress:   accountAddress,
		AppointeeAddress: appointeeAddress,
		Target:           target,
		Selector:         selector,
	}

	removePermissionRequest := elcontracts.PermissionRemoveRequest{
		AccountAddress:   accountAddress,
		AppointeeAddress: appointeeAddress,
		Target:           target,
		Selector:         selector,
	}

	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	opts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	txOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        opts,
	}

	t.Run("set permission to account", func(t *testing.T) {
		receipt, err := chainWriter.SetPermission(context.Background(), setPermissionRequest, txOptions)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		canCall, err := chainReader.CanCall(context.Background(), accountAddress, appointeeAddress, target, selector)
		require.NoError(t, err)
		require.True(t, canCall)
	})

	t.Run("set permission to account when already set", func(t *testing.T) {
		_, err := chainWriter.SetPermission(context.Background(), setPermissionRequest, txOptions)
		require.Error(t, err, "cannot set a permission that has already been set")
	})

	t.Run("remove permission from account", func(t *testing.T) {
		receipt, err := chainWriter.RemovePermission(context.Background(), removePermissionRequest, txOptions)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		canCall, err := chainReader.CanCall(context.Background(), accountAddress, appointeeAddress, target, selector)
		require.NoError(t, err)
		require.False(t, canCall)
	})

	t.Run("remove permission from account when not set", func(t *testing.T) {
		_, err := chainWriter.RemovePermission(context.Background(), removePermissionRequest, txOptions)
		require.Error(t, err, "cannot remove a permission that has not been set")
	})
}

func TestModifyAllocations(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	config := elcontracts.Config{
		DelegationManagerAddress: contractAddrs.DelegationManager,
	}

	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	opts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	txOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        opts,
	}

	strategyAddr := contractAddrs.Erc20MockStrategy
	avsAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	operatorSetId := uint32(1)

	operatorSet := allocationmanager.OperatorSet{
		Avs: avsAddr,
		Id:  operatorSetId,
	}
	newAllocation := uint64(100)
	allocateParams := []allocationmanager.IAllocationManagerTypesAllocateParams{
		{
			OperatorSet:   operatorSet,
			Strategies:    []common.Address{strategyAddr},
			NewMagnitudes: []uint64{newAllocation},
		},
	}

	requestAllocation := elcontracts.AllocationModifyRequest{
		OperatorAddress: operatorAddr,
		Allocations:     allocateParams,
	}

	_, err = chainWriter.ModifyAllocations(context.Background(), requestAllocation, txOptions)
	require.Error(t, err, "cannot modify allocations without initializing the allocation delay")

	delay := uint32(1)
	requestDelay := elcontracts.AllocationDelayRequest{
		OperatorAddress: operatorAddr,
		Delay:           delay,
	}
	// The allocation delay must be initialized before modifying the allocations
	receipt, err := chainWriter.SetAllocationDelay(context.Background(), requestDelay, txOptions)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	allocationConfigurationDelay := 1200
	// Advance the chain by the required number of blocks
	// (ALLOCATION_CONFIGURATION_DELAY) to apply the allocation delay
	testutils.AdvanceChainByNBlocksExecInContainer(context.Background(), allocationConfigurationDelay+1, anvilC)

	// Retrieve the allocation delay so that the delay is applied
	_, err = chainReader.GetAllocationDelay(context.Background(), operatorAddr)
	require.NoError(t, err)

	err = createOperatorSet(anvilHttpEndpoint, privateKeyHex, avsAddr, operatorSetId, strategyAddr)
	require.NoError(t, err)

	receipt, err = chainWriter.ModifyAllocations(context.Background(), requestAllocation, txOptions)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Check that the new allocation is pending and the current magnitude is zero
	allocationInfo, err := chainReader.GetAllocationInfo(context.Background(), operatorAddr, strategyAddr)
	require.NoError(t, err)
	pendingDiff := allocationInfo[0].PendingDiff
	require.Equal(t, big.NewInt(int64(newAllocation)), pendingDiff)
	require.Equal(t, allocationInfo[0].CurrentMagnitude, big.NewInt(0))

	// Retrieve the allocation delay and advance the chain
	allocationDelay, err := chainReader.GetAllocationDelay(context.Background(), operatorAddr)
	require.NoError(t, err)
	testutils.AdvanceChainByNBlocksExecInContainer(context.Background(), int(allocationDelay), anvilC)

	// Check the new allocation has been updated after the delay
	allocationInfo, err = chainReader.GetAllocationInfo(context.Background(), operatorAddr, strategyAddr)
	require.NoError(t, err)

	currentMagnitude := allocationInfo[0].CurrentMagnitude
	require.Equal(t, big.NewInt(int64(newAllocation)), currentMagnitude)
}

func TestAddAndRemovePendingAdmin(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	permissionControllerAddr := common.HexToAddress(testutils.PERMISSION_CONTROLLER_ADDRESS)

	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)
	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	opts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	txOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        opts,
	}

	pendingAdmin := common.HexToAddress(testutils.ANVIL_THIRD_ADDRESS)
	acceptPendingAdminRequest := elcontracts.PendingAdminAcceptRequest{
		AccountAddress: operatorAddr,
		AdminAddress:   pendingAdmin,
	}

	removePendingAdminRequest := elcontracts.PendingAdminRemoveRequest{
		AccountAddress: operatorAddr,
		AdminAddress:   pendingAdmin,
	}

	t.Run("remove pending admin when not added", func(t *testing.T) {
		_, err := chainWriter.RemovePendingAdmin(context.Background(), removePendingAdminRequest, txOptions)
		require.Error(t, err, "cannot remove a pending admin that has not been added")
	})

	t.Run("add pending admin", func(t *testing.T) {
		receipt, err := chainWriter.AddPendingAdmin(context.Background(), acceptPendingAdminRequest, txOptions)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdmin)
		require.NoError(t, err)
		require.True(t, isPendingAdmin)
	})

	t.Run("add pending admin when already added", func(t *testing.T) {
		_, err := chainWriter.AddPendingAdmin(context.Background(), acceptPendingAdminRequest, txOptions)
		require.Error(t, err, "cannot add a pending admin that has already been added")
	})

	t.Run("remove pending admin", func(t *testing.T) {
		receipt, err := chainWriter.RemovePendingAdmin(context.Background(), removePendingAdminRequest, txOptions)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdmin)
		require.NoError(t, err)
		require.False(t, isPendingAdmin)
	})
}

func TestAcceptAdmin(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	permissionControllerAddr := common.HexToAddress(testutils.PERMISSION_CONTROLLER_ADDRESS)

	accountAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	accountPrivateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	accountChainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, accountPrivateKeyHex, config)
	require.NoError(t, err)

	pendingAdminPrivateKeyHex := testutils.ANVIL_SECOND_PRIVATE_KEY
	adminChainWriter, err := testclients.NewTestChainWriterFromConfig(
		anvilHttpEndpoint,
		pendingAdminPrivateKeyHex,
		config,
	)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	txManagerAccount, err := testclients.NewTestTxManager(anvilHttpEndpoint, accountPrivateKeyHex)
	require.NoError(t, err)
	optsAccount, err := txManagerAccount.GetNoSendTxOpts()
	require.NoError(t, err)

	txManagerPendingAdmin, err := testclients.NewTestTxManager(anvilHttpEndpoint, pendingAdminPrivateKeyHex)
	require.NoError(t, err)
	optsPendingAdmin, err := txManagerPendingAdmin.GetNoSendTxOpts()
	require.NoError(t, err)

	txOptionsAccount := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        optsAccount,
	}

	txOptionsPendingAdmin := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        optsPendingAdmin,
	}

	pendingAdminAddr := common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
	request := elcontracts.PendingAdminAcceptRequest{
		AccountAddress: accountAddr,
		AdminAddress:   pendingAdminAddr,
	}
	receipt, err := accountChainWriter.AddPendingAdmin(context.Background(), request, txOptionsAccount)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	acceptAdminRequest := elcontracts.AdminAcceptRequest{
		AccountAddress: accountAddr,
	}
	t.Run("accept admin", func(t *testing.T) {
		receipt, err = adminChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest, txOptionsPendingAdmin)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		isAdmin, err := chainReader.IsAdmin(context.Background(), accountAddr, pendingAdminAddr)
		require.NoError(t, err)
		require.True(t, isAdmin)
	})

	t.Run("accept admin when already accepted", func(t *testing.T) {
		_, err = adminChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest, txOptionsPendingAdmin)
		require.Error(t, err, "cannot accept an admin that has already been accepted")
	})
}

func TestRemoveAdmin(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	permissionControllerAddr := common.HexToAddress(testutils.PERMISSION_CONTROLLER_ADDRESS)

	accountAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	accountPrivateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	accountChainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, accountPrivateKeyHex, config)
	require.NoError(t, err)

	// Adding two admins and removing one. Cannot remove the last admin, so one must remain
	admin1 := common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
	admin1PrivateKeyHex := testutils.ANVIL_SECOND_PRIVATE_KEY

	admin2 := common.HexToAddress(testutils.ANVIL_THIRD_ADDRESS)
	admin2PrivateKeyHex := testutils.ANVIL_THIRD_PRIVATE_KEY

	admin1ChainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, admin1PrivateKeyHex, config)
	require.NoError(t, err)

	admin2ChainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, admin2PrivateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	accountTxManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, accountPrivateKeyHex)
	require.NoError(t, err)
	optsAccount, err := accountTxManager.GetNoSendTxOpts()
	require.NoError(t, err)

	admin1TxManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, admin1PrivateKeyHex)
	require.NoError(t, err)
	optsAdmin1, err := admin1TxManager.GetNoSendTxOpts()
	require.NoError(t, err)

	admin2TxManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, admin2PrivateKeyHex)
	require.NoError(t, err)
	optsAdmin2, err := admin2TxManager.GetNoSendTxOpts()
	require.NoError(t, err)

	accountTxOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        optsAccount,
	}

	admin1TxOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        optsAdmin1,
	}

	admin2TxOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        optsAdmin2,
	}
	addAdmin1Request := elcontracts.PendingAdminAcceptRequest{
		AccountAddress: accountAddr,
		AdminAddress:   admin1,
	}
	addAdmin2Request := elcontracts.PendingAdminAcceptRequest{
		AccountAddress: accountAddr,
		AdminAddress:   admin2,
	}
	acceptAdminRequest := elcontracts.AdminAcceptRequest{
		AccountAddress: accountAddr,
	}

	// Add and accept admin 1
	receipt, err := accountChainWriter.AddPendingAdmin(context.Background(), addAdmin1Request, accountTxOptions)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	receipt, err = admin1ChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest, admin1TxOptions)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Add and accept admin 2
	receipt, err = admin1ChainWriter.AddPendingAdmin(context.Background(), addAdmin2Request, admin1TxOptions)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	receipt, err = admin2ChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest, admin2TxOptions)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	removeAdminRequest := elcontracts.AdminRemoveRequest{
		AccountAddress: accountAddr,
		AdminAddress:   admin2,
	}

	t.Run("remove admin 2", func(t *testing.T) {
		receipt, err = admin1ChainWriter.RemoveAdmin(context.Background(), removeAdminRequest, admin1TxOptions)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		isAdmin, err := chainReader.IsAdmin(context.Background(), accountAddr, admin2)
		require.NoError(t, err)
		require.False(t, isAdmin)
	})

	t.Run("remove admin 2 when already removed", func(t *testing.T) {
		_, err := admin1ChainWriter.RemoveAdmin(context.Background(), removeAdminRequest, admin1TxOptions)
		require.Error(t, err, "cannot remove an admin that has already been removed")
	})
}

func TestProcessClaim(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	opts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	txOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        opts,
	}

	activationDelay := uint32(0)
	// Set activation delay to zero so that the earnings can be claimed right after submitting the root
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	cumulativeEarnings := int64(42)
	recipient := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	claim, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings, privateKeyHex)
	require.NoError(t, err)

	request := elcontracts.ClaimProcessRequest{
		Claim:            *claim,
		RecipientAddress: recipient,
	}
	receipt, err = chainWriter.ProcessClaim(context.Background(), request, txOptions)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)
}

func TestProcessClaims(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, testutils.ANVIL_FIRST_PRIVATE_KEY)
	require.NoError(t, err)

	opts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	txOptions := &elcontracts.TxOptions{
		WaitForReceipt: true,
		Options:        opts,
	}

	activationDelay := uint32(0)
	// Set activation delay to zero so that the earnings can be claimed right after submitting the root
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	recipient := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	cumulativeEarnings1 := int64(42)
	cumulativeEarnings2 := int64(4256)

	emptyClaims := []rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim{}
	request := elcontracts.ClaimsProcessRequest{
		Claims:           emptyClaims,
		RecipientAddress: recipient,
	}
	_, err = chainWriter.ProcessClaims(context.Background(), request, txOptions)
	require.Error(t, err, "cannot process empty claims")

	// Generate 2 claims
	claim1, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings1, privateKeyHex)
	require.NoError(t, err)

	claim2, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings2, privateKeyHex)
	require.NoError(t, err)

	claims := []rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim{
		*claim1, *claim2,
	}
	request = elcontracts.ClaimsProcessRequest{
		Claims:           claims,
		RecipientAddress: recipient,
	}
	receipt, err = chainWriter.ProcessClaims(context.Background(), request, txOptions)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)
}

// Creates an operator set with `avsAddress`, `operatorSetId` and `erc20MockStrategyAddr`.
func createOperatorSet(
	anvilHttpEndpoint string,
	privateKeyHex string,
	avsAddress common.Address,
	operatorSetId uint32,
	erc20MockStrategyAddr common.Address,
) error {
	testConfig := testutils.GetDefaultTestConfig()
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	config := elcontracts.Config{
		DelegationManagerAddress: contractAddrs.DelegationManager,
	}
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})
	ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
	if err != nil {
		return err
	}

	elBindings, err := elcontracts.NewBindingsFromConfig(config, ethHttpClient, logger)
	if err != nil {
		return err
	}

	allocationManager := elBindings.AllocationManager
	registryCoordinatorAddress := contractAddrs.RegistryCoordinator
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(
		registryCoordinatorAddress,
		ethHttpClient,
	)
	if err != nil {
		return err
	}
	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, privateKeyHex)
	if err != nil {
		return err
	}
	noSendTxOpts, err := txManager.GetNoSendTxOpts()
	if err != nil {
		return err
	}

	tx, err := allocationManager.SetAVSRegistrar(noSendTxOpts, avsAddress, registryCoordinatorAddress)
	if err != nil {
		return err
	}

	waitForReceipt := true

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return err
	}

	tx, err = registryCoordinator.EnableOperatorSets(noSendTxOpts)
	if err != nil {
		return err
	}

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
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

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
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

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	return err
}

// Sets the testing RewardsCoordinator's activationDelay.
// This is useful to test ChainWriter setter functions that depend on activationDelay.
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

	txManager, err := testclients.NewTestTxManager(httpEndpoint, privateKeyHex)
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

// Returns a (test) claim for the given cumulativeEarnings, whose earner is
// the account given by the testutils.ANVIL_FIRST_ADDRESS address.
// This was taken from the eigensdk-rs
// https://github.com/Layr-Labs/eigensdk-rs/blob/d79b3672584b92f3c5fb204fde6bea394fbf0f12/crates/chainio/clients/elcontracts/src/lib.rs#L146
func newTestClaim(
	chainReader *elcontracts.ChainReader,
	httpEndpoint string,
	cumulativeEarnings int64,
	privateKeyHex string,
) (*rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim, error) {
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(httpEndpoint)
	mockStrategyAddr := contractAddrs.Erc20MockStrategy
	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	waitForReceipt := true

	ethClient, err := ethclient.Dial(httpEndpoint)
	if err != nil {
		return nil, utils.WrapError("Failed to create eth client", err)
	}

	txManager, err := testclients.NewTestTxManager(httpEndpoint, privateKeyHex)
	if err != nil {
		return nil, utils.WrapError("Failed to create tx manager", err)
	}

	contractStrategy, err := strategy.NewContractIStrategy(mockStrategyAddr, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch strategy contract", err)
	}

	tokenAddr, err := contractStrategy.UnderlyingToken(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return nil, utils.WrapError("Failed to fetch token address", err)
	}

	token, err := mockerc20.NewContractMockERC20(tokenAddr, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create token contract", err)
	}

	noSendTxOpts, err := txManager.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("Failed to get NoSend tx opts", err)
	}

	// Mint tokens for the RewardsCoordinator
	tx, err := token.Mint(noSendTxOpts, rewardsCoordinatorAddr, big.NewInt(cumulativeEarnings))
	if err != nil {
		return nil, utils.WrapError("Failed to create Mint tx", err)
	}

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("Failed to mint tokens for RewardsCoordinator", err)
	}

	// Generate token tree leaf
	// For the tree structure, see
	// https://github.com/Layr-Labs/eigenlayer-contracts/blob/a888a1cd1479438dda4b138245a69177b125a973/docs/core/RewardsCoordinator.md#rewards-merkle-tree-structure
	earnerAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	tokenLeaf := rewardscoordinator.IRewardsCoordinatorTypesTokenTreeMerkleLeaf{
		Token:              tokenAddr,
		CumulativeEarnings: big.NewInt(cumulativeEarnings),
	}
	encodedTokenLeaf := []byte{}
	tokenLeafSalt := uint8(1)

	// Write the *big.Int to a 32-byte sized buffer to match the uint256 length
	cumulativeEarningsBytes := [32]byte{}
	tokenLeaf.CumulativeEarnings.FillBytes(cumulativeEarningsBytes[:])

	encodedTokenLeaf = append(encodedTokenLeaf, tokenLeafSalt)
	encodedTokenLeaf = append(encodedTokenLeaf, tokenLeaf.Token.Bytes()...)
	encodedTokenLeaf = append(encodedTokenLeaf, cumulativeEarningsBytes[:]...)

	// Hash token tree leaf to get root
	earnerTokenRoot := crypto.Keccak256(encodedTokenLeaf)

	// Generate earner tree leaf
	earnerLeaf := rewardscoordinator.IRewardsCoordinatorTypesEarnerTreeMerkleLeaf{
		Earner:          earnerAddr,
		EarnerTokenRoot: [32]byte(earnerTokenRoot),
	}
	// Encode earner leaft
	encodedEarnerLeaf := []byte{}
	earnerLeafSalt := uint8(0)
	encodedEarnerLeaf = append(encodedEarnerLeaf, earnerLeafSalt)
	encodedEarnerLeaf = append(encodedEarnerLeaf, earnerLeaf.Earner.Bytes()...)
	encodedEarnerLeaf = append(encodedEarnerLeaf, earnerTokenRoot...)

	// Hash encoded earner tree leaf to get root
	earnerTreeRoot := crypto.Keccak256(encodedEarnerLeaf)

	// Fetch the next root index from contract
	nextRootIndex, err := chainReader.GetDistributionRootsLength(context.Background())
	if err != nil {
		return nil, utils.WrapError("Failed to call GetDistributionRootsLength", err)
	}

	tokenLeaves := []rewardscoordinator.IRewardsCoordinatorTypesTokenTreeMerkleLeaf{tokenLeaf}
	// Construct the claim
	claim := rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim{
		RootIndex:   uint32(nextRootIndex.Uint64()),
		EarnerIndex: 0,
		// Empty proof because leaf == root
		EarnerTreeProof: []byte{},
		EarnerLeaf:      earnerLeaf,
		TokenIndices:    []uint32{0},
		// Empty proof because leaf == root
		TokenTreeProofs: [][]byte{{}},
		TokenLeaves:     tokenLeaves,
	}

	root := [32]byte(earnerTreeRoot)
	// Fetch the current timestamp to increase it
	currRewardsCalculationEndTimestamp, err := chainReader.CurrRewardsCalculationEndTimestamp(context.Background())
	if err != nil {
		return nil, utils.WrapError("Failed to call CurrRewardsCalculationEndTimestamp", err)
	}

	rewardsCoordinator, err := rewardscoordinator.NewContractIRewardsCoordinator(rewardsCoordinatorAddr, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create rewards coordinator contract", err)
	}

	rewardsUpdater := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	// Change the rewards updater to be able to submit the new root
	tx, err = rewardsCoordinator.SetRewardsUpdater(noSendTxOpts, rewardsUpdater)
	if err != nil {
		return nil, utils.WrapError("Failed to create SetRewardsUpdater tx", err)
	}

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("Failed to setRewardsUpdate", err)
	}

	tx, err = rewardsCoordinator.SubmitRoot(noSendTxOpts, root, currRewardsCalculationEndTimestamp+1)
	if err != nil {
		return nil, utils.WrapError("Failed to create SubmitRoot tx", err)
	}

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("Failed to submit root", err)
	}

	return &claim, nil
}
