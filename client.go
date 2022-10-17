package sdk

import (
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

	key := key.NewClient(cfg.Algo, cfg.FileManager)
	controller, err := controller.NewClient(cfg.GrpcAddr)
	if err != nil {
		return nil, err
	}
	executor, err := executor.NewClient(cfg.GrpcAddr)
	if err != nil {
		return nil, err
	}
	rivSpace := rivSpace.NewClient(cfg.RivSpaceAddress)

	return &Client{
		Key:        key,
		Controller: controller,
		Executor:   executor,
		RivSpace:   rivSpace,
	}, nil
}
