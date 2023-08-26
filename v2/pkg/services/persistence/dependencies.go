package persistence

import (
	"github.com/gravestench/runtime"

	"github.com/owncast/owncast/v2/pkg/services/database"
)

func (s *Service) DependenciesResolved() bool {
	if s.db == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(rt runtime.R) {
	for _, service := range rt.Services() {
		switch candidate := service.(type) {
		case database.Dependency:
			if candidate.GetDatabase() != nil {
				s.db = candidate
			}
		}
	}
}
