package handlers

import (
	"encoding/binary"
	"go-mq/internal/service"
	"net"

	"go.uber.org/zap"
)

var connHandler Handler // this will be initialized in the main function

type Handler interface {
	Ping(conn net.Conn) error

	Publish(conn net.Conn) error
	Listen(conn net.Conn) error
	DequeueOne(conn net.Conn) error
}

type handler struct {
	service service.Service
}

func New(service service.Service) {
	connHandler = &handler{service: service}
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
