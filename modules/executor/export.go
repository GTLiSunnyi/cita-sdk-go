package executor

import (
	"github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/contract"
)

type Client interface {
	Call(header types.GrpcRequestHeader, contract *contract.Contract, fromAddress string, funcName string, params []interface{}, res interface{}) error
}
