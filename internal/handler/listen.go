package handler

import (
	"encoding/json"
	"go-mq/internal/entities"
	"net/http"

	"go.uber.org/zap"
)

func (h *handler) Listen(w http.ResponseWriter, r *http.Request) {
	// Parsing the request
	var lisReq *entities.ListenerRequest
	err := json.NewDecoder(r.Body).Decode(&lisReq)
	if err != nil {
		zap.L().Debug("Error decoding the message", zap.Any("error", err))
	}

	// Validation of the ListenerRequest

	// service call
}
