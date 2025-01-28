package avsregistry

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	blsapkregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	"github.com/Layr-Labs/eigensdk-go/logging"
)

type eLReader interface {
	CalculateOperatorAVSRegistrationDigestHash(
		ctx context.Context,
		operatorAddr gethcommon.Address,
		serviceManagerAddr gethcommon.Address,
		operatorToAvsRegistrationSigSalt [32]byte,
		operatorToAvsRegistrationSigExpiry *big.Int,
	) ([32]byte, error)
}

type ChainWriter struct {
	serviceManagerAddr     gethcommon.Address
	registryCoordinator    *regcoord.ContractRegistryCoordinator
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever
	stakeRegistry          *stakeregistry.ContractStakeRegistry
	blsApkRegistry         *blsapkregistry.ContractBLSApkRegistry
	elReader               eLReader
	logger                 logging.Logger
	ethClient              eth.HttpBackend
	txMgr                  txmgr.TxManager
}

func NewChainWriter(
	serviceManagerAddr gethcommon.Address,
	registryCoordinator *regcoord.ContractRegistryCoordinator,
	operatorStateRetriever *opstateretriever.ContractOperatorStateRetriever,
	stakeRegistry *stakeregistry.ContractStakeRegistry,
	blsApkRegistry *blsapkregistry.ContractBLSApkRegistry,
	elReader eLReader,
	logger logging.Logger,
	ethClient eth.HttpBackend,
	txMgr txmgr.TxManager,
) *ChainWriter {
	logger = logger.With(logging.ComponentKey, "avsregistry/ChainWriter")

	return &ChainWriter{
		serviceManagerAddr:     serviceManagerAddr,
		registryCoordinator:    registryCoordinator,
		operatorStateRetriever: operatorStateRetriever,
		stakeRegistry:          stakeRegistry,
		blsApkRegistry:         blsApkRegistry,
		elReader:               elReader,
		logger:                 logger,
		ethClient:              ethClient,
		txMgr:                  txMgr,
	}
}

// NewWriterFromConfig creates a new ChainWriter from the provided config
func NewWriterFromConfig(
	cfg Config,
	client eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ChainWriter, error) {
	bindings, err := NewBindingsFromConfig(cfg, client, logger)
	if err != nil {
		return nil, err
	}
	elReader, err := elcontracts.NewReaderFromConfig(elcontracts.Config{
		DelegationManagerAddress: bindings.DelegationManagerAddr,
		AvsDirectoryAddress:      bindings.AvsDirectoryAddr,
	}, client, logger)
	if err != nil {
		return nil, err
	}

	return NewChainWriter(
		bindings.ServiceManagerAddr,
		bindings.RegistryCoordinator,
		bindings.OperatorStateRetriever,
		bindings.StakeRegistry,
		bindings.BlsApkRegistry,
		elReader,
		logger,
		client,
		txMgr,
	), nil
}

