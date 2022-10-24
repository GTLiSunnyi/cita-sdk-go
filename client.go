package sdk

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/GTLiSunnyi/cita-sdk-go/modules/controller"
	"github.com/GTLiSunnyi/cita-sdk-go/modules/evm"
	"github.com/GTLiSunnyi/cita-sdk-go/modules/executor"
	"github.com/GTLiSunnyi/cita-sdk-go/modules/key"
	"github.com/GTLiSunnyi/cita-sdk-go/modules/rivSpace"
	sdktypes "github.com/GTLiSunnyi/cita-sdk-go/types"
)

type Client struct {
	Key        key.Client
	Controller controller.Client
	Executor   executor.Client
	Evm        evm.Client
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
	grpcClient, err := grpc.Dial(cfg.GrpcAddr, dialOpts...)
	if err != nil {
		return nil, err
	}

	key := key.NewClient(cfg.Algo, cfg.FileManager)
	controller := controller.NewClient(grpcClient)
	executor := executor.NewClient(grpcClient)
	rivSpace := rivSpace.NewClient(cfg.RivSpaceAddress)
	evm := evm.NewClient(grpcClient)

	return &Client{
		Key:        key,
		Controller: controller,
		Executor:   executor,
		Evm:        evm,
		RivSpace:   rivSpace,
	}, nil
}
