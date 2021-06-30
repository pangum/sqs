package sqs

import (
	`time`
)

var _ optionReceive = (*optionVisibilityTimeout)(nil)

type optionVisibilityTimeout struct {
	visibilityTimeout time.Duration
}

func VisibilityTimeout(visibilityTimeout time.Duration) *optionVisibilityTimeout {
	return &optionVisibilityTimeout{visibilityTimeout: visibilityTimeout}
}

func (vt *optionVisibilityTimeout) applyReceive(options *optionsReceive) {
	options.visibilityTimeout = int32(vt.visibilityTimeout)
}
