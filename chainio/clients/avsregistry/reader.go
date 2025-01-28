package avsregistry

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	apkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

// DefaultQueryBlockRange different node providers have different eth_getLogs range limits.
// 10k is an arbitrary choice that should work for most
var DefaultQueryBlockRange uint64 = 10_000

type Config struct {
	RegistryCoordinatorAddress    common.Address
	OperatorStateRetrieverAddress common.Address
}

type ChainReader struct {
	logger                  logging.Logger
	blsApkRegistryAddr      common.Address
	registryCoordinatorAddr common.Address
	registryCoordinator     *regcoord.ContractRegistryCoordinator
	operatorStateRetriever  *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry           *stakeregistry.ContractStakeRegistry
	ethClient               eth.HttpBackend
}

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
		return nil, err
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

// GetQuorumCount returns the total number of quorums the registry coordinator has created
func (r *ChainReader) GetQuorumCount(ctx context.Context, request QuorumCountRequest) (QuorumCountResponse, error) {
	if r.registryCoordinator == nil {
		return QuorumCountResponse{}, errors.New("RegistryCoordinator contract not provided")
	}

	quorumCount, err := r.registryCoordinator.QuorumCount(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
	)
	if err != nil {
		return QuorumCountResponse{}, utils.WrapError("Failed to get quorum count", err)
	}

	return QuorumCountResponse{QuorumCount: quorumCount}, nil
}

// GetOperatorsStakeInQuorumsAtCurrentBlock returns the list of operators and their stakes for the specified quorums at
// the current block
func (r *ChainReader) GetOperatorsStakeInQuorumsAtCurrentBlock(
	ctx context.Context,
	request OperatorsStakeInQuorumAtCurrentBlockRequest,
) (OperatorsStakeInQuorumResponse, error) {
	if r.operatorStateRetriever == nil {
		return OperatorsStakeInQuorumResponse{}, errors.New("OperatorStateRetriever contract not provided")
	}

	curBlock, err := r.ethClient.BlockNumber(ctx)
	if err != nil {
		return OperatorsStakeInQuorumResponse{}, utils.WrapError("Cannot get current block number", err)
	}

	operatorStakes, err := r.operatorStateRetriever.GetOperatorState(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		r.registryCoordinatorAddr,
		request.QuorumNumbers.UnderlyingType(),
		uint32(curBlock),
	)
	if err != nil {
		return OperatorsStakeInQuorumResponse{}, utils.WrapError("Cannot get operators stake", err)
	}

	return OperatorsStakeInQuorumResponse{OperatorsStakeInQuorum: operatorStakes}, nil
}

// GetOperatorsStakeInQuorumsAtBlock returns the list of operators and their stakes for the specified quorums at a
// historical block
func (r *ChainReader) GetOperatorsStakeInQuorumsAtBlock(
	ctx context.Context,
	request OperatorsStakeInQuorumAtBlockRequest,
) (OperatorsStakeInQuorumResponse, error) {
	if r.operatorStateRetriever == nil {
		return OperatorsStakeInQuorumResponse{}, errors.New("OperatorStateRetriever contract not provided")
	}

	operatorStakes, err := r.operatorStateRetriever.GetOperatorState(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		r.registryCoordinatorAddr,
		request.QuorumNumbers.UnderlyingType(),
		request.HistoricalBlockNumber)
	if err != nil {
		return OperatorsStakeInQuorumResponse{}, utils.WrapError("Failed to get operators state", err)
	}
	return OperatorsStakeInQuorumResponse{OperatorsStakeInQuorum: operatorStakes}, nil
}

// GetOperatorAddrsInQuorumsAtCurrentBlock returns the addresses of operators registered in the specified quorums at the
// current block
func (r *ChainReader) GetOperatorAddrsInQuorumsAtCurrentBlock(
	ctx context.Context,
	request OperatorAddrsInQuorumsAtCurrentBlockRequest,
) (OperatorAddrsInQuorumsAtCurrentBlockResponse, error) {
	if r.operatorStateRetriever == nil {
		return OperatorAddrsInQuorumsAtCurrentBlockResponse{}, errors.New(
			"OperatorStateRetriever contract not provided",
		)
	}

	curBlock, err := r.ethClient.BlockNumber(ctx)
	if err != nil {
		return OperatorAddrsInQuorumsAtCurrentBlockResponse{}, utils.WrapError(
			"Failed to get current block number",
			err,
		)
	}

	operatorStakes, err := r.operatorStateRetriever.GetOperatorState(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		r.registryCoordinatorAddr,
		request.QuorumNumbers.UnderlyingType(),
		uint32(curBlock),
	)
	if err != nil {
		return OperatorAddrsInQuorumsAtCurrentBlockResponse{}, utils.WrapError("Failed to get operators state", err)
	}
	var quorumOperatorAddrs [][]common.Address
	for _, quorum := range operatorStakes {
		var operatorAddrs []common.Address
		for _, operator := range quorum {
			operatorAddrs = append(operatorAddrs, operator.Operator)
		}
		quorumOperatorAddrs = append(quorumOperatorAddrs, operatorAddrs)
	}
	return OperatorAddrsInQuorumsAtCurrentBlockResponse{OperatorAddrsInQuorums: quorumOperatorAddrs}, nil

}

