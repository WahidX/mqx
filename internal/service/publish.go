package service

import (
	"context"
	"go-mq/internal/entities"
)

func (s *service) Publish(ctx context.Context, msg *entities.Message) error {
	// Validation of the message
	// service call
	return nil
}
