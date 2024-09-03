package main

import (
	"fmt"
	"go-mq/internal/entities"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	// var cfg entities.Config
	// readFile(&cfg)
	// readEnv(&cfg)
	// fmt.Printf("%+v", cfg)
	// TODO: Reading a config will take place here. Need to decide the configurable vars

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", pingHandler)

	http.ListenAndServe(":4000", mux)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *entities.Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

// func readEnv(cfg *entities.Config) {
// err := envconfig.Process("", cfg)
// 	if err != nil {
// 		processError(err)
// 	}
// }
