package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs/types`
)

type optionsReceive struct {
	optionsBase

	waitTimeSeconds       int32
	visibilityTimeout     int32
	maxNumberOfMessages   int32
	attributeNames        []types.QueueAttributeName
	messageAttributeNames []string
}

func defaultOptionsReceive(label string, waitTimeSeconds int32) *optionsReceive {
	return &optionsReceive{
		optionsBase: optionsBase{
			label: label,
		},
		waitTimeSeconds:     waitTimeSeconds,
		maxNumberOfMessages: 1,
	}
}
