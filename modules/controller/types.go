package controller

/*
@To, 交易的目的地址
@Data, 交易的数据
@Value, 交易的 value
@Quota, 交易的 quota
@ValidUntilBlock, 交易生效的区块高度，默认+95
*/
type SendRequest struct {
	To              string
	Data            string
	Value           string
	Quota           uint64
	ValidUntilBlock string
}
