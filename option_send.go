package sqs

type optionSend interface {
	applySend(options *optionsSend)
}
