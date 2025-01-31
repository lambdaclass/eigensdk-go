package ecdsa

import (
	"bufio"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

type ECDSAKey struct {
	key *ecdsa.PrivateKey
}

// Save encrypts and stores the private key at the given path
func (e ECDSAKey) Save(path string, password string) error {
	UUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	// We are using https://github.com/ethereum/go-ethereum/blob/master/accounts/keystore/key.go#L41
	// to store the keys which requires us to have random UUID for encryption
	key := &keystore.Key{
		Id:         UUID,
		Address:    crypto.PubkeyToAddress(e.key.PublicKey),
		PrivateKey: e.key,
	}

	encryptedBytes, err := keystore.EncryptKey(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return err
	}

	return writeBytesToFile(path, encryptedBytes)
}

// Read loads an encrypted private key from the given path
func (e *ECDSAKey) Read(path string, password string) (ECDSAKey, error) {
	keyStoreContents, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return ECDSAKey{}, err
	}

	sk, err := keystore.DecryptKey(keyStoreContents, password)
	if err != nil {
		return ECDSAKey{}, err
	}

	return ECDSAKey{key: sk.PrivateKey}, nil
}

// CreateNewEcdsaKey generates a new ECDSA private key
func CreateNewEcdsaKey() (ECDSAKey, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return ECDSAKey{}, err
	}
	return ECDSAKey{key: privateKey}, nil
}

// CreateNewEcdsaKey generates a new ECDSA private key
func CreateNewEcdsaKeyFromHex(privateKeyHex string) (ECDSAKey, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return ECDSAKey{}, err
	}
	return ECDSAKey{key: privateKey}, nil
}

func (e ECDSAKey) GetPrivateKey() *ecdsa.PrivateKey {
	return e.key
}

func writeBytesToFile(path string, data []byte) error {
	dir := filepath.Dir(path)

	// create the directory if it doesn't exist. If exists, it does nothing
	if err := os.MkdirAll(dir, 0750); err != nil {
		fmt.Println("Error creating directories:", err)
		return err
	}

	file, err := os.Create(filepath.Clean(path))
	if err != nil {
		fmt.Println("file create error")
		return err
	}
	// remember to close the file
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	_, err = file.Write(data)

	return err
}

// GetAddressFromKeyStoreFile We are using Web3 format defined by
// https://ethereum.org/en/developers/docs/data-structures-and-encoding/web3-secret-storage/
func GetAddressFromKeyStoreFile(keyStoreFile string) (common.Address, error) {
	keyJson, err := os.ReadFile(filepath.Clean(keyStoreFile))
	if err != nil {
		return common.Address{}, err
	}

	// The reason we have map[string]interface{} is because `address` is string but the `crypto` field is an object
	// we don't care about the object in this method, but we still need to unmarshal it
	m := make(map[string]interface{})
	if err := json.Unmarshal(keyJson, &m); err != nil {
		return common.Address{}, err
	}

	if address, ok := m["address"].(string); !ok {
		return common.Address{}, fmt.Errorf("address not found in key file")
	} else {
		return common.HexToAddress(address), nil
	}
}

func KeyAndAddressFromHexKey(hexkey string) (*ecdsa.PrivateKey, common.Address, error) {
	hexkey = strings.TrimPrefix(hexkey, "0x")
	ecdsaSk, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to convert hexkey %s to ecdsa key: %w", hexkey, err)
	}
	pk := ecdsaSk.Public()
	address := crypto.PubkeyToAddress(*pk.(*ecdsa.PublicKey))
	return ecdsaSk, address, nil
}
