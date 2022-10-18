package contract

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/GTLiSunnyi/cita-sdk-go/types"
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

func (contract Contract) GetEvent(receipt *types.Receipt, funcSignature, eventName string) ([]byte, error) {
	topicHash := crypto.Keccak256Hash([]byte(funcSignature)).Hex()

	var data = make(map[string]interface{})
	for _, log := range receipt.Data.Logs {
		if log.Topics[0] != topicHash {
			continue
		}

		logBytes, err := hex.DecodeString(log.Data[2:])
		if err != nil {
			return nil, err
		}

		if len(logBytes)%32 != 0 {
			return nil, fmt.Errorf("logBytes 的长度不是32的倍数, length: %d", len(logBytes))
		}

		err = contract.Abi.UnpackIntoMap(data, eventName, logBytes)
		if err != nil {
			return nil, err
		}
	}

	return json.Marshal(data)
}
