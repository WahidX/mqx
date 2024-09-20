package handlers

import (
	"bufio"
	"context"
	"net"
)

// Means we need to deenqueue a message
func (h *handler) Listen(ctx context.Context, reader *bufio.Reader, conn net.Conn) error {
	// topic := r.URL.Query().Get("topic")
	// if topic == "" {
	// 	sendResponse(w, http.StatusBadRequest, "topic is required")
	// 	return
	// }

	// // service call
	// h.service.Listen(r.Context(), w, topic)
	return nil
}

func (h *handler) DequeueOne(ctx context.Context, reader *bufio.Reader, conn net.Conn) error {
	// Parsing the request
	// topic := r.URL.Query().Get("topic")

	// // service call
	// msg, err := h.service.DequeueOne(r.Context(), topic)
	// if err != nil {
	// 	sendResponse(w, http.StatusInternalServerError, err.Error())
	// }

	return nil
}
