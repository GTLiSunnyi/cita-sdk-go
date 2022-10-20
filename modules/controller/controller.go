package controller

import (
	"encoding/hex"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	grpcproto "google.golang.org/protobuf/proto"

	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/utils"
)

type controllerClient struct {
	client *grpc.ClientConn
}

func NewClient(client *grpc.ClientConn) Client {
	return &controllerClient{
		client: client,
	}
}

// 获取区块高度
// if set for_padding, get block number of the pending block
func (client controllerClient) GetBlockNumber(for_padding bool, header sdktypes.GrpcRequestHeader) (uint64, error) {
	// 设置 grpc context
	ctx, cancel := sdktypes.MakeGrpcRequestCtx(header)
	defer cancel()

	flag := &Flag{Flag: for_padding}

	callRes, err := NewRPCServiceClient(client.client).GetBlockNumber(ctx, flag)
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
func (client controllerClient) GetSystemConfig(header sdktypes.GrpcRequestHeader) (*SystemConfig, error) {
	// 设置 grpc context
	ctx, cancel := sdktypes.MakeGrpcRequestCtx(header)
	defer cancel()

	callRes, err := NewRPCServiceClient(client.client).GetSystemConfig(ctx, &sdktypes.Empty{})
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
func (client controllerClient) SendTx(keypair types.KeyPair, req SendRequest, header sdktypes.GrpcRequestHeader) (string, error) {
	to, err := utils.ParseAddress(req.To)
	if err != nil {
		return "", err
	}
	value, err := utils.ParseValue(req.Value)
	if err != nil {
		return "", err
	}
	if req.Quota == 0 {
		req.Quota = 200000
	}
	validUntilBlock, err := client.getValidUntilBlock(req.ValidUntilBlock, header)
	if err != nil {
		return "", err
	}

	systemConfig, err := client.GetSystemConfig(header)
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().Unix())
	nonce := strconv.FormatUint(rand.Uint64(), 10)

	rawTx := sdktypes.Transaction{
		Version:         systemConfig.Version,
		To:              to,
		Data:            req.Data,
		Value:           value,
		Nonce:           nonce,
		Quota:           req.Quota,
		ValidUntilBlock: validUntilBlock,
		ChainId:         systemConfig.ChainId,
	}

	hash, err := client.sendRawTx(&rawTx, keypair, header)
	if err != nil {
		return "", err
	}

	return "0x" + hex.EncodeToString(hash), nil
}

func (client controllerClient) sendRawTx(rawTx *sdktypes.Transaction, keypair types.KeyPair, header sdktypes.GrpcRequestHeader) ([]byte, error) {
	tx, err := client.signRawTx(rawTx, keypair)
	if err != nil {
		return nil, err
	}

	return client.sendRaw(tx, header)
}

func (client controllerClient) signRawTx(rawTx *sdktypes.Transaction, keypair types.KeyPair) (*sdktypes.RawTransaction, error) {
	// 序列化
	buf, err := grpcproto.Marshal(rawTx)
	if err != nil {
		return nil, err
	}

	// 哈希
	tx_hash := utils.Sm3Hash(buf)

	// 签名
	signature, err := keypair.Sign(tx_hash)
	if err != nil {
		return nil, err
	}

	var sender = keypair.GetAddressBytes()

	witness := &sdktypes.Witness{
		Signature: signature,
		Sender:    sender[:],
	}

	normalTx := &sdktypes.RawTransaction_NormalTx{
		NormalTx: &sdktypes.UnverifiedTransaction{
			Transaction:     rawTx,
			TransactionHash: tx_hash,
			Witness:         witness,
		},
	}

	return &sdktypes.RawTransaction{
		Tx: normalTx,
	}, nil
}

func (client controllerClient) sendRaw(tx *sdktypes.RawTransaction, header sdktypes.GrpcRequestHeader) ([]byte, error) {
	// 设置 grpc context
	ctx, cancel := sdktypes.MakeGrpcRequestCtx(header)
	defer cancel()

	callRes, err := NewRPCServiceClient(client.client).SendRawTransaction(ctx, tx)
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

	return callRes.GetHash(), nil
}

func (client controllerClient) getValidUntilBlock(validUntilBlock string, header sdktypes.GrpcRequestHeader) (uint64, error) {
	if validUntilBlock == "" {
		validUntilBlock = "+95"
	}
	blockNumber, err := client.GetBlockNumber(false, header)
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

func (client controllerClient) GetTransaction(header sdktypes.GrpcRequestHeader, tx_hash []byte) (*sdktypes.UnverifiedTransaction, error) {
	// 设置 grpc context
	ctx, cancel := sdktypes.MakeGrpcRequestCtx(header)
	defer cancel()

	callRes, err := NewRPCServiceClient(client.client).GetTransaction(ctx, &sdktypes.Hash{Hash: tx_hash})
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

	return callRes.GetNormalTx(), nil
}
