package service

import (
	"bufio"
	"context"
	"mqx/internal/entities"
	"mqx/internal/repository"
	"mqx/internal/topichub"
	"net"
	"net/http"
)

type Service interface {
	Enqueue(ctx context.Context, msg *entities.Message) error
	DequeueOne(ctx context.Context, topic string) (*entities.Message, error)
	DequeueStream(ctx context.Context, w http.ResponseWriter, topic string)

	Listen(ctx context.Context, topic string, reader *bufio.Reader, conn net.Conn)
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
