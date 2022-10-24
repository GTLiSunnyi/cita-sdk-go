package evm

import "github.com/GTLiSunnyi/cita-sdk-go/types"

type Client interface {
	GetAbi(header types.GrpcRequestHeader, contractAddress string) ([]byte, error)
}
