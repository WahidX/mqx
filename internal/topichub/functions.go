package topichub

import "net/http"

func GetTopicConns(topic string) []http.ResponseWriter {
	return topicHub[topic]
}

func AddConnection(topic string, conn http.ResponseWriter) {

	topicHub[topic] = append(topicHub[topic], conn)
}

func RemoveConn(topic string, conn http.ResponseWriter) {
	conns := topicHub[topic]
	for i, c := range conns {
		if c == conn {
			topicHub[topic] = append(conns[:i], conns[i+1:]...)
			return
		}
	}
}
