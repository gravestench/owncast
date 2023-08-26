package persistence

import (
	"path/filepath"
	"time"

	"github.com/owncast/owncast/config"
	"github.com/owncast/owncast/utils"
)

// SetupPersistence will open the datastore and make it available.
func (s *Service) SetupPersistence() error {
	s.Setup()

	dbBackupTicker := time.NewTicker(1 * time.Hour)
	go func() {
		backupFile := filepath.Join(config.BackupDirectory, "owncastdb.bak")
		for range dbBackupTicker.C {
			utils.Backup(s.db.GetDatabase(), backupFile)
		}
	}()

	return nil
}
