package elcontracts_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChainReader(t *testing.T) {
	read_clients, anvilHttpEndpoint := testclients.BuildTestReadClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	operator := types.Operator{
		Address: testutils.ANVIL_FIRST_ADDRESS,
	}

	t.Run("is operator registered", func(t *testing.T) {
		isOperator, err := read_clients.ElChainReader.IsOperatorRegistered(ctx, operator)
		assert.NoError(t, err)
		assert.Equal(t, isOperator, true)
	})

	t.Run("get operator details", func(t *testing.T) {
		operatorDetails, err := read_clients.ElChainReader.GetOperatorDetails(ctx, operator)
		assert.NoError(t, err)
		assert.NotNil(t, operatorDetails)
		assert.Equal(t, operator.Address, operatorDetails.Address)
	})

	t.Run("get strategy and underlying token", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategy, underlyingTokenAddr, err := read_clients.ElChainReader.GetStrategyAndUnderlyingToken(
			ctx,
			strategyAddr,
		)
		assert.NoError(t, err)
		assert.NotNil(t, strategy)
		assert.NotEqual(t, common.Address{}, underlyingTokenAddr)

		erc20Token, err := erc20.NewContractIERC20(underlyingTokenAddr, read_clients.EthHttpClient)
		assert.NoError(t, err)

		tokenName, err := erc20Token.Name(&bind.CallOpts{})
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenName)
	})

	t.Run("get strategy and underlying ERC20 token", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategy, contractUnderlyingToken, underlyingTokenAddr, err := read_clients.ElChainReader.GetStrategyAndUnderlyingERC20Token(
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
		shares, err := read_clients.ElChainReader.GetOperatorSharesInStrategy(
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
		digest, err := read_clients.ElChainReader.CalculateDelegationApprovalDigestHash(
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
		digest, err := read_clients.ElChainReader.CalculateOperatorAVSRegistrationDigestHash(
			ctx,
			common.HexToAddress(operator.Address),
			avs,
			salt,
			expiry,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, digest)
	})

	t.Run("get staker shares", func(t *testing.T) {
		strategies, shares, err := read_clients.ElChainReader.GetStakerShares(
			ctx,
			common.HexToAddress(operator.Address),
		)
		assert.NotZero(t, len(strategies), "Strategies has at least one element")
		assert.NotZero(t, len(shares), "Shares has at least one element")
		assert.Equal(t, len(strategies), len(shares), "Strategies has the same ammount of elements as shares")
		assert.NoError(t, err)
	})

	t.Run("get delegated operator", func(t *testing.T) {
		val := big.NewInt(0)
		address, err := read_clients.ElChainReader.GetDelegatedOperator(
			ctx,
			common.HexToAddress(operator.Address),
			val,
		)

		assert.NoError(t, err)
		// The delegated operator of an operator is the operator itself
		assert.Equal(t, address.String(), operator.Address)
	})

	t.Run("GetOperatorShares", func(t *testing.T) {
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategies := []common.Address{strategyAddr}
		shares, err := read_clients.ElChainReader.GetOperatorShares(
			ctx,
			common.HexToAddress(operator.Address),
			strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 1)

		// with n strategies, response's list length is n
		strategies = []common.Address{strategyAddr, strategyAddr, strategyAddr}
		shares, err = read_clients.ElChainReader.GetOperatorShares(
			ctx,
			common.HexToAddress(operator.Address),
			strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 3)

		// We could test modify the shares and verify the diff is the expected
	})

	t.Run("GetOperatorsShares", func(t *testing.T) {
		operatorAddr := common.HexToAddress(operator.Address)
		operators := []common.Address{operatorAddr}
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategies := []common.Address{strategyAddr}
		shares, err := read_clients.ElChainReader.GetOperatorsShares(
			ctx,
			operators,
			strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 1)

		// with n strategies, response's list length is [1][n]
		mult_strategies := []common.Address{strategyAddr, strategyAddr, strategyAddr}
		shares, err = read_clients.ElChainReader.GetOperatorsShares(
			ctx,
			operators,
			mult_strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 1)
		assert.Len(t, shares[0], 3)

		// with n strategies, response's list length is [n][1]
		mult_operators := []common.Address{operatorAddr, operatorAddr, operatorAddr}
		shares, err = read_clients.ElChainReader.GetOperatorsShares(
			ctx,
			mult_operators,
			strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 3)
		assert.Len(t, shares[0], 1)

		// with n strategies and n operators, response's list length is [n][n]
		shares, err = read_clients.ElChainReader.GetOperatorsShares(
			ctx,
			mult_operators,
			mult_strategies,
		)
		assert.NoError(t, err)
		assert.Len(t, shares, 3)
		assert.Len(t, shares[2], 3)
	})
}

func TestGetCurrentClaimableDistributionRoot(t *testing.T) {
	// Verifies GetCurrentClaimableDistributionRoot returns 0 if no root and the root if there's one
	_, anvilHttpEndpoint := testclients.BuildTestClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

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

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	// Create and configure rewards coordinator
	ethClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)
	rewardsCoordinator, err := rewardscoordinator.NewContractIRewardsCoordinator(rewardsCoordinatorAddr, ethClient)
	require.NoError(t, err)

	ecdsaPrivKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	// Set delay to zero to inmediatly operate with coordinator
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, ecdsaPrivKeyHex, uint32(0))
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// Create txManager to send transactions to the Ethereum node
	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, ecdsaPrivKeyHex)
	require.NoError(t, err)
	noSendTxOpts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	rewardsUpdater := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

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
}

