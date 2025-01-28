package avsregistry

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	opinfoservice "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type avsRegistryReader interface {
	GetOperatorsStakeInQuorumsAtBlock(
		ctx context.Context,
		request avsregistry.OperatorsStakeInQuorumAtBlockRequest,
	) (avsregistry.OperatorsStakeInQuorumResponse, error)

	GetOperatorFromId(
		ctx context.Context,
		request avsregistry.OperatorFromIdRequest,
	) (avsregistry.OperatorFromIdResponse, error)

	GetCheckSignaturesIndices(
		ctx context.Context,
		request avsregistry.SignaturesIndicesRequest,
	) (avsregistry.SignaturesIndicesResponse, error)
}

// AvsRegistryServiceChainCaller is a wrapper around Reader that transforms the data into
// nicer golang types that are easier to work with
type AvsRegistryServiceChainCaller struct {
	avsRegistryReader
	operatorInfoService opinfoservice.OperatorsInfoService
	logger              logging.Logger
}

var _ AvsRegistryService = (*AvsRegistryServiceChainCaller)(nil)

func NewAvsRegistryServiceChainCaller(
	reader avsRegistryReader,
	operatorInfoService opinfoservice.OperatorsInfoService,
	logger logging.Logger,
) *AvsRegistryServiceChainCaller {
	return &AvsRegistryServiceChainCaller{
		avsRegistryReader:   reader,
		operatorInfoService: operatorInfoService,
		logger:              logger,
	}
}

func (ar *AvsRegistryServiceChainCaller) GetOperatorsAvsStateAtBlock(
	ctx context.Context,
	quorumNumbers types.QuorumNums,
	blockNumber types.BlockNum,
) (map[types.OperatorId]types.OperatorAvsState, error) {
	operatorsAvsState := make(map[types.OperatorId]types.OperatorAvsState)
	// Get operator state for each quorum by querying BLSOperatorStateRetriever (this call is why this service
	// implementation is called ChainCaller)
	response, err := ar.avsRegistryReader.GetOperatorsStakeInQuorumsAtBlock(
		ctx,
		avsregistry.OperatorsStakeInQuorumAtBlockRequest{
			QuorumNumbers:         quorumNumbers,
			BlockNumber:           nil,
			HistoricalBlockNumber: blockNumber,
		},
	)
	if err != nil {
		return nil, utils.WrapError("Failed to get operator state", err)
	}
	numquorums := len(quorumNumbers)
	if len(response.OperatorsStakeInQuorum) != numquorums {
		ar.logger.Error(
			"Number of quorums returned from GetOperatorsStakeInQuorumsAtBlock does not match number of quorums requested. Probably pointing to old contract or wrong implementation.",
			"service",
			"AvsRegistryServiceChainCaller",
			"operatorsStakesInQuorums",
			response.OperatorsStakeInQuorum,
			"numquorums",
			numquorums,
		)
		return nil,
			utils.WrapError(
				"number of quorums returned from GetOperatorsStakeInQuorumsAtBlock does not match number of quorums requested. Probably pointing to old contract or wrong implementation",
				nil,
			)
	}

	for quorumIdx, quorumNum := range quorumNumbers {
		for _, operator := range response.OperatorsStakeInQuorum[quorumIdx] {
			info, err := ar.getOperatorInfo(ctx, operator.OperatorId)
			if err != nil {
				return nil, utils.WrapError("Failed to find pubkeys for operator while building operatorsAvsState", err)
			}
			if operatorAvsState, ok := operatorsAvsState[operator.OperatorId]; ok {
				operatorAvsState.StakePerQuorum[quorumNum] = operator.Stake
			} else {
				stakePerQuorum := make(map[types.QuorumNum]types.StakeAmount)
				stakePerQuorum[quorumNum] = operator.Stake
				operatorsAvsState[operator.OperatorId] = types.OperatorAvsState{
					OperatorId:     operator.OperatorId,
					OperatorInfo:   info,
					StakePerQuorum: stakePerQuorum,
					BlockNumber:    blockNumber,
				}
				operatorsAvsState[operator.OperatorId].StakePerQuorum[quorumNum] = operator.Stake
			}
		}
	}

	return operatorsAvsState, nil
}

func (ar *AvsRegistryServiceChainCaller) GetQuorumsAvsStateAtBlock(
	ctx context.Context,
	quorumNumbers types.QuorumNums,
	blockNumber types.BlockNum,
) (map[types.QuorumNum]types.QuorumAvsState, error) {
	operatorsAvsState, err := ar.GetOperatorsAvsStateAtBlock(ctx, quorumNumbers, blockNumber)
	if err != nil {
		return nil, utils.WrapError("Failed to get quorum state", err)
	}
	quorumsAvsState := make(map[types.QuorumNum]types.QuorumAvsState)
	for _, quorumNum := range quorumNumbers {
		aggPubkeyG1 := bls.NewG1Point(big.NewInt(0), big.NewInt(0))
		totalStake := big.NewInt(0)
		for _, operator := range operatorsAvsState {
			// only include operators that have a stake in this quorum
			if stake, ok := operator.StakePerQuorum[quorumNum]; ok {
				aggPubkeyG1.Add(operator.OperatorInfo.Pubkeys.G1Pubkey)
				totalStake.Add(totalStake, stake)
			}
		}
		quorumsAvsState[quorumNum] = types.QuorumAvsState{
			QuorumNumber: quorumNum,
			AggPubkeyG1:  aggPubkeyG1,
			TotalStake:   totalStake,
			BlockNumber:  blockNumber,
		}
	}
	return quorumsAvsState, nil
}

func (ar *AvsRegistryServiceChainCaller) getOperatorInfo(
	ctx context.Context,
	operatorId types.OperatorId,
) (types.OperatorInfo, error) {
	response, err := ar.avsRegistryReader.GetOperatorFromId(
		context.Background(),
		avsregistry.OperatorFromIdRequest{OperatorId: operatorId, BlockNumber: nil},
	)
	if err != nil {
		return types.OperatorInfo{}, utils.WrapError("Failed to get operator address from pubkey hash", err)
	}
	info, ok := ar.operatorInfoService.GetOperatorInfo(ctx, response.OperatorAddress)
	if !ok {
		return types.OperatorInfo{}, fmt.Errorf(
			"failed to get operator info from operatorInfoService (operatorAddr: %v, operatorId: %v)",
			response.OperatorAddress,
			operatorId,
		)
	}
	return info, nil
}
