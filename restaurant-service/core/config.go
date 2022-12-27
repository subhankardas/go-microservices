package core

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
)

// Returns config properties from given path, profile (filename) and extension.
func LoadConfig(path, profile, extension string) *models.Config {
	config := readConfigFromFile(path, profile, extension)

	// Update properties of any type to required values
	config.Log.LogLevel = getLogLevel(config.Log.Level)
	config.Server.RequestTimeoutDuration = time.Duration(time.Duration(config.Server.RequestTimeoutInMillisecond) * time.Millisecond)

	return config
}

// Loads config properties from YAML file.
func readConfigFromFile(path, filename, extension string) *models.Config {
	var config = &models.Config{}

	viper.AddConfigPath(path)      // Path to look for the config file
	viper.SetConfigName(filename)  // Name of config file (without extension)
	viper.SetConfigType(extension) // REQUIRED if the config file does not have the extension in the name

	// Find and read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error: %s, cause: %v", UNABLE_TO_LOAD_CONFIG_FILE, err)
	}

	// Parse config properties
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("error: %s, cause: %v", UNABLE_TO_READ_CONFIG_FILE, err)
	}

	return config
}

// Maps log level code to actual log level type.
func getLogLevel(level string) Level {
	switch level {
	case ErrorLevel.String():
		return ErrorLevel
	case FatalLevel.String():
		return FatalLevel
	case InfoLevel.String():
		return InfoLevel
	case PanicLevel.String():
		return PanicLevel
	case WarnLevel.String():
		return WarnLevel
	}
	return DebugLevel
}
