package sqs

var (
	_ optionSend    = (*optionQueueUrl)(nil)
	_ optionReceive = (*optionQueueUrl)(nil)
)

type optionQueueUrl struct {
	url string
}

func QueueUrl(url string) *optionQueueUrl {
	return &optionQueueUrl{url: url}
}

func (qu *optionQueueUrl) applySend(options *optionsSend) {
	options.url = qu.url
}

func (qu *optionQueueUrl) applyReceive(options *optionsReceive) {
	options.url = qu.url
}
