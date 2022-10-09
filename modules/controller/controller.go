package controller

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
	"github.com/GTLiSunnyi/cita-sdk-go/protos/proto"
	sdk "github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/utils"
)

type controllerClient struct {
	client *grpc.ClientConn
}

func NewClient(controller_addr string) (Client, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client, err := grpc.Dial(controller_addr, dialOpts...)
	if err != nil {
		return nil, err
	}

	return controllerClient{
		client: client,
	}, nil
}

// 获取区块高度
// if set for_padding, get block number of the pending block
func (controller controllerClient) GetBlockNumber(for_padding bool) (uint64, error) {
	flag := &proto.Flag{Flag: for_padding}

	gRpcClient := proto.NewRPCServiceClient(controller.client)

	// 设置 grpc 超时时间
	clientDeadline := time.Now().Add(sdk.GRPC_TIMEOUT)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	callRes, err := gRpcClient.GetBlockNumber(ctx, flag)
	if err != nil {
		//获取错误状态
		statu, ok := status.FromError(err)
		if ok {
			//判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded {
				return 0, errors.New("请求超时")
			}
		}
		return 0, err
	}

	return callRes.GetBlockNumber(), nil
}

// 获取系统配置
func (controller controllerClient) GetSystemConfig() (*proto.SystemConfig, error) {
	gRpcClient := proto.NewRPCServiceClient(controller.client)

	// 设置 grpc 超时时间
	clientDeadline := time.Now().Add(sdk.GRPC_TIMEOUT)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	callRes, err := gRpcClient.GetSystemConfig(ctx, &proto.Empty{})
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

	return callRes, nil
}

// 发送交易
func (controller controllerClient) SendTx(keypair types.KeyPair, req SendRequest) error {
	to, err := utils.ParseData(req.To)
	if err != nil {
		return err
	}
	data, err := utils.ParseData(req.Data)
	if err != nil {
		return err
	}
	value, err := utils.ParseData(req.Value)
	if err != nil {
		return err
	}
	validUntilBlock, err := controller.getValidUntilBlock(req.ValidUntilBlock)
	if err != nil {
		return err
	}

	systemConfig, err := controller.GetSystemConfig()
	if err != nil {
		return err
	}

	rand.Seed(time.Now().Unix())
	nonce := strconv.FormatUint(rand.Uint64(), 10)

	rawTx := proto.Transaction{
		Version:         systemConfig.Version,
		To:              to,
		Data:            data,
		Value:           value,
		Nonce:           nonce,
		Quota:           req.Quota,
		ValidUntilBlock: validUntilBlock,
		ChainId:         systemConfig.ChainId,
	}

	controller.SendRawTx(rawTx, keypair)

	return nil
}

func (controller controllerClient) SendRawTx(rawTx proto.Transaction, keypair types.KeyPair) error {
	// keypair.Sign()
	return nil
}

func (controller controllerClient) SignRawTx() {

}

func (controller controllerClient) getValidUntilBlock(validUntilBlock string) (uint64, error) {
	blockNumber, err := controller.GetBlockNumber(false)
	if err != nil {
		return 0, err
	}

	if validUntilBlock[0] == '+' {
		num, err := strconv.ParseUint(strings.Split(validUntilBlock, "+")[1], 10, 64)
		if err != nil {
			return 0, err
		}

		return blockNumber + num, nil
	} else if validUntilBlock[0] == '-' {
		num, err := strconv.ParseUint(strings.Split(validUntilBlock, "-")[1], 10, 64)
		if err != nil {
			return 0, err
		}

		return blockNumber - num, nil
	} else {
		num, err := strconv.ParseUint(validUntilBlock, 10, 64)
		if err != nil {
			return 0, err
		}

		return num, nil
	}
}
