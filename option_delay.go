package sqs

import (
	`time`
)

var _ optionSend = (*optionDelay)(nil)

type optionDelay struct {
	delay int32
}

// Delay 配置延迟
func Delay(delay time.Duration) *optionDelay {
	return &optionDelay{delay: int32(delay / time.Second)}
}

func (wts *optionDelay) applySend(options *optionsSend) {
	options.delaySeconds = wts.delay
}
