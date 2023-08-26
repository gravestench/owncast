package persistence

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"sync"
	"time"

	"github.com/gravestench/runtime/pkg"
	// sqlite requires a blank import.
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	"github.com/owncast/owncast/config"
	owncastDatabase "github.com/owncast/owncast/db"
	"github.com/owncast/owncast/v2/pkg/services/database"
)

// Service is the global key/value store for configuration values.
type Service struct {
	logger *logrus.Logger

	db    database.Dependency
	muxDb sync.Mutex

	cache    map[string][]byte
	muxCache sync.Mutex

	ready bool
}

func (s *Service) Init(rt pkg.IsRuntime) {
	go s.Setup()
}

func (s *Service) Name() string {
	return "Persistence"
}

func (s *Service) BindLogger(logger *logrus.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *logrus.Logger {
	return s.logger
}

// GetQueries will return the shared instance of the SQL query generator.
func (s *Service) GetQueries() *owncastDatabase.Queries {
	return owncastDatabase.New(s.db.GetDatabase())
}

// Get will query the database for the key and return the entry.
func (s *Service) Get(key string) (KeyVal, error) {
	cachedValue, err := s.GetCachedValue(key)
	if err == nil {
		return KeyVal{
			Key:   key,
			Value: cachedValue,
		}, nil
	}

	var resultKey string
	var resultValue []byte

	row := s.db.GetDatabase().QueryRow("SELECT key, value FROM datastore WHERE key = ? LIMIT 1", key)

	if err = row.Scan(&resultKey, &resultValue); err != nil {
		return KeyVal{}, err
	}

	result := KeyVal{
		Key:   resultKey,
		Value: resultValue,
	}
	s.SetCachedValue(resultKey, resultValue)

	return result, nil
}

// Save will save the KeyVal to the s.db.
func (s *Service) Save(e KeyVal) error {
	s.muxDb.Lock()
	defer s.muxDb.Unlock()

	var dataGob bytes.Buffer
	enc := gob.NewEncoder(&dataGob)
	if err := enc.Encode(e.Value); err != nil {
		return err
	}

	tx, err := s.db.GetDatabase().Begin()
	if err != nil {
		return err
	}
	var stmt *sql.Stmt
	stmt, err = tx.Prepare("INSERT INTO datastore (key, value) VALUES(?, ?) ON CONFLICT(key) DO UPDATE SET value=excluded.value")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(e.Key, dataGob.Bytes())

	if err != nil {
		return err
	}
	defer stmt.Close()

	if err = tx.Commit(); err != nil {
		s.Logger().Fatalln(err)
	}

	s.SetCachedValue(e.Key, dataGob.Bytes())

	return nil
}

// Setup will create the datastore table and perform initial initialization.
func (s *Service) Setup() {
	s.Logger().Info("setting up persistence...")

	s.cache = make(map[string][]byte)

	createTableSQL := `CREATE TABLE IF NOT EXISTS datastore (
		"key" string NOT NULL PRIMARY KEY,
		"value" BLOB,
		"timestamp" DATE DEFAULT CURRENT_TIMESTAMP NOT NULL
	);`

	s.MustExec(createTableSQL)

	if !s.HasPopulatedDefaults() {
		s.PopulateDefaults()
	}

	if !s.hasPopulatedFederationDefaults() {
		if err := s.SetFederationGoLiveMessage(config.GetDefaults().FederationGoLiveMessage); err != nil {
			s.Logger().Errorln(err)
		}
		if err := s.SetBool("HAS_POPULATED_FEDERATION_DEFAULTS", true); err != nil {
			s.Logger().Errorln(err)
		}
	}

	// Set the server initialization date if needed.
	if hasSetInitDate, _ := s.GetServerInitTime(); hasSetInitDate == nil || !hasSetInitDate.Valid {
		_ = s.SetServerInitTime(time.Now())
	}

	s.migrateDatastoreValues(s)

	s.ready = true
}

// Reset will delete all config entries in the datastore and start over.
func (s *Service) Reset() {
	stmt, err := s.db.GetDatabase().Prepare("DELETE FROM datastore")
	if err != nil {
		s.Logger().Fatalln(err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(); err != nil {
		s.Logger().Fatalln(err)
	}

	s.PopulateDefaults()
}

func (s *Service) Ready() bool {
	return s.ready
}

// MustExec will execute a SQL statement on a provided database instance.
func (s *Service) MustExec(command string) {
	stmt, err := s.db.GetDatabase().Prepare(command)
	if err != nil {
		s.Logger().Panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		s.Logger().Warnln(err)
	}
}
