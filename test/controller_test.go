package test

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	sdk "github.com/GTLiSunnyi/cita-sdk-go"
	"github.com/GTLiSunnyi/cita-sdk-go/modules/controller"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
)

func TestController(t *testing.T) {
	// 创建客户端
	cfg, err := sdktypes.NewClientConfig(GrpcAddress)
	if err != nil {
		t.Fatal(err.Error())
	}

	client, err := sdk.NewClient(cfg)
	if err != nil {
		t.Fatal(err.Error())
	}

	// 构造请求头
	header := sdktypes.GrpcRequestHeader{
		XAuthorization: Authorization,
		ChainCode:      ChainCode,
		AppUserCode:    AppUserCode,
	}

	// 测试获取区块高度
	blockNumber, err := client.Controller.GetBlockNumber(false, header)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(blockNumber)

	// 测试获取系统配置
	systemConfig, err := client.Controller.GetSystemConfig(header)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("block interval: ", systemConfig.BlockInterval)

	// 测试发送交易
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	name := strconv.Itoa(i)
	keypair, err := client.Key.Generate(name, "123")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(keypair.GetAddressString())

	req := controller.SendRequest{
		To:   ContractAddress,
		Data: Data,
	}
	hash, err := client.Controller.SendTx(keypair, req, header)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(hash))
}
