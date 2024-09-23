package handlers

import (
	"bufio"
	"context"
	"encoding/binary"
	"mqx/internal/service"
	"net"

	"go.uber.org/zap"
)

var connHandler Handler // this will be initialized in the main function

// reader is a buffered reader to read the incoming data from the client
// conn is the connection object, will be used to write the response back to the client
// Not every handler will require the reader and conn object
type Handler interface {
	Ping(ctx context.Context, conn net.Conn) error

	Publish(ctx context.Context, reader *bufio.Reader, conn net.Conn) error
	Listen(ctx context.Context, reader *bufio.Reader, conn net.Conn) error
}

type handler struct {
	service service.Service
}

func New(service service.Service) {
	connHandler = &handler{service: service}
}

func (h *handler) Ping(ctx context.Context, conn net.Conn) error {
	zap.L().Info("Ping received from: ", zap.String("remote_addr", conn.RemoteAddr().String()))
	err := binary.Write(conn, binary.BigEndian, uint32(1))
	if err != nil {
		zap.L().Warn("Error writing response", zap.Error(err))
		return err
	}

	return nil
}
