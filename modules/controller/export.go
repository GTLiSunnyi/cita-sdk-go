package controller

type Client interface {
	GetBlockNumber(for_padding bool) (uint64, error)
}
