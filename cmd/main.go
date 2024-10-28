package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mqx/internal/db"
	"mqx/internal/handlers"
	"mqx/internal/repository"
	"mqx/internal/service"
	"mqx/internal/topichub"
	"mqx/internal/utils"
	"mqx/pkg/logger"
	"mqx/pkg/store"

	"go.uber.org/zap"
)

func main() {
	// Loading config
	utils.LoadConfig()

	// Logger setup
	l := logger.Init(utils.Conf.Env)
	defer func() {
		_ = l.Sync() // Dosent work
		// if err != nil && !errors.Is(err, syscall.ENOTTY) {
		// 	zap.L().Warn("Error syncing logger: ", zap.Any("error", err))
		// }
	}()

	tempMain()
	return

	sqliteDB := db.Connect()
	defer sqliteDB.Close()

	// Setting up the layers
	repository := repository.New(sqliteDB)
	service := service.New(repository)
	handlers.New(service)

	c := net.ListenConfig{
		KeepAliveConfig: net.KeepAliveConfig{Enable: true},
	}
	ctx := context.Background()

	// Start listening for incoming TCP connections on a particular port
	server, err := c.Listen(ctx, "tcp", ":"+utils.Conf.Server.Port)
	if err != nil {
		zap.L().Fatal("Error starting server: %v", zap.Error(err))
	}
	zap.L().Info("Message Queue Server listening on port " + utils.Conf.Server.Port)

	// Start a new goroutine to accept and handle incoming connections concurrently
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				zap.L().Warn("Error accepting connection", zap.Error(err))
				continue
			}

			zap.L().Info("New connection accepted", zap.String("remote_addr", conn.RemoteAddr().String()))
			go handlers.HandleRawConn(ctx, conn)
		}
	}()

	// Channel to listen for interrupt or terminate signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM) //Ctrl+C (SIGINT) or other Interrupt signals

	// Waiting for
	<-stopChan

	zap.L().Info("Shutting down server...")

	// Context with timeout to allow graceful shutdown
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	topichub.CloseAllConns(ctx)

	zap.L().Info("Server stopped gracefully.")
}

func tempMain() {
	fmt.Println("starting temp script..............")

	s := "\n"
	b := []byte(s)
	fmt.Println(b)
	fmt.Println(len(b))

	store := store.New()

	topic := "tpc"
	ctx := context.Background()

	// err := store.Enqueue(ctx, topic, []byte("HelloThere:"+time.Now().String()))
	// if err != nil {
	// 	fmt.Println("ERR:", err)
	// 	return
	// }

	data, err := store.Dequeue(ctx, topic)
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}

	fmt.Println(string(data))

	data, err = store.Dequeue(ctx, topic)
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}

	fmt.Println(string(data))

	data, err = store.Dequeue(ctx, topic)
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}

	fmt.Println(string(data))

	data, err = store.Dequeue(ctx, topic)
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}

	fmt.Println(string(data))
}
