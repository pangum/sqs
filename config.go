package sqs

import (
	`time`
)

type config struct {
	// 区域
	Region string `default:"ap-east-1" json:"region" yaml:"region" xml:"region" validate:"required"`
	// 授权
	Credentials credentialsConfig `json:"credentialsConfig" yaml:"credentialsConfig" xml:"credentialsConfig" validate:"structonly"`
	// 队列名称
	Queue string `json:"queue" yaml:"queue" xml:"queue" validate:"required"`
	// 长轮询，减少费用
	Wait time.Duration `json:"wait" yaml:"wait" xml:"wait" validate:"omitempty,max=20000"`
}
