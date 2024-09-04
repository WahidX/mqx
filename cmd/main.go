package main

import (
	"go-mq/internal/entities"
	"io/ioutil"
	"log"
	"net/http"

	"go-mq/pkg/logger"

	"gopkg.in/yaml.v2"
)

func main() {
	cfg := loadConfig()
	logger.Init(cfg.Env)
	defer logger.L.Sync()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", pingHandler)
	http.ListenAndServe(":4000", mux)
}

func loadConfig() *entities.Config {
	yamlFile, err := ioutil.ReadFile("config.yml")
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

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

// func readEnv(cfg *entities.Config) {
// err := envconfig.Process("", cfg)
// 	if err != nil {
// 		processError(err)
// 	}
// }
