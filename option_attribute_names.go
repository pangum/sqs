package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs/types`
)

var _ optionReceive = (*optionAttributeNames)(nil)

type optionAttributeNames struct {
	attributeNames []types.QueueAttributeName
}

func AttributeNames(names ...types.QueueAttributeName) *optionAttributeNames {
	return &optionAttributeNames{attributeNames: names}
}

func (an *optionAttributeNames) applyReceive(options *optionsReceive) {
	options.attributeNames = an.attributeNames
}
