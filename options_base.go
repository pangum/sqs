package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs`
)

type optionsBase struct {
	url string
	fns []func(*sqs.Options)
}
