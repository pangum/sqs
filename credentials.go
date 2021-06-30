package sqs

type credentialsConfig struct {
	// 授权，相当于用户名
	Access string `json:"access" yaml:"access" xml:"access" validate:"required"`
	// 授权，相当于密码
	Secret string `json:"secret" yaml:"secret" xml:"secret" validate:"required"`
}
