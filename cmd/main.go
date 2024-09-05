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
	cfg := utils.LoadConfig()

	// Logger setup
	logger.Init(cfg.Env)
	defer func() {
		err := logger.L.Sync()
		if err != nil {
			logger.L.Fatal("Error syncing logger: ", zap.Any("error", err))
		}
	}()

	repository := repository.New()
	services := service.New(repository)
	handlers := handler.New(services)

	mux := apis.RestMux(handlers)

	logger.L.Info("Server listening port " + cfg.Server.Port)
	logger.L.Fatal("Server error: ", zap.Any("error", http.ListenAndServe(":"+cfg.Server.Port, mux)))
}
