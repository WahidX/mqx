package main

import (
	"errors"
	"net/http"
	"syscall"

	"go-mq/internal/apis"
	"go-mq/internal/db"
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
	l := logger.Init(utils.Conf.Env)
	defer func() {
		err := l.Sync()
		if err != nil && !errors.Is(err, syscall.ENOTTY) {
			zap.L().Fatal("Error syncing logger: ", zap.Any("error", err))
		}
	}()

	sqliteDB := db.Connect()

	// Setting up the layers
	repository := repository.New(sqliteDB)
	services := service.New(repository)
	handlers := handler.New(services)
	mux := apis.RestMux(handlers)

	// Starting the rest server
	zap.L().Info("Server listening port " + utils.Conf.Server.Port)
	zap.L().Fatal("Server error: ", zap.Any("error", http.ListenAndServe(":"+utils.Conf.Server.Port, mux)))
}
