package main

import (
	"s3/internal/bucket"
	"s3/internal/pkg/cloud/aws"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	ses, err := aws.New(aws.Config{
		Address: "http://localhost:4566",
		Region:  "us-east-1",
		Profile: "localstack",
		ID:      "test",
		Secret:  "test",
	})
	if err != nil {
		logrus.Fatalln(err)
	}

	// TODO: to use flag params by command line
	bucket.Bucket(aws.NewS3(ses, time.Second*5))
}
