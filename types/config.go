package types

import (
	"errors"

	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/store"
	"github.com/mitchellh/go-homedir"
)

const (
	defaultPath = ".cita-sdk-go"
	defaultAlgo = types.KeyType("sm2")
)

type ClientConfig struct {
	Controller_addr string
	Executor_addr   string
	// PrivKeyArmor DAO Implements
	FileManager store.FileManager
	Algo        types.KeyType
}

func NewClientConfig(controller_addr, executor_addr string, options ...Option) (ClientConfig, error) {
	cfg := ClientConfig{
		Controller_addr: controller_addr,
		Executor_addr:   executor_addr,
	}
	for _, optionFn := range options {
		if err := optionFn(&cfg); err != nil {
			return ClientConfig{}, err
		}
	}

	if err := cfg.checkAndSetDefault(); err != nil {
		return ClientConfig{}, err
	}
	return cfg, nil
}

func (cfg *ClientConfig) checkAndSetDefault() error {
	if len(cfg.Controller_addr) == 0 {
		return errors.New("controller_addr can not be empty")
	}
	if len(cfg.Executor_addr) == 0 {
		return errors.New("executor_addr can not be empty")
	}
	if err := AlgoOption(cfg.Algo)(cfg); err != nil {
		return err
	}
	if err := FileManagerOption(cfg.FileManager)(cfg); err != nil {
		return err
	}

	return nil
}

type Option func(cfg *ClientConfig) error

func AlgoOption(algo types.KeyType) Option {
	return func(cfg *ClientConfig) error {
		if algo == "" {
			algo = defaultAlgo
		}
		cfg.Algo = algo
		return nil
	}
}

func FileManagerOption(fm store.FileManager) Option {
	return func(cfg *ClientConfig) error {
		home, err := homedir.Dir()
		if err != nil {
			return err
		}
		if fm.Path == "" {
			fm.Path = home + "/" + defaultPath
		}
		cfg.FileManager = fm
		return nil
	}
}
