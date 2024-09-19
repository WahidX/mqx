package apis

import (
	"go-mq/internal/handler"
	"net/http"
)

func RestMux(handlers handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", handlers.Ping)

	mux.HandleFunc("POST /publish", handlers.Publish) // this will enqueue published messages
	mux.HandleFunc("GET /listen", handlers.Listen)    // this will dequeue available messages
	mux.HandleFunc("GET /message", handlers.DequeueOne)

	return mux
}
