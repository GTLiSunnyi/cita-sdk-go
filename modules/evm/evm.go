package evm

import (
	"errors"

	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/utils"
)

type evmClient struct {
	client *grpc.ClientConn
}

func NewClient(client *grpc.ClientConn) Client {
	return &evmClient{
		client: client,
	}
}

func (client evmClient) GetAbi(header types.GrpcRequestHeader, contractAddress string) ([]byte, error) {
	address, err := utils.ParseData(contractAddress)
	if err != nil {
		return nil, err
	}
	// 设置 grpc context
	ctx, cancel := types.MakeGrpcRequestCtx(header)
	defer cancel()

	res, err := NewRPCServiceClient(client.client).GetAbi(ctx, &types.Address{Address: address})
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

	return res.GetBytesAbi(), nil
}
