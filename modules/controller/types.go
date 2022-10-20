package controller

import "github.com/GTLiSunnyi/cita-sdk-go/types/contract"

/*
Contract, 合约结构体
FuncName, 方法的名称
Params, 方法的参数
Value(可选), 交易的 value
Quota(可选), 交易的 quota
ValidUntilBlock(可选), 交易生效的区块高度，默认+95
*/
type SendRequest struct {
	Contract        *contract.Contract
	FuncName        string
	Params          []interface{}
	Value           string
	Quota           uint64
	ValidUntilBlock string
}
