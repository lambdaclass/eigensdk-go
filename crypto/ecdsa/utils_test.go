package ecdsa

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestECDSAPrivateKey(t *testing.T) {
	var tests = map[string]struct {
		keyPath  string
		password string
		wantErr  bool
	}{
		"valid ecdsa key save": {
			keyPath:  "./operator_keys_test_directory/test.ecdsa.key.json",
			password: "test",
			wantErr:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Cleanup(func() {
				dir := filepath.Dir(tt.keyPath)
				_ = os.RemoveAll(dir)
			})
			key, err := CreateNewEcdsaKey()
			assert.NoError(t, err)

			err = key.Save(tt.keyPath, tt.password)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			readKeyPair, err := key.Read(tt.keyPath, tt.password)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, key, readKeyPair)
			}
		})
	}
}
