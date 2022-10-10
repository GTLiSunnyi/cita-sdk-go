package test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	sdk "github.com/GTLiSunnyi/cita-sdk-go"
	"github.com/GTLiSunnyi/cita-sdk-go/protos/proto"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
)

func TestXxx(t *testing.T) {
	// 序列化
	s1 := proto.Transaction{
		Version:         1,
		To:              []byte{1},
		Nonce:           "1",
		Quota:           1,
		ValidUntilBlock: 1,
		Data:            []byte{1},
		Value:           []byte{1},
		ChainId:         []byte{1},
	}
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer) //创建编码器
	err1 := encoder.Encode(&s1)        //编码
	if err1 != nil {
		log.Panic(err1)
	}
	fmt.Printf("序列化后：%x\n", buffer.Bytes())

	data, _ := json.Marshal(&s1)
	fmt.Printf("序列化后：%x\n", data)

}

func TestKey(t *testing.T) {
	// 创建客户端
	cfg, err := sdktypes.NewClientConfig("controller_addr", "executor_addr")
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
	t.Log(keypair1.GetAddress())

	// 测试读取密钥
	keypair2, err := client.Key.Get("test", "123")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(keypair2.GetAddress())

	// 测试导入密钥
	keypair3, err := client.Key.Import("changename1", "1234", keypair1.GetPrivateKey())
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
