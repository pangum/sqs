package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs/types`
)

var _ optionSend = (*optionAttributes)(nil)

type optionAttributes struct {
	attributes map[string]types.MessageAttributeValue
}

// Attributes 配置消息属性
func Attributes(attributes map[string]types.MessageAttributeValue) *optionAttributes {
	return &optionAttributes{attributes: attributes}
}

func (a *optionAttributes) applySend(options *optionsSend) {
	options.messageAttributes = a.attributes
}
