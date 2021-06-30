package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs`
)

// Client Sqs客户端封装
type Client struct {
	client   *sqs.Client
	queueUrl string
}
