package kms

import (
	"errors"

	"github.com/GTLiSunnyi/cita-sdk-go/crypto"
	"github.com/GTLiSunnyi/cita-sdk-go/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type kmsClient struct {
	client *grpc.ClientConn
}

func NewClient(client *grpc.ClientConn) Client {
	return &kmsClient{
		client: client,
	}
}

func (client kmsClient) SignMessage(header types.GrpcRequestHeader, tx_hash []byte) ([]byte, error) {
	// 设置 grpc context
	ctx, cancel := types.MakeGrpcRequestCtx(header)
	defer cancel()

	callRes, err := crypto.NewKmsServiceClient(client.client).SignMessage(ctx, &crypto.SignMessageRequest{Msg: tx_hash})
	if err != nil {
		//获取错误状态
		statu, ok := status.FromError(err)
		if ok {
			//判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded {
				return nil, errors.New("请求超时")
			}
		}
		return nil, err
	}

	return callRes.GetSignature(), nil
}
