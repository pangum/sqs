package sqs

type awsConfig struct {
	// 区域
	Region string `default:"ap-east-1" json:"region" yaml:"region" xml:"region" toml:"region" validate:"required"`
	// 授权
	Credentials credentialsConfig `json:"credentials" yaml:"credentials" xml:"credentials" toml:"credentials" validate:"structonly"`
	// Sqs配置
	Sqs config `json:"sqs" yaml:"sqs" xml:"sqs" toml:"sqs" validate:"required"`
}
