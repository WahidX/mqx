package service

import (
	"context"
	"go-mq/internal/entities"
	"go-mq/internal/repository"
	"go-mq/internal/topichub"
	"net/http"
)

type Service interface {
	Enqueue(ctx context.Context, msg *entities.Message) error
	DequeueOne(ctx context.Context, topic string) (*entities.Message, error)
	DequeueStream(ctx context.Context, w http.ResponseWriter, topic string)
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
