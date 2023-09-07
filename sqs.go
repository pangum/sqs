package sqs

import (
	"github.com/goexl/sqs"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
	"github.com/pangum/sqs/internal/config"
	"github.com/pangum/sqs/internal/core"
)

func newSqs(conf *pangu.Config, logger logging.Logger) (client *Client, err error) {
	wrapper := new(core.Wrapper)
	if err = conf.Load(wrapper); nil != err {
		return
	}

	aws := wrapper.Aws
	self := aws.Sqs
	builder := sqs.New().Region(aws.RealRegion()).Wait(self.Wait).Logger(logger)

	// 配置授权
	credential := aws.RealCredential()
	if nil != credential {
		builder = builder.Credential().Default(credential.Id, credential.Key).Build()
	}

	if "" != self.Queue {
		queue := new(config.Queue)
		queue.Name = self.Queue
		queue.Label = "default"
		self.Queues = append(self.Queues, queue)
	}
	for _, queue := range self.Queues {
		builder.Queue(queue.Label, queue.Name)
	}
	client = builder.Build()

	return
}
