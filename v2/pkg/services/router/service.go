package router

import (
	"github.com/gravestench/runtime/pkg"
	"github.com/sirupsen/logrus"
)

type Service struct {
	logger *logrus.Logger
}

func (s *Service) BindLogger(logger *logrus.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *logrus.Logger {
	return s.logger
}

func (s *Service) Init(rt pkg.IsRuntime) {
	s.listenForNewRouteInitializers(rt)
	s.bindExistingServices(rt)
}

func (s *Service) Name() string {
	return "Router"
}
