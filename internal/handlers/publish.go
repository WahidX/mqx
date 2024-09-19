package handlers

import (
	"net"

	"go.uber.org/zap"
)

// Means we need to enqueue the message
func (h *handler) Publish(conn net.Conn) error {
	// Parsing the request

	// read json body into []byte type
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	sendResponse(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// topic := r.URL.Query().Get("topic")

	// err = h.service.Publish(r.Context(), &entities.Message{
	// 	Header:    nil,
	// 	Data:      body,
	// 	Timestamp: time.Now().Unix(),
	// 	Topic:     topic,
	// })

	// if err != nil {
	// 	sendResponse(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	zap.L().Info("Message stored successfully")
	return nil
}
