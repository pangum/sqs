package sqs

type queue struct {
	// 名称
	Name string `json:"name" yaml:"name" xml:"name" validate:"required"`
	// 标签
	Label string `json:"label" yaml:"label" xml:"label" validate:"required"`
}
