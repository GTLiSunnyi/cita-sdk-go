package types

import (
	"encoding/hex"
	"math/big"
)

type KeyType string

const (
	Sm2Type     = KeyType("sm2")
	AddressSize = 20
)

type Address [AddressSize]byte

func (a Address) String() string {
	return hex.EncodeToString(a[:])
}

type KeyPair interface {
	GetPrivateKey() string
	GetPrivateKeyBytes() []byte
	GetPrivateKeyBigint() *big.Int
	GetPublicKey() string
	GetPublicKeyBytes() []byte
	GetAddressBytes() Address
	GetAddressString() string
	Type() KeyType
	Sign(msg []byte) ([]byte, error)
	Verify(msg []byte, sig []byte) bool
}
