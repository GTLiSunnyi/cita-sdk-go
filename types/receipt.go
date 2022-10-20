package types

type ReceiptWrap struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    MyReceipt `json:"data"`
}

type MyReceipt struct {
	TransactionHash     string   `json:"transactionHash"`
	TransactionIndex    uint64   `json:"transactionIndex"`
	BlockHash           string   `json:"blockHash"`
	BlockNumber         uint64   `json:"blockNumber"`
	CumulativeGasUsed   string   `json:"cumulativeGasUsed"`
	CumulativeQuotaUsed string   `json:"cumulativeQuotaUsed"`
	GasUsed             string   `json:"gasUsed"`
	QuotaUsed           uint64   `json:"quotaUsed"`
	ContractAddress     string   `json:"contractAddress"`
	Root                string   `json:"root"`
	Status              string   `json:"status"`
	From                string   `json:"from"`
	To                  string   `json:"to"`
	Logs                []*MyLog `json:"logs"`
	LogsBloom           string   `json:"logsBloom"`
	ErrorMessage        string   `json:"errorMessage"`
}

type MyLog struct {
	Address             string   `json:"address"`
	Topics              []string `json:"topics"`
	Data                string   `json:"data"`
	BlockHash           string   `json:"blockHash"`
	BlockNumber         uint64   `json:"blockNumber"`
	TransactionHash     string   `json:"transactionHash"`
	TransactionIndex    uint64   `json:"transactionIndex"`
	LogIndex            uint64   `json:"logIndex"`
	TransactionLogIndex uint64   `json:"transactionLogIndex"`
}