// GetOperatorsStakeInQuorumsOfOperatorAtBlock returns the stakes of a specific operator in the quorums at a historical
// block
func (r *ChainReader) GetOperatorsStakeInQuorumsOfOperatorAtBlock(
	ctx context.Context,
	request OperatorsStakeInQuorumsOfOperatorAtBlockRequest,
) (OperatorsStakeInQuorumsOfOperatorResponse, error) {
	if r.operatorStateRetriever == nil {
		return OperatorsStakeInQuorumsOfOperatorResponse{}, errors.New(
			"OperatorStateRetriever contract not provided",
		)
	}

	quorumBitmap, operatorStakes, err := r.operatorStateRetriever.GetOperatorState0(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		r.registryCoordinatorAddr,
		request.OperatorId,
		request.HistoricalBlockNumber)
	if err != nil {
		return OperatorsStakeInQuorumsOfOperatorResponse{}, utils.WrapError("Failed to get operators state", err)
	}

	quorums := types.BitmapToQuorumIds(quorumBitmap)
	return OperatorsStakeInQuorumsOfOperatorResponse{
		QuorumNumbers:           quorums,
		OperatorsStakesInQuorum: operatorStakes,
	}, nil
}

// GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock returns the stakes of a specific operator in the quorums at the
// current block
func (r *ChainReader) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(
	ctx context.Context,
	request OperatorsStakeInQuorumsOfOperatorAtCurrentBlockRequest,
) (OperatorsStakeInQuorumsOfOperatorResponse, error) {
	curBlock, err := r.ethClient.BlockNumber(ctx)
	if err != nil {
		return OperatorsStakeInQuorumsOfOperatorResponse{}, utils.WrapError("Failed to get current block number", err)
	}

	return r.GetOperatorsStakeInQuorumsOfOperatorAtBlock(ctx, OperatorsStakeInQuorumsOfOperatorAtBlockRequest{
		BlockNumber:           request.BlockNumber,
		HistoricalBlockNumber: uint32(curBlock),
		OperatorId:            request.OperatorId,
	})
}

// GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock returns the current stake of an operator in the quorums where they
// are registered
func (r *ChainReader) GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(
	ctx context.Context,
	request OperatorStakeInQuorumsOfOperatorAtCurrentBlockRequest,
) (OperatorStakeInQuorumsOfOperatorResponse, error) {
	if r.registryCoordinator == nil {
		return OperatorStakeInQuorumsOfOperatorResponse{}, errors.New(
			"registryCoordinator contract not provided",
		)
	}
	if r.stakeRegistry == nil {
		return OperatorStakeInQuorumsOfOperatorResponse{}, errors.New(
			"stakeRegistry contract not provided",
		)
	}

	if request.BlockNumber == nil {
		latestBlock, err := r.ethClient.BlockNumber(ctx)
		if err != nil {
			return OperatorStakeInQuorumsOfOperatorResponse{},
				utils.WrapError("failed to get latest block number", err)
		}

		request.BlockNumber = big.NewInt(int64(latestBlock))
	}

	callOpts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: request.BlockNumber,
	}

	quorumBitmap, err := r.registryCoordinator.GetCurrentQuorumBitmap(callOpts, request.OperatorId)
	if err != nil {
		return OperatorStakeInQuorumsOfOperatorResponse{},
			utils.WrapError("failed to get operator quorums", err)
	}
	quorums := types.BitmapToQuorumIds(quorumBitmap)

	quorumStakes := make(map[types.QuorumNum]types.StakeAmount)
	for _, quorum := range quorums {
		stake, err := r.stakeRegistry.GetCurrentStake(
			callOpts,
			request.OperatorId,
			uint8(quorum),
		)
		if err != nil {
			return OperatorStakeInQuorumsOfOperatorResponse{},
				utils.WrapError("failed to get operator stake", err)
		}
		quorumStakes[quorum] = stake
	}

	// 6. Devolver la respuesta como un struct
	return OperatorStakeInQuorumsOfOperatorResponse{
		QuorumStakes: quorumStakes,
	}, nil
}

