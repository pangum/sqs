package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs/types`
)

type optionsSend struct {
	queueUrl                string
	delaySeconds            int32
	messageAttributes       map[string]types.MessageAttributeValue
	messageSystemAttributes map[string]types.MessageSystemAttributeValue
}

func defaultOptionsSend(url string) *optionsSend {
	return &optionsSend{
		queueUrl: url,
	}
}
