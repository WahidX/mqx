package handler

import (
	"encoding/json"
	"go-mq/internal/entities"
	"net/http"

	"go.uber.org/zap"
)

func (h *handler) Publish(w http.ResponseWriter, r *http.Request) {
	// Parsing the request
	var msg *entities.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		zap.L().Debug("Error decoding the message", zap.Any("error", err))
	}

	// Validation of the message

	// service call
}
