package store

import "context"

func (s *store) Enqueue(ctx context.Context, topic string, data []byte) error {
	// Write data into a file for the topic.

	return nil
}

func (s *store) Dequeue(ctx context.Context, topic string) (data []byte, err error) {
	// Do something
	return nil, nil
}
