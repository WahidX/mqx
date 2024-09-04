package handler

import (
	"go-mq/internal/service"
	"net/http"
)

type Handler interface {
	Ping(http.ResponseWriter, *http.Request)
}

type handler struct {
	service service.Service
}

func New(service service.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}
