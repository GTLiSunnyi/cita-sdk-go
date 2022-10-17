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
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	grpcproto "google.golang.org/protobuf/proto"

	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
	"github.com/GTLiSunnyi/cita-sdk-go/protos/proto"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
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
func (client controllerClient) GetBlockNumber(for_padding bool, header sdktypes.GrpcRequestHeader) (uint64, error) {
	flag := &Flag{Flag: for_padding}

	gRpcClient := NewRPCServiceClient(client.client)

	// 设置 grpc context
	ctx, cancel := sdktypes.MakeGrpcRequestCtx(header)
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
func (client controllerClient) GetSystemConfig(header sdktypes.GrpcRequestHeader) (*SystemConfig, error) {
	gRpcClient := NewRPCServiceClient(client.client)

	// 设置 grpc context
	ctx, cancel := sdktypes.MakeGrpcRequestCtx(header)
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
func (client controllerClient) SendTx(keypair types.KeyPair, req SendRequest, header sdktypes.GrpcRequestHeader) (string, error) {
	to, err := utils.ParseAddress(req.To)
	if err != nil {
		return "", err
	}
	data, err := utils.ParseData(req.Data)
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
		Data:            data,
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

	// dialOpts := []grpc.DialOption{
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// }

	// const grpc_adress = "121.36.209.102:18987"
	// const authorization = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJXZ2lDZVNIbVVoRTJqbmFxM0QwX2JOeXM4amdFdDM0cjJIMGdMdzFyOGdZIn0.eyJleHAiOjE2NjY2NzMzMzksImlhdCI6MTY2NTgwOTMzOSwianRpIjoiNWVkNmJlNTMtNTc0ZS00MmM5LTk4ZDktOWMxY2NhMjliNzk3IiwiaXNzIjoiaHR0cDovL3JpdnNwYWNlLWtleWNsb2FrOjgwODAvYXV0aC9yZWFsbXMvcml2c3BhY2UiLCJzdWIiOiJlZGM5ZjM5Yi01NDE2LTRkMDktYmQwYi1mN2FkMjlhNDM3YTYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJhZG1pbi1jbGkiLCJzZXNzaW9uX3N0YXRlIjoiOGJmZTg3NjgtM2E5OS00YThlLWEyM2MtMjc1ZTVkZTA4MzA3IiwiYWNyIjoiMSIsInNjb3BlIjoicHJvZmlsZSBlbWFpbCBwaG9uZSByaXZzcGFjZS1hcHAiLCJzaWQiOiI4YmZlODc2OC0zYTk5LTRhOGUtYTIzYy0yNzVlNWRlMDgzMDciLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm9yZ0NvZGUiOiJvcmctMDAxIiwibmFtZSI6IjIwMjIwNzA3MTUyMzE1VlQ1cFB4IiwicHJlZmVycmVkX3VzZXJuYW1lIjoiMjAyMjA3MDcxNTIzMTV2dDVwcHgiLCJ1c2VyTmFtZSI6InJpdnNwYWNlIiwiZ2l2ZW5fbmFtZSI6IjIwMjIwNzA3MTUyMzE1VlQ1cFB4IiwiYXBwbGljYXRpb25Db2RlIjoiYXBwLTc0MjQ0MjY0ODk0OTg4Njk3NiIsInVzZXJDb2RlIjoicml2c3BhY2UiLCJhcHBsaWNhdGlvbk5hbWUiOiLlpKnkuZDmtYvor5UiLCJlbWFpbCI6ImFwcC03NDI0NDI2NDg5NDk4ODY5NzZAZXhhbXBsZS5jb20ifQ.DbmxMkLuPg7XZJuf98U6ev4Ae4oA3rBsntSJwXJVc53frNYe1nUKheKO1XygmDNF_0jvsPXuABQv3ESKvMi5aPhswpcWVeNeVELx1D-o-B1JqX4LhS3Bp4qiJgS8hhhtYRc0odZ6Fz_mOZTy7G5zQyWwa7LpGUdSXpcCi-be0Iegdv4p5fDN_3AIcYVtSB2Olr8x1p7ZFY69AZQg32KnWEOgKybIdrq9mTpPukyv-zBK5D2yRD3OD05019laUXdNeNSbN6eiRIFow__84ctOfuV3-WneM4C9-nJHsIxmDclinOLZXa7JzSbow6xmXBRCL0b533DHBmNGHQdlPCfHsw"
	// const chain_code = "chain-762429652294832128"
	// const app_user_code = "tianle"
	// grpcclient, err := grpc.Dial(grpc_adress, dialOpts...)
	// if err != nil {
	// 	return nil, err
	// }

	// gRpcClient := crypto.NewKmsServiceClient(grpcclient)

	// // 构造请求头
	// header := sdktypes.GrpcRequestHeader{
	// 	XAuthorization: authorization,
	// 	ChainCode:      chain_code,
	// 	AppUserCode:    app_user_code,
	// }

	// // 设置 grpc context
	// ctx, cancel := sdktypes.MakeGrpcRequestCtx(header)
	// defer cancel()

	// callRes, err := gRpcClient.SignMessage(ctx, &crypto.SignMessageRequest{Msg: tx_hash})
	// if err != nil {
	// 	//获取错误状态
	// 	statu, ok := status.FromError(err)
	// 	if ok {
	// 		//判断是否为调用超时
	// 		if statu.Code() == codes.DeadlineExceeded {
	// 			return nil, errors.New("请求超时")
	// 		}
	// 	}
	// 	return nil, err
	// }

	// fmt.Println(callRes.GetStatus())
	// fmt.Println(callRes.GetSignature())
	// fmt.Println(signature)

	witness := &sdktypes.Witness{
		Signature: signature,
		Sender:    keypair.GetAddressBytes(),
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
	gRpcClient := NewRPCServiceClient(client.client)

	// 设置 grpc context
	ctx, cancel := sdktypes.MakeGrpcRequestCtx(header)
	defer cancel()

	callRes, err := gRpcClient.SendRawTransaction(ctx, tx)
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
