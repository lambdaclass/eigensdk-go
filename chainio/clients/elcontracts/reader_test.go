package elcontracts_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
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
