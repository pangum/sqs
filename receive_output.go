package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs`
)

type ReceiveOutput struct {
	*sqs.ReceiveMessageOutput
}
