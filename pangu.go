package sqs

import (
	"github.com/pangum/pangu"
	"github.com/pangum/sqs/internal/plugin"
)

func init() {
	creator := new(plugin.Creator)
	pangu.New().Get().Dependency().Put(
		creator.New,
	).Build().Build().Apply()
}
