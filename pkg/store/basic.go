package store

import (
	"context"
	"fmt"
	"io"
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
	file, err := openFile(getMessageFileName(topic), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		zap.L().Warn("Failed to create file", zap.Error(err))
		return err
	}
	defer file.Close()

	// length is stored using 4 bytes
	lenBytes := []byte(fmt.Sprint(len(data)))
	for i := 0; i <= 4-len(lenBytes); i++ {
		lenBytes = append([]byte{'0'}, lenBytes...)
	}

	data = append(lenBytes, data...)
	data = append(data, '\n')

	// TODO: Write needs to be atomic
	_, err = file.Write(data)
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

	file, err := os.Open("files/" + getMessageFileName(topic))
	if err != nil {
		if err == os.ErrNotExist {
			return nil, nil
		}
		zap.L().Warn("Failed to open file", zap.Error(err))
		return nil, err
	}
	defer file.Close()

	// Length is stored using 4 bytes
	lengthBytes := make([]byte, 4)
	topicData := s.tmap.getTopic(topic)
	topicData.Mu.RLock()

	_, err = file.ReadAt(lengthBytes, topicData.Roffset)
	if err != nil {
		if err == io.EOF { // no new messages
			topicData.Mu.RUnlock()
			return nil, nil
		}

		zap.L().Warn("Failed to read length bytes from file", zap.Error(err))
		topicData.Mu.RUnlock()
		return nil, err
	}

	zap.L().Debug("Read length bytes", zap.ByteString("length", lengthBytes))

	length, _ := strconv.Atoi(string(lengthBytes))
	zap.L().Debug("Read length", zap.Int("length", length))
	length++ // +1 for \n

	contentBytes := make([]byte, length)

	_, err = file.ReadAt(contentBytes, s.tmap.getTopic(topic).Roffset+4)

	contentBytes = contentBytes[:len(contentBytes)-1] // remove \n

	zap.L().Debug("Read content bytes", zap.ByteString("content", contentBytes))

	topicData.Mu.RUnlock()
	topicData.Mu.Lock()
	defer topicData.Mu.Unlock()

	topicData.Roffset += int64(length) + 4

	return nil, nil
}
