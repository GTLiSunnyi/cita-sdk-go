package controller

import (
	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
)

type Client interface {
	GetBlockNumber(header sdktypes.GrpcRequestHeader, for_padding bool) (uint64, error)
	GetSystemConfig(header sdktypes.GrpcRequestHeader) (*SystemConfig, error)
	Send(header sdktypes.GrpcRequestHeader, keypair types.KeyPair, req SendRequest) (string, error)
	GetTransaction(header sdktypes.GrpcRequestHeader, tx_hash []byte) (*sdktypes.UnverifiedTransaction, error)
}
