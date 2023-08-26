package yp

import (
	"net/http"

	"github.com/gravestench/runtime"

	"github.com/owncast/owncast/v2/pkg/services/common_logging"
)

var (
	_ runtime.Service          = &Service{}
	_ runtime.HasDependencies  = &Service{}
	_ common_logging.HasLogger = &Service{}
)

// RouteInitializer describes a runtime service that has un-protected routes.
type RouteInitializer interface {
	runtime.Service
	Routes() map[string]func(http.ResponseWriter, *http.Request)
}
