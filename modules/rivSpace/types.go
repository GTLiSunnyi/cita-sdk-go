package rivSpace

const GetReceiptTimeout = 10000 // 单位：毫秒
const GetReceiptInterval = 500

type SendResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    data   `json:"data"`
}

type data struct {
	Version int    `json:" version"`
	TxHash  string `json:"txHash"`
}

type CreateAccountResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    createAccountData `json:"data"`
}

type createAccountData struct {
	AppUserCode string `json:"appUserCode"`
	Address     string `json:"address"`
}
