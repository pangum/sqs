package sqs

import (
	`github.com/aws/aws-sdk-go-v2/service/sqs`
)

var (
	_ optionSend    = (*optionOptFns)(nil)
	_ optionReceive = (*optionOptFns)(nil)
)

type optionOptFns struct {
	fns []func(*sqs.Options)
}

func OptFns(fns ...func(*Options)) *optionOptFns {
	sqsFns := make([]func(options *sqs.Options), 0, len(fns))
	for _, fn := range fns {
		sqsFns = append(sqsFns, func(options *sqs.Options) {
			fn(&Options{Options: options})
		})
	}

	return &optionOptFns{fns: sqsFns}
}

func (of *optionOptFns) applySend(options *optionsSend) {
	options.fns = of.fns
}

func (of *optionOptFns) applyReceive(options *optionsReceive) {
	options.fns = of.fns
}
