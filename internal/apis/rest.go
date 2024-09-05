package apis

import (
	"go-mq/internal/handler"
	"net/http"
)

func RestMux(handlers handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", handlers.Ping)
	mux.HandleFunc("POST /publish", handlers.Publish)
	mux.HandleFunc("GET /listen", handlers.Listen)

	return mux
}