// GetCheckSignaturesIndices returns the indices needed to verify operator signatures
func (r *ChainReader) GetCheckSignaturesIndices(
	ctx context.Context,
	request SignaturesIndicesRequest,
) (SignaturesIndicesResponse, error) {
	if r.operatorStateRetriever == nil {
		return SignaturesIndicesResponse{}, errors.New(
			"OperatorStateRetriever contract not provided",
		)
	}

	nonSignerOperatorIdsBytes := make([][32]byte, len(request.NonSignerOperatorIds))
	for i, id := range request.NonSignerOperatorIds {
		nonSignerOperatorIdsBytes[i] = id
	}
	checkSignatureIndices, err := r.operatorStateRetriever.GetCheckSignaturesIndices(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		r.registryCoordinatorAddr,
		request.ReferenceBlockNumber,
		request.QuorumNumbers.UnderlyingType(),
		nonSignerOperatorIdsBytes,
	)
	if err != nil {
		return SignaturesIndicesResponse{}, utils.WrapError(
			"Failed to get check signatures indices",
			err,
		)
	}
	return SignaturesIndicesResponse{SignaturesIndices: checkSignatureIndices}, nil
}

// GetOperatorId returns the ID of an operator given their address
func (r *ChainReader) GetOperatorId(
	ctx context.Context,
	request OperatorIdRequest,
) (OperatorIdResponse, error) {
	if r.registryCoordinator == nil {
		return OperatorIdResponse{}, errors.New("RegistryCoordinator contract not provided")
	}

	operatorId, err := r.registryCoordinator.GetOperatorId(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return OperatorIdResponse{}, utils.WrapError("Failed to get operator id", err)
	}
	return OperatorIdResponse{OperatorId: operatorId}, nil
}

// GetOperatorFromId returns the address of an operator given their ID
func (r *ChainReader) GetOperatorFromId(
	ctx context.Context,
	request OperatorFromIdRequest,
) (OperatorFromIdResponse, error) {
	if r.registryCoordinator == nil {
		return OperatorFromIdResponse{}, errors.New("RegistryCoordinator contract not provided")
	}

	operatorAddress, err := r.registryCoordinator.GetOperatorFromId(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorId,
	)
	if err != nil {
		return OperatorFromIdResponse{}, utils.WrapError("Failed to get operator address", err)
	}
	return OperatorFromIdResponse{OperatorAddress: operatorAddress}, nil
}

func (r *ChainReader) QueryRegistrationDetail(
	ctx context.Context,
	request RegistrationDetailRequest,
) (RegistrationDetailResponse, error) {
	if r.registryCoordinator == nil {
		return RegistrationDetailResponse{}, errors.New("RegistryCoordinator contract not provided")
	}

	operatorIdResponse, err := r.GetOperatorId(ctx, OperatorIdRequest(request))
	if err != nil {
		return RegistrationDetailResponse{}, utils.WrapError("Failed to get operator id", err)
	}

	value, err := r.registryCoordinator.GetCurrentQuorumBitmap(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		operatorIdResponse.OperatorId)
	if err != nil {
		return RegistrationDetailResponse{}, utils.WrapError("Failed to get operator quorums", err)
	}

	numBits := value.BitLen()
	var quorums []bool
	for i := 0; i < numBits; i++ {
		quorums = append(quorums, value.Int64()&(1<<i) != 0)
	}

	if len(quorums) == 0 {
		numQuorumsRequest := QuorumCountRequest{BlockNumber: request.BlockNumber}
		numQuorumsResponse, err := r.GetQuorumCount(ctx, numQuorumsRequest)
		if err != nil {
			return RegistrationDetailResponse{}, utils.WrapError("Failed to get quorum count", err)
		}
		for i := uint8(0); i < numQuorumsResponse.QuorumCount; i++ {
			quorums = append(quorums, false)
		}
	}
	return RegistrationDetailResponse{Quorums: quorums}, nil
}

// IsOperatorRegistered checks if an operator is registered with an AVS
func (r *ChainReader) IsOperatorRegistered(
	ctx context.Context,
	request OperatorRegisteredRequest,
) (OperatorRegisteredResponse, error) {
	if r.registryCoordinator == nil {
		return OperatorRegisteredResponse{}, errors.New("RegistryCoordinator contract not provided")
	}

	operatorStatus, err := r.registryCoordinator.GetOperatorStatus(
		&bind.CallOpts{Context: ctx, BlockNumber: request.BlockNumber},
		request.OperatorAddress,
	)
	if err != nil {
		return OperatorRegisteredResponse{}, utils.WrapError("Failed to get operator status", err)
	}

	// 0 = NEVER_REGISTERED, 1 = REGISTERED, 2 = DEREGISTERED
	registeredWithAvs := operatorStatus == 1
	return OperatorRegisteredResponse{IsRegistered: registeredWithAvs}, nil
}

