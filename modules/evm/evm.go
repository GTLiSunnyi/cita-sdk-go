package evm

import "google.golang.org/grpc"

type evmClient struct {
	client *grpc.ClientConn
}

func NewClient(client *grpc.ClientConn) Client {
	return &evmClient{
		client: client,
	}
}
