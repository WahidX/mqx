package handler

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-mq/internal/service"
	"io"
	"net"
	"net/http"

	"go.uber.org/zap"
)

type Handler interface {
	Ping(http.ResponseWriter, *http.Request)
	Publish(http.ResponseWriter, *http.Request)
	Listen(http.ResponseWriter, *http.Request)
	DequeueOne(http.ResponseWriter, *http.Request)
}

type handler struct {
	service service.Service
}

func New(service service.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong")) // nolint: errcheck
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) error {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err := w.Write(response)

	return err
}

func Handle(conn net.Conn) error {
	reader := bufio.NewReader(conn)

	for {
		commandByte, err := reader.ReadByte()
		if err == io.EOF { // It will keep looping for any new messages on the same connection until the client disconnects
			zap.L().Warn("Client disconnected", zap.Error(err))
			return err
		}
		if err != nil {
			fmt.Println(err == io.EOF)
			zap.L().Warn("Error reading command byte", zap.Error(err))
			return err
		}

		switch Command(commandByte) {
		case Publish:
			// Reading the message length
			var msgLen uint32
			err := binary.Read(reader, binary.BigEndian, &msgLen)
			if err != nil {
				zap.L().Warn("Error reading message length", zap.Error(err))
				return err
			}

			// read full message
			msg := make([]byte, msgLen)
			_, err = io.ReadFull(reader, msg)
			if err != nil {
				zap.L().Warn("Error reading message body", zap.Error(err))
				return err
			}

			fmt.Println("Received message:", string(msg))
			// enqueue it

			// Acknowledge the message
			conn.Write([]byte("Message received"))

		case Listen:
			// read the topic
			topic, err := reader.ReadString('\n')
			if err != nil {
				zap.L().Warn("Error reading topic", zap.Error(err))
				return err
			}

			zap.L().Debug("Need message", zap.String("topic", topic))
			// add the connection to the topic

		default:
			zap.L().Info("Unknown command", zap.Any("incoming command", commandByte))
			conn.Write([]byte("Unknown command"))

		}
	}
}
