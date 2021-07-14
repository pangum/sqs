package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs/types`
)

type optionsSend struct {
	optionsBase

	delaySeconds            int32
	messageAttributes       map[string]types.MessageAttributeValue
	messageSystemAttributes map[string]types.MessageSystemAttributeValue
}

func defaultOptionsSend(label string) *optionsSend {
	return &optionsSend{
		optionsBase: optionsBase{
			label: label,
		},
	}
}
