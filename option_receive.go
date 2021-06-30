package sqs

type optionReceive interface {
	applyReceive(options *optionsReceive)
}
