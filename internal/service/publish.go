package service

import (
	"context"
	"mqx/internal/entities"
	"mqx/internal/topichub"

	"go.uber.org/zap"
)

func (s *service) Publish(ctx context.Context, msg *entities.Message) error {
	enqueueMesasge := func() error {
		// Enqueue the message
		_, err := s.Repository.EnqueueMessage(ctx, &entities.MessageRow{
			Data:      msg.Data,
			Timestamp: msg.Timestamp,
			Topic:     msg.Topic,
		})
		if err == nil {
			zap.L().Debug("Message stored successfully")
		}
		return err
	}

	// Start pushing the message to connected listeners if there's any
	conns := topichub.GetTopicConns(msg.Topic)
	if len(conns) == 0 {
		return enqueueMesasge()
	}

	sent := 0

	for _, conn := range conns {
		_, err := conn.Write(msg.Data)
		if err != nil {
			conn.Close()
			zap.L().Debug("Failed to write message to listener", zap.Error(err))
			continue
		} else {
			sent++
		}
	}

	// Failed to send message to any of the listeners
	if sent == 0 {
		return enqueueMesasge()
	}

	return nil
}
