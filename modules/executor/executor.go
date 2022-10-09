package executor

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type executorClient struct {
	client *grpc.ClientConn
}

func NewClient(executor_addr string) (Client, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	client, err := grpc.Dial(executor_addr, dialOpts...)
	if err != nil {
		return nil, err
	}

	return executorClient{
		client: client,
	}, nil
}
