package handler

import (
	"go-mq/internal/entities"
	"io"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

func (h *handler) Publish(w http.ResponseWriter, r *http.Request) {
	// Parsing the request

	// read json body into []byte type
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	topic := r.URL.Query().Get("topic")
	partition, err := strconv.Atoi(r.URL.Query().Get("partition"))
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, "Invalid partition")
		return
	}

	err = h.service.Publish(r.Context(), &entities.Message{
		Header:    nil,
		Data:      body,
		Timestamp: time.Now().Unix(),
		Topic:     topic,
		Partition: partition,
	})

	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	zap.L().Info("Message stored successfully")
	sendResponse(w, http.StatusOK, "Message stored successfully")
}
