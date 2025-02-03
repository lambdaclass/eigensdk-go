package avsregistry

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	apkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
)

// DefaultQueryBlockRange different node providers have different eth_getLogs range limits.
// 10k is an arbitrary choice that should work for most
var DefaultQueryBlockRange = big.NewInt(10_000)

type Config struct {
	RegistryCoordinatorAddress    common.Address
	OperatorStateRetrieverAddress common.Address
}

// The ChainReader provides methods to call the
// AVS registry contract's view functions.
type ChainReader struct {
	logger                  logging.Logger
	blsApkRegistryAddr      common.Address
	registryCoordinatorAddr common.Address
	registryCoordinator     *regcoord.ContractRegistryCoordinator
	operatorStateRetriever  *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry           *stakeregistry.ContractStakeRegistry
	ethClient               eth.HttpBackend
}

// Creates a new instance  of the ChainReader.
func NewChainReader(
	registryCoordinatorAddr common.Address,
	blsApkRegistryAddr common.Address,
	registryCoordinator *regcoord.ContractRegistryCoordinator,
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	logger logging.Logger,
	ethClient eth.HttpBackend,
) *ChainReader {
	logger = logger.With(logging.ComponentKey, "avsregistry/ChainReader")

	return &ChainReader{
		blsApkRegistryAddr:      blsApkRegistryAddr,
		registryCoordinatorAddr: registryCoordinatorAddr,
		registryCoordinator:     registryCoordinator,
		operatorStateRetriever:  operatorStateRetriever,
		stakeRegistry:           stakeRegistry,
		logger:                  logger,
		ethClient:               ethClient,
	}
}

// NewReaderFromConfig creates a new ChainReader
func NewReaderFromConfig(
	cfg Config,
	client eth.HttpBackend,
	logger logging.Logger,
) (*ChainReader, error) {
	bindings, err := NewBindingsFromConfig(cfg, client, logger)
	if err != nil {
		wrappedError := elcontracts.CreateForNestedError("NewBindingsFromConfig", err)
		return nil, wrappedError
	}

	return NewChainReader(
		bindings.RegistryCoordinatorAddr,
		bindings.BlsApkRegistryAddr,
		bindings.RegistryCoordinator,
		bindings.OperatorStateRetriever,
		bindings.StakeRegistry,
		logger,
		client,
	), nil
}

// Returns the total quorum count read from the RegistryCoordinator
func (r *ChainReader) GetQuorumCount(opts *bind.CallOpts) (uint8, error) {
	if r.registryCoordinator == nil {
		wrappedError := elcontracts.CreateErrorForMissingContract("RegistryCoordinator")
		return 0, wrappedError
	}

	cont, err := r.registryCoordinator.QuorumCount(opts)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("registryCoordinator.QuorumCount", err)
		return 0, wrappedError
	}

	return cont, nil
}

// Returns, for each quorum in `quorumNumbers`, a vector of the operators registered for
// that quorum at the current block, containing each operator's `operatorId` and `stake`.
func (r *ChainReader) GetOperatorsStakeInQuorumsAtCurrentBlock(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
) ([][]opstateretriever.OperatorStateRetrieverOperator, error) {
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	curBlock, err := r.ethClient.BlockNumber(opts.Context)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("ethClient.BlockNumber", err)
		return nil, wrappedError
	}
	if curBlock > math.MaxUint32 {
		wrappedError := elcontracts.CreateForOtherError("Current block number is too large to fit into an uint32", err)
		return nil, wrappedError
	}

	operatorStakes, err := r.GetOperatorsStakeInQuorumsAtBlock(opts, quorumNumbers, uint32(curBlock))
	if err != nil {
		wrappedError := elcontracts.CreateForNestedError("GetOperatorsStakeInQuorumsAtBlock", err)
		return nil, wrappedError
	}

	return operatorStakes, nil
}

