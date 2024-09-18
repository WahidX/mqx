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
	// 	zap.L().Debug("Error decoding the message", zap.Any("error", err))
	// }

	// Validation of the ListenerRequest

	// service call
	h.service.Listen(r.Context(), w, lisReq)
}

func (h *handler) GetSingleMessage(w http.ResponseWriter, r *http.Request) {
	// Parsing the request
	topic := r.URL.Query().Get("topic")

	// service call
	msg, err := h.service.GetSingleMessage(r.Context(), topic)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err.Error())
	}

	// Response
	sendResponse(w, http.StatusOK, msg)
	return
}
