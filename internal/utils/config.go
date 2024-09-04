package utils

import (
	"go-mq/internal/entities"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

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

	return config
}
