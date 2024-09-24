package service

import (
	"bufio"
	"context"
	"fmt"
	"mqx/internal/entities"
	"mqx/internal/topichub"
	"net"

	"go.uber.org/zap"
)

func (s *service) DequeueOne(ctx context.Context, topic string) (*entities.Message, error) {
	msgRow, err := s.Repository.DequeueMessage(ctx, topic)
	if err != nil {
		return nil, fmt.Errorf("failed to get message: %v", err)
	}

	if msgRow == nil {
		return nil, nil
	}

	msg := &entities.Message{
		Data:      msgRow.Data,
		Timestamp: msgRow.Timestamp,
		Topic:     msgRow.Topic,
	}

	return msg, nil
}

func (s *service) Listen(ctx context.Context, topic string, reader *bufio.Reader, conn net.Conn) {
	// First keep dequeuing messages and write in conn
	// When there's no messages; store the conn in topicHub
	// and keep reading the connection for exit signal
	// From topicHub messages will be sent to respective conn.

	errCount := 0

	for {
		msg, err := s.DequeueOne(ctx, topic)
		if err != nil {
			if errCount > 5 {
				return
			}
			errCount++
			continue
		}

		errCount = 0

		if msg != nil {
			_, err := conn.Write(msg.Data)
			if err != nil {
				zap.L().Warn("Failed to write message", zap.Error(err))
				conn.Close()
				return
			}
			continue
		}

		//  msg == nil
		topichub.AddConnection(topic, conn)
		zap.L().Debug("New connection added in topicHub")

		_, _ = reader.ReadByte() // the program will block in this line
		zap.L().Debug("Listener disconnects, closing connection...", zap.String("remote_addr", conn.RemoteAddr().String()))
		return
	}
}
