package handlers

import (
	"net/http"
)

// Means we need to deenqueue a message
func (h *handler) Listen(w http.ResponseWriter, r *http.Request) {
	// Parsing the request
	// var lisReq *entities.ListenerRequest
	// err := json.NewDecoder(r.Body).Decode(&lisReq)
	// if err != nil {
	// 	zap.L().Debug("Error decoding the message", zap.Any("error", err))
	// }

	// Validation of the ListenerRequest

	topic := r.URL.Query().Get("topic")
	if topic == "" {
		sendResponse(w, http.StatusBadRequest, "topic is required")
		return
	}

	// service call
	h.service.Listen(r.Context(), w, topic)
}

func (h *handler) DequeueOne(w http.ResponseWriter, r *http.Request) {
	// Parsing the request
	topic := r.URL.Query().Get("topic")

	// service call
	msg, err := h.service.DequeueOne(r.Context(), topic)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err.Error())
	}

	// Response
	sendResponse(w, http.StatusOK, msg)
	return
}
