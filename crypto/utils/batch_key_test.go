package utils

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadBatchKeys(t *testing.T) {

	var tests = []struct {
		keyFolder string
		isEcdsa   bool
	}{
		{
			keyFolder: "test_data/bls_keys",
			isEcdsa:   false,
		},
		{
			keyFolder: "test_data/ecdsa_keys",
			isEcdsa:   true,
		},
	}

	for _, test := range tests {
		readKeys, err := ReadBatchKeys(test.keyFolder, test.isEcdsa)
		if err != nil {
			t.Errorf("Error reading batch keys: %s", err)
		}

		if test.isEcdsa {
			for _, key := range readKeys {
				fmt.Printf("Valdiate ecdsa key: %s with password %s\n", key.FilePath, key.Password)

				ecdsaKey, err := ecdsa.CreateNewEcdsaKey()
				require.NoError(t, err)
				pk, err := ecdsaKey.Read(key.FilePath, key.Password)
				if err != nil {
					assert.Fail(t, "Error reading ecdsa key")
					break
				}
				assert.Equal(t, key.PrivateKey, "0x"+hex.EncodeToString(pk.GetPrivateKey().D.Bytes()))
			}
		} else {
			for _, key := range readKeys {
				fmt.Printf("Valdiate bls key: %s with password %s\n", key.FilePath, key.Password)

				blsKey, err := bls.CreateNewBLSKey()
				require.NoError(t, err)

				pk, err := blsKey.Read(key.FilePath, key.Password)
				if err != nil {
					assert.Fail(t, "Error reading bls key")
					break
				}
				assert.Equal(t, key.PrivateKey, pk.GetKeyPair().PrivKey.String())
			}
		}
	}
}
