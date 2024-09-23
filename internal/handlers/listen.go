package handlers

import (
	"bufio"
	"context"
	"net"

	"go.uber.org/zap"
)

// Means we need to dequeue a message
func (h *handler) Listen(ctx context.Context, reader *bufio.Reader, conn net.Conn) error {
	topic, err := reader.ReadString('\n')
	if err != nil {
		zap.L().Warn("Error reading topic name", zap.Error(err))
		return err
	}

	h.service.Listen(ctx, topic, reader, conn)

	return nil
}
