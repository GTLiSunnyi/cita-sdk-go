package sdk

import (
	"github.com/GTLiSunnyi/cita-sdk-go/modules/controller"
	"github.com/GTLiSunnyi/cita-sdk-go/modules/executor"
	"github.com/GTLiSunnyi/cita-sdk-go/modules/key"
	sdk "github.com/GTLiSunnyi/cita-sdk-go/types"
)

type Client struct {
	Key        key.Client
	Controller controller.Client
	Executor   executor.Client
}

func NewClient(cfg sdk.ClientConfig) (*Client, error) {
	// 初始化目录
	if err := cfg.FileManager.CreateRootDir(); err != nil {
		return nil, err
	}

	key := key.NewClient(cfg.Algo, cfg.FileManager)
	controller, err := controller.NewClient(cfg.Controller_addr)
	if err != nil {
		return nil, err
	}
	executor, err := executor.NewClient(cfg.Executor_addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		Key:        key,
		Controller: controller,
		Executor:   executor,
	}, nil
}
