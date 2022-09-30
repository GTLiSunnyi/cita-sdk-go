package types

type KeyType string

const (
	Sm2Type = KeyType("sm2")
)

type KeyPair interface {
	GetPrivateKey() string
	GetPublicKey() string
	GetAddress() string
	Type() KeyType
	Sign(msg []byte) ([]byte, error)
	Verify(msg []byte, sig []byte) bool
}