// RegisterOperatorInQuorumWithAVSRegistryCoordinator
// TODO(samlaf): an operator that is already registered in a quorum can register with another quorum without passing
// signatures perhaps we should add another sdk function for this purpose, that just takes in a quorumNumber and
// socket? RegisterOperatorInQuorumWithAVSRegistryCoordinator is used to register a single operator with the AVS's
// registry coordinator. - operatorEcdsaPrivateKey is the operator's ecdsa private key (used to sign a message to
// register operator in eigenlayer's delegation manager)
//   - operatorToAvsRegistrationSigSalt is a random salt used to prevent replay attacks
//   - operatorToAvsRegistrationSigExpiry is the expiry time of the signature
//
// Deprecated: use RegisterOperator instead.
// We will only keep high-level functionality such as RegisterOperator, and low level functionality
// such as this function should eventually all be done with bindings directly instead.
func (w *ChainWriter) RegisterOperatorInQuorumWithAVSRegistryCoordinator(
	ctx context.Context,
	// we need to pass the private key explicitly and can't use the signer because registering requires signing a
	// message which isn't a transaction and the signer can only signs transactions see operatorSignature in
	// https://github.com/Layr-Labs/eigenlayer-middleware/blob/m2-mainnet/docs/RegistryCoordinator.md#registeroperator
	// TODO(madhur): check to see if we can make the signer and txmgr more flexible so we can use them (and remote
	// signers) to sign non txs
	request OperatorRegisterInQuorumWithAVSRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	operatorAddr := crypto.PubkeyToAddress(request.OperatorEcdsaPrivateKey.PublicKey)
	w.logger.Info(
		"registering operator with the AVS's registry coordinator",
		"avs-service-manager",
		w.serviceManagerAddr,
		"operator",
		operatorAddr,
		"quorumNumbers",
		request.QuorumNumbers,
		"socket",
		request.Socket,
	)
	// params to register bls pubkey with bls apk registry
	g1HashedMsgToSign, err := w.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, operatorAddr)
	if err != nil {
		return nil, err
	}
	signedMsg := chainioutils.ConvertToBN254G1Point(
		request.BlsKeyPair.SignHashedToCurveMessage(chainioutils.ConvertBn254GethToGnark(g1HashedMsgToSign)).G1Point,
	)
	G1pubkeyBN254 := chainioutils.ConvertToBN254G1Point(request.BlsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := chainioutils.ConvertToBN254G2Point(request.BlsKeyPair.GetPubKeyG2())
	pubkeyRegParams := regcoord.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: signedMsg,
		PubkeyG1:                    G1pubkeyBN254,
		PubkeyG2:                    G2pubkeyBN254,
	}

	// params to register operator in delegation manager's operator-avs mapping
	msgToSign, err := w.elReader.CalculateOperatorAVSRegistrationDigestHash(
		ctx,
		operatorAddr,
		w.serviceManagerAddr,
		request.OperatorToAvsRegistrationSigSalt,
		request.OperatorToAvsRegistrationSigExpiry,
	)
	if err != nil {
		return nil, err
	}
	operatorSignature, err := crypto.Sign(msgToSign[:], request.OperatorEcdsaPrivateKey)
	if err != nil {
		return nil, err
	}
	// the crypto library is low level and deals with 0/1 v values, whereas ethereum expects 27/28, so we add 27
	// see https://github.com/ethereum/go-ethereum/issues/28757#issuecomment-1874525854
	// and https://twitter.com/pcaversaccio/status/1671488928262529031
	operatorSignature[64] += 27
	operatorSignatureWithSaltAndExpiry := regcoord.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: operatorSignature,
		Salt:      request.OperatorToAvsRegistrationSigSalt,
		Expiry:    request.OperatorToAvsRegistrationSigExpiry,
	}

	// TODO: this call will fail if max number of operators are already registered
	// in that case, need to call churner to kick out another operator. See eigenDA's node/operator.go implementation
	tx, err := w.registryCoordinator.RegisterOperator(
		txOptions.Options,
		request.QuorumNumbers.UnderlyingType(),
		request.Socket,
		pubkeyRegParams,
		operatorSignatureWithSaltAndExpiry,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully registered operator with AVS registry coordinator",
		"txHash",
		receipt.TxHash.String(),
		"avs-service-manager",
		w.serviceManagerAddr,
		"operator",
		operatorAddr,
		"quorumNumbers",
		request.QuorumNumbers,
	)
	return receipt, nil
}

// RegisterOperator is similar to RegisterOperatorInQuorumWithAVSRegistryCoordinator but
// generates a random salt and expiry for the signature.
func (w *ChainWriter) RegisterOperator(
	ctx context.Context,
	// we need to pass the private key explicitly and can't use the signer because registering requires signing a
	// message which isn't a transaction and the signer can only signs transactions. See operatorSignature in
	// https://github.com/Layr-Labs/eigenlayer-middleware/blob/m2-mainnet/docs/RegistryCoordinator.md#registeroperator
	// TODO(madhur): check to see if we can make the signer and txmgr more flexible so we can use them (and remote
	// signers) to sign non txs
	request OperatorRegisterRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	operatorAddr := crypto.PubkeyToAddress(request.OperatorEcdsaPrivateKey.PublicKey)
	w.logger.Info(
		"registering operator with the AVS's registry coordinator",
		"avs-service-manager",
		w.serviceManagerAddr,
		"operator",
		operatorAddr,
		"quorumNumbers",
		request.QuorumNumbers,
		"socket",
		request.Socket,
	)
	// params to register bls pubkey with bls apk registry
	g1HashedMsgToSign, err := w.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, operatorAddr)
	if err != nil {
		return nil, err
	}
	signedMsg := chainioutils.ConvertToBN254G1Point(
		request.BlsKeyPair.SignHashedToCurveMessage(chainioutils.ConvertBn254GethToGnark(g1HashedMsgToSign)).G1Point,
	)
	G1pubkeyBN254 := chainioutils.ConvertToBN254G1Point(request.BlsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := chainioutils.ConvertToBN254G2Point(request.BlsKeyPair.GetPubKeyG2())
	pubkeyRegParams := regcoord.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: signedMsg,
		PubkeyG1:                    G1pubkeyBN254,
		PubkeyG2:                    G2pubkeyBN254,
	}

	// generate a random salt and 1 hour expiry for the signature
	var operatorToAvsRegistrationSigSalt [32]byte
	_, err = rand.Read(operatorToAvsRegistrationSigSalt[:])
	if err != nil {
		return nil, err
	}

	curBlockNum, err := w.ethClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	curBlock, err := w.ethClient.BlockByNumber(context.Background(), new(big.Int).SetUint64(curBlockNum))
	if err != nil {
		return nil, err
	}
	sigValidForSeconds := int64(60 * 60) // 1 hour
	operatorToAvsRegistrationSigExpiry := new(
		big.Int,
	).Add(new(big.Int).SetUint64(curBlock.Time()), big.NewInt(sigValidForSeconds))

	// params to register operator in delegation manager's operator-avs mapping
	msgToSign, err := w.elReader.CalculateOperatorAVSRegistrationDigestHash(
		ctx,
		operatorAddr,
		w.serviceManagerAddr,
		operatorToAvsRegistrationSigSalt,
		operatorToAvsRegistrationSigExpiry,
	)
	if err != nil {
		return nil, err
	}
	operatorSignature, err := crypto.Sign(msgToSign[:], request.OperatorEcdsaPrivateKey)
	if err != nil {
		return nil, err
	}
	// the crypto library is low level and deals with 0/1 v values, whereas ethereum expects 27/28, so we add 27
	// see https://github.com/ethereum/go-ethereum/issues/28757#issuecomment-1874525854
	// and https://twitter.com/pcaversaccio/status/1671488928262529031
	operatorSignature[64] += 27
	operatorSignatureWithSaltAndExpiry := regcoord.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: operatorSignature,
		Salt:      operatorToAvsRegistrationSigSalt,
		Expiry:    operatorToAvsRegistrationSigExpiry,
	}

	// TODO: this call will fail if max number of operators are already registered
	// in that case, need to call churner to kick out another operator. See eigenDA's node/operator.go implementation
	tx, err := w.registryCoordinator.RegisterOperator(
		txOptions.Options,
		request.QuorumNumbers.UnderlyingType(),
		request.Socket,
		pubkeyRegParams,
		operatorSignatureWithSaltAndExpiry,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully registered operator with AVS registry coordinator",
		"txHash",
		receipt.TxHash.String(),
		"avs-service-manager",
		w.serviceManagerAddr,
		"operator",
		operatorAddr,
		"quorumNumbers",
		request.QuorumNumbers,
	)
	return receipt, nil
}