// the contract stores historical state, so blockNumber should be the block number of the state you want to query
// and the blockNumber in opts should be the block number of the latest block (or set to nil, which is equivalent)
func (r *ChainReader) GetOperatorsStakeInQuorumsAtBlock(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
	blockNumber uint32,
) ([][]opstateretriever.OperatorStateRetrieverOperator, error) {
	if r.operatorStateRetriever == nil {
		wrappedError := elcontracts.CreateErrorForMissingContract("OperatorStateRetriever")
		return nil, wrappedError
	}

	operatorStakes, err := r.operatorStateRetriever.GetOperatorState(
		opts,
		r.registryCoordinatorAddr,
		quorumNumbers.UnderlyingType(),
		blockNumber)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("operatorStateRetriever.GetOperatorState", err)
		return nil, wrappedError
	}
	return operatorStakes, nil
}

// Returns, for each quorum in `quorumNumbers`, a vector of the addresses of the
// operators registered for that quorum at the current block.
func (r *ChainReader) GetOperatorAddrsInQuorumsAtCurrentBlock(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
) ([][]common.Address, error) {
	if r.operatorStateRetriever == nil {
		wrappedError := elcontracts.CreateErrorForMissingContract("OperatorStateRetriever")
		return nil, wrappedError
	}
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	curBlock, err := r.ethClient.BlockNumber(opts.Context)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("ethClient.BlockNumber", err)
		return nil, wrappedError
	}
	if curBlock > math.MaxUint32 {
		wrappedError := elcontracts.CreateForOtherError("Current block number is too large to fit into an uint32", err)
		return nil, wrappedError
	}
	operatorStakes, err := r.operatorStateRetriever.GetOperatorState(
		opts,
		r.registryCoordinatorAddr,
		quorumNumbers.UnderlyingType(),
		uint32(curBlock),
	)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("operatorStateRetriever.GetOperatorState", err)
		return nil, wrappedError
	}
	var quorumOperatorAddrs [][]common.Address
	for _, quorum := range operatorStakes {
		var operatorAddrs []common.Address
		for _, operator := range quorum {
			operatorAddrs = append(operatorAddrs, operator.Operator)
		}
		quorumOperatorAddrs = append(quorumOperatorAddrs, operatorAddrs)
	}
	return quorumOperatorAddrs, nil
}

// Returns a tuple containing
//   - An array with the quorum IDs in which the given operator is registered at the given block
//   - An array that contains, for each quorum, an array with the address, id and stake
//     of each operator registered in that quorum.
func (r *ChainReader) GetOperatorsStakeInQuorumsOfOperatorAtBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
	blockNumber uint32,
) (types.QuorumNums, [][]opstateretriever.OperatorStateRetrieverOperator, error) {
	if r.operatorStateRetriever == nil {
		wrappedError := elcontracts.CreateErrorForMissingContract("OperatorStateRetriever")
		return nil, nil, wrappedError
	}

	quorumBitmap, operatorStakes, err := r.operatorStateRetriever.GetOperatorState0(
		opts,
		r.registryCoordinatorAddr,
		operatorId,
		blockNumber)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("operatorStateRetriever.GetOperatorState0", err)
		return nil, nil, wrappedError
	}
	quorums := types.BitmapToQuorumIds(quorumBitmap)
	return quorums, operatorStakes, nil
}

// opts will be modified to have the latest blockNumber, so make sure not to reuse it
// blockNumber in opts will be ignored, and the chain will be queried to get the latest blockNumber
func (r *ChainReader) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (types.QuorumNums, [][]opstateretriever.OperatorStateRetrieverOperator, error) {
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	curBlock, err := r.ethClient.BlockNumber(opts.Context)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("ethClient.BlockNumber", err)
		return nil, nil, wrappedError
	}
	if curBlock > math.MaxUint32 {
		wrappedError := elcontracts.CreateForOtherError("Current block number is too large to fit into an uint32", err)
		return nil, nil, wrappedError
	}
	opts.BlockNumber = big.NewInt(int64(curBlock))
	quorums, operatorsStake, err := r.GetOperatorsStakeInQuorumsOfOperatorAtBlock(opts, operatorId, uint32(curBlock))
	if err != nil {
		wrappedError := elcontracts.CreateForNestedError("GetOperatorsStakeInQuorumsOfOperatorAtBlock", err)
		return nil, nil, wrappedError
	}

	return quorums, operatorsStake, nil
}

