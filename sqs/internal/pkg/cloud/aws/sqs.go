package aws

import (
	"context"
	"fmt"
	"sqs/internal/pkg/cloud"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/aws/aws-sdk-go/aws/session"
)

var messageRetentionPeriod = "345600"
var _ cloud.MessageClient = SQS{}

type SQS struct {
	timeout time.Duration
	client  *sqs.SQS
}

func NewSQS(session *session.Session, timeout time.Duration) SQS {
	return SQS{
		timeout: timeout,
		client:  sqs.New(session),
	}
}

func (s SQS) CreateQueue(ctx context.Context, queueName string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err := s.client.CreateQueueWithContext(ctx, &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
		Attributes: map[string]*string{
			"MessageRetentionPeriod":        aws.String(string(messageRetentionPeriod)),
			"VisibilityTimeout":             aws.String("5"),
			"ReceiveMessageWaitTimeSeconds": aws.String("20"),
		},
	})
	if err != nil {
		return "", fmt.Errorf("create: %w", err)
	}

	return *res.QueueUrl, nil
}

func (s SQS) List(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	list, err := s.client.ListQueuesWithContext(ctx, &sqs.ListQueuesInput{
		MaxResults: aws.Int64(10),
	})
	if err != nil {
		return "", err
	}

	return list.GoString(), nil
}

func (s SQS) Send(ctx context.Context, req *cloud.SendRequest) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	attrs := make(map[string]*sqs.MessageAttributeValue, len(req.Attributes))
	for _, attr := range req.Attributes {
		attrs[attr.Key] = &sqs.MessageAttributeValue{
			StringValue: aws.String(attr.Value),
			DataType:    aws.String(attr.Type),
		}
	}

	res, err := s.client.SendMessageWithContext(ctx, &sqs.SendMessageInput{
		MessageAttributes: attrs,
		MessageBody:       aws.String(req.Body),
		QueueUrl:          aws.String(req.QueueURL),
	})
	if err != nil {
		return "", fmt.Errorf("send: %w", err)
	}

	return *res.MessageId, nil
}

func (s SQS) Receive(ctx context.Context, queueURL string) (*cloud.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*(20*5))
	defer cancel()

	res, err := s.client.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:              aws.String(queueURL),
		MaxNumberOfMessages:   aws.Int64(1),
		WaitTimeSeconds:       aws.Int64(20),
		MessageAttributeNames: aws.StringSlice([]string{"ALL"}),
	})
	if err != nil {
		return nil, fmt.Errorf("receive: %w", err)
	}

	if len(res.Messages) == 0 {
		return nil, nil
	}

	attrs := make(map[string]string)
	for key, attr := range res.Messages[0].MessageAttributes {
		attrs[key] = *attr.StringValue
	}

	return &cloud.Message{
		ID:            *res.Messages[0].MessageId,
		ReceiptHandle: *res.Messages[0].ReceiptHandle,
		Body:          *res.Messages[0].Body,
		Attributes:    attrs,
	}, nil
}

func (s SQS) Delete(ctx context.Context, queueURL, rcvHandle string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if _, err := s.client.DeleteMessageWithContext(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(rcvHandle),
	}); err != nil {
		return err
	}

	return nil
}
