package rivSpace

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/GTLiSunnyi/cita-sdk-go/types"
	"github.com/GTLiSunnyi/cita-sdk-go/types/contract"
)

type rivSpaceClient struct {
	types.RivSpaceAddress
}

func NewClient(rivSpaceAddress types.RivSpaceAddress) Client {
	return rivSpaceClient{
		RivSpaceAddress: types.RivSpaceAddress{
			RunAddress:           rivSpaceAddress.RunAddress,
			ReceiptAddress:       rivSpaceAddress.ReceiptAddress,
			CreateAccountAddress: rivSpaceAddress.CreateAccountAddress,
		},
	}
}

func (client rivSpaceClient) Send(params map[string]interface{}, header types.GrpcRequestHeader) (Receipt, error) {
	data, err := SendRivSpaceRequest(params, client.RunAddress, header)
	if err != nil {
		return Receipt{}, err
	}

	var res SendResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return Receipt{}, err
	}

	if res.Code != 200 {
		return Receipt{}, errors.New(res.Message)
	}

	var receipt Receipt
	ch := make(chan error, 1)
	go func() {
		var err error
		receipt, err = client.GetReceipt(res.Data.TxHash, header)
		ch <- err
	}()

	select {
	case <-time.After(GetReceiptTimeout * time.Millisecond):
		return receipt, errors.New("请求交易receipt超时, txHash: " + res.Data.TxHash)
	case err = <-ch:
		return receipt, err
	}
}

func (client rivSpaceClient) CreateAccount(name, appId, appSecret string, header types.GrpcRequestHeader) (string, error) {
	params := map[string]interface{}{
		"appUserCode": name,
		"appId":       appId,
		"appSecret":   appSecret,
	}

	data, err := SendRivSpaceRequest(params, client.CreateAccountAddress, header)
	if err != nil {
		return "", err
	}

	var res CreateAccountResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return "", err
	}

	if res.Code != 200 {
		return "", errors.New(res.Message)
	}

	return res.Data.Address, nil
}

func (client rivSpaceClient) GetReceipt(tx_hash string, header types.GrpcRequestHeader) (Receipt, error) {
	params := map[string]interface{}{
		"txHash": tx_hash,
	}

	data, err := SendRivSpaceRequest(params, client.ReceiptAddress, header)
	if err != nil {
		return Receipt{}, err
	}

	var receipt Receipt
	err = json.Unmarshal(data, &receipt)
	if err != nil {
		return receipt, err
	}

	// 交易没有上链的话，间隔 250ms 继续查询
	if receipt.Data.TransactionHash == "" {
		time.Sleep(time.Duration(250) * time.Millisecond)
		return client.GetReceipt(tx_hash, header)
	}

	if receipt.Data.ErrorMessage != "" {
		return receipt, errors.New(receipt.Data.ErrorMessage)
	}
	return receipt, nil
}

func (client rivSpaceClient) GetEvent(contract contract.Contract, receipt Receipt, eventName string) (map[string]interface{}, error) {
	var m = map[string]interface{}{}
	for _, log := range receipt.Data.Logs {
		logBytes, err := hex.DecodeString(log.Data[2:])
		if err != nil {
			return nil, err
		}

		err = contract.Abi.UnpackIntoMap(m, eventName, logBytes)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}
