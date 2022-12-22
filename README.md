# Localstack

A fully functional local AWS cloud stack. Develop and test your cloud & Serverless apps offline!

* [Install](https://localstack.cloud/)

### aws
  * Tool to manage AWS services from command calling http api.
  * It work on local and production env
### awslocal
  * Wrapper replacement for the aws command that runs at localstack.
  * It work locally
  * It already has defined the endpoint-url

### How to use aws at localstack?

```sh
# s3 service
$ aws --endpoint-url=http://localhost:4566 s3 mb s3://brauliobucket
$ aws --endpoint-url=http://localhost:4566 s3 cp README.md s3://brauliobucket
$ aws --endpoint-url=http://localhost:4566 s3 ls s3://brauliobucket
$ aws --endpoint-url=http://localhost:4566 s3 mv test.json s3://test/test1.json
$ aws --endpoint-url=http://localhost:4566 s3 rm s3://test/test1.json

# sqs service
$ aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name example-1
$ aws --endpoint-url=http://localhost:4566 sqs list-queues
$ aws --endpoint-url=http://localhost:4566 sqs send-message --queue-url http://localhost:4566/000000000000/example-1 --message-body "Hello world"
$ aws --endpoint-url=http://localhost:4566 sqs receive-message --queue-url http://localhost:4566/000000000000/example-1
$ aws --endpoint-url=http://localhost:4566 sqs delete-queue --queue-url http://localhost:4566/000000000000/example-1

# awslocal
$ awslocal sqs create-queue --queue-name example-2
$ awslocal sqs list-queues
$ awslocal sqs send-message --queue-url http://localhost:4566/000000000000/example-2 --message-body "I am working with localstack"
$ awslocal sqs receive-message --queue-url http://localhost:4566/000000000000/example-2
```

## SQS
* [Example](/sqs/)

## S3
* [Example](/s3/)

## CloudWatch

Cloudwatch supervises and monitors AWS resources in real-time, it can add or centralize metrics in the same place, 
also send alerts and notify us when any of our services has exceeded a limit that previously was configured.

Note: LocalStack currently supports metric-alarm evaluation with statistic and comparison-operator.

### What services I can monitor with cloudwatch?
  - Computing services
  - Database services (Aurora, DynamoDB, etc)
  - Volumes EBS and EFS (Disk I/O)
  - Load Balancers (type request)
  - Billing
  - Containers like EKS & ECS
