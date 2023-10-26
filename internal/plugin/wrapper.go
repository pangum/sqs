package plugin

import (
	"github.com/pangum/sqs/internal/config"
)

type Wrapper struct {
	// 亚马逊云顶层配置
	Aws config.Aws `json:"aws" yaml:"aws" xml:"aws" validate:"structonly"`
}
