package elcontracts

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	chainioutils "github.com/Layr-Labs/eigensdk-go/chainio/utils"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	erc20 "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IERC20"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	strategy "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IStrategy"
	permissioncontroller "github.com/Layr-Labs/eigensdk-go/contracts/bindings/PermissionController"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type Reader interface {
	GetStrategyAndUnderlyingERC20Token(
		ctx context.Context, strategyAddr gethcommon.Address,
	) (*strategy.ContractIStrategy, erc20.ContractIERC20Methods, gethcommon.Address, error)
}

type ChainWriter struct {
	delegationManager    *delegationmanager.ContractDelegationManager
	strategyManager      *strategymanager.ContractStrategyManager
	rewardsCoordinator   *rewardscoordinator.ContractIRewardsCoordinator
	avsDirectory         *avsdirectory.ContractAVSDirectory
	allocationManager    *allocationmanager.ContractAllocationManager
	permissionController *permissioncontroller.ContractPermissionController
	strategyManagerAddr  gethcommon.Address
	elChainReader        Reader
	ethClient            eth.HttpBackend
	logger               logging.Logger
	txMgr                txmgr.TxManager
}

func NewChainWriter(
	delegationManager *delegationmanager.ContractDelegationManager,
	strategyManager *strategymanager.ContractStrategyManager,
	rewardsCoordinator *rewardscoordinator.ContractIRewardsCoordinator,
	avsDirectory *avsdirectory.ContractAVSDirectory,
	allocationManager *allocationmanager.ContractAllocationManager,
	permissionController *permissioncontroller.ContractPermissionController,
	strategyManagerAddr gethcommon.Address,
	elChainReader Reader,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) *ChainWriter {
	logger = logger.With(logging.ComponentKey, "elcontracts/writer")

	return &ChainWriter{
		delegationManager:    delegationManager,
		strategyManager:      strategyManager,
		strategyManagerAddr:  strategyManagerAddr,
		rewardsCoordinator:   rewardsCoordinator,
		allocationManager:    allocationManager,
		permissionController: permissionController,
		avsDirectory:         avsDirectory,
		elChainReader:        elChainReader,
		logger:               logger,
		ethClient:            ethClient,
		txMgr:                txMgr,
	}
}

