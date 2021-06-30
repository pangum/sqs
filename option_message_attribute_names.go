package sqs

var _ optionReceive = (*optionMessageAttributeNames)(nil)

type optionMessageAttributeNames struct {
	messageAttributeNames []string
}

func MessageAttributeNames(names ...string) *optionMessageAttributeNames {
	return &optionMessageAttributeNames{messageAttributeNames: names}
}

func (man *optionMessageAttributeNames) applyReceive(options *optionsReceive) {
	options.messageAttributeNames = man.messageAttributeNames
}
