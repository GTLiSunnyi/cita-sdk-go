package sdk

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/GTLiSunnyi/cita-sdk-go/modules/controller"
	"github.com/GTLiSunnyi/cita-sdk-go/modules/executor"
	"github.com/GTLiSunnyi/cita-sdk-go/modules/key"
	"github.com/GTLiSunnyi/cita-sdk-go/modules/rivSpace"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
)

type Client struct {
	Key        key.Client
	Controller controller.Client
	Executor   executor.Client
	RivSpace   rivSpace.Client
}

func NewClient(cfg sdktypes.ClientConfig) (*Client, error) {
	// 初始化目录
	if err := cfg.FileManager.CreateRootDir(); err != nil {
		return nil, err
	}

	// 创建 grpc 客户端
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client, err := grpc.Dial(cfg.GrpcAddr, dialOpts...)
	if err != nil {
		return nil, err
	}

	key := key.NewClient(cfg.Algo, cfg.FileManager)
	controller := controller.NewClient(client)
	executor := executor.NewClient(client)
	rivSpace := rivSpace.NewClient(cfg.RivSpaceAddress)

	return &Client{
		Key:        key,
		Controller: controller,
		Executor:   executor,
		RivSpace:   rivSpace,
	}, nil
}
