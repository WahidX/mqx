package store

import (
	"context"
)

func init() {
	// Find the index file and load the topics into memory by filling the store struct.
}

type Store interface {
	Enqueue(ctx context.Context, topic string, data []byte) error
	Dequeue(ctx context.Context, topic string) (data []byte, err error)
}

type store struct {
	tmap topicMap
}

func New() Store {
	m := make(topicMap)

	return &store{
		tmap: m,
	}
}
