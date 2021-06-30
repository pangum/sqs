package sqs

const (
	// ConsumeStatusSuccess 消费成功，会从队列中删除消息
	ConsumeStatusSuccess ConsumeStatus = 1
	// ConsumeStatusLater 延迟消费，等待下一次消费
	ConsumeStatusLater ConsumeStatus = 2
)

// ConsumeStatus 消费状态
type ConsumeStatus int
