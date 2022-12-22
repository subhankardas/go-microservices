package core

import (
	"log"
	"os"

	"github.com/subhankardas/go-microservices/restaurant-service/models"
	"gopkg.in/yaml.v2"
)

// Reads and return config properties.
func LoadConfig(path string) *models.Config {
	config := readConfigFromYamlFile(path)

	// Update properties of any type to required values
	config.Log.Level = getLogLevel(config.Log.LevelName)

	return config
}

// Loads config properties from YAML file.
func readConfigFromYamlFile(path string) *models.Config {
	var config = &models.Config{}

	// Open config YAML file in read only mode, close in defer
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("error: %s, cause: %v", UNABLE_TO_LOAD_CONFIG_FILE, err)
	}
	defer file.Close()

	// Read config properties and return
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
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
