package sqs

import (
	`time`
)

var _ optionReceive = (*optionRetryDuration)(nil)

type optionRetryDuration struct {
	duration time.Duration
}

// RetryDuration 配置重试间隔
func RetryDuration(duration time.Duration) *optionRetryDuration {
	return &optionRetryDuration{duration: duration}
}

func (wts *optionRetryDuration) applyReceive(options *optionsReceive) {
	options.retryDuration = wts.duration
}
