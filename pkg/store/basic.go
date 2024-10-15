package store

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"

	"go.uber.org/zap"
)

func (s *store) Enqueue(ctx context.Context, topic string, data []byte) error {
	// Check if the topic name is valid
	// Get the file
	// get the offset for writing data
	// prepare the message
	// Write the message

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
		offset = int(t.Roffset)
	}

	data = append([]byte(fmt.Sprint(len(data))), data...)

	// TODO: Write needs to be atomic
	_, err = file.WriteAt(data, int64(offset))
	if err != nil {
		zap.L().Warn("Failed to write to file", zap.Error(err))
		return err
	}

	return nil
}

func (s *store) Dequeue(ctx context.Context, topic string) (data []byte, err error) {
	// validate the topic name
	// Open the file to read
	// Get the offset to read from
	// Decide how long it needs to read ()
	// Read the message

	if err := topicNameIsValid(topic); err != nil {
		return nil, err
	}

	file, err := os.Open(getMessageFileName(topic))
	if err != nil {
		if err == os.ErrNotExist {
			return nil, nil
		}
		zap.L().Warn("Failed to open file", zap.Error(err))
		return nil, err
	}

	// Length is stored using 4 bytes
	lengthBytes := make([]byte, 4)

	_, err = file.ReadAt(lengthBytes, s.topicMap[topic].Roffset)
	if err != nil {
		if err == io.EOF { // no new messages
			return nil, nil
		}

		zap.L().Warn("Failed to read length from file", zap.Error(err))
		return nil, err
	}

	zap.L().Debug("Read length bytes", zap.ByteString("length", lengthBytes))

	length, _ := strconv.Atoi(string(lengthBytes))
	zap.L().Debug("Read length", zap.Int("length", length))

	contentBytes := make([]byte, length) // +1 for \n

	_, err = file.ReadAt(contentBytes, s.topicMap[topic].Roffset)

	zap.L().Debug("Read content bytes", zap.ByteString("content", contentBytes))

	return nil, nil
}
