package database

import (
	log "github.com/sirupsen/logrus"
)

// MustExec will execute a SQL statement on a provided database instance.
func (s *Service) MustExec(command string) {
	stmt, err := s.db.Prepare(command)
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Warnln(err)
	}
}
