package sqs

var _ optionReceive = (*optionMaxRetry)(nil)

type optionMaxRetry struct {
	maxRetry int
}

// MaxRetry 配置最大重试次数
func MaxRetry(maxRetry int) *optionMaxRetry {
	return &optionMaxRetry{maxRetry: maxRetry}
}

func (mr *optionMaxRetry) applyReceive(options *optionsReceive) {
	options.maxRetry = mr.maxRetry
}
