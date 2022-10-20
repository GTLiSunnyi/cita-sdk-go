package contract

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/GTLiSunnyi/cita-sdk-go/types"
)

type Contract struct {
	Abi     abi.ABI
	Address string
}

func NewContract(address string, abiPath string) (*Contract, error) {
	abiFile, err := os.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}
	abiStr := string(abiFile)

	ethAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, err
	}

	return &Contract{
		Abi:     ethAbi,
		Address: address,
	}, nil
}

func (contract Contract) GetEvent(receipt *types.MyReceipt, eventName string, event interface{}) error {
	for _, log := range receipt.Logs {
		if log.Topics[0] != contract.Abi.Events[eventName].ID.String() {
			continue
		}

		logBytes, err := hex.DecodeString(log.Data[2:])
		if err != nil {
			return err
		}

		if len(logBytes)%32 != 0 {
			return fmt.Errorf("logBytes 的长度不是32的倍数, length: %d", len(logBytes))
		}

		err = contract.Abi.Unpack(event, eventName, logBytes)
		if err != nil {
			return err
		}
	}

	return nil
}
