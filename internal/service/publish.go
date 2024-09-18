package service

import (
	"context"
	"go-mq/internal/entities"
)

func (s *service) Publish(ctx context.Context, msg *entities.Message) error {
	// Enqueue the message
	_, err := s.Repository.EnqueueMessage(ctx, &entities.MessageRow{
		Data:      msg.Data,
		Timestamp: msg.Timestamp,
		Topic:     msg.Topic,
		Partition: msg.Partition,
	})

	return err
}
