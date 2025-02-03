package avsregistry

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
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
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
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

// The ChainWriter provides methods to call the
// AVS registry contract's state-changing functions.
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

// Returns a new instance of ChainWriter.
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
		wrappedError := elcontracts.CreateForNestedError("NewBindingsFromConfig", err)
		return nil, wrappedError
	}
	elReader, err := elcontracts.NewReaderFromConfig(elcontracts.Config{
		DelegationManagerAddress: bindings.DelegationManagerAddr,
		AvsDirectoryAddress:      bindings.AvsDirectoryAddr,
	}, client, logger)
	if err != nil {
		wrappedError := elcontracts.CreateForNestedError("elcontracts.NewReaderFromConfig", err)
		return nil, wrappedError
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
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	operatorToAvsRegistrationSigSalt [32]byte,
	operatorToAvsRegistrationSigExpiry *big.Int,
	blsKeyPair *bls.KeyPair,
	quorumNumbers types.QuorumNums,
	socket string,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	operatorAddr := crypto.PubkeyToAddress(operatorEcdsaPrivateKey.PublicKey)
	w.logger.Info(
		"registering operator with the AVS's registry coordinator",
		"avs-service-manager",
		w.serviceManagerAddr,
		"operator",
		operatorAddr,
		"quorumNumbers",
		quorumNumbers,
		"socket",
		socket,
	)
	// params to register bls pubkey with bls apk registry
	g1HashedMsgToSign, err := w.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, operatorAddr)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("registryCoordinator.PubkeyRegistrationMessageHash", err)
		return nil, wrappedError
	}
	signedMsg := chainioutils.ConvertToBN254G1Point(
		blsKeyPair.SignHashedToCurveMessage(chainioutils.ConvertBn254GethToGnark(g1HashedMsgToSign)).G1Point,
	)
	G1pubkeyBN254 := chainioutils.ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := chainioutils.ConvertToBN254G2Point(blsKeyPair.GetPubKeyG2())
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
		operatorToAvsRegistrationSigSalt,
		operatorToAvsRegistrationSigExpiry,
	)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("elReader.CalculateOperatorAVSRegistrationDigestHash", err)
		return nil, wrappedError
	}
	operatorSignature, err := crypto.Sign(msgToSign[:], operatorEcdsaPrivateKey)
	if err != nil {
		wrappedError := elcontracts.CreateForOtherError("Failed to sign msg with private key", err)
		return nil, wrappedError
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

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := elcontracts.CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}
	// TODO: this call will fail if max number of operators are already registered
	// in that case, need to call churner to kick out another operator. See eigenDA's node/operator.go implementation
	tx, err := w.registryCoordinator.RegisterOperator(
		noSendTxOpts,
		quorumNumbers.UnderlyingType(),
		socket,
		pubkeyRegParams,
		operatorSignatureWithSaltAndExpiry,
	)
	if err != nil {
		wrappedError := elcontracts.CreateForTxGenerationError("registryCoordinator.RegisterOperator", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := elcontracts.CreateForSendError(err)
		return nil, wrappedError
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
		quorumNumbers,
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
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	blsKeyPair *bls.KeyPair,
	quorumNumbers types.QuorumNums,
	socket string,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	operatorAddr := crypto.PubkeyToAddress(operatorEcdsaPrivateKey.PublicKey)
	w.logger.Info(
		"registering operator with the AVS's registry coordinator",
		"avs-service-manager",
		w.serviceManagerAddr,
		"operator",
		operatorAddr,
		"quorumNumbers",
		quorumNumbers,
		"socket",
		socket,
	)
	// params to register bls pubkey with bls apk registry
	g1HashedMsgToSign, err := w.registryCoordinator.PubkeyRegistrationMessageHash(&bind.CallOpts{}, operatorAddr)
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("registryCoordinator.PubkeyRegistrationMessageHash", err)
		return nil, wrappedError
	}
	signedMsg := chainioutils.ConvertToBN254G1Point(
		blsKeyPair.SignHashedToCurveMessage(chainioutils.ConvertBn254GethToGnark(g1HashedMsgToSign)).G1Point,
	)
	G1pubkeyBN254 := chainioutils.ConvertToBN254G1Point(blsKeyPair.GetPubKeyG1())
	G2pubkeyBN254 := chainioutils.ConvertToBN254G2Point(blsKeyPair.GetPubKeyG2())
	pubkeyRegParams := regcoord.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: signedMsg,
		PubkeyG1:                    G1pubkeyBN254,
		PubkeyG2:                    G2pubkeyBN254,
	}

	// generate a random salt and 1 hour expiry for the signature
	var operatorToAvsRegistrationSigSalt [32]byte
	_, err = rand.Read(operatorToAvsRegistrationSigSalt[:])
	if err != nil {
		wrappedError := elcontracts.CreateForOtherError("Failed to generate a random salt for signature", err)
		return nil, wrappedError
	}

	curBlockNum, err := w.ethClient.BlockNumber(context.Background())
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("ethClient.BlockNumber", err)
		return nil, wrappedError
	}
	curBlock, err := w.ethClient.BlockByNumber(context.Background(), new(big.Int).SetUint64(curBlockNum))
	if err != nil {
		wrappedError := elcontracts.CreateForBindingError("ethClient.BlockByNumber", err)
		return nil, wrappedError
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
		wrappedError := elcontracts.CreateForBindingError("elReader.CalculateOperatorAVSRegistrationDigestHash", err)
		return nil, wrappedError
	}
	operatorSignature, err := crypto.Sign(msgToSign[:], operatorEcdsaPrivateKey)
	if err != nil {
		wrappedError := elcontracts.CreateForOtherError("Failed to sign msg with private key", err)
		return nil, wrappedError
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

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := elcontracts.CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}
	// TODO: this call will fail if max number of operators are already registered
	// in that case, need to call churner to kick out another operator. See eigenDA's node/operator.go implementation
	tx, err := w.registryCoordinator.RegisterOperator(
		noSendTxOpts,
		quorumNumbers.UnderlyingType(),
		socket,
		pubkeyRegParams,
		operatorSignatureWithSaltAndExpiry,
	)
	if err != nil {
		wrappedError := elcontracts.CreateForTxGenerationError("registryCoordinator.RegisterOperator", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := elcontracts.CreateForSendError(err)
		return nil, wrappedError
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
		quorumNumbers,
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
	operatorsPerQuorum [][]gethcommon.Address,
	quorumNumbers types.QuorumNums,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("updating stakes for entire operator set", "quorumNumbers", quorumNumbers)
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := elcontracts.CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}
	tx, err := w.registryCoordinator.UpdateOperatorsForQuorum(
		noSendTxOpts,
		operatorsPerQuorum,
		quorumNumbers.UnderlyingType(),
	)
	if err != nil {
		wrappedError := elcontracts.CreateForTxGenerationError("registryCoordinator.UpdateOperatorsForQuorum", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := elcontracts.CreateForSendError(err)
		return nil, wrappedError
	}
	w.logger.Info(
		"successfully updated stakes for entire operator set",
		"txHash",
		receipt.TxHash.String(),
		"quorumNumbers",
		quorumNumbers,
	)
	return receipt, nil

}

// Updates the stakes of a the given `operators` for all the quorums.
// On success, returns the receipt of the transaction.
func (w *ChainWriter) UpdateStakesOfOperatorSubsetForAllQuorums(
	ctx context.Context,
	operators []gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("updating stakes of operator subset for all quorums", "operators", operators)
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := elcontracts.CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}
	tx, err := w.registryCoordinator.UpdateOperators(noSendTxOpts, operators)
	if err != nil {
		wrappedError := elcontracts.CreateForTxGenerationError("registryCoordinator.UpdateOperators", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := elcontracts.CreateForSendError(err)
		return nil, wrappedError
	}
	w.logger.Info(
		"successfully updated stakes of operator subset for all quorums",
		"txHash",
		receipt.TxHash.String(),
		"operators",
		operators,
	)
	return receipt, nil
}

// Deregisters the caller from the quorums given by `quorumNumbers`.
// On success, returns the receipt of the transaction.
func (w *ChainWriter) DeregisterOperator(
	ctx context.Context,
	quorumNumbers types.QuorumNums,
	pubkey regcoord.BN254G1Point,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("deregistering operator with the AVS's registry coordinator")
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := elcontracts.CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}
	tx, err := w.registryCoordinator.DeregisterOperator0(noSendTxOpts, quorumNumbers.UnderlyingType())
	if err != nil {
		wrappedError := elcontracts.CreateForTxGenerationError("registryCoordinator.DeregisterOperator0", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := elcontracts.CreateForSendError(err)
		return nil, wrappedError
	}
	w.logger.Info(
		"successfully deregistered operator with the AVS's registry coordinator",
		"txHash",
		receipt.TxHash.String(),
	)
	return receipt, nil
}

// Deregisters an operator from the given operator sets.
// On success, returns the receipt of the transaction.
func (w *ChainWriter) DeregisterOperatorOperatorSets(
	ctx context.Context,
	operatorSetIds types.OperatorSetIds,
	operator types.Operator,
	pubkey regcoord.BN254G1Point,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	w.logger.Info("deregistering operator with the AVS's registry coordinator")

	operatorAddress := gethcommon.HexToAddress(operator.Address)
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := elcontracts.CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}
	tx, err := w.registryCoordinator.DeregisterOperator(noSendTxOpts, operatorAddress, operatorSetIds.UnderlyingType())
	if err != nil {
		wrappedError := elcontracts.CreateForTxGenerationError("registryCoordinator.DeregisterOperator", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := elcontracts.CreateForSendError(err)
		return nil, wrappedError
	}
	w.logger.Info(
		"successfully deregistered operator with the AVS's registry coordinator",
		"txHash",
		receipt.TxHash.String(),
	)
	return receipt, nil
}

// Updates the socket of the sender (if it is a registered operator).
// On success, returns the receipt of the transaction.
func (w *ChainWriter) UpdateSocket(
	ctx context.Context,
	socket types.Socket,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := elcontracts.CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}
	tx, err := w.registryCoordinator.UpdateSocket(noSendTxOpts, socket.String())
	if err != nil {
		wrappedError := elcontracts.CreateForTxGenerationError("registryCoordinator.UpdateSocket", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := elcontracts.CreateForSendError(err)
		return nil, wrappedError
	}
	return receipt, nil
}
