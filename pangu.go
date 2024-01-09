package sqs

import (
	"github.com/pangum/pangu"
	"github.com/pangum/sqs/internal/plugin"
)

func init() {
	ctor := new(plugin.Constructor)
	pangu.New().Get().Dependency().Put(
		ctor.New,
	).Build().Build().Apply()
}
