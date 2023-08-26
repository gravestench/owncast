package common_logging

import (
	"github.com/gravestench/runtime"
	"github.com/sirupsen/logrus"
)

var (
	_ runtime.Service = &Service{}
)

type HasLogger interface {
	BindLogger(logger *logrus.Logger)
	Logger() *logrus.Logger
}
