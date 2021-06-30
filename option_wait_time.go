package sqs

import (
	`time`
)

var _ optionReceive = (*optionWaitTime)(nil)

type optionWaitTime struct {
	waitTimeSeconds time.Duration
}

func WaitTime(waitTime time.Duration) *optionWaitTime {
	return &optionWaitTime{waitTimeSeconds: waitTime}
}

func (wt *optionWaitTime) applyReceive(options *optionsReceive) {
	options.waitTimeSeconds = int32(wt.waitTimeSeconds)
}
