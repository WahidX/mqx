package service

import (
	"context"
	"fmt"
	"go-mq/internal/entities"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func (s *service) Listen(ctx context.Context, w http.ResponseWriter, lReq *entities.ListenerRequest) {
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

	var i = 0

	for {
		select {
		case <-ctx.Done():
			zap.L().Info("Listener disconnected")
			return
		default:
			fmt.Fprintf(w, "data: Message %d\n\n", i)
			flusher.Flush() // Flush the data to the client

			i++
			time.Sleep(1 * time.Second)
		}
	}

}
