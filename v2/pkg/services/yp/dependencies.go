package yp

import (
	"github.com/gravestench/runtime/pkg"

	"github.com/owncast/owncast/v2/pkg/services/persistence"
)

func (s *Service) DependenciesResolved() bool {
	if s.persistence == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(runtime pkg.IsRuntime) {
	for _, service := range runtime.Services() {
		if candidate, ok := service.(persistence.Dependency); ok {
			if candidate.Ready() {
				s.persistence = candidate
			}
		}
	}
}
