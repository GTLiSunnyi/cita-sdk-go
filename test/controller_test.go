package test

import (
	"testing"

	sdk "github.com/GTLiSunnyi/cita-sdk-go"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
)

func TestGetBlockNumber(t *testing.T) {
	// 创建客户端
	cfg, err := sdktypes.NewClientConfig("121.36.209.102:18987", "121.36.209.102:18987")
	if err != nil {
		t.Fatal(err.Error())
	}

	client, err := sdk.NewClient(cfg)
	if err != nil {
		t.Fatal(err.Error())
	}

	// 测试获取区块高度
	blockNumber, err := client.Controller.GetBlockNumber(false)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(blockNumber)
}
