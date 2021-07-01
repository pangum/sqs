package sqs

var (
	_ optionSend    = (*optionQueueUrl)(nil)
	_ optionReceive = (*optionQueueUrl)(nil)
)

type optionQueueUrl struct {
	url string
}

// QueueUrl 配置队列地址
func QueueUrl(url string) *optionQueueUrl {
	return &optionQueueUrl{url: url}
}

func (qu *optionQueueUrl) applySend(options *optionsSend) {
	options.url = qu.url
}

func (qu *optionQueueUrl) applyReceive(options *optionsReceive) {
	options.url = qu.url
}
