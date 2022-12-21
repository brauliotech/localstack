package message

import (
	"context"
	"sqs/internal/pkg/cloud"

	"github.com/sirupsen/logrus"
)

func Message(client cloud.MessageClient) {
	ctx := context.Background()
	queueURL, _ := createQueue(ctx, client)
	logrus.Println("Queue generated with URL: ", queueURL)

	err := send(ctx, client, queueURL)
	logrus.Println("Message sent by queue with error: ", err)

	msg, _ := receive(ctx, client, queueURL)
	logrus.Println("Message received: ", msg)

	err = delete(ctx, client, queueURL, msg.ReceiptHandle)
	logrus.Println("Queue deleted with error: ", err)

	queues, _ := list(ctx, client)
	logrus.Println("List of queues: ", queues)
}

func createQueue(ctx context.Context, client cloud.MessageClient) (string, error) {
	// TODO: queue name parametisable
	url, err := client.CreateQueue(ctx, "queue-name-example")
	if err != nil {
		return "", err
	}

	return url, nil
}

func send(ctx context.Context, client cloud.MessageClient, queueURL string) error {
	_, err := client.Send(ctx, &cloud.SendRequest{
		QueueURL: queueURL,
		Body:     "In Golang, flag is a built-in package shipped with Go standard library",
		Attributes: []cloud.Attribute{
			{
				Key:   "Title",
				Value: "Flag in Go",
				Type:  "String",
			},
			{
				Key:   "Year",
				Value: "2022",
				Type:  "Number",
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func receive(ctx context.Context, client cloud.MessageClient, queueURL string) (*cloud.Message, error) {
	res, err := client.Receive(ctx, queueURL)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func delete(ctx context.Context, client cloud.MessageClient, queueURL, rcvHnd string) error {
	if err := client.Delete(ctx, queueURL, rcvHnd); err != nil {
		return err
	}

	return nil
}

func list(ctx context.Context, client cloud.MessageClient) (string, error) {
	list, err := client.List(ctx)
	if err != nil {
		return "", err
	}

	return list, nil
}
