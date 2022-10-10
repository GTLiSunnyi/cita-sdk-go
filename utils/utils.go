package utils

import (
	"context"
	"encoding/hex"
	"strings"

	"github.com/tjfoc/gmsm/sm3"
	"google.golang.org/grpc/metadata"
)

func ParseData(str string) ([]byte, error) {
	return hex.DecodeString(strings.Split(str, "0x")[1])
}

func Sm3Hash(data []byte) []byte {
	return sm3.New().Sum(data)
}

func MakeCtxWithHeader(authorization, chain_code string) context.Context {
	md := metadata.New(map[string]string{
		"x-authorization": authorization,
		"chain_code":      chain_code,
	})

	return metadata.NewOutgoingContext(context.Background(), md)
}
