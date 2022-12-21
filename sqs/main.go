package main

import (
	"sqs/config"
	"sqs/internal/message"
	"sqs/internal/pkg/cloud/aws"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	ses, err := aws.New(aws.Config{
		Address: config.Config.Address,
		Region:  config.Config.Region,
		Profile: config.Config.Profile,
		ID:      config.Config.ID,
		Secret:  config.Config.Secret,
	})
	if err != nil {
		logrus.Fatalln(err)
	}

	// TODO: move to flag params
	message.Message(aws.NewSQS(ses, time.Second*5))
}
