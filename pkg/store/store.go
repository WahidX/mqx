package store

import (
	"context"
	"mqx/pkg/store/types"
)

var storeSnapShot map[string]*types.Topic

func init() {
	// Find the index file and load the topics into memory by filling the store struct.

}

type Store interface {
	Enqueue(ctx context.Context, topic string, data []byte) error
	Dequeue(ctx context.Context, topic string) (data []byte, err error)
}

type store struct {
	topicMap map[string]*types.Topic
}

func New() Store {
	return &store{}
}
