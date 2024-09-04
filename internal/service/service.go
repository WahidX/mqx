package service

type Service interface {
}

type service struct {
}

func New() Service {
	return &service{}
}
