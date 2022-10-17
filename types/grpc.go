package types

import (
	"context"
	"time"

	"google.golang.org/grpc/metadata"
)

const GRPC_TIMEOUT = time.Duration(10 * time.Second)

type GrpcRequestHeader struct {
	XAuthorization string
	ChainCode      string
	AppUserCode    string
}

func MakeGrpcRequestCtx(c GrpcRequestHeader) (context.Context, context.CancelFunc) {
	// 设置请求头
	md := metadata.New(map[string]string{
		"x-authorization": c.XAuthorization,
		"chain_code":      c.ChainCode,
		"app-user-code":   c.AppUserCode,
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 设置 grpc 超时时间
	clientDeadline := time.Now().Add(GRPC_TIMEOUT)
	return context.WithDeadline(ctx, clientDeadline)
}
