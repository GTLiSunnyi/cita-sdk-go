package executor

import (
	"errors"
	"fmt"

	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	status "google.golang.org/grpc/status"

	"github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/contract"
	"github.com/GTLiSunnyi/cita-sdk-go/utils"
)

type executorClient struct {
	client *grpc.ClientConn
}

func NewClient(grpc_addr string) (Client, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client, err := grpc.Dial(grpc_addr, dialOpts...)
	if err != nil {
		return nil, err
	}

	return executorClient{
		client: client,
	}, nil
}

func (client executorClient) Call(header types.GrpcRequestHeader, contract *contract.Contract, fromAddress string, funcName string, params []interface{}, res interface{}) error {
	data, err := contract.Abi.Pack(funcName, params...)
	if err != nil {
		fmt.Printf("abi.Pack failed: %v\n", err)
		return err
	}

	from, err := utils.ParseAddress(fromAddress)
	if err != nil {
		return err
	}
	to, err := utils.ParseAddress(contract.Address)
	if err != nil {
		return err
	}

	callReq := &CallRequest{
		To:     to,
		From:   from,
		Method: data,
		Args:   [][]byte{},
	}

	// 设置 grpc context
	ctx, cancel := types.MakeGrpcRequestCtx(header)
	defer cancel()

	callRes, err := NewExecutorServiceClient(client.client).Call(ctx, callReq)
	if err != nil {
		//获取错误状态
		statu, ok := status.FromError(err)
		if ok {
			//判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded {
				return errors.New("请求超时")
			}
		}
		return err
	}

	data = callRes.GetValue()[2:]

	return contract.Abi.UnpackIntoInterface(res, funcName, data)
}
