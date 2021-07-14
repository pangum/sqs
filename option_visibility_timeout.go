package sqs

import (
	`time`
)

var _ optionReceive = (*optionVisibilityTimeout)(nil)

type optionVisibilityTimeout struct {
	visibilityTimeout int32
}

// VisibilityTimeout 配置消息可见性
func VisibilityTimeout(visibilityTimeout time.Duration) *optionVisibilityTimeout {
	return &optionVisibilityTimeout{visibilityTimeout: int32(visibilityTimeout / time.Second)}
}

func (vt *optionVisibilityTimeout) applyReceive(options *optionsReceive) {
	options.visibilityTimeout = int32(vt.visibilityTimeout)
}
