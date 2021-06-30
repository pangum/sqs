package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs/types`
)

var _ optionSend = (*optionSystemAttributes)(nil)

type optionSystemAttributes struct {
	attributes map[string]types.MessageSystemAttributeValue
}

func SystemAttributes(attributes map[string]types.MessageSystemAttributeValue) *optionSystemAttributes {
	return &optionSystemAttributes{attributes: attributes}
}

func (a *optionSystemAttributes) applySend(options *optionsSend) {
	options.messageSystemAttributes = a.attributes
}
