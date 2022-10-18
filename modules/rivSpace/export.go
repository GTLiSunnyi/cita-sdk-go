package rivSpace

import (
	"github.com/GTLiSunnyi/cita-sdk-go/types"
)

type Client interface {
	Send(params map[string]interface{}, header types.GrpcRequestHeader) (*types.Receipt, error)
	CreateAccount(name, appId, appSecret string, header types.GrpcRequestHeader) (string, error)
	GetReceipt(tx_hash string, header types.GrpcRequestHeader) (*types.Receipt, error)
}
