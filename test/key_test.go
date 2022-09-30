package test

import (
	"testing"

	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
	"github.com/GTLiSunnyi/cita-sdk-go/modules/key"
	sdk "github.com/GTLiSunnyi/cita-sdk-go/types"
)

func TestKey(t *testing.T) {
	// 测试密钥生成
	cfg, err := sdk.NewClientConfig("controller_addr", "executor_addr")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = cfg.FileManager.CreateRootDir()
	if err != nil {
		t.Fatal(err.Error())
	}
	keyClient := key.NewClient(types.Sm2Type, cfg.FileManager)
	keypair1, err := keyClient.Generate("test", "123")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(keypair1.GetAddress())

	// 测试读取密钥
	keypair2, err := keyClient.Get("test", "123")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(keypair2.GetAddress())

	// 测试导入密钥
	keypair3, err := keyClient.Import("changename1", "1234", keypair1.GetPrivateKey())
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(keypair3.GetAddress())

	// 测试签名、验证
	msg := []byte("testmsg")
	sig, err := keypair2.Sign(msg)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(keypair3.Verify(msg, sig))
}
