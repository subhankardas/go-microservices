package models

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type LogLevel = zapcore.Level

type Config struct {
	Server   Server
	Log      Log
	Database Database
}

type Server struct {
	Port                        string
	RequestTimeoutInMillisecond uint16
	RequestTimeoutDuration      time.Duration
}

type Log struct {
	Filepath string
	Level    string
	LogLevel LogLevel
}

type Database struct {
	Name     string
	Host     string
	Port     uint16
	Username string
	Password string
}
