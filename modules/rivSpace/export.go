package rivSpace

import (
	"github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/contract"
)

type Client interface {
	Send(header types.GrpcRequestHeader, params map[string]interface{}) (*types.MyReceipt, error)
	SendAndGetEvent(header types.GrpcRequestHeader, contract *contract.Contract, params map[string]interface{}, eventName string, event interface{}) error
	CreateAccount(header types.GrpcRequestHeader, name, appId, appSecret string) (string, error)
	GetReceipt(header types.GrpcRequestHeader, tx_hash string) (*types.MyReceipt, error)
}