// QueryExistingRegisteredOperatorPubKeys returns the public keys of registered operators
func (r *ChainReader) QueryExistingRegisteredOperatorPubKeys(
	ctx context.Context,
	request OperatorQueryRequest,
) (OperatorPubKeysResponse, error) {
	blsApkRegistryAbi, err := apkreg.ContractBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return OperatorPubKeysResponse{}, utils.WrapError("Cannot get Abi", err)
	}

	if request.StopBlock == 0 {
		curBlockNum, err := r.ethClient.BlockNumber(ctx)
		if err != nil {
			return OperatorPubKeysResponse{}, utils.WrapError("Cannot get current block number", err)
		}
		request.StopBlock = curBlockNum
	}
	if request.BlockRange == 0 {
		request.BlockRange = DefaultQueryBlockRange
	}

	operatorAddresses := make([]types.OperatorAddr, 0)
	operatorPubkeys := make([]types.OperatorPubkeys, 0)
	// QueryExistingRegisteredOperatorPubKeys and QueryExistingRegisteredOperatorSockets
	// both run in parallel and they read and mutate the same variable startBlock,
	// so we clone it to prevent the race condition.
	for i := request.StartBlock; i <= request.StopBlock; i += request.BlockRange {
		// Subtract 1 since FilterQuery is inclusive
		toBlock := i + request.BlockRange - 1
		if toBlock > request.StopBlock {
			toBlock = request.StopBlock
		}

		// FilterQuery needs big.Int
		fromBlockBig := new(big.Int).SetUint64(i)
		toBlockBig := new(big.Int).SetUint64(toBlock)

		query := ethereum.FilterQuery{
			FromBlock: fromBlockBig,
			ToBlock:   toBlockBig,
			Addresses: []common.Address{r.blsApkRegistryAddr},
			Topics:    [][]common.Hash{{blsApkRegistryAbi.Events["NewPubkeyRegistration"].ID}},
		}

		logs, err := r.ethClient.FilterLogs(ctx, query)
		if err != nil {
			return OperatorPubKeysResponse{}, utils.WrapError("Cannot filter logs", err)
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
				return OperatorPubKeysResponse{}, utils.WrapError("Cannot unpack event data", err)
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

	return OperatorPubKeysResponse{
		OperatorAddresses: operatorAddresses,
		OperatorsPubkeys:  operatorPubkeys,
	}, nil
}

// QueryExistingRegisteredOperatorSockets returns the sockets of registered operators
func (r *ChainReader) QueryExistingRegisteredOperatorSockets(
	ctx context.Context,
	request OperatorQueryRequest,
) (OperatorSocketsResponse, error) {
	if r.registryCoordinator == nil {
		return OperatorSocketsResponse{}, errors.New("RegistryCoordinator contract not provided")
	}

	if request.StopBlock == 0 {
		curBlockNum, err := r.ethClient.BlockNumber(ctx)
		if err != nil {
			return OperatorSocketsResponse{}, utils.WrapError("Cannot get current block number", err)
		}
		request.StopBlock = curBlockNum
	}
	if request.BlockRange == 0 {
		request.BlockRange = DefaultQueryBlockRange
	}

	operatorIdToSocketMap := make(map[types.OperatorId]types.Socket)
	// QueryExistingRegisteredOperatorPubKeys and QueryExistingRegisteredOperatorSockets
	// both run in parallel and they read and mutate the same variable startBlock,
	// so we clone it to prevent the race condition.
	for i := request.StartBlock; i <= request.StopBlock; i += request.BlockRange {
		// Subtract 1 since FilterQuery is inclusive
		toBlock := i + request.BlockRange - 1
		if toBlock > request.StopBlock {
			toBlock = request.StopBlock
		}

		end := toBlock

		filterOpts := &bind.FilterOpts{
			Start: i,
			End:   &end,
		}
		socketUpdates, err := r.registryCoordinator.FilterOperatorSocketUpdate(filterOpts, nil)
		if err != nil {
			return OperatorSocketsResponse{}, utils.WrapError("Cannot filter operator socket updates", err)
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
	return OperatorSocketsResponse{Sockets: operatorIdToSocketMap}, nil
}