// UpdateStakesOfEntireOperatorSetForQuorums is used by avs teams running https://github.com/Layr-Labs/avs-sync
// to updates the stake of their entire operator set.
// Because of high gas costs of this operation, it typically needs to be called for every quorum, or perhaps for a
// small grouping of quorums
// (highly dependent on number of operators per quorum)
func (w *ChainWriter) UpdateStakesOfEntireOperatorSetForQuorums(
	ctx context.Context,
	request StakesOfEntireOperatorSetForQuorumsRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	w.logger.Info("updating stakes for entire operator set", "quorumNumbers", request.QuorumNumbers)

	tx, err := w.registryCoordinator.UpdateOperatorsForQuorum(
		txOptions.Options,
		request.OperatorsPerQuorum,
		request.QuorumNumbers.UnderlyingType(),
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully updated stakes for entire operator set",
		"txHash",
		receipt.TxHash.String(),
		"quorumNumbers",
		request.QuorumNumbers,
	)
	return receipt, nil

}

func (w *ChainWriter) UpdateStakesOfOperatorSubsetForAllQuorums(
	ctx context.Context,
	request StakesOfOperatorSubsetForAllQuorumsRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	w.logger.Info("updating stakes of operator subset for all quorums", "operators", request.OperatorsAddresses)

	tx, err := w.registryCoordinator.UpdateOperators(txOptions.Options, request.OperatorsAddresses)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully updated stakes of operator subset for all quorums",
		"txHash",
		receipt.TxHash.String(),
		"operators",
		request.OperatorsAddresses,
	)
	return receipt, nil
}

// DeregisterOperator is used to deregister an operator from one or more quorums.
func (w *ChainWriter) DeregisterOperator(
	ctx context.Context,
	request OperatorDeregisterRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	w.logger.Info("deregistering operator with the AVS's registry coordinator")

	tx, err := w.registryCoordinator.DeregisterOperator0(txOptions.Options, request.QuorumNumbers.UnderlyingType())
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully deregistered operator with the AVS's registry coordinator",
		"txHash",
		receipt.TxHash.String(),
	)
	return receipt, nil
}

// DeregisterOperatorOperatorSets is used to deregister an operator from one or more operator sets.
func (w *ChainWriter) DeregisterOperatorOperatorSets(
	ctx context.Context,
	request OperatorDeregisterOperatorSetsRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	w.logger.Info("deregistering operator with the AVS's registry coordinator")

	operatorAddress := gethcommon.HexToAddress(request.Operator.Address)

	tx, err := w.registryCoordinator.DeregisterOperator(
		txOptions.Options,
		operatorAddress,
		request.OperatorSetIds.UnderlyingType(),
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully deregistered operator with the AVS's registry coordinator",
		"txHash",
		receipt.TxHash.String(),
	)
	return receipt, nil
}

// UpdateSocket is used to update the socket of the msg.sender given they are a registered operator
func (w *ChainWriter) UpdateSocket(
	ctx context.Context,
	request SocketUpdateRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	tx, err := w.registryCoordinator.UpdateSocket(txOptions.Options, request.Socket.String())
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send UpdateSocket tx with err: " + err.Error())
	}
	return receipt, nil
}
