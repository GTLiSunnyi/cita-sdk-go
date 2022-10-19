package test

import (
	"testing"

	sdk "github.com/GTLiSunnyi/cita-sdk-go"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/contract"
)

func TestExecutor(t *testing.T) {
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

	// 测试 call 方法
	contract, err := contract.NewContract(ContractAddress, "./abi.json")
	if err != nil {
		t.Fatal(err.Error())
	}

	// 获取结构体
	var res1 GetStruct
	err = client.Executor.Call(header, contract, UserAddress, GetStructFuncName, []interface{}{}, &res1)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(res1)

	// 获取数组
	var res2 []string
	err = client.Executor.Call(header, contract, UserAddress, GetArrFuncName, []interface{}{}, &res2)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(len(res2))
}
