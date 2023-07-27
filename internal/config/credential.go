package config

type Credential struct {
	// 授权，相当于用户名
	Id string `json:"id" yaml:"id" xml:"id" toml:"id" validate:"required"`
	// 授权，相当于密码
	Key string `json:"key" yaml:"key" xml:"key" toml:"key" validate:"required"`
	// 授权验证（策略）
	Session string `json:"session" yaml:"session" xml:"session" toml:"session" validate:"omitempty"`
}
