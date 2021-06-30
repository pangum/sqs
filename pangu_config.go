package sqs

type panguConfig struct {
	// 亚马逊云顶层配置
	Aws awsConfig `json:"aws" yaml:"aws" xml:"aws" validate:"structonly"`
}
