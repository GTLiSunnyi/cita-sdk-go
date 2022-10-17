package types

import (
	"errors"

	"github.com/mitchellh/go-homedir"

	"github.com/GTLiSunnyi/cita-sdk-go/crypto/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/store"
)

const (
	defaultPath                 = ".cita-sdk-go"
	defaultAlgo                 = types.KeyType("sm2")
	defaultRunAddress           = "http://portal-api-test.taidihub.com/contract-management-service-chain-762429652294832128/api/v1/manage/run"
	defaultReceiptAddress       = "http://portal-api-test.taidihub.com/contract-management-service-chain-762429652294832128/api/v1/manage/transaction"
	defaultCreateAccountAddress = "http://portal-api-test.taidihub.com/management-platform/api/v1/application/user/create"
)

type ClientConfig struct {
	GrpcAddr string

	// PrivKeyArmor DAO Implements
	FileManager store.FileManager
	Algo        types.KeyType

	RivSpaceAddress
}

func NewClientConfig(grpc_addr string, options ...Option) (ClientConfig, error) {
	cfg := ClientConfig{
		GrpcAddr: grpc_addr,
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
	if len(cfg.GrpcAddr) == 0 {
		return errors.New("GrpcAddr can not be empty")
	}
	if err := AlgoOption(cfg.Algo)(cfg); err != nil {
		return err
	}
	if err := FileManagerOption(cfg.FileManager)(cfg); err != nil {
		return err
	}
	if err := RivSpaceAddressOption(cfg.RivSpaceAddress)(cfg); err != nil {
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

func RivSpaceAddressOption(rivSpaceAddress RivSpaceAddress) Option {
	return func(cfg *ClientConfig) error {
		if rivSpaceAddress.RunAddress == "" {
			rivSpaceAddress = RivSpaceAddress{
				defaultRunAddress,
				defaultReceiptAddress,
				defaultCreateAccountAddress,
			}
		}
		cfg.RivSpaceAddress = rivSpaceAddress
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