func TestGetRootIndexFromRootHash(t *testing.T) {
	_, anvilHttpEndpoint := testclients.BuildTestClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	// Create and configure rewards coordinator
	ethClient, err := ethclient.Dial(anvilHttpEndpoint)
	require.NoError(t, err)
	rewardsCoordinator, err := rewardscoordinator.NewContractIRewardsCoordinator(rewardsCoordinatorAddr, ethClient)
	require.NoError(t, err)
	ecdsaPrivKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	// Set delay to zero to inmediatly operate with coordinator
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, ecdsaPrivKeyHex, uint32(0))
	require.NoError(t, err)
	require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

	// Create txManager to send transactions to the Ethereum node
	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, ecdsaPrivKeyHex)
	require.NoError(t, err)
	noSendTxOpts, err := txManager.GetNoSendTxOpts()
	require.NoError(t, err)

	rewardsUpdater := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

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
	assert.Equal(t, err.Error(), "execution reverted: custom error 0x504570e3",
		"GetRootIndexFromHash should return an InvalidRoot() error",
	)
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
}

func TestGetCumulativeClaimedRewards(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}
	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	activationDelay := uint32(0)
	// Set activation delay to zero so that the earnings can be claimed right after submitting the root
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.True(t, receipt.Status == gethtypes.ReceiptStatusSuccessful)

	strategyAddr := contractAddrs.Erc20MockStrategy
	strategy, contractUnderlyingToken, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingERC20Token(
		ctx,
		strategyAddr,
	)
	assert.NoError(t, err)
	assert.NotNil(t, strategy)
	assert.NotEqual(t, common.Address{}, underlyingTokenAddr)
	assert.NotNil(t, contractUnderlyingToken)

	anvil_address := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	// This tests that without claims result is zero
	claimed, err := chainReader.GetCumulativeClaimed(ctx, anvil_address, underlyingTokenAddr)
	assert.Zero(t, claimed.Cmp(big.NewInt(0)))
	assert.NoError(t, err)

	cumulativeEarnings := int64(45)
	claim, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings, privateKeyHex)
	require.NoError(t, err)

	receipt, err = chainWriter.ProcessClaim(context.Background(), *claim, rewardsCoordinatorAddr, true)
	require.NoError(t, err)
	require.True(t, receipt.Status == gethtypes.ReceiptStatusSuccessful)

	// This tests that with a claim result is cumulativeEarnings
	claimed, err = chainReader.GetCumulativeClaimed(ctx, anvil_address, underlyingTokenAddr)
	assert.Equal(t, claimed, big.NewInt(cumulativeEarnings))
	assert.NoError(t, err)
}

