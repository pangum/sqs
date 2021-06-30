package sqs

type Handler interface {
	OnMessage() (err error)
}
