package core

type Wrapper struct {
	// 亚马逊云顶层配置
	Aws Aws `json:"aws" yaml:"aws" xml:"aws" validate:"structonly"`
}
