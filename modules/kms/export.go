package kms

import "github.com/GTLiSunnyi/cita-sdk-go/types"

type Client interface {
	SignMessage(header types.GrpcRequestHeader, tx_hash []byte) ([]byte, error)
}
