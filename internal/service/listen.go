package service

import (
	"context"
	"fmt"
	"go-mq/internal/entities"
	"go-mq/pkg/logger"
	"net/http"
	"time"
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
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "data: Message %d\n\n", i+1)
		flusher.Flush() // Flush the data to the client

		time.Sleep(1 * time.Second)
	}

	logger.Info("Stream closed")
}
