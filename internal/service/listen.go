package service

import (
	"bufio"
	"context"
	"fmt"
	"mqx/internal/entities"
	"mqx/internal/topichub"
	"net"
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

func (s *service) Listen(ctx context.Context, topic string, reader *bufio.Reader, conn net.Conn) {
	// First keep dequeuing messages and write in conn
	// When there's no messages; store the conn in topicHub
	// and keep reading the connection for exit signal
	// From topicHub messages will be sent to respective conn.

	errCount := 0

	for {
		msg, err := s.DequeueOne(ctx, topic)
		if err != nil {
			if errCount > 5 {
				return
			}
			errCount++
			continue
		}

		errCount = 0

		if msg != nil {
			_, err := conn.Write(msg.Data)
			if err != nil {
				zap.L().Warn("Failed to write message", zap.Error(err))
				conn.Close()
				return
			}
			continue
		}

		//  msg == nil
		topichub.AddConnection(topic, conn)
		zap.L().Debug("New connection added in topicHub")

		_, _ = reader.ReadByte() // the program will block in this line
		zap.L().Debug("Listener disconnects, closing connection...", zap.String("remote_addr", conn.RemoteAddr().String()))
		return
	}
}

func (s *service) DequeueStream(ctx context.Context, w http.ResponseWriter, topic string) {
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
				// topichub.AddConnection(topic, w)
				continue
			}

			fmt.Fprintf(w, "New message %s\n\n", msg.Data)
			flusher.Flush() // Flush the data to the client

			time.Sleep(3 * time.Second)
		}
	}

}
