package main

import (
	"github.com/gravestench/runtime"

	"github.com/owncast/owncast/v2/pkg/services/common_logging"
	"github.com/owncast/owncast/v2/pkg/services/database"
	"github.com/owncast/owncast/v2/pkg/services/persistence"
	"github.com/owncast/owncast/v2/pkg/services/router"
	"github.com/owncast/owncast/v2/pkg/services/yp"
)

func main() {
	rt := runtime.New("Owncast")

	rt.Add(&common_logging.Service{})
	rt.Add(&database.Service{})
	rt.Add(&persistence.Service{})
	rt.Add(&router.Service{})
	rt.Add(&yp.Service{})

	rt.Run()
}
