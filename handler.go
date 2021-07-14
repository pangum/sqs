package sqs

import (
	`context`
)

// Handler 消费处理器
type Handler interface {
	// OnMessage 处理消息
	OnMessage(context context.Context, message *Message) (status ConsumeStatus, err error)
}
