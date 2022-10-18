package rivSpace

import (
	"github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/contract"
)

type Client interface {
	Send(params map[string]interface{}, header types.GrpcRequestHeader) (Receipt, error)
	CreateAccount(name, appId, appSecret string, header types.GrpcRequestHeader) (string, error)
	GetReceipt(tx_hash string, header types.GrpcRequestHeader) (Receipt, error)
	GetEvent(contract contract.Contract, receipt Receipt, funcSignature, eventName string) ([]byte, error)
}
