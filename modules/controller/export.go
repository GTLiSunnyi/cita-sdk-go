package controller

import "github.com/GTLiSunnyi/cita-sdk-go/crypto/types"

type Client interface {
	GetBlockNumber(for_padding bool, authorization, chain_code string) (uint64, error)
	GetSystemConfig(authorization, chain_code string) (*SystemConfig, error)
	SendTx(keypair types.KeyPair, req SendRequest, authorization, chain_code string) ([]byte, error)
}
