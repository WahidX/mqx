package cmd

import (
	"fmt"
	"go-mq/internal/entities"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	var cfg entities.Config
	readFile(&cfg)
	readEnv(&cfg)
	fmt.Printf("%+v", cfg)
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

func readEnv(cfg *entities.Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
