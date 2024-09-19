package service

import (
	"context"
	"fmt"
	"go-mq/internal/entities"
	"go-mq/internal/topichub"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func (s *service) DequeueOne(ctx context.Context, topic string) (*entities.Message, error) {
	msgRow, err := s.Repository.DequeueMessage(ctx, topic)
	if err != nil {
		return nil, fmt.Errorf("failed to get message: %v", err)
	}

	if msgRow == nil {
		return nil, nil
	}

	msg := &entities.Message{
		Data:      msgRow.Data,
		Timestamp: msgRow.Timestamp,
		Topic:     msgRow.Topic,
	}

	return msg, nil
}

func (s *service) Listen(ctx context.Context, w http.ResponseWriter, topic string) {
	// Set headers to indicate that the connection should remain open
	w.Header().Set("Content-Type", "text/event-stream") // For SSE, you can also use "text/plain"
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Flusher is required to flush the response writer buffer manually
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Stop the stream in case context cancelled
	// Simulate a stream of messages

	errCount := 0

	for {
		select {
		case <-ctx.Done():
			zap.L().Info("Listener disconnected")
			return

		default:
			// Keep Getting messages from the topic
			msg, err := s.DequeueOne(ctx, topic)
			if err != nil {
				errCount++
				if errCount > 5 {
					zap.L().Warn("Error getting message", zap.Any("error", err))
					return
				}
			}

			errCount = 0

			if msg == nil {
				// save the connection in topicHub and wait for any new message
				topichub.AddConnection(topic, w)
				continue
			}

			fmt.Fprintf(w, "New message %s\n\n", msg.Data)
			flusher.Flush() // Flush the data to the client

			time.Sleep(3 * time.Second)
		}
	}

}
