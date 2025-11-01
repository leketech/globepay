package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger creates a new configured logger
func NewLogger(level logrus.Level, debug bool) *logrus.Logger {
	logger := logrus.New()
	
	// Set log level
	logger.SetLevel(level)
	
	// Set output
	logger.SetOutput(os.Stdout)
	
	// Set formatter
	if debug {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	} else {
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
	
	return logger
}