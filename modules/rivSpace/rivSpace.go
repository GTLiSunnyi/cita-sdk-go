package rivSpace

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/GTLiSunnyi/cita-sdk-go/types"
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

func (client rivSpaceClient) Send(params map[string]interface{}, header types.GrpcRequestHeader) (*types.MyReceipt, error) {
	data, err := SendRivSpaceRequest(params, client.RunAddress, header)
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

func (client rivSpaceClient) GetReceipt(tx_hash string, header types.GrpcRequestHeader) (*types.MyReceipt, error) {
	params := map[string]interface{}{
		"txHash": tx_hash,
	}

	data, err := SendRivSpaceRequest(params, client.ReceiptAddress, header)
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
		return client.GetReceipt(tx_hash, header)
	}

	if receiptWrap.Data.ErrorMessage != "" {
		return nil, errors.New(receiptWrap.Data.ErrorMessage)
	}
	return &receiptWrap.Data, nil
}
