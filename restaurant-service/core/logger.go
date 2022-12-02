package core

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Wrapper interface for our core logger based on zap logger.
type Logger interface {
	Debug(msg string, values ...interface{})
	Error(msg string, values ...interface{})
	Fatal(msg string, values ...interface{})
	Info(msg string, values ...interface{})
	Panic(msg string, values ...interface{})
	Warn(msg string, values ...interface{})
}

// Internal struct to wrap around zap logger reference.
type loggerImpl struct {
	zap *zap.Logger
}

// Implementations for Logger interface //

// Debug logs a message at DebugLevel. The message includes any fields passed at the log
// site, as well as any fields accumulated on the logger.
func (log *loggerImpl) Debug(msg string, values ...interface{}) {
	log.zap.Debug(fmt.Sprintf(msg, values...))
}

// Error logs a message at ErrorLevel. The message includes any fields passed at the log
// site, as well as any fields accumulated on the logger.
func (log *loggerImpl) Error(msg string, values ...interface{}) {
	log.zap.Error(fmt.Sprintf(msg, values...))
}

// Fatal logs a message at FatalLevel. The message includes any fields passed at the log
// site, as well as any fields accumulated on the logger. The logger then calls os.Exit(1),
// even if logging at FatalLevel is disabled.
func (log *loggerImpl) Fatal(msg string, values ...interface{}) {
	log.zap.Fatal(fmt.Sprintf(msg, values...))
}

// Info logs a message at InfoLevel. The message includes any fields passed at the log
// site, as well as any fields accumulated on the logger.
func (log *loggerImpl) Info(msg string, values ...interface{}) {
	log.zap.Info(fmt.Sprintf(msg, values...))
}

// Panic logs a message at PanicLevel. The message includes any fields passed at the log
// site, as well as any fields accumulated on the logger. The logger then panics, even if
// logging at PanicLevel is disabled.
func (log *loggerImpl) Panic(msg string, values ...interface{}) {
	log.zap.Panic(fmt.Sprintf(msg, values...))
}

// Warn logs a message at WarnLevel. The message includes any fields passed at the log
// site, as well as any fields accumulated on the logger.
func (log *loggerImpl) Warn(msg string, values ...interface{}) {
	log.zap.Warn(fmt.Sprintf(msg, values...))
}

var lock = &sync.Mutex{}
var logger *loggerImpl // Singleton logger instance

// Constructor for logger that ensures only single instance is created for core logger.
func NewLogger(filepath string) Logger {
	if logger == nil { // Need to create new instance
		lock.Lock()
		defer lock.Unlock()

		if logger == nil { // Instance not created by other routines even after locking
			logger = newLoggerImpl(filepath) // Hence create logger instance
		}
	}
	return logger
}

// Constructor to create new core logger.
func newLoggerImpl(filepath string) *loggerImpl {
	// Create logger config and encoders
	config := getLoggerConfig()
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	// Create log file with filepath
	logFile, _ := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// Create zap core with file and console encoders with default log level
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	// Create new logger with add caller and skip 1 level and return
	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	return &loggerImpl{log}
}

// Returns logger config with default settings.
func getLoggerConfig() zapcore.EncoderConfig {
	var config = zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.MessageKey = "message"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.ConsoleSeparator = " | "

	return config
}
