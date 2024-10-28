package store

import (
	"mqx/pkg/store/types"
	"sync"
)

type topicMap map[string]*types.Topic

// getTopic returns the topic data for the given topic name.
// If the topic does not exist, it creates a new topic by default.
func (t topicMap) getTopic(topic string) *types.Topic {
	topicData, ok := t[topic]

	if !ok {
		topicData = &types.Topic{
			Mu:      sync.RWMutex{},
			Name:    topic,
			Roffset: 0,
			Woffset: 0,
		}

		t[topic] = topicData
	}

	return topicData
}
