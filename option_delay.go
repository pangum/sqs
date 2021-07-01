package sqs

import (
	`time`
)

var _ optionSend = (*optionDelay)(nil)

type optionDelay struct {
	delay time.Duration
}

// Delay 配置延迟
func Delay(delay time.Duration) *optionDelay {
	return &optionDelay{delay: delay}
}

func (wts *optionDelay) applySend(options *optionsSend) {
	options.delaySeconds = int32(wts.delay)
}
