package service

import "go-mq/internal/repository"

type Service interface {
}

type service struct {
	Repository repository.Repository
}

func New(repository repository.Repository) Service {
	return &service{
		Repository: repository,
	}
}
