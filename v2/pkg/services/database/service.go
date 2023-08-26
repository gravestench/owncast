package database

import (
	"database/sql"

	"github.com/gravestench/runtime/pkg"
	"github.com/sirupsen/logrus"
)

type Service struct {
	logger *logrus.Logger
	db     *sql.DB
}

func (s *Service) BindLogger(logger *logrus.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *logrus.Logger {
	return s.logger
}

func (s *Service) Init(rt pkg.IsRuntime) {
	s.setup(":memory:")
}

func (s *Service) Name() string {
	return "Database"
}
