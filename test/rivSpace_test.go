package test

import (
	"encoding/json"
	"testing"

	sdk "github.com/GTLiSunnyi/cita-sdk-go"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/contract"
)

func TestRivSpace(t *testing.T) {
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

	// 测试创建账户
	address, err := client.RivSpace.CreateAccount("test", AppId, AppSecret, header)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(address)

	// 测试发送交易
	params := map[string]interface{}{
		"contractAddress": ContractAddress,
		"method":          MethodName,
		"param":           []map[string]interface{}{},
		"response":        []string{},
	}

	receipt, err := client.RivSpace.Send(params, header)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(receipt)

	// 测试获取事件
	contract, err := contract.NewContract(ContractAddress, "./abi.json")
	if err != nil {
		t.Fatal(err.Error())
	}

	data, err := contract.GetEvent(receipt, EventName)
	if err != nil {
		t.Fatal(err.Error())
	}

	var res = &AddEvent{}
	err = json.Unmarshal(data, res)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(res)
}
