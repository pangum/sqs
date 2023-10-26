package config

import (
	"github.com/goexl/gox"
)

type Aws struct {
	// 区域
	Region string `default:"ap-east-1" json:"region" yaml:"region" xml:"region" toml:"region"`
	// 授权
	Credential *Credential `json:"credential" yaml:"credential" xml:"credential" toml:"credential"`
	// Sqs配置
	Sqs Sqs `json:"sqs" yaml:"sqs" xml:"sqs" toml:"sqs" validate:"required"`
}

func (a *Aws) RealRegion() string {
	return gox.Ift("" != a.Sqs.Region, a.Sqs.Region, a.Region)
}

func (a *Aws) RealCredential() *Credential {
	return gox.Ift(nil != a.Sqs.Credential, a.Sqs.Credential, a.Credential)
}
