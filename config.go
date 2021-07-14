package sqs

import (
	`time`
)

type config struct {
	// 区域
	Region string `default:"ap-east-1" json:"region" yaml:"region" xml:"region" validate:"required"`
	// 授权
	Credentials credentialsConfig `json:"credentials" yaml:"credentials" xml:"credentials" validate:"structonly"`
	// 队列
	Queue string `json:"queue" yaml:"queue" xml:"queue" validate:"required_without=Queues"`
	// 队列列表
	Queues []queue `json:"queues" yaml:"queues" xml:"queues" validate:"required_without=Queue"`
	// 授权验证（策略）
	Session string `json:"session" yaml:"session" xml:"session" validate:"omitempty"`
	// 长轮询，减少费用
	Wait time.Duration `json:"wait" yaml:"wait" xml:"wait" validate:"omitempty,max=20000"`
}
