package router

import (
	"net/http"

	"github.com/gravestench/runtime"
)

func (s *Service) bindExistingServices(rt runtime.R) {
	for _, service := range rt.Services() {
		s.bind(service)
	}
}

func (s *Service) bind(service runtime.Service) {
	if candidate, ok := service.(RouteInitializer); ok {
		s.Logger().Infof("binding un-protected routes for %s", s.Name())
		for route, handler := range candidate.Routes() {
			http.HandleFunc(route, handler)
		}
	}
}
