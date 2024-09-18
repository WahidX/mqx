package service

import (
	"context"
	"go-mq/internal/entities"
	"go-mq/internal/repository"
	"net/http"
)

type Service interface {
	Publish(ctx context.Context, msg *entities.Message) error
	GetSingleMessage(ctx context.Context, topic string) (*entities.Message, error)
	Listen(ctx context.Context, w http.ResponseWriter, lReq *entities.ListenerRequest)
}

type service struct {
	Repository repository.Repository
}

func New(repository repository.Repository) Service {
	return &service{
		Repository: repository,
	}
}
