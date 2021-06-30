package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs`
)

type SendOutput struct {
	*sqs.SendMessageOutput
}
