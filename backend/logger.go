package backend

import (
	"time"

	"github.com/sirupsen/logrus"
)

// Logger is the backend logger
var Logger *logrus.Logger
var logger = Logger

func init() {
	Logger = logrus.StandardLogger()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})
}
