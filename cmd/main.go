package main

import (
	"net/http"

	"go-mq/internal/apis"
	"go-mq/internal/handler"
	"go-mq/internal/repository"
	"go-mq/internal/service"
	"go-mq/internal/utils"
	"go-mq/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// Loading config
	utils.LoadConfig()

	// Logger setup
	logger.Init(utils.Conf.Env)
	defer func() {
		err := logger.L.Sync()
		if err != nil {
			logger.L.Fatal("Error syncing logger: ", zap.Any("error", err))
		}
	}()

	// Setting up the layers
	repository := repository.New()
	services := service.New(repository)
	handlers := handler.New(services)
	mux := apis.RestMux(handlers)

	// Starting the rest server
	logger.L.Info("Server listening port " + utils.Conf.Server.Port)
	logger.L.Fatal("Server error: ", zap.Any("error", http.ListenAndServe(":"+utils.Conf.Server.Port, mux)))
}
