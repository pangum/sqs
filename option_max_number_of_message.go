package sqs

var _ optionReceive = (*optionMaxNumberOfMessages)(nil)

type optionMaxNumberOfMessages struct {
	maxNumberOfMessage int32
}

func MaxNumberOfMessages(maxNumberOfMessages int32) *optionMaxNumberOfMessages {
	return &optionMaxNumberOfMessages{maxNumberOfMessage: maxNumberOfMessages}
}

func (mnm *optionMaxNumberOfMessages) applyReceive(options *optionsReceive) {
	options.maxNumberOfMessages = mnm.maxNumberOfMessage
}
