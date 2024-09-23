package topichub

import (
	"context"
	"net"
	"net/http"
)

func GetTopicConns(topic string) []net.Conn {
	return topicHub[topic]
}

func AddConnection(topic string, conn net.Conn) {

	topicHub[topic] = append(topicHub[topic], conn)
}

func RemoveConn(topic string, conn net.Conn) {
	conns := topicHub[topic]
	for i, c := range conns {
		if c == conn {
			topicHub[topic] = append(conns[:i], conns[i+1:]...)
			return
		}
	}
}

func CloseAllConns(ctx context.Context) {
	for _, conns := range topicHub {
		for _, conn := range conns {
			conn.Write([]byte("Server is shutting down"))
			conn.(http.Flusher).Flush()
		}
	}
}
