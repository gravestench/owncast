package database

import (
	"database/sql"

	"github.com/gravestench/runtime"

	"github.com/owncast/owncast/v2/pkg/services/common_logging"
)

var (
	_ runtime.Service          = &Service{}
	_ common_logging.HasLogger = &Service{}
)

type Dependency = HasDatabase

type HasDatabase interface {
	GetDatabase() *sql.DB
}

type HasDefaultsInDatabase interface {
	HasPopulatedDefaults(*sql.DB) bool
	PopulateDefaults(*sql.DB)
}
