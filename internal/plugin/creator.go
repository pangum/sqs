package plugin

import (
	"github.com/goexl/log"
	"github.com/goexl/sqs"
	"github.com/pangum/pangu"
	"github.com/pangum/sqs/internal/config"
)

type Creator struct {
	// 解决命名空间问题
}

func (c *Creator) New(config *pangu.Config, logger log.Logger) (client *sqs.Client, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else {
		client, err = c.new(&wrapper.Aws, logger)
	}

	return
}

func (c *Creator) new(aws *config.Aws, logger log.Logger) (client *sqs.Client, err error) {
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
