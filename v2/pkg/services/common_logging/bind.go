package common_logging

import (
	"github.com/gravestench/runtime"
)

func (s *Service) bindExistingServices(rt runtime.R) {
	for _, service := range rt.Services() {
		if service.Name() == s.Name() {
			continue
		}

		s.bind(service)
	}
}

func (s *Service) bind(service runtime.Service) {
	if candidate, ok := service.(HasLogger); ok {
		candidate.BindLogger(createLogger(service.Name()))
	}
}