// To avoid a possible race condition, this method must assure that all the calls
// are made with the same blockNumber.
// So, if the blockNumber and blockHash are not set in opts, blockNumber will be set
// to the latest block.
// All calls to the chain use `opts` parameter.
func (r *ChainReader) GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (map[types.QuorumNum]types.StakeAmount, error) {
	if r.registryCoordinator == nil {
		wrappedError := elcontracts.CreateErrorForMissingContract("RegistryCoordinator")
		return nil, wrappedError
	}

	if r.stakeRegistry == nil {
		wrappedError := elcontracts.CreateErrorForMissingContract("StakeRegistry")
		return nil, wrappedError
	}

	// check if opts parameter has not a block number set (BlockNumber)
	var defaultHash common.Hash
	if opts.BlockNumber == nil && opts.BlockHash == defaultHash {
		// if not, set the block number to the latest block
		if opts.Context == nil {
			opts.Context = context.Background()
		}
		latestBlock, err := r.ethClient.BlockNumber(opts.Context)
		if err != nil {
			wrappedError := elcontracts.CreateForBindingError("ethClient.BlockNumber", err)
			return nil, wrappedError
		}
		opts.BlockNumber = big.NewInt(int64(latestBlock))
	}

	quorumBitmap, err := r.registryCoordinator.GetCurrentQuorumBitmap(opts, operatorId)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("registryCoordinator.GetCurrentQuorumBitmap", err)
		return nil, wrappedError
	}
	quorums := types.BitmapToQuorumIds(quorumBitmap)
	quorumStakes := make(map[types.QuorumNum]types.StakeAmount)
	for _, quorum := range quorums {
		stake, err := r.stakeRegistry.GetCurrentStake(
			opts,
			operatorId,
			uint8(quorum),
		)
		if err != nil {
			wrappedError := elcontracts.CreateForBindingError("stakeRegistry.GetCurrentStake", err)
			return nil, wrappedError
		}
		quorumStakes[quorum] = stake
	}
	return quorumStakes, nil
}

// Returns a struct containing the indices of the quorum members that signed,
// and the ones that didn't.
func (r *ChainReader) GetCheckSignaturesIndices(
	opts *bind.CallOpts,
	referenceBlockNumber uint32,
	quorumNumbers types.QuorumNums,
	nonSignerOperatorIds []types.OperatorId,
) (opstateretriever.OperatorStateRetrieverCheckSignaturesIndices, error) {
	if r.operatorStateRetriever == nil {
		wrappedError := elcontracts.CreateErrorForMissingContract("OperatorStateRetriever")
		return opstateretriever.OperatorStateRetrieverCheckSignaturesIndices{}, wrappedError
	}

	nonSignerOperatorIdsBytes := make([][32]byte, len(nonSignerOperatorIds))
	for i, id := range nonSignerOperatorIds {
		nonSignerOperatorIdsBytes[i] = id
	}
	checkSignatureIndices, err := r.operatorStateRetriever.GetCheckSignaturesIndices(
		opts,
		r.registryCoordinatorAddr,
		referenceBlockNumber,
		quorumNumbers.UnderlyingType(),
		nonSignerOperatorIdsBytes,
	)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("operatorStateRetriever.GetCheckSignaturesIndices", err)
		return opstateretriever.OperatorStateRetrieverCheckSignaturesIndices{}, wrappedError
	}
	return checkSignatureIndices, nil
}

// Given an operator address, returns its ID.
func (r *ChainReader) GetOperatorId(
	opts *bind.CallOpts,
	operatorAddress common.Address,
) ([32]byte, error) {
	if r.registryCoordinator == nil {
		wrappedError := elcontracts.CreateErrorForMissingContract("RegistryCoordinator")
		return [32]byte{}, wrappedError
	}

	operatorId, err := r.registryCoordinator.GetOperatorId(
		opts,
		operatorAddress,
	)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("registryCoordinator.GetOperatorId", err)
		return [32]byte{}, wrappedError
	}
	return operatorId, nil
}

// Given an operator ID, returns its address.
func (r *ChainReader) GetOperatorFromId(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (common.Address, error) {
	if r.registryCoordinator == nil {
		wrappedError := elcontracts.CreateErrorForMissingContract("RegistryCoordinator")
		return common.Address{}, wrappedError
	}

	operatorAddress, err := r.registryCoordinator.GetOperatorFromId(
		opts,
		operatorId,
	)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("registryCoordinator.GetOperatorFromId", err)
		return common.Address{}, wrappedError
	}
	return operatorAddress, nil
}

