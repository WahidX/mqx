package handler

import (
	"go-mq/internal/entities"
	"net/http"
)

func (h *handler) Listen(w http.ResponseWriter, r *http.Request) {
	// Parsing the request
	var lisReq *entities.ListenerRequest
	// err := json.NewDecoder(r.Body).Decode(&lisReq)
	// if err != nil {
	// 	logger.Debug("Error decoding the message", zap.Any("error", err))
	// }

	// Validation of the ListenerRequest

	// service call
	h.service.Listen(r.Context(), w, lisReq)
}
