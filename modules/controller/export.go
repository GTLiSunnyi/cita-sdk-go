package controller

type Client interface {
	GetBlockNumber(for_padding bool, authorization, chain_code string) (uint64, error)
}
