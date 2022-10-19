package executor

import (
	"errors"
	"fmt"

	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/contract"
	"github.com/GTLiSunnyi/cita-sdk-go/utils"
)

type executorClient struct {
	client *grpc.ClientConn
}

func NewClient(client *grpc.ClientConn) Client {
	return &executorClient{
		client: client,
	}
}

// params：填函数的参数，可以传入 big.Int\[]byte\string，例如：[]interface{}{big.NewInt(10), []byte{1}}
func (client executorClient) Call(header types.GrpcRequestHeader, contract contract.Contract, userAddress string, funcName string, params []interface{}, res interface{}) error {
	data, err := contract.Abi.Pack(funcName, params...)
	if err != nil {
		return err
	}

	from, err := utils.ParseAddress(userAddress)
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

	data = callRes.GetValue()

	if len(data)%32 != 0 {
		return fmt.Errorf("data 的长度不是32的倍数, length: %d", len(data))
	}

	err = contract.Abi.Unpack(res, funcName, data)
	return err
}
