package store

import (
	"context"
	"io/fs"
	"os"

	"go.uber.org/zap"
)

func (s *store) Enqueue(ctx context.Context, topic string, data []byte) error {
	if err := topicNameIsValid(topic); err != nil {
		return err
	}

	// Get the existing file or create a new one
	file, err := os.OpenFile(getMessageFileName(topic), os.O_CREATE|os.O_APPEND, fs.ModeAppend)
	if err != nil {
		zap.L().Warn("Failed to open file", zap.Error(err))
		return err
	}

	offset := 0 // 0 if its a new topic
	if t, ok := s.topicMap[topic]; ok {
		offset = int(t.Offset)
	}

	// TODO: Write needs to be atomic
	_, err = file.WriteAt(data, int64(offset))
	if err != nil {
		zap.L().Warn("Failed to write to file", zap.Error(err))
		return err
	}

	return nil
}

func (s *store) Dequeue(ctx context.Context, topic string) (data []byte, err error) {
	// Do something
	return nil, nil
}
