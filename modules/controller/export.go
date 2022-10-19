package controller

import (
	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
)

type Client interface {
	GetBlockNumber(for_padding bool, header sdktypes.GrpcRequestHeader) (uint64, error)
	GetSystemConfig(header sdktypes.GrpcRequestHeader) (*SystemConfig, error)
	SendTx(keypair types.KeyPair, req SendRequest, header sdktypes.GrpcRequestHeader) (string, error)
	GetTransaction(header sdktypes.GrpcRequestHeader, tx_hash []byte) (*sdktypes.UnverifiedTransaction, error)
}
