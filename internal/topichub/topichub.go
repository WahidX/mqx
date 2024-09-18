package topichub

import (
	"net"
	"sync"

	"go.uber.org/zap"
)

var (
	topicHub map[string][]net.Conn
	once     sync.Once
)

// InitTopicHub initializes the topicHub map; will run only once
func InitTopicHub() {
	if topicHub != nil {
		zap.L().Info("TopicHub is already initialized")
		return
	}

	once.Do(func() {
		topicHub = make(map[string][]net.Conn)
		zap.L().Info("TopicHub is initialized")
	})
}
