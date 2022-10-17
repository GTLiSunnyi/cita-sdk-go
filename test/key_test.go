package test

import (
	"testing"

	sdk "github.com/GTLiSunnyi/cita-sdk-go"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
)

func TestKey(t *testing.T) {
	// 创建客户端
	cfg, err := sdktypes.NewClientConfig(GrpcAddress)
	if err != nil {
		t.Fatal(err.Error())
	}

	client, err := sdk.NewClient(cfg)
	if err != nil {
		t.Fatal(err.Error())
	}

	// 测试密钥生成
	keypair1, err := client.Key.Generate("test", "123")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(keypair1.GetAddressString())

	// 测试读取密钥
	keypair2, err := client.Key.Get("test", "123")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(keypair2.GetAddressString())

	// 测试导入密钥
	keypair3, err := client.Key.Import("changename1", "1234", keypair1.GetPrivateKey())
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(keypair3.GetAddressString())

	// 测试签名、验证
	msg := []byte("testmsg")
	sig, err := keypair2.Sign(msg)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(keypair3.Verify(msg, sig))
}
