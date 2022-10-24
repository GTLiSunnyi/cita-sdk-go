package evm

import (
	"github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/contract"
)

type Client interface {
	NewContract(header types.GrpcRequestHeader, contractAddress string) (*contract.Contract, error)
	GetAbi(header types.GrpcRequestHeader, contractAddress string) ([]byte, error)
}
