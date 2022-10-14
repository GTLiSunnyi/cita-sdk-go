package types

type KeyType string

const (
	Sm2Type     = KeyType("sm2")
	AddressSize = 20
)

type KeyPair interface {
	GetPrivateKey() string
	GetPublicKey() string
	GetAddress() []byte
	GetAddressString() string
	Type() KeyType
	Sign(msg []byte) ([]byte, error)
	Verify(msg []byte, sig []byte) bool
}
