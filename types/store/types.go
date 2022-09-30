package store

import "github.com/GTLiSunnyi/cita-sdk-go/crypto/types"

// KeyInfo saves the basic information of the key
type KeyInfo struct {
	Name       string        `json:"name"`
	IsLocked   bool          `json:"is_locked"`
	CryptoType types.KeyType `json:"crypto_type"`
	Address    string        `json:"address"`
	PublicKey  string        `json:"public_key"`
	PrivateKey string        `json:"private_key"`
}

type Crypto interface {
	Encrypt(data string, password string) (string, error)
	Decrypt(data string, password string) (string, error)
}
