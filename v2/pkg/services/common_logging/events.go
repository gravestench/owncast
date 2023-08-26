package common_logging

import (
	"github.com/gravestench/runtime"
	runtimeEvents "github.com/gravestench/runtime/pkg/events"
)

func (s *Service) listenForNewRouteInitializers(rt runtime.R) {
	rt.Events().On(runtimeEvents.EventServiceAdded, s.onServiceAdded)
}

func (s *Service) onServiceAdded(args ...any) {
	if len(args) < 1 {
		return
	}

	if candidate, ok := args[0].(runtime.Service); ok {
		go s.bind(candidate)
	}
}
