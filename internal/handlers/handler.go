package handlers

import (
	"encoding/binary"
	"encoding/json"
	"go-mq/internal/service"
	"net"
	"net/http"

	"go.uber.org/zap"
)

var connHandler Handler // this will be initialized in the main function

type Handler interface {
	// Ping(http.ResponseWriter, *http.Request)
	// Publish(http.ResponseWriter, *http.Request)
	// Listen(http.ResponseWriter, *http.Request)
	// DequeueOne(http.ResponseWriter, *http.Request)

	Ping(conn net.Conn) error
}

type handler struct {
	service service.Service
}

func New(service service.Service) {
	connHandler = &handler{service: service}
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) error {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err := w.Write(response)

	return err
}

func (h *handler) Ping(conn net.Conn) error {
	zap.L().Info("Ping received from: ", zap.String("remote_addr", conn.RemoteAddr().String()))
	err := binary.Write(conn, binary.BigEndian, uint32(1))
	if err != nil {
		zap.L().Warn("Error writing response", zap.Error(err))
		return err
	}

	return nil
}
