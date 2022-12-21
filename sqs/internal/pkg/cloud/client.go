package cloud

import "context"

type MessageClient interface {
	CreateQueue(ctx context.Context, queueName string) (string, error)
	Send(ctx context.Context, req *SendRequest) (string, error)
	Receive(ctx context.Context, queueURL string) (*Message, error)
	Delete(ctx context.Context, queueURL, rcvHandle string) error
	List(ctx context.Context) (string, error)
}