// Returns an array of booleans, where the boolean at index i represents
// whether the operator is registered for the quorum i.
func (r *ChainReader) QueryRegistrationDetail(
	opts *bind.CallOpts,
	operatorAddress common.Address,
) ([]bool, error) {
	operatorId, err := r.GetOperatorId(opts, operatorAddress)
	if err != nil {
		wrappedError := elcontracts.CreateForNestedError("GetOperatorId", err)
		return nil, wrappedError
	}
	value, err := r.registryCoordinator.GetCurrentQuorumBitmap(opts, operatorId)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("registryCoordinator.GetCurrentQuorumBitmap", err)
		return nil, wrappedError
	}
	numBits := value.BitLen()
	var quorums []bool
	for i := 0; i < numBits; i++ {
		quorums = append(quorums, value.Int64()&(1<<i) != 0)
	}
	if len(quorums) == 0 {
		numQuorums, err := r.GetQuorumCount(opts)
		if err != nil {
			wrappedError := elcontracts.CreateForNestedError("GetQuorumCount", err)
			return nil, wrappedError
		}
		for i := uint8(0); i < numQuorums; i++ {
			quorums = append(quorums, false)
		}
	}
	return quorums, nil
}

// Returns true if the operator is registered, false otherwise.
func (r *ChainReader) IsOperatorRegistered(
	opts *bind.CallOpts,
	operatorAddress common.Address,
) (bool, error) {
	if r.registryCoordinator == nil {
		wrappedError := elcontracts.CreateErrorForMissingContract("RegistryCoordinator")
		return false, wrappedError
	}

	operatorStatus, err := r.registryCoordinator.GetOperatorStatus(opts, operatorAddress)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("registryCoordinator.GetOperatorStatus", err)
		return false, wrappedError
	}

	// 0 = NEVER_REGISTERED, 1 = REGISTERED, 2 = DEREGISTERED
	registeredWithAvs := operatorStatus == 1
	return registeredWithAvs, nil
}

// Queries existing operators for a particular block range.
// Returns two arrays. The first one contains the addresses
// of the operators, and the second contains their corresponding public keys.
func (r *ChainReader) QueryExistingRegisteredOperatorPubKeys(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
	blockRange *big.Int,
) ([]types.OperatorAddr, []types.OperatorPubkeys, error) {
	blsApkRegistryAbi, err := apkreg.ContractBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		wrappedError := elcontracts.CreateForOtherError("Failed to get bls apk registry ABI", err)
		return nil, nil, wrappedError
	}

	if startBlock == nil {
		startBlock = big.NewInt(0)
	}
	if stopBlock == nil {
		curBlockNum, err := r.ethClient.BlockNumber(ctx)
		if err != nil {
			wrappedError := elcontracts.CreateForBindingError("ethClient.BlockNumber", err)
			return nil, nil, wrappedError
		}
		stopBlock = new(big.Int).SetUint64(curBlockNum)
	}
	if blockRange == nil {
		blockRange = DefaultQueryBlockRange
	}

	operatorAddresses := make([]types.OperatorAddr, 0)
	operatorPubkeys := make([]types.OperatorPubkeys, 0)
	// QueryExistingRegisteredOperatorPubKeys and QueryExistingRegisteredOperatorSockets
	// both run in parallel and they read and mutate the same variable startBlock,
	// so we clone it to prevent the race condition.
	// TODO: we might want to eventually change the function signature to pass a uint,
	// but that would be a breaking change
	for i := new(big.Int).Set(startBlock); i.Cmp(stopBlock) <= 0; i.Add(i, blockRange) {
		// Subtract 1 since FilterQuery is inclusive
		toBlock := big.NewInt(0).Add(i, big.NewInt(0).Sub(blockRange, big.NewInt(1)))
		if toBlock.Cmp(stopBlock) > 0 {
			toBlock = stopBlock
		}
		query := ethereum.FilterQuery{
			FromBlock: i,
			ToBlock:   toBlock,
			Addresses: []common.Address{
				r.blsApkRegistryAddr,
			},
			Topics: [][]common.Hash{{blsApkRegistryAbi.Events["NewPubkeyRegistration"].ID}},
		}

		logs, err := r.ethClient.FilterLogs(ctx, query)
		if err != nil {
			wrappedError := elcontracts.CreateForBindingError("ethClient.FilterLogs", err)
			return nil, nil, wrappedError
		}
		r.logger.Debug(
			"avsRegistryChainReader.QueryExistingRegisteredOperatorPubKeys",
			"numTransactionLogs",
			len(logs),
			"fromBlock",
			i,
			"toBlock",
			toBlock,
		)

		for _, vLog := range logs {
			// get the operator address
			operatorAddr := common.HexToAddress(vLog.Topics[1].Hex())
			operatorAddresses = append(operatorAddresses, operatorAddr)

			event, err := blsApkRegistryAbi.Unpack("NewPubkeyRegistration", vLog.Data)
			if err != nil {
				wrappedError := elcontracts.CreateForOtherError("Failed to unpack event data", err)
				return nil, nil, wrappedError
			}

			G1Pubkey := event[0].(struct {
				X *big.Int "json:\"X\""
				Y *big.Int "json:\"Y\""
			})

			G2Pubkey := event[1].(struct {
				X [2]*big.Int "json:\"X\""
				Y [2]*big.Int "json:\"Y\""
			})

			operatorPubkey := types.OperatorPubkeys{
				G1Pubkey: bls.NewG1Point(
					G1Pubkey.X,
					G1Pubkey.Y,
				),
				G2Pubkey: bls.NewG2Point(
					G2Pubkey.X,
					G2Pubkey.Y,
				),
			}

			operatorPubkeys = append(operatorPubkeys, operatorPubkey)
		}
	}

	return operatorAddresses, operatorPubkeys, nil
}

