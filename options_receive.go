package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs/types`
)

type optionsReceive struct {
	queueUrl              string
	waitTimeSeconds       int32
	visibilityTimeout     int32
	maxNumberOfMessages   int32
	attributeNames        []types.QueueAttributeName
	messageAttributeNames []string
}

func defaultOptionsReceive(url string, waitTimeSeconds int32) *optionsReceive {
	return &optionsReceive{
		queueUrl:            url,
		waitTimeSeconds:     waitTimeSeconds,
		maxNumberOfMessages: 1,
	}
}
