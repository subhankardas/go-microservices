package core

import (
	"fmt"
	"os"
	"sync"

	"github.com/subhankardas/go-microservices/restaurant-service/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field

var (
	Int    = zap.Int
	String = zap.String
	Error  = zap.Error
	Bool   = zap.Bool
	Any    = zap.Any
)

type Level = zapcore.Level

var (
	DebugLevel = zapcore.DebugLevel
	ErrorLevel = zapcore.ErrorLevel
	FatalLevel = zapcore.FatalLevel
	InfoLevel  = zapcore.InfoLevel
	PanicLevel = zapcore.PanicLevel
	WarnLevel  = zapcore.WarnLevel
)

// Wrapper interface for our core logger based on zap logger.
type Logger interface {
	Debugf(id string, msg string, values ...interface{})
	Debug(id string, msg string, fields ...Field)
	Errorf(id string, msg string, values ...interface{})
	Error(id string, msg string, fields ...Field)
	Fatalf(id string, msg string, values ...interface{})
	Fatal(id string, msg string, fields ...Field)
	Infof(id string, msg string, values ...interface{})
	Info(id string, msg string, fields ...Field)
	Panicf(id string, msg string, values ...interface{})
	Panic(id string, msg string, fields ...Field)
	Warnf(id string, msg string, values ...interface{})
	Warn(id string, msg string, fields ...Field)
}

// Internal struct to wrap around zap logger reference.
type loggerImpl struct {
	zap *zap.Logger
}

// Implementations for Logger interface //

func (log *loggerImpl) Debugf(id string, msg string, values ...interface{}) {
	log.zap.Debug(fmt.Sprintf(msg, values...), String(ID, id))
}

func (log *loggerImpl) Debug(id string, msg string, fields ...Field) {
	fields = append(fields, String(ID, id))
	log.zap.Debug(msg, fields...)
}

func (log *loggerImpl) Errorf(id string, msg string, values ...interface{}) {
	log.zap.Error(fmt.Sprintf(msg, values...), String(ID, id))
}

func (log *loggerImpl) Error(id string, msg string, fields ...Field) {
	fields = append(fields, String(ID, id))
	log.zap.Error(msg, fields...)
}

func (log *loggerImpl) Fatalf(id string, msg string, values ...interface{}) {
	log.zap.Fatal(fmt.Sprintf(msg, values...), String(ID, id))
}

func (log *loggerImpl) Fatal(id string, msg string, fields ...Field) {
	fields = append(fields, String(ID, id))
	log.zap.Fatal(msg, fields...)
}

func (log *loggerImpl) Infof(id string, msg string, values ...interface{}) {
	log.zap.Info(fmt.Sprintf(msg, values...), String(ID, id))
}

func (log *loggerImpl) Info(id string, msg string, fields ...Field) {
	fields = append(fields, String(ID, id))
	log.zap.Info(msg, fields...)
}

func (log *loggerImpl) Panicf(id string, msg string, values ...interface{}) {
	log.zap.Panic(fmt.Sprintf(msg, values...), String(ID, id))
}

func (log *loggerImpl) Panic(id string, msg string, fields ...Field) {
	fields = append(fields, String(ID, id))
	log.zap.Panic(msg, fields...)
}

func (log *loggerImpl) Warnf(id string, msg string, values ...interface{}) {
	log.zap.Warn(fmt.Sprintf(msg, values...), String(ID, id))
}

func (log *loggerImpl) Warn(id string, msg string, fields ...Field) {
	fields = append(fields, String(ID, id))
	log.zap.Warn(msg, fields...)
}

var lock = &sync.Mutex{}
var logger *loggerImpl // Singleton logger instance

// Constructor for core logger with configurable file path and default log level.
func NewLogger(config models.Log) Logger {
	if logger == nil { // Need to create new instance
		lock.Lock()
		defer lock.Unlock()

		if logger == nil { // Instance not created by other routines even after locking
			logger = newLoggerImpl(config) // Hence create singleton logger instance
		}
	}
	return logger
}

// Creates new logger implementation instance.
func newLoggerImpl(cfg models.Log) *loggerImpl {
	// Create logger config and encoders
	config := getLogEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	// Create log file with filepath
	logFile, _ := os.OpenFile(cfg.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// Create zap core with file and console encoders with default log level
	defaultLogLevel := cfg.LogLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	// Create new logger with add caller and skip 1 level and return
	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	return &loggerImpl{log}
}

// Returns logger config with default settings.
func getLogEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:       "message",
		LevelKey:         "level",
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		TimeKey:          "timestamp",
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		CallerKey:        "caller",
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: "	|	",
	}
}
