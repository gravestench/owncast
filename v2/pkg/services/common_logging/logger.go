package common_logging

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type customFormatter struct {
	serviceName string
}

func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("3:04PM")
	return []byte(fmt.Sprintf("%s %s [%s]: %s\n", timestamp, strings.ToUpper(entry.Level.String()[:3]), f.serviceName, entry.Message)), nil
}

func createLogger(serviceName string) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&customFormatter{
		serviceName: serviceName,
	})

	return logger
}
