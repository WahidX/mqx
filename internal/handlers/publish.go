package handlers

import (
	"bufio"
	"context"
	"encoding/binary"
	"go-mq/internal/entities"
	"io"
	"net"
	"time"

	"go.uber.org/zap"
)

// Means we need to enqueue the message
func (h *handler) Publish(ctx context.Context, reader *bufio.Reader, conn net.Conn) error {
	// Parsing the request

	// Read topic name
	// Read message length
	// Read message body

	topic, err := reader.ReadString('\n')
	if err != nil {
		zap.L().Warn("Error reading topic", zap.Error(err))
		return err
	}

	var msgLen uint32
	err = binary.Read(reader, binary.BigEndian, &msgLen)
	if err != nil {
		zap.L().Warn("Error reading message length", zap.Error(err))
		return err
	}

	msg := make([]byte, msgLen)
	_, err = io.ReadFull(reader, msg)
	if err != nil {
		zap.L().Warn("Error reading message body", zap.Error(err))
		return err
	}

	zap.L().Debug("Message received", zap.String("topic", topic), zap.String("message", string(msg)))

	err = h.service.Enqueue(ctx, &entities.Message{
		Header:    nil,
		Data:      msg,
		Timestamp: time.Now().Unix(),
		Topic:     topic,
	})
	if err != nil {
		return err
	}

	zap.L().Debug("Message stored successfully")
	return nil
}
