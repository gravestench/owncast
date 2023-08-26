package database

import (
	"database/sql"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/owncast/owncast/utils"
)

const (
	schemaVersion = 7
)

func (s *Service) setup(file string) error {
	// Allow support for in-memory databases for tests.

	if file == ":memory:" {
		inMemoryDb, err := sql.Open("sqlite3", file)
		if err != nil {
			log.Fatal(err.Error())
		}
		s.db = inMemoryDb
	} else {
		// Create empty DB file if it doesn't exist.
		if !utils.DoesFileExists(file) {
			log.Traceln("Creating new database at", file)

			_, err := os.Create(file) //nolint:gosec
			if err != nil {
				log.Fatal(err.Error())
			}
		}

		onDiskDb, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_cache_size=10000&cache=shared&_journal_mode=WAL", file))
		if err != nil {
			return err
		}
		s.db = onDiskDb
		s.db.SetMaxOpenConns(1)
	}

	// Some SQLite optimizations
	_, _ = s.db.Exec("pragma journal_mode = WAL")
	_, _ = s.db.Exec("pragma synchronous = normal")
	_, _ = s.db.Exec("pragma temp_store = memory")
	_, _ = s.db.Exec("pragma wal_checkpoint(full)")

	s.createWebhooksTable()
	s.createUsersTable()
	s.createAccessTokenTable()

	if _, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS config (
		"key" string NOT NULL PRIMARY KEY,
		"value" TEXT
	);`); err != nil {
		return err
	}

	var version int
	err := s.db.QueryRow("SELECT value FROM config WHERE key='version'").
		Scan(&version)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}

		// fresh database: initialize it with the current schema version
		_, err := s.db.Exec("INSERT INTO config(key, value) VALUES(?, ?)", "version", schemaVersion)
		if err != nil {
			return err
		}
		version = schemaVersion
	}

	// is database from a newer Owncast version?
	if version > schemaVersion {
		return fmt.Errorf("incompatible database version %d (versions up to %d are supported)",
			version, schemaVersion)
	}

	// is database schema outdated?
	if version < schemaVersion {
		if err = s.migrateDatabaseSchema(s.db, version, schemaVersion); err != nil {
			return err
		}
	}

	return nil
}
