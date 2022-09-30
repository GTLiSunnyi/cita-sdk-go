package key

import (
	"fmt"

	"github.com/GTLiSunnyi/cita-sdk-go/crypto/keys/sm2"
	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/store"
)

type keyClient struct {
	FileManager store.FileManager
	Algo        types.KeyType
	store.Crypto
}

func NewClient(algo types.KeyType, fm store.FileManager) Client {
	return keyClient{
		FileManager: fm,
		Algo:        algo,
		Crypto:      store.SM4{},
	}
}

func (k keyClient) Generate(name, password string) (types.KeyPair, error) {
	if k.FileManager.Has(name) {
		return nil, fmt.Errorf("name %s has existed", name)
	}

	var keypair types.KeyPair
	var err error
	switch k.Algo {
	default:
		keypair, err = sm2.NewKeyPair()
	}
	if err != nil {
		return nil, err
	}

	privateKey := keypair.GetPrivateKey()

	var isLocked bool
	if password != "" {
		var err error
		privateKey, err = k.Encrypt(privateKey, password)
		if err != nil {
			return nil, err
		}
		isLocked = true
	}

	keyInfo := store.KeyInfo{
		Name:       name,
		IsLocked:   isLocked,
		CryptoType: k.Algo,
		Address:    keypair.GetAddress(),
		PublicKey:  keypair.GetPublicKey(),
		PrivateKey: privateKey,
	}

	err = k.FileManager.Write(name, keyInfo)
	return keypair, err
}

func (k keyClient) Get(name, password string) (types.KeyPair, error) {
	keyInfo, err := k.FileManager.Read(name)
	if err != nil {
		return nil, err
	}

	if keyInfo.IsLocked {
		keyInfo.PrivateKey, err = k.Decrypt(keyInfo.PrivateKey, password)
		if err != nil {
			return nil, err
		}
	}

	var keypair types.KeyPair
	switch k.Algo {
	default:
		keypair, err = sm2.ImportKeyPair(keyInfo.PrivateKey)
	}
	if err != nil {
		return nil, err
	}

	return keypair, nil
}

func (k keyClient) Import(name, password, privKeyStr string) (types.KeyPair, error) {
	if k.FileManager.Has(name) {
		return nil, fmt.Errorf("%s has existed", name)
	}

	var keypair types.KeyPair
	var err error
	switch k.Algo {
	default:
		keypair, err = sm2.ImportKeyPair(privKeyStr)
	}
	if err != nil {
		return nil, err
	}

	privateKey := keypair.GetPrivateKey()

	var isLocked bool
	if password != "" {
		var err error
		privateKey, err = k.Encrypt(privateKey, password)
		if err != nil {
			return nil, err
		}
		isLocked = true
	}

	keyInfo := store.KeyInfo{
		Name:       name,
		IsLocked:   isLocked,
		CryptoType: k.Algo,
		Address:    keypair.GetAddress(),
		PublicKey:  keypair.GetPublicKey(),
		PrivateKey: privateKey,
	}

	err = k.FileManager.Write(name, keyInfo)
	if err != nil {
		return nil, err
	}

	return keypair, nil
}

// func (k keyClient) Recover(name, password, mnemonic string) (string, sdk.Error) {
// 	address, err := k.KeyManager.Recover(name, password, mnemonic)
// 	return address, sdk.Wrap(err)
// }

// func (k keyClient) Export(name, password string) (string, sdk.Error) {
// 	keystore, err := k.KeyManager.Export(name, password)
// 	return keystore, sdk.Wrap(err)
// }

// func (k keyClient) Delete(name, password string) sdk.Error {
// 	err := k.KeyManager.Delete(name, password)
// 	return sdk.Wrap(err)
// }
