package elcontracts

import (
	"context"

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
		wrappedError := CreateForNestedError("NewBindingsFromConfig", err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("DelegationManager")
		return nil, wrappedError
	}

	w.logger.Infof("registering operator %s to EigenLayer", operator.Address)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}
	tx, err := w.delegationManager.RegisterAsOperator(
		noSendTxOpts,
		gethcommon.HexToAddress(operator.DelegationApproverAddress),
		operator.AllocationDelay,
		operator.MetadataUrl,
	)
	if err != nil {
		wrappedError := CreateForTxGenerationError("delegationManager.RegisterAsOperator", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("DelegationManager")
		return nil, wrappedError
	}

	w.logger.Infof("updating operator details of operator %s to EigenLayer", operator.Address)

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.delegationManager.ModifyOperatorDetails(
		noSendTxOpts,
		gethcommon.HexToAddress(operator.Address),
		gethcommon.HexToAddress(operator.DelegationApproverAddress),
	)
	if err != nil {
		wrappedError := CreateForTxGenerationError("delegationManager.ModifyOperatorDetails", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("DelegationManager")
		return nil, wrappedError
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.delegationManager.UpdateOperatorMetadataURI(noSendTxOpts, operatorAddress, uri)
	if err != nil {
		wrappedError := CreateForTxGenerationError("delegationManager.UpdateOperatorMetadataURI", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("StrategyManager")
		return nil, wrappedError
	}

	w.logger.Infof("depositing %s tokens into strategy %s", amount.String(), strategyAddr)
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}
	_, underlyingTokenContract, underlyingTokenAddr, err := w.elChainReader.GetStrategyAndUnderlyingERC20Token(
		ctx,
		strategyAddr,
	)
	if err != nil {
		wrappedError := CreateForNestedError("elChainReader.GetStrategyAndUnderlyingERC20Token", err)
		return nil, wrappedError
	}

	tx, err := underlyingTokenContract.Approve(noSendTxOpts, w.strategyManagerAddr, amount)
	if err != nil {
		wrappedError := Error{3, "Other errors", "failed to approve token transfer", err}
		return nil, wrappedError
	}
	_, err = w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
	}

	tx, err = w.strategyManager.DepositIntoStrategy(noSendTxOpts, strategyAddr, underlyingTokenAddr, amount)
	if err != nil {
		wrappedError := CreateForTxGenerationError("strategyManager.DepositIntoStrategy", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return nil, wrappedError
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.rewardsCoordinator.SetClaimerFor(noSendTxOpts, claimer)
	if err != nil {
		wrappedError := CreateForTxGenerationError("rewardsCoordinator.SetClaimerFor", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return nil, wrappedError
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.rewardsCoordinator.ProcessClaim(noSendTxOpts, claim, recipientAddress)
	if err != nil {
		wrappedError := CreateForTxGenerationError("rewardsCoordinator.ProcessClaim", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return nil, wrappedError
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.rewardsCoordinator.SetOperatorAVSSplit(noSendTxOpts, operator, avs, split)
	if err != nil {
		wrappedError := CreateForTxGenerationError("rewardsCoordinator.SetOperatorAVSSplit", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return nil, wrappedError
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.rewardsCoordinator.SetOperatorPISplit(noSendTxOpts, operator, split)
	if err != nil {
		wrappedError := CreateForTxGenerationError("rewardsCoordinator.SetOperatorPISplit", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("RewardsCoordinator")
		return nil, wrappedError
	}

	if len(claims) == 0 {
		wrappedError := Error{3, "Other errors", "No claims were found to process, at least one claim must be provided", nil}
		return nil, wrappedError
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.rewardsCoordinator.ProcessClaims(noSendTxOpts, claims, recipientAddress)
	if err != nil {
		wrappedError := CreateForTxGenerationError("rewardsCoordinator.ProcessClaims", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return nil, wrappedError
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
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
		wrappedError := CreateForTxGenerationError("allocationManager.DeregisterFromOperatorSets", err)
		return nil, wrappedError
	}

	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return nil, wrappedError
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.allocationManager.ModifyAllocations(noSendTxOpts, operatorAddress, allocations)
	if err != nil {
		wrappedError := CreateForTxGenerationError("allocationManager.ModifyAllocations", err)
		return nil, wrappedError
	}

	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return nil, wrappedError
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.allocationManager.SetAllocationDelay(noSendTxOpts, operatorAddress, delay)
	if err != nil {
		wrappedError := CreateForTxGenerationError("allocationManager.SetAllocationDelay", err)
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
	}

	return receipt, nil
}

func (w *ChainWriter) DeregisterFromOperatorSets(
	ctx context.Context,
	operator gethcommon.Address,
	request DeregistrationRequest,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return nil, wrappedError
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.allocationManager.DeregisterFromOperatorSets(
		noSendTxOpts,
		allocationmanager.IAllocationManagerTypesDeregisterParams{
			Operator:       operator,
			Avs:            request.AVSAddress,
			OperatorSetIds: request.OperatorSetIds,
		})
	if err != nil {
		wrappedError := CreateForTxGenerationError("allocationManager.DeregisterFromOperatorSets", err)
		return nil, wrappedError
	}

	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
	}

	return receipt, nil
}

func (w *ChainWriter) RegisterForOperatorSets(
	ctx context.Context,
	registryCoordinatorAddr gethcommon.Address,
	request RegistrationRequest,
) (*gethtypes.Receipt, error) {
	if w.allocationManager == nil {
		wrappedError := CreateErrorForMissingContract("AllocationManager")
		return nil, wrappedError
	}

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	pubkeyRegParams, err := getPubkeyRegistrationParams(
		w.ethClient,
		registryCoordinatorAddr,
		request.OperatorAddress,
		request.BlsKeyPair,
	)
	if err != nil {
		wrappedError := Error{
			2,
			"Nested error",
			"Failed to get public key registration params",
			err,
		}
		return nil, wrappedError
	}

	data, err := abiEncodeRegistrationParams(request.Socket, *pubkeyRegParams)
	if err != nil {
		wrappedError := Error{
			2,
			"Nested error",
			"Failed to encode registration params",
			err,
		}
		return nil, wrappedError
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
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
	}

	return receipt, nil
}

func (w *ChainWriter) RemovePermission(
	ctx context.Context,
	request RemovePermissionRequest,
) (*gethtypes.Receipt, error) {
	txOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}
	tx, err := w.NewRemovePermissionTx(txOpts, request)
	if err != nil {
		wrappedError := Error{
			2,
			"Nested error",
			"Failed to create a new remove permission Tx",
			err,
		}
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
	}

	return receipt, nil
}

func (w *ChainWriter) NewRemovePermissionTx(
	txOpts *bind.TransactOpts,
	request RemovePermissionRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return nil, wrappedError
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
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return nil, wrappedError
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
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.NewSetPermissionTx(txOpts, request)
	if err != nil {
		wrappedError := Error{
			2,
			"Nested error",
			"Failed to create a new set permission Tx",
			err,
		}
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
	}

	return receipt, nil
}

func (w *ChainWriter) NewAcceptAdminTx(
	txOpts *bind.TransactOpts,
	request AcceptAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return nil, wrappedError
	}
	return w.permissionController.AcceptAdmin(txOpts, request.AccountAddress)
}

func (w *ChainWriter) AcceptAdmin(
	ctx context.Context,
	request AcceptAdminRequest,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.NewAcceptAdminTx(noSendTxOpts, request)
	if err != nil {
		wrappedError := Error{
			2,
			"Nested error",
			"Failed to create a new accept admin Tx",
			err,
		}
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
	}

	return receipt, nil
}

func (w *ChainWriter) NewAddPendingAdminTx(
	txOpts *bind.TransactOpts,
	request AddPendingAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return nil, wrappedError
	}
	return w.permissionController.AddPendingAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

func (w *ChainWriter) AddPendingAdmin(ctx context.Context, request AddPendingAdminRequest) (*gethtypes.Receipt, error) {
	txOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}
	tx, err := w.NewAddPendingAdminTx(txOpts, request)
	if err != nil {
		wrappedError := Error{
			2,
			"Nested error",
			"Failed to create a new add pending admin Tx",
			err,
		}
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
	}

	return receipt, nil
}

func (w *ChainWriter) NewRemoveAdminTx(
	txOpts *bind.TransactOpts,
	request RemoveAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return nil, wrappedError
	}
	return w.permissionController.RemoveAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

func (w *ChainWriter) RemoveAdmin(
	ctx context.Context,
	request RemoveAdminRequest,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.NewRemoveAdminTx(noSendTxOpts, request)
	if err != nil {
		wrappedError := Error{
			2,
			"Nested error",
			"Failed to create a new remove admin Tx",
			err,
		}
		return nil, wrappedError
	}
	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
	}

	return receipt, nil
}

func (w *ChainWriter) NewRemovePendingAdminTx(
	txOpts *bind.TransactOpts,
	request RemovePendingAdminRequest,
) (*gethtypes.Transaction, error) {
	if w.permissionController == nil {
		wrappedError := CreateErrorForMissingContract("PermissionController")
		return nil, wrappedError
	}
	return w.permissionController.RemovePendingAdmin(txOpts, request.AccountAddress, request.AdminAddress)
}

func (w *ChainWriter) RemovePendingAdmin(
	ctx context.Context,
	request RemovePendingAdminRequest,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		wrappedError := CreateNoSendTxOptsFailedError(err)
		return nil, wrappedError
	}

	tx, err := w.NewRemovePendingAdminTx(noSendTxOpts, request)
	if err != nil {
		wrappedError := Error{
			2,
			"Nested error",
			"Failed to create a new remove pending admin Tx",
			err,
		}
		return nil, wrappedError
	}

	receipt, err := w.txMgr.Send(ctx, tx, request.WaitForReceipt)
	if err != nil {
		wrappedError := CreateForSendError(err)
		return nil, wrappedError
	}

	return receipt, nil
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
		wrappedError := Error{3, "Other errors", "Failed to encode abi params", err}
		return nil, wrappedError
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
		wrappedError := Error{3, "Other errors", "Failed to pack arguments", err}
		return nil, wrappedError
	}
	// The encoder is prepending 32 bytes to the data as if it was used in a dynamic function parameter.
	// This is not used when decoding the bytes directly, so we need to remove it.
	return data[32:], nil
}
