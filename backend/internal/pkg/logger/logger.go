package logger

import (
	"log"

	"go.uber.org/zap"
)

// Logger wraps zap logger
type Logger struct {
	logger *zap.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("Failed to create logger", err)
	}
	return &Logger{logger: logger}
}

// Info logs an info message
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

// Error logs an error message
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

// Sync flushes any buffered log entries
func (l *Logger) Sync() error {
	return l.logger.Sync()
}
