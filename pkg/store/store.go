package store

import "context"

func Init() {
	// Do something
}

type Store interface {
	Enqueue(ctx context.Context, topic string, data []byte) error
	Dequeue(ctx context.Context, topic string) (data []byte, err error)
}

type store struct {
}

func New() Store {
	return &store{}
}
