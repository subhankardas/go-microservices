package models

import "go.uber.org/zap/zapcore"

type LogLevel = zapcore.Level

type Config struct {
	Log      Log      `yaml:"log"`
	Database Database `yaml:"database"`
}

type Log struct {
	Filepath  string   `yaml:"filepath"`
	LevelName string   `yaml:"level"`
	Level     LogLevel `yaml:"-"`
}

type Database struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
