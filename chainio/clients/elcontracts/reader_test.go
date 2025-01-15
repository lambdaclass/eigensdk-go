package elcontracts_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	mockerc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockERC20"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

	t.Run("get staker shares", func(t *testing.T) {
		strategies, shares, err := clients.ElChainReader.GetStakerShares(
			ctx,
			common.HexToAddress(operator.Address),
		)
		assert.NotZero(t, len(strategies))            // Strategies has at least one element
		assert.NotZero(t, len(shares))                // Shares has at least one element
		assert.Equal(t, len(strategies), len(shares)) // Strategies has the same ammount of elements as shares
		assert.NoError(t, err)
	})

	t.Run("get delegated operator", func(t *testing.T) {
		val := big.NewInt(0)
		address, err := clients.ElChainReader.GetDelegatedOperator(
			ctx,
			common.HexToAddress(operator.Address),
			val,
		)

		assert.NoError(t, err)
		// The delegated operator of an operator is the operator itself
		assert.Equal(t, address.String(), operator.Address)
	})

	t.Run("get Allocatable Magnitude", func(t *testing.T) {
		// Without changes, Allocable magnitude is max magnitude

		rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
		config := elcontracts.Config{
			DelegationManagerAddress:  contractAddrs.DelegationManager,
			RewardsCoordinatorAddress: rewardsCoordinatorAddr,
		}

		chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
		require.NoError(t, err)

		strategyAddr := contractAddrs.Erc20MockStrategy

		strategies := []common.Address{strategyAddr}
		maxmagnitude, err := chainReader.GetMaxMagnitudes(ctx, common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS), strategies)
		assert.NoError(t, err)

		allocable, err := chainReader.GetAllocatableMagnitude(ctx, common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS), strategyAddr)
		assert.NoError(t, err)

		assert.Equal(t, maxmagnitude[0], allocable)
	})

	t.Run("GetOperatorShares", func(t *testing.T) {
		// Maybe could add more strategies
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategies := []common.Address{strategyAddr}
		shares, err := clients.ElChainReader.GetOperatorShares(
			ctx,
			common.HexToAddress(operator.Address),
			strategies,
		)
		assert.NoError(t, err)
		assert.NotZero(t, shares)
	})

	t.Run("GetOperatorsShares", func(t *testing.T) {
		// Maybe could add more operators and strategies
		operatorAddrs := []common.Address{common.HexToAddress(operator.Address)}
		strategyAddr := contractAddrs.Erc20MockStrategy
		strategies := []common.Address{strategyAddr}
		shares, err := clients.ElChainReader.GetOperatorsShares(
			ctx,
			operatorAddrs,
			strategies,
		)
		assert.NoError(t, err)
		assert.NotZero(t, shares)
	})

	t.Run("GetOperatorSetsForOperator", func(t *testing.T) {
		// GetOperatorSetsForOperator with an operator without sets returns an empty list
		operatorSet, err := clients.ElChainReader.GetOperatorSetsForOperator(
			ctx,
			common.HexToAddress(operator.Address),
		)
		assert.NoError(t, err)
		assert.Empty(t, operatorSet)
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
	require.Equal(t, receipt.Status, uint64(1))

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
	require.Equal(t, receipt.Status, uint64(1))

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
	require.True(t, receipt.Status == uint64(1))

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
	assert.True(t, claimed.Cmp(big.NewInt(0)) == 0)
	assert.NoError(t, err)

	cumulativeEarnings := int64(45)
	claim, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings, privateKeyHex)
	require.NoError(t, err)

	earner := common.HexToAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")
	receipt, err = chainWriter.ProcessClaim(context.Background(), *claim, earner, true)
	require.NoError(t, err)
	require.True(t, receipt.Status == uint64(1))

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
	require.True(t, receipt.Status == uint64(1))

	cumulativeEarnings := int64(45)
	claim, err := newTestClaim(chainReader, anvilHttpEndpoint, cumulativeEarnings, privateKeyHex)
	require.NoError(t, err)

	earner := common.HexToAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")
	receipt, err = chainWriter.ProcessClaim(context.Background(), *claim, earner, true)
	require.NoError(t, err)
	require.True(t, receipt.Status == uint64(1))

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

// The functions below will be replaced for those placed in testutils/testclients/testclients.go

// Returns a (test) claim for the given cumulativeEarnings, whose earner is
// the account given by the testutils.ANVIL_FIRST_ADDRESS address.
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
