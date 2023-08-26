package router

import (
	"github.com/gravestench/runtime"
)

func (s *Service) DependenciesResolved() bool {
	if s.Logger() == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(rt runtime.R) {
	// noop
}
