package contract

import (
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Contract struct {
	Abi     abi.ABI
	Address string
}

func NewContract(address string, abiPath string) (Contract, error) {
	abiFile, err := os.ReadFile(abiPath)
	if err != nil {
		return Contract{}, err
	}
	abiStr := string(abiFile)

	ethAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return Contract{}, err
	}

	return Contract{
		Abi:     ethAbi,
		Address: address,
	}, nil
}