// Queries existing operator sockets for a particular block range.
// Returns a mapping containing operator IDs as keys and their
// corresponding sockets as values.
func (r *ChainReader) QueryExistingRegisteredOperatorSockets(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
	blockRange *big.Int,
) (map[types.OperatorId]types.Socket, error) {
	if r.registryCoordinator == nil {
		wrappedError := elcontracts.CreateErrorForMissingContract("RegistryCoordinator")
		return nil, wrappedError
	}

	if startBlock == nil {
		startBlock = big.NewInt(0)
	}
	if stopBlock == nil {
		curBlockNum, err := r.ethClient.BlockNumber(ctx)
		if err != nil {
			wrappedError := elcontracts.CreateForBindingError("ethClient.BlockNumber", err)
			return nil, wrappedError
		}
		stopBlock = new(big.Int).SetUint64(curBlockNum)
	}
	if blockRange == nil {
		blockRange = DefaultQueryBlockRange
	}

	operatorIdToSocketMap := make(map[types.OperatorId]types.Socket)
	// QueryExistingRegisteredOperatorPubKeys and QueryExistingRegisteredOperatorSockets
	// both run in parallel and they read and mutate the same variable startBlock,
	// so we clone it to prevent the race condition.
	// TODO: we might want to eventually change the function signature to pass a uint,
	// but that would be a breaking change
	for i := new(big.Int).Set(startBlock); i.Cmp(stopBlock) <= 0; i.Add(i, blockRange) {
		// Subtract 1 since FilterQuery is inclusive
		toBlock := big.NewInt(0).Add(i, big.NewInt(0).Sub(blockRange, big.NewInt(1)))
		if toBlock.Cmp(stopBlock) > 0 {
			toBlock = stopBlock
		}

		end := toBlock.Uint64()

		filterOpts := &bind.FilterOpts{
			Start: i.Uint64(),
			End:   &end,
		}
		socketUpdates, err := r.registryCoordinator.FilterOperatorSocketUpdate(filterOpts, nil)
		if err != nil {
			wrappedError := elcontracts.CreateForBindingError("registryCoordinator.FilterOperatorSocketUpdate", err)
			return nil, wrappedError
		}

		numSocketUpdates := 0
		for socketUpdates.Next() {
			operatorIdToSocketMap[socketUpdates.Event.OperatorId] = types.Socket(socketUpdates.Event.Socket)
			numSocketUpdates++
		}
		r.logger.Debug(
			"avsRegistryChainReader.QueryExistingRegisteredOperatorSockets",
			"numTransactionLogs",
			numSocketUpdates,
			"fromBlock",
			i,
			"toBlock",
			toBlock,
		)
	}
	return operatorIdToSocketMap, nil
}
