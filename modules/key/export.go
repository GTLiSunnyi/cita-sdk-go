package key

import (
	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
)

type Client interface {
	Generate(name, password string) (types.KeyPair, error)
	Get(name, password string) (types.KeyPair, error)
	Import(name, password, privKeyStr string) (types.KeyPair, error)
}
