package clients_test

import (
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestEcClient(t *testing.T) {
	var config BuildAllConfig

	/*
		type BuildAllConfig struct {
		EthHttpUrl                 string
		EthWsUrl                   string
		RegistryCoordinatorAddr    string
		OperatorStateRetrieverAddr string
		AvsName                    string
		PromMetricsIpPortAddress   string
		}
	*/

	config.validate(logger)

	// create a new BaseClient
	avsRegistryChainReader, avsRegistryChainSubscriber, avsRegistryContractBindings, err := avsregistry.BuildClientsForEcMetrics(
		avsregistry.Config{
			RegistryCoordinatorAddress:    gethcommon.HexToAddress(config.RegistryCoordinatorAddr),
			OperatorStateRetrieverAddress: gethcommon.HexToAddress(config.OperatorStateRetrieverAddr),
		},
		ethHttpClient,
		ethWsClient,
		logger,
	)

}

/*
func TestValidateRawGithubUrl(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		expectedErr error
	}{
		{
			name: "valid raw github url",
			url:  "https://raw.githubusercontent.com/Layr-Labs/eigensdk-go/main/README.md",
		},
		{
			name:        "invalid raw github url",
			url:         "https://facebook.com",
			expectedErr: ErrInvalidGithubRawUrl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRawGithubUrl(tt.url)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
*/
