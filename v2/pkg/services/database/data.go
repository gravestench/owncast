// This is a centralized place to connect to the database, and hold a reference to it.
// Other packages can share this reference.  This package would also be a place to add any kind of
// persistence-related convenience methods or migrations.

package database

import (
	"database/sql"
)

// GetDatabase will return the shared instance of the actual database.
func (s *Service) GetDatabase() *sql.DB {
	return s.db
}
