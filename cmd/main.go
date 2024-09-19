package main

import (
	"errors"
	"net"
	"syscall"

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
			zap.L().Warn("Error syncing logger: ", zap.Any("error", err))
		}
	}()

	sqliteDB := db.Connect()
	defer sqliteDB.Close()

	// Setting up the layers
	repository := repository.New(sqliteDB)
	services := service.New(repository)
	_ = handler.New(services)
	// mux := apis.RestMux(handlers)

	// Start listening for incoming TCP connections in a goroutine

	server, err := net.Listen("tcp", ":"+utils.Conf.Server.Port)
	if err != nil {
		zap.L().Fatal("Error starting server: %v", zap.Error(err))
	}

	zap.L().Info("Message Queue Server listening on port " + utils.Conf.Server.Port)

	// Start a new goroutine to accept and handle incoming connections
	// keeping the main thread free to handle signals
	for {
		conn, err := server.Accept()
		if err != nil {
			zap.L().Warn("Error accepting connection", zap.Error(err))
			continue
		}

		go handler.Handle(conn)
	}

}
