package types

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
