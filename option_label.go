package sqs

var (
	_ optionSend    = (*optionLabel)(nil)
	_ optionReceive = (*optionLabel)(nil)
)

type optionLabel struct {
	label string
}

// Label 配置标签
func Label(label string) *optionLabel {
	return &optionLabel{
		label: label,
	}
}

func (of *optionLabel) applySend(options *optionsSend) {
	options.label = of.label
}

func (of *optionLabel) applyReceive(options *optionsReceive) {
	options.label = of.label
}
