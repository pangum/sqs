package plugin

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/goexl/sqs"
	"github.com/pangum/pangu"
	"github.com/pangum/sqs/internal/config"
)

type Constructor struct {
	// 解决命名空间问题
}

func (c *Constructor) New(config *pangu.Config, http *http.Client, logger log.Logger) (client *sqs.Client, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else {
		client, err = c.new(&wrapper.Aws, http, logger)
	}

	return
}

func (c *Constructor) new(aws *config.Aws, http *http.Client, logger log.Logger) (client *sqs.Client, err error) {
	self := aws.Sqs
	builder := sqs.New().Region(aws.RealRegion()).Wait(self.Wait).Http(http).Logger(logger)

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
