package handlers

import (
	"bufio"
	"context"
	"net"
	"time"

	"go.uber.org/zap"
)

// Means we need to deenqueue a message
func (h *handler) Listen(ctx context.Context, reader *bufio.Reader, conn net.Conn) error {
	// Flow:
	// - Read the topic name
	// - Keep doing (Dequeue one message from the topic and send it to the client)

	topic, err := reader.ReadString('\n')
	if err != nil {
		zap.L().Warn("Error reading topic name", zap.Error(err))
		return err
	}

	errCount := 0

	for {
		msg, err := h.service.DequeueOne(ctx, topic)
		if err != nil {
			if errCount > 5 {
				return err
			}
			errCount++
			continue
		}
		errCount = 0

		if msg == nil {
			zap.L().Debug("Waiting for message on topic:" + topic)
			time.Sleep(3 * time.Second)
			continue
		}

		zap.L().Debug("Message dequeued", zap.String("topic", topic), zap.Any("message", msg))
		_, err = conn.Write(msg.Data)
		if err != nil {
			return err
		}

	}
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
