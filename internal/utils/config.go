package utils

import (
	"mqx/internal/entities"
	"log"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var Conf *entities.Config

// LoadConfig reads the config.yml file and unmarshals it into Config
func LoadConfig() *entities.Config {
	yamlFile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Error reading config.yml file: %s\n", err)
	}

	var config *entities.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error unmarshaling config.yml: %s\n", err)
	}

	Conf = config

	zap.L().Info("Config reloaded") // this will not work at server.go start time since then logger is not initialized
	return config
}
