package elcontracts

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"math/big"

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
	"github.com/Layr-Labs/eigensdk-go/types"
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
	operator types.Operator,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("registering operator %s to EigenLayer", operator.Address)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.delegationManager.RegisterAsOperator(
		noSendTxOpts,
		gethcommon.HexToAddress(operator.DelegationApproverAddress),
		operator.AllocationDelay,
		operator.MetadataUrl,
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info("tx successfully included", "txHash", receipt.TxHash.String())

	return receipt, nil
}

func (w *ChainWriter) UpdateOperatorDetails(
	ctx context.Context,
	operator types.Operator,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	w.logger.Infof("updating operator details of operator %s to EigenLayer", operator.Address)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.delegationManager.ModifyOperatorDetails(
		noSendTxOpts,
		gethcommon.HexToAddress(operator.Address),
		gethcommon.HexToAddress(operator.DelegationApproverAddress),
	)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Info(
		"successfully updated operator details",
		"txHash",
		receipt.TxHash.String(),
		"operator",
		operator.Address,
	)

	return receipt, nil
}

func (w *ChainWriter) UpdateMetadataURI(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	uri string,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.delegationManager == nil {
		return nil, errors.New("DelegationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.delegationManager.UpdateOperatorMetadataURI(noSendTxOpts, operatorAddress, uri)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
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
	strategyAddr gethcommon.Address,
	amount *big.Int,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.strategyManager == nil {
		return nil, errors.New("StrategyManager contract not provided")
	}

	w.logger.Infof("depositing %s tokens into strategy %s", amount.String(), strategyAddr)
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	_, underlyingTokenContract, underlyingTokenAddr, err := w.elChainReader.GetStrategyAndUnderlyingERC20Token(
		ctx,
		strategyAddr,
	)
	if err != nil {
		return nil, err
	}

	tx, err := underlyingTokenContract.Approve(noSendTxOpts, w.strategyManagerAddr, amount)
	if err != nil {
		return nil, errors.Join(errors.New("failed to approve token transfer"), err)
	}
	_, err = w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	tx, err = w.strategyManager.DepositIntoStrategy(noSendTxOpts, strategyAddr, underlyingTokenAddr, amount)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}

	w.logger.Infof("deposited %s into strategy %s", amount.String(), strategyAddr)
	return receipt, nil
}

func (w *ChainWriter) SetClaimerFor(
	ctx context.Context,
	claimer gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}

	tx, err := w.rewardsCoordinator.SetClaimerFor(noSendTxOpts, claimer)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ProcessClaim(
	ctx context.Context,
	claim rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
	recipientAddress gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.ProcessClaim(noSendTxOpts, claim, recipientAddress)
	if err != nil {
		return nil, utils.WrapError("failed to create ProcessClaim tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) SetOperatorAVSSplit(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
	split uint16,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.SetOperatorAVSSplit(noSendTxOpts, operator, avs, split)
	if err != nil {
		return nil, utils.WrapError("failed to create SetOperatorAVSSplit tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) SetOperatorPISplit(
	ctx context.Context,
	operator gethcommon.Address,
	split uint16,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.SetOperatorPISplit(noSendTxOpts, operator, split)
	if err != nil {
		return nil, utils.WrapError("failed to create SetOperatorAVSSplit tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ProcessClaims(
	ctx context.Context,
	claims []rewardscoordinator.IRewardsCoordinatorTypesRewardsMerkleClaim,
	recipientAddress gethcommon.Address,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.rewardsCoordinator == nil {
		return nil, errors.New("RewardsCoordinator contract not provided")
	}

	if len(claims) == 0 {
		return nil, errors.New("claims is empty, at least one claim must be provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.rewardsCoordinator.ProcessClaims(noSendTxOpts, claims, recipientAddress)
	if err != nil {
		return nil, utils.WrapError("failed to create ProcessClaims tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ForceDeregisterFromOperatorSets(
	ctx context.Context,
	operator gethcommon.Address,
	avs gethcommon.Address,
	operatorSetIds []uint32,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AVSDirectory contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.allocationManager.DeregisterFromOperatorSets(
		noSendTxOpts,
		allocationmanager.IAllocationManagerTypesDeregisterParams{
			Operator:       operator,
			Avs:            avs,
			OperatorSetIds: operatorSetIds,
		},
	)

	if err != nil {
		return nil, utils.WrapError("failed to create ForceDeregisterFromOperatorSets tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) ModifyAllocations(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	allocations []allocationmanager.IAllocationManagerTypesAllocateParams,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.allocationManager.ModifyAllocations(noSendTxOpts, operatorAddress, allocations)
	if err != nil {
		return nil, utils.WrapError("failed to create ModifyAllocations tx", err)
	}

	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) SetAllocationDelay(
	ctx context.Context,
	operatorAddress gethcommon.Address,
	delay uint32,
	waitForReceipt bool,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.allocationManager.SetAllocationDelay(noSendTxOpts, operatorAddress, delay)
	if err != nil {
		return nil, utils.WrapError("failed to create InitializeAllocationDelay tx", err)
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		return nil, utils.WrapError("failed to send tx", err)
	}

	return receipt, nil
}

func (w *ChainWriter) DeregisterFromOperatorSets(
	ctx context.Context,
	operator gethcommon.Address,
	request DeregistrationRequest,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.allocationManager.DeregisterFromOperatorSets(
		noSendTxOpts,
		allocationmanager.IAllocationManagerTypesDeregisterParams{
			Operator:       operator,
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
	registryCoordinatorAddr gethcommon.Address,
	request RegistrationRequest,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		return nil, errors.New("AllocationManager contract not provided")
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	pubkeyRegParams, err := getPubkeyRegistrationParams(
		w.ethClient,
		registryCoordinatorAddr,
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
		noSendTxOpts,
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
	request RemovePermissionRequest,
) (*gethtypes.Receipt, error) {
	txOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no-send tx opts", err)
	}
	tx, err := w.NewRemovePermissionTx(txOpts, request)
	if err != nil {
		return nil, utils.WrapError("failed to create NewRemovePermissionTx", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

func (w *ChainWriter) NewRemovePermissionTx(
	txOpts *bind.TransactOpts,
	request RemovePermissionRequest,
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

func (w *ChainWriter) NewSetPermissionTx(
	txOpts *bind.TransactOpts,
	request SetPermissionRequest,
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
	request SetPermissionRequest,
) (*gethtypes.Receipt, error) {
	txOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no-send tx opts", err)
	}

	tx, err := w.NewSetPermissionTx(txOpts, request)
	if err != nil {
		return nil, utils.WrapError("failed to create NewSetPermissionTx", err)
	}

	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

func (w *ChainWriter) NewAcceptAdminTx(
	txOpts *bind.TransactOpts,
	request AcceptAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.AcceptAdmin(txOpts, request.AccountAddress)
}

func (w *ChainWriter) AcceptAdmin(
	ctx context.Context,
	request AcceptAdminRequest,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.NewAcceptAdminTx(noSendTxOpts, request)
	if err != nil {
		return nil, utils.WrapError("failed to create AcceptAdmin transaction", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

func (w *ChainWriter) NewAddPendingAdminTx(
	txOpts *bind.TransactOpts,
	request AddPendingAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.AddPendingAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

func (w *ChainWriter) AddPendingAdmin(ctx context.Context, request AddPendingAdminRequest) (*gethtypes.Receipt, error) {
	txOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}
	tx, err := w.NewAddPendingAdminTx(txOpts, request)
	if err != nil {
		return nil, utils.WrapError("failed to create AddPendingAdminTx", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

func (w *ChainWriter) NewRemoveAdminTx(
	txOpts *bind.TransactOpts,
	request RemoveAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.RemoveAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

func (w *ChainWriter) RemoveAdmin(
	ctx context.Context,
	request RemoveAdminRequest,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.NewRemoveAdminTx(noSendTxOpts, request)
	if err != nil {
		return nil, utils.WrapError("failed to create RemoveAdmin transaction", err)
	}
	return w.txMgr.Send(ctx, tx, request.WaitForReceipt)
}

func (w *ChainWriter) NewRemovePendingAdminTx(
	txOpts *bind.TransactOpts,
	request RemovePendingAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		return nil, errors.New("permission contract not provided")
	}
	return w.permissionController.RemovePendingAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

func (w *ChainWriter) RemovePendingAdmin(
	ctx context.Context,
	request RemovePendingAdminRequest,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("failed to get no send tx opts", err)
	}

	tx, err := w.NewRemovePendingAdminTx(noSendTxOpts, request)
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
