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

const grpc_adress = "121.36.209.102:18987"
const authorization = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJXZ2lDZVNIbVVoRTJqbmFxM0QwX2JOeXM4amdFdDM0cjJIMGdMdzFyOGdZIn0.eyJleHAiOjE2NjYyNTg1MzUsImlhdCI6MTY2NTM5NDUzNSwianRpIjoiOGY1Njc4NWItMmMzMS00NjY0LTk2NTQtMjFhOGZmMDBhMzc2IiwiaXNzIjoiaHR0cDovL3JpdnNwYWNlLWtleWNsb2FrOjgwODAvYXV0aC9yZWFsbXMvcml2c3BhY2UiLCJzdWIiOiJlZGM5ZjM5Yi01NDE2LTRkMDktYmQwYi1mN2FkMjlhNDM3YTYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJhZG1pbi1jbGkiLCJzZXNzaW9uX3N0YXRlIjoiNjIyMzk1YzAtMjdmNy00MzczLTk0NjYtMWFhNmNhMThkMDFlIiwiYWNyIjoiMSIsInNjb3BlIjoicHJvZmlsZSBlbWFpbCBwaG9uZSByaXZzcGFjZS1hcHAiLCJzaWQiOiI2MjIzOTVjMC0yN2Y3LTQzNzMtOTQ2Ni0xYWE2Y2ExOGQwMWUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm9yZ0NvZGUiOiJvcmctMDAxIiwibmFtZSI6IjIwMjIwNzA3MTUyMzE1VlQ1cFB4IiwicHJlZmVycmVkX3VzZXJuYW1lIjoiMjAyMjA3MDcxNTIzMTV2dDVwcHgiLCJ1c2VyTmFtZSI6InJpdnNwYWNlIiwiZ2l2ZW5fbmFtZSI6IjIwMjIwNzA3MTUyMzE1VlQ1cFB4IiwiYXBwbGljYXRpb25Db2RlIjoiYXBwLTc0MjQ0MjY0ODk0OTg4Njk3NiIsInVzZXJDb2RlIjoicml2c3BhY2UiLCJhcHBsaWNhdGlvbk5hbWUiOiLlpKnkuZDmtYvor5UiLCJlbWFpbCI6ImFwcC03NDI0NDI2NDg5NDk4ODY5NzZAZXhhbXBsZS5jb20ifQ.UuWeDEANCB0DPT-EkFXeA-3KM7buNajP9T3_Chi6BfPoo_teJEctDwoJ497-SPCZWa0Aq5DkOyDsgKtIiAlEq8gL0nw4OKFXzi6KgJO4Kx9M1NzCJI2xaZGxzM_oXhH4P1CDg6CPrXx3rWSPWpL7L5OcmxBezDNL17YHDGI9eP_QxYS3px-L17rBXdRIA01hXcVN4CG0HOhUn2i06mGqCrSy7dnYk_rv1udetSJWBWskHfD-8270V_tnoy9BbFptAPS_IGkNMuL7CcdcdfkFdRwCiVotwyEyOEX2plUv24NlDQlPBhAEV1-KIo4YR1NjxR63mWp3aMRoZAhdcPsWTQ"
const chain_id = "chain-762429652294832128"

const contract_address = "0xc86de7ea867fc549548d90a1aa13fc23f42da3bd"
const data = "0xe00fe2eb"

func TestGetBlockNumber(t *testing.T) {
	// 创建客户端
	cfg, err := sdktypes.NewClientConfig(grpc_adress, grpc_adress)
	if err != nil {
		t.Fatal(err.Error())
	}

	client, err := sdk.NewClient(cfg)
	if err != nil {
		t.Fatal(err.Error())
	}

	// 测试获取区块高度
	blockNumber, err := client.Controller.GetBlockNumber(false, authorization, chain_id)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(blockNumber)

	// 测试获取系统配置
	systemConfig, err := client.Controller.GetSystemConfig(authorization, chain_id)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(systemConfig.BlockInterval)

	// 测试发送交易
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	name := strconv.Itoa(i)
	keypair, err := client.Key.Generate(name, "123")
	if err != nil {
		t.Fatal(err.Error())
	}
	req := controller.SendRequest{
		To:   contract_address,
		Data: data,
	}
	hash, err := client.Controller.SendTx(keypair, req, authorization, chain_id)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(hash)
}
