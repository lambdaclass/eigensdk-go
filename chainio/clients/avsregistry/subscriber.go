package avsregistry

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsapkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/logging"
)

type ChainSubscriber struct {
	logger         logging.Logger
	regCoord       regcoord.ContractRegistryCoordinatorFilters
	blsApkRegistry blsapkreg.ContractBLSApkRegistryFilters
}

// NewChainSubscriber creates a new instance of ChainSubscriber
// The bindings must be created using websocket ETH Client
func NewChainSubscriber(
	regCoord regcoord.ContractRegistryCoordinatorFilters,
	blsApkRegistry blsapkreg.ContractBLSApkRegistryFilters,
	logger logging.Logger,
) *ChainSubscriber {
	logger = logger.With(logging.ComponentKey, "avsregistry/ChainSubscriber")

	return &ChainSubscriber{
		regCoord:       regCoord,
		blsApkRegistry: blsApkRegistry,
		logger:         logger,
	}
}

// NewSubscriberFromConfig creates a new instance of ChainSubscriber
// A websocket ETH Client must be provided
func NewSubscriberFromConfig(
	cfg Config,
	wsClient eth.WsBackend,
	logger logging.Logger,
) (*ChainSubscriber, error) {
	bindings, err := NewBindingsFromConfig(cfg, wsClient, logger)
	if err != nil {
		wrappedError := elcontracts.CreateForNestedError("NewBindingsFromConfig", err)
		return nil, wrappedError
	}

	return NewChainSubscriber(bindings.RegistryCoordinator, bindings.BlsApkRegistry, logger), nil
}

func (s *ChainSubscriber) SubscribeToNewPubkeyRegistrations() (chan *blsapkreg.ContractBLSApkRegistryNewPubkeyRegistration, event.Subscription, error) {
	newPubkeyRegistrationChan := make(chan *blsapkreg.ContractBLSApkRegistryNewPubkeyRegistration)
	sub, err := s.blsApkRegistry.WatchNewPubkeyRegistration(
		&bind.WatchOpts{}, newPubkeyRegistrationChan, nil,
	)
	if err != nil {
		wrappedError := elcontracts.CreateForOtherError("Failed to subscribe to NewPubkeyRegistration events", err)
		return nil, nil, wrappedError
	}
	return newPubkeyRegistrationChan, sub, nil
}

func (s *ChainSubscriber) SubscribeToOperatorSocketUpdates() (chan *regcoord.ContractRegistryCoordinatorOperatorSocketUpdate, event.Subscription, error) {
	operatorSocketUpdateChan := make(chan *regcoord.ContractRegistryCoordinatorOperatorSocketUpdate)
	sub, err := s.regCoord.WatchOperatorSocketUpdate(
		&bind.WatchOpts{}, operatorSocketUpdateChan, nil,
	)
	if err != nil {
		wrappedError := elcontracts.CreateForOtherError("Failed to subscribe to OperatorSocketUpdate events", err)
		return nil, nil, wrappedError
	}
	return operatorSocketUpdateChan, sub, nil
}
