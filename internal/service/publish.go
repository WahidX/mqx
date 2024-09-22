package service

import (
	"context"
	"fmt"
	"mqx/internal/entities"
	"mqx/internal/topichub"
	"net/http"
)

func (s *service) Enqueue(ctx context.Context, msg *entities.Message) error {
	enqueueMesasge := func() error {
		// Enqueue the message
		_, err := s.Repository.EnqueueMessage(ctx, &entities.MessageRow{
			Data:      msg.Data,
			Timestamp: msg.Timestamp,
			Topic:     msg.Topic,
		})
		return err
	}

	// Start pushing the message to connected listeners if there's any
	resWriters := topichub.GetTopicConns(msg.Topic)
	if len(resWriters) == 0 {
		return enqueueMesasge()
	}

	sent := 0

	for _, w := range resWriters {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming not supported", http.StatusInternalServerError)
			// !! Need to close the connection somehow

			continue
		}

		// Write the message to the response writer
		_, e := fmt.Fprintln(w, msg)
		if e != nil {
			// !! Need to close the connection somehow
			continue
		} else {
			sent++
			flusher.Flush()
		}
	}

	// Failed to send message to any of the listeners
	if sent == 0 {
		return enqueueMesasge()
	}

	return nil
}
