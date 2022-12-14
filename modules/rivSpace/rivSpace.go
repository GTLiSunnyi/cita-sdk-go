package rivSpace

import (
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

func (client rivSpaceClient) Send(header types.GrpcRequestHeader, params map[string]interface{}) (*types.MyReceipt, error) {
	data, err := SendRivSpaceRequest(header, params, client.RunAddress)
	if err != nil {
		return nil, err
	}

	var res SendResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}

	var receipt = &types.MyReceipt{}
	ch := make(chan error, 1)
	go func() {
		var err error
		receipt, err = client.GetReceipt(header, res.Data.TxHash)
		ch <- err
	}()

	select {
	case <-time.After(GetReceiptTimeout * time.Millisecond):
		return receipt, errors.New("请求交易receipt超时, txHash: " + res.Data.TxHash)
	case err = <-ch:
		return receipt, err
	}
}

func (client rivSpaceClient) SendAndGetEvent(header types.GrpcRequestHeader, contract *contract.Contract, params map[string]interface{}, eventName string, event interface{}) error {
	receipt, err := client.Send(header, params)
	if err != nil {
		return err
	}

	return contract.GetEvent(receipt, eventName, event)
}

func (client rivSpaceClient) CreateAccount(header types.GrpcRequestHeader, name, appId, appSecret string) (string, error) {
	params := map[string]interface{}{
		"appUserCode": name,
		"appId":       appId,
		"appSecret":   appSecret,
	}

	data, err := SendRivSpaceRequest(header, params, client.CreateAccountAddress)
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

func (client rivSpaceClient) GetReceipt(header types.GrpcRequestHeader, tx_hash string) (*types.MyReceipt, error) {
	params := map[string]interface{}{
		"txHash": tx_hash,
	}

	data, err := SendRivSpaceRequest(header, params, client.ReceiptAddress)
	if err != nil {
		return nil, err
	}

	var receiptWrap = &types.ReceiptWrap{}
	err = json.Unmarshal(data, &receiptWrap)
	if err != nil {
		return nil, err
	}

	// 交易没有上链的话，间隔 继续查询
	if receiptWrap.Data.TransactionHash == "" {
		time.Sleep(time.Duration(GetReceiptInterval) * time.Millisecond)
		return client.GetReceipt(header, tx_hash)
	}

	if receiptWrap.Data.ErrorMessage != "" {
		return nil, errors.New(receiptWrap.Data.ErrorMessage)
	}
	return &receiptWrap.Data, nil
}
