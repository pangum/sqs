package sqs

import (
	`github.com/storezhang/pangu`
)

func init() {
	if err := pangu.New().Provides(newSqs); nil != err {
		panic(err)
	}
}
