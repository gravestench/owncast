package common_logging

import (
	"github.com/gravestench/runtime/pkg"
	"github.com/sirupsen/logrus"
)

type Service struct {
	logger *logrus.Logger
}

func (s *Service) Init(rt pkg.IsRuntime) {
	s.logger = createLogger(s.Name())
	s.bindExistingServices(rt)
	s.listenForNewRouteInitializers(rt)
}

func (s *Service) Name() string {
	return "Common Logging"
}
