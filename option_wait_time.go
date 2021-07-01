package sqs

import (
	`time`
)

var _ optionReceive = (*optionWaitTime)(nil)

type optionWaitTime struct {
	waitTimeSeconds time.Duration
}

// WaitTime 配置拉取消息最大等待时间
func WaitTime(waitTime time.Duration) *optionWaitTime {
	return &optionWaitTime{waitTimeSeconds: waitTime}
}

func (wt *optionWaitTime) applyReceive(options *optionsReceive) {
	options.waitTimeSeconds = int32(wt.waitTimeSeconds)
}
