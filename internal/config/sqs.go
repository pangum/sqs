package config

import (
	"time"
)

type Sqs struct {
	// 区域
	Region string `json:"region" yaml:"region" xml:"region" toml:"region"`
	// 授权
	Credential *Credential `json:"credential" yaml:"credential" xml:"credential" toml:"credential"`
	// 队列
	Queue string `json:"queue" yaml:"queue" xml:"queue" toml:"queue" validate:"required_without=Queues"`
	// 队列列表
	Queues []*Queue `json:"queues" yaml:"queues" xml:"queues" toml:"queues" validate:"required_without=Queue"`
	// 长轮询，减少费用
	Wait time.Duration `json:"wait" yaml:"wait" xml:"wait" toml:"wait" validate:"omitempty,max=20000"`
}
