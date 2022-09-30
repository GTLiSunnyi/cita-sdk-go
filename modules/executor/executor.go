package executor

import "google.golang.org/grpc"

type executorClient struct {
	client *grpc.ClientConn
}

func NewClient(executor_addr string) (Client, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	client, err := grpc.Dial(executor_addr, dialOpts...)
	if err != nil {
		return nil, err
	}

	return executorClient{
		client: client,
	}, nil
}
