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

type Receipt struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    receiptData `json:"data"`
}
type receiptData struct {
	TransactionHash string `json:"transactionHash"`
	ErrorMessage    string `json:"errorMessage"`
	Logs            []logs `json:"logs"`
}
type logs struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
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
