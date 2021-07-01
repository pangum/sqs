package sqs

import (
	`time`
)

var _ optionReceive = (*optionVisibilityTimeout)(nil)

type optionVisibilityTimeout struct {
	visibilityTimeout time.Duration
}

// VisibilityTimeout 配置消息可见性
func VisibilityTimeout(visibilityTimeout time.Duration) *optionVisibilityTimeout {
	return &optionVisibilityTimeout{visibilityTimeout: visibilityTimeout}
}

func (vt *optionVisibilityTimeout) applyReceive(options *optionsReceive) {
	options.visibilityTimeout = int32(vt.visibilityTimeout)
}