func NewWriterFromConfig(
	cfg Config,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) (*ChainWriter, error) {
	elContractBindings, err := NewBindingsFromConfig(
		cfg,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	elChainReader := NewChainReader(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AllocationManager,
		elContractBindings.PermissionController,
		logger,
		ethClient,
	)
	return NewChainWriter(
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AvsDirectory,
		elContractBindings.AllocationManager,
		elContractBindings.PermissionController,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		ethClient,
		logger,
		eigenMetrics,
		txMgr,
	), nil
}

func (w *ChainWriter) RegisterAsOperator(
	ctx context.Context,
	request RegisterOperatorRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("registering operator %s to EigenLayer", request.Operator.Address)

	tx, err := w.delegationManager.RegisterAsOperator(
		txOptions.Options,
		gethcommon.HexToAddress(request.Operator.DelegationApproverAddress),
		request.Operator.AllocationDelay,
		request.Operator.MetadataUrl,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info("tx successfully included", "txHash", receipt.TxHash.String())

	return receipt, nil
}

func (w *ChainWriter) UpdateOperatorDetails(
	ctx context.Context,
	request OperatorDetailsRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("updating operator details of operator %s to EigenLayer", request.Operator.Address)

	tx, err := w.delegationManager.ModifyOperatorDetails(
		txOptions.Options,
		gethcommon.HexToAddress(request.Operator.Address),
		gethcommon.HexToAddress(request.Operator.DelegationApproverAddress),
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully updated operator details",
		"txHash",
		receipt.TxHash.String(),
		"operator",
		request.Operator.Address,
	)

	return receipt, nil
}

func (w *ChainWriter) UpdateMetadataURI(
	ctx context.Context,
	request MetadataURIRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	tx, err := w.delegationManager.UpdateOperatorMetadataURI(txOptions.Options, request.OperatorAddress, request.Uri)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully updated operator metadata uri",
		"txHash",
		receipt.TxHash.String(),
	)

	return receipt, nil
}

func (w *ChainWriter) DepositERC20IntoStrategy(
	ctx context.Context,
	request ERC20IntoStrategyRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.strategyManager == nil {
		return nil, errors.New("StrategyManager contract not provided")
	}

	w.logger.Infof("depositing %s tokens into strategy %s", request.Amount.String(), request.StrategyAddress)

	_, underlyingTokenContract, underlyingTokenAddr, err := w.elChainReader.GetStrategyAndUnderlyingERC20Token(
		ctx,
		request.StrategyAddress,
	)
	if err != nil {
		return nil, err
	}

	tx, err := underlyingTokenContract.Approve(txOptions.Options, w.strategyManagerAddr, request.Amount)
	if err != nil {
		return nil, errors.Join(errors.New("failed to approve token transfer"), err)
	}
	_, err = w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	tx, err = w.strategyManager.DepositIntoStrategy(
		txOptions.Options,
		request.StrategyAddress,
		underlyingTokenAddr,
		request.Amount,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	w.logger.Infof("deposited %s into strategy %s", request.Amount.String(), request.StrategyAddress)
	return receipt, nil
}

func (w *ChainWriter) SetClaimerFor(
	ctx context.Context,
	request ClaimForRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	tx, err := w.rewardsCoordinator.SetClaimerFor(txOptions.Options, request.Claimer)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ProcessClaim(
	ctx context.Context,
	request ClaimProcessRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	tx, err := w.rewardsCoordinator.ProcessClaim(txOptions.Options, request.Claim, request.RecipientAddress)
	if err != nil {
		return nil, utils.WrapError("failed to create ProcessClaim tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) SetOperatorAVSSplit(
	ctx context.Context,
	request OperatorAVSSplitRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	tx, err := w.rewardsCoordinator.SetOperatorAVSSplit(
		txOptions.Options,
		request.OperatorAddress,
		request.AVSAddress,
		request.Split,
	)
	if err != nil {
		return nil, utils.WrapError("failed to create SetOperatorAVSSplit tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) SetOperatorPISplit(
	ctx context.Context,
	request OperatorPISplitRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	tx, err := w.rewardsCoordinator.SetOperatorPISplit(txOptions.Options, request.OperatorAddress, request.Split)
	if err != nil {
		return nil, utils.WrapError("failed to create SetOperatorAVSSplit tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ProcessClaims(
	ctx context.Context,
	request ClaimsProcessRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	if len(request.Claims) == 0 {
		return nil, errors.New("claims is empty, at least one claim must be provided")
	}

	tx, err := w.rewardsCoordinator.ProcessClaims(txOptions.Options, request.Claims, request.RecipientAddress)
	if err != nil {
		return nil, utils.WrapError("failed to create ProcessClaims tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ForceDeregisterFromOperatorSets(
	ctx context.Context,
	request OperatorSetDeregisterRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AVSDirectory contract not provided")
	}

	tx, err := w.allocationManager.DeregisterFromOperatorSets(
		txOptions.Options,
		allocationmanager.IAllocationManagerTypesDeregisterParams{
			Operator:       request.OperatorAddress,
			Avs:            request.AVSAddress,
			OperatorSetIds: request.OperatorSetIds,
		},
	)

	if err != nil {
		return nil, utils.WrapError("failed to create ForceDeregisterFromOperatorSets tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ModifyAllocations(
	ctx context.Context,
	request AllocationModifyRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	tx, err := w.allocationManager.ModifyAllocations(txOptions.Options, request.OperatorAddress, request.Allocations)
	if err != nil {
		return nil, utils.WrapError("failed to create ModifyAllocations tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) SetAllocationDelay(
	ctx context.Context,
	request AllocationDelayRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	tx, err := w.allocationManager.SetAllocationDelay(txOptions.Options, request.OperatorAddress, request.Delay)
	if err != nil {
		return nil, utils.WrapError("failed to create InitializeAllocationDelay tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) DeregisterFromOperatorSets(
	ctx context.Context,
	request OperatorSetDeregisterRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	tx, err := w.allocationManager.DeregisterFromOperatorSets(
		txOptions.Options,
		allocationmanager.IAllocationManagerTypesDeregisterParams{
			Operator:       request.OperatorAddress,
			Avs:            request.AVSAddress,
			OperatorSetIds: request.OperatorSetIds,
		})
	if err != nil {
		return nil, utils.WrapError("failed to create DeregisterFromOperatorSets tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) RegisterForOperatorSets(
	ctx context.Context,
	request RegisterOperatorSetsRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	pubkeyRegParams, err := getPubkeyRegistrationParams(
		w.ethClient,
		request.RegistryCoordinatorAddress,
		request.OperatorAddress,
		request.BlsKeyPair,
	)
	if err != nil {
		return nil, utils.WrapError("failed to get public key registration params", err)
	}

	data, err := abiEncodeRegistrationParams(request.Socket, *pubkeyRegParams)
	if err != nil {
		return nil, utils.WrapError("failed to encode registration params", err)
	}
	tx, err := w.allocationManager.RegisterForOperatorSets(
		txOptions.Options,
		request.OperatorAddress,
		allocationmanager.IAllocationManagerTypesRegisterParams{
			Avs:            request.AVSAddress,
			OperatorSetIds: request.OperatorSetIds,
			Data:           data,
		})
	if err != nil {
		return nil, utils.WrapError("failed to create RegisterForOperatorSets tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) RemovePermission(
	ctx context.Context,
	request PermissionRemoveRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	tx, err := w.NewRemovePermissionTx(request, txOptions.Options)
	if err != nil {
		return nil, utils.WrapError("failed to create NewRemovePermissionTx", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

// Should be a public or private method?
func (w *ChainWriter) NewRemovePermissionTx(
	request PermissionRemoveRequest,
	txOpts *bind.TransactOpts,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}

	return w.permissionController.RemoveAppointee(
		txOpts,
		request.AccountAddress,
		request.AppointeeAddress,
		request.Target,
		request.Selector,
	)
}

// Should be a public or private method?
func (w *ChainWriter) NewSetPermissionTx(
	request PermissionSetRequest,
	txOpts *bind.TransactOpts,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.SetAppointee(
		txOpts,
		request.AccountAddress,
		request.AppointeeAddress,
		request.Target,
		request.Selector,
	)
}

func (w *ChainWriter) SetPermission(
	ctx context.Context,
	request PermissionSetRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	tx, err := w.NewSetPermissionTx(request, txOptions.Options)
	if err != nil {
		return nil, utils.WrapError("failed to create NewSetPermissionTx", err)
	}

	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

// Should be a public or private method?
func (w *ChainWriter) NewAcceptAdminTx(
	request AdminAcceptRequest,
	txOpts *bind.TransactOpts,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.AcceptAdmin(txOpts, request.AccountAddress)
}

func (w *ChainWriter) AcceptAdmin(
	ctx context.Context,
	request AdminAcceptRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	tx, err := w.NewAcceptAdminTx(request, txOptions.Options)
	if err != nil {
		return nil, utils.WrapError("failed to create AcceptAdmin transaction", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

// Should be a public or private method?
func (w *ChainWriter) NewAddPendingAdminTx(
	request PendingAdminAcceptRequest,
	txOpts *bind.TransactOpts,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.AddPendingAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

func (w *ChainWriter) AddPendingAdmin(
	ctx context.Context,
	request PendingAdminAcceptRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	tx, err := w.NewAddPendingAdminTx(request, txOptions.Options)
	if err != nil {
		return nil, utils.WrapError("failed to create AddPendingAdminTx", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

// Should be a public or private method?
func (w *ChainWriter) NewRemoveAdminTx(
	request AdminRemoveRequest,
	txOpts *bind.TransactOpts,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.RemoveAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

func (w *ChainWriter) RemoveAdmin(
	ctx context.Context,
	request AdminRemoveRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	tx, err := w.NewRemoveAdminTx(request, txOptions.Options)
	if err != nil {
		return nil, utils.WrapError("failed to create RemoveAdmin transaction", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

// Should be a public or private method?
func (w *ChainWriter) NewRemovePendingAdminTx(
	request PendingAdminRemoveRequest,
	txOpts *bind.TransactOpts,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.RemovePendingAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

func (w *ChainWriter) RemovePendingAdmin(
	ctx context.Context,
	request PendingAdminRemoveRequest,
	txOptions *TxOptions,
) (*gethtypes.Receipt, error) {
	tx, err := w.NewRemovePendingAdminTx(request, txOptions.Options)
	if err != nil {
		return nil, utils.WrapError("failed to create RemovePendingAdmin transaction", err)
	}

	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

func getPubkeyRegistrationParams(
	ethClient bind.ContractBackend,
	registryCoordinatorAddr, operatorAddress gethcommon.Address,
	blsKeyPair *bls.KeyPair,
) (*regcoord.IBLSApkRegistryPubkeyRegistrationParams, error) {
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethClient)
	if err != nil {
		return nil, utils.WrapError("failed to create registry coordinator", err)
	}
	// params to register bls pubkey with bls apk registry
	g1HashedMsgToSign, err := registryCoordinator.PubkeyRegistrationMessageHash(
		&bind.CallOpts{},
		operatorAddress,
	)
	if err != nil {
		return nil, err
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
	return &pubkeyRegParams, nil
}

func abiEncodeRegistrationParams(
	socket string,
	pubkeyRegistrationParams regcoord.IBLSApkRegistryPubkeyRegistrationParams,
) ([]byte, error) {
	registrationParamsType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "Socket", Type: "string"},
		{Name: "PubkeyRegParams", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "PubkeyRegistrationSignature", Type: "tuple", Components: []abi.ArgumentMarshaling{
				{Name: "X", Type: "uint256"},
				{Name: "Y", Type: "uint256"},
			}},
			{Name: "PubkeyG1", Type: "tuple", Components: []abi.ArgumentMarshaling{
				{Name: "X", Type: "uint256"},
				{Name: "Y", Type: "uint256"},
			}},
			{Name: "PubkeyG2", Type: "tuple", Components: []abi.ArgumentMarshaling{
				{Name: "X", Type: "uint256[2]"},
				{Name: "Y", Type: "uint256[2]"},
			}},
		}},
	})
	if err != nil {
		return nil, err
	}

	registrationParams := struct {
		Socket          string
		PubkeyRegParams regcoord.IBLSApkRegistryPubkeyRegistrationParams
	}{
		socket,
		pubkeyRegistrationParams,
	}

	args := abi.Arguments{
		{Type: registrationParamsType, Name: "registrationParams"},
	}

	data, err := args.Pack(&registrationParams)
	if err != nil {
		return nil, err
	}
	// The encoder is prepending 32 bytes to the data as if it was used in a dynamic function parameter.
	// This is not used when decoding the bytes directly, so we need to remove it.
	return data[32:], nil
}
