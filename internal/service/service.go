package service

import (
	"bufio"
	"context"
	"mqx/internal/entities"
	"mqx/internal/repository"
	"mqx/internal/topichub"
	"net"
)

type Service interface {
	Publish(ctx context.Context, msg *entities.Message) error
	Listen(ctx context.Context, topic string, reader *bufio.Reader, conn net.Conn)
	DequeueOne(ctx context.Context, topic string) (*entities.Message, error)
}

type service struct {
	Repository repository.Repository
}

func New(repository repository.Repository) Service {
	topichub.InitTopicHub()

	return &service{
		Repository: repository,
	}
}
