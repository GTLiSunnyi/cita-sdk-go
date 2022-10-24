package evm

import (
	"errors"
	"strings"

	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/contract"
	"github.com/GTLiSunnyi/cita-sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type evmClient struct {
	grpcClient *grpc.ClientConn
}

func NewClient(grpcClient *grpc.ClientConn) Client {
	return &evmClient{
		grpcClient: grpcClient,
	}
}

func (client evmClient) NewContract(header types.GrpcRequestHeader, contractAddress string) (*contract.Contract, error) {
	bytesAbi, err := client.GetAbi(header, contractAddress)
	if err != nil {
		return nil, err
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(bytesAbi)))
	if err != nil {
		return nil, err
	}

	return &contract.Contract{
		Address: contractAddress,
		Abi:     contractAbi,
	}, nil
}

func (client evmClient) GetAbi(header types.GrpcRequestHeader, contractAddress string) ([]byte, error) {
	address, err := utils.ParseData(contractAddress)
	if err != nil {
		return nil, err
	}
	// 设置 grpc context
	ctx, cancel := types.MakeGrpcRequestCtx(header)
	defer cancel()

	res, err := NewRPCServiceClient(client.grpcClient).GetAbi(ctx, &types.Address{Address: address})
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
