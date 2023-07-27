package sqs

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependency(newSqs)
}
