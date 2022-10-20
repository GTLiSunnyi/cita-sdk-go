package rivSpace

import (
	"github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/contract"
)

type Client interface {
	Send(params map[string]interface{}, header types.GrpcRequestHeader) (*types.MyReceipt, error)
	SendAndGetEvent(contract *contract.Contract, params map[string]interface{}, header types.GrpcRequestHeader, eventName string, res interface{}) error
	CreateAccount(name, appId, appSecret string, header types.GrpcRequestHeader) (string, error)
	GetReceipt(tx_hash string, header types.GrpcRequestHeader) (*types.MyReceipt, error)
}
