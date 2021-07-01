package sqs

var _ optionReceive = (*optionMaxNumberOfMessages)(nil)

type optionMaxNumberOfMessages struct {
	maxNumberOfMessage int32
}

// MaxNumberOfMessages 配置每次拉取的最大消息数量
func MaxNumberOfMessages(maxNumberOfMessages int32) *optionMaxNumberOfMessages {
	return &optionMaxNumberOfMessages{maxNumberOfMessage: maxNumberOfMessages}
}

func (mnm *optionMaxNumberOfMessages) applyReceive(options *optionsReceive) {
	options.maxNumberOfMessages = mnm.maxNumberOfMessage
}
