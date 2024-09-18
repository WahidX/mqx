package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
			zap.L().Warn("Error syncing logger: ", zap.Any("error", err))
		}
	}()

	sqliteDB := db.Connect()
	defer sqliteDB.Close()

	// Setting up the layers
	repository := repository.New(sqliteDB)
	services := service.New(repository)
	handlers := handler.New(services)
	mux := apis.RestMux(handlers)

	server := &http.Server{
		Addr:    ":" + utils.Conf.Server.Port,
		Handler: mux,
	}

	// Channel to listen for interrupt or terminate signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		zap.L().Info("Server listening on port " + server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("Could not listen on " + server.Addr + err.Error())
		}
	}()

	// Waiting for Ctrl+C (SIGINT) or other termination signals
	<-stopChan
	zap.L().Info("Shutting down server...")

	// Context with timeout to allow graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server forced to shutdown: " + err.Error())
	}

	zap.L().Info("Server stopped gracefully.")
}