func TestCheckClaim(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	ctx := context.Background()

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}
	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	// Create ChainWriter and chain reader
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	activationDelay := uint32(0)
	// Set activation delay to zero so that the earnings can be claimed right after submitting the root
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.True(t, receipt.Status == gethtypes.ReceiptStatusSuccessful)

	cumulativeEarnings := int64(45)
	claim, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings, privateKeyHex)
	require.NoError(t, err)

	receipt, err = chainWriter.ProcessClaim(context.Background(), *claim, rewardsCoordinatorAddr, true)
	require.NoError(t, err)
	require.True(t, receipt.Status == gethtypes.ReceiptStatusSuccessful)

	strategyAddr := contractAddrs.Erc20MockStrategy
	strategy, contractUnderlyingToken, underlyingTokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingERC20Token(
		ctx,
		strategyAddr,
	)
	assert.NoError(t, err)
	assert.NotNil(t, strategy)
	assert.NotEqual(t, common.Address{}, underlyingTokenAddr)
	assert.NotNil(t, contractUnderlyingToken)

	checked, err := chainReader.CheckClaim(ctx, *claim)
	require.NoError(t, err)
	assert.True(t, checked)
}

func TestGetAllocatableMagnitudeAndGetMaxMagnitudes(t *testing.T) {
	// Without changes, Allocable magnitude is max magnitude

	// Test setup
	ctx := context.Background()

	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	config := elcontracts.Config{
		DelegationManagerAddress: contractAddrs.DelegationManager,
	}

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	strategyAddr := contractAddrs.Erc20MockStrategy
	testAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	operatorSetId := uint32(1)

	strategies := []common.Address{strategyAddr}
	maxMagnitudes, err := chainReader.GetMaxMagnitudes(ctx, testAddr, strategies)
	assert.NoError(t, err)

	// Assert that at the beginning, Allocatable Magnitude is Max allocatable magnitude
	allocable, err := chainReader.GetAllocatableMagnitude(ctx, testAddr, strategyAddr)
	assert.NoError(t, err)

	assert.Equal(t, maxMagnitudes[0], allocable)

	// Reduce allocatable magnitude for testAddr
	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY

	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	waitForReceipt := true
	delay := uint32(1)
	receipt, err := chainWriter.SetAllocationDelay(context.Background(), operatorAddr, delay, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	allocationConfigurationDelay := 1200
	testutils.AdvanceChainByNBlocksExecInContainer(context.Background(), allocationConfigurationDelay+1, anvilC)

	// Check that Allocation delay has been applied
	_, err = chainReader.GetAllocationDelay(context.Background(), operatorAddr)
	require.NoError(t, err)

	err = createOperatorSet(anvilHttpEndpoint, privateKeyHex, testAddr, operatorSetId, strategyAddr)
	require.NoError(t, err)

	operatorSet := allocationmanager.OperatorSet{
		Avs: testAddr,
		Id:  operatorSetId,
	}
	allocatable_reduction := uint64(100)
	allocateParams := []allocationmanager.IAllocationManagerTypesAllocateParams{
		{
			OperatorSet:   operatorSet,
			Strategies:    []common.Address{strategyAddr},
			NewMagnitudes: []uint64{allocatable_reduction},
		},
	}

	receipt, err = chainWriter.ModifyAllocations(context.Background(), operatorAddr, allocateParams, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Assert that after slashing, Allocatable Magnitude + slash ammount equals Max allocatable magnitude
	allocable, err = chainReader.GetAllocatableMagnitude(ctx, testAddr, strategyAddr)
	assert.NoError(t, err)

	assert.Equal(t, maxMagnitudes[0], allocable+allocatable_reduction)
}

func TestAdminFunctions(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	assert.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	assert.NoError(t, err)

	permissionControllerAddr := common.HexToAddress(testutils.PERMISSION_CONTROLLER_ADDRESS)
	config := elcontracts.Config{
		PermissionsControllerAddress: permissionControllerAddr,
	}

	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	accountChainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	assert.NoError(t, err)

	pendingAdminAddr := common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
	pendingAdminPrivateKeyHex := testutils.ANVIL_SECOND_PRIVATE_KEY
	adminChainWriter, err := testclients.NewTestChainWriterFromConfig(
		anvilHttpEndpoint,
		pendingAdminPrivateKeyHex,
		config,
	)
	assert.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	assert.NoError(t, err)

	t.Run("non-existent pending admin", func(t *testing.T) {
		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdminAddr)
		assert.NoError(t, err)
		assert.False(t, isPendingAdmin)
	})

	t.Run("list pending admins when empty", func(t *testing.T) {
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

		receipt, err := accountChainWriter.AddPendingAdmin(context.Background(), request)
		assert.NoError(t, err)
		assert.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdminAddr)
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

	t.Run("list admins", func(t *testing.T) {
		acceptAdminRequest := elcontracts.AcceptAdminRequest{
			AccountAddress: operatorAddr,
			WaitForReceipt: true,
		}

		receipt, err := adminChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest)
		assert.NoError(t, err)
		assert.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

		listAdmins, err := chainReader.ListAdmins(context.Background(), operatorAddr)
		assert.NoError(t, err)
		assert.Len(t, listAdmins, 1)

		admin := listAdmins[0]
		isAdmin, err := chainReader.IsAdmin(context.Background(), operatorAddr, admin)
		assert.NoError(t, err)
		assert.True(t, isAdmin)
	})
}

func TestAppointeesFunctions(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	assert.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	assert.NoError(t, err)

	permissionControllerAddr := common.HexToAddress(testutils.PERMISSION_CONTROLLER_ADDRESS)
	config := elcontracts.Config{
		PermissionsControllerAddress: permissionControllerAddr,
	}

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	assert.NoError(t, err)

	privateKey := testutils.ANVIL_FIRST_PRIVATE_KEY
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKey, config)
	assert.NoError(t, err)

	accountAddress := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)

	appointeeAddress := common.HexToAddress(testutils.ANVIL_SECOND_ADDRESS)
	target := common.HexToAddress(testutils.ANVIL_THIRD_ADDRESS)
	selector := [4]byte{0, 1, 2, 3}

	t.Run("list appointees when empty", func(t *testing.T) {
		appointees, err := chainReader.ListAppointees(context.Background(), accountAddress, target, selector)
		assert.NoError(t, err)
		assert.Empty(t, appointees)
	})

	t.Run("list appointees", func(t *testing.T) {
		setPermissionRequest := elcontracts.SetPermissionRequest{
			AccountAddress:   accountAddress,
			AppointeeAddress: appointeeAddress,
			Target:           target,
			Selector:         selector,
			WaitForReceipt:   true,
		}

		receipt, err := chainWriter.SetPermission(context.Background(), setPermissionRequest)
		require.NoError(t, err)
		require.Equal(t, receipt.Status, gethtypes.ReceiptStatusSuccessful)

		canCall, err := chainReader.CanCall(context.Background(), accountAddress, appointeeAddress, target, selector)
		require.NoError(t, err)
		require.True(t, canCall)

		appointees, err := chainReader.ListAppointees(context.Background(), accountAddress, target, selector)
		assert.NoError(t, err)
		assert.NotEmpty(t, appointees)
	})

	t.Run("list appointees permissions", func(t *testing.T) {
		appointeesPermission, _, err := chainReader.ListAppointeePermissions(
			context.Background(),
			accountAddress,
			appointeeAddress,
		)
		assert.NoError(t, err)
		assert.NotEmpty(t, appointeesPermission)
	})
}
