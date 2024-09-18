package handler

import (
	"encoding/json"
	"go-mq/internal/service"
	"net/http"
)

type Handler interface {
	Ping(http.ResponseWriter, *http.Request)
	Publish(http.ResponseWriter, *http.Request)
	Listen(http.ResponseWriter, *http.Request)
	GetSingleMessage(http.ResponseWriter, *http.Request)
}

type handler struct {
	service service.Service
}

func New(service service.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong")) // nolint: errcheck
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) error {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err := w.Write(response)

	return err
}
