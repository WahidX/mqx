package entities

import "time"

type Message struct {
	Header    map[string]any
	Data      []byte
	Timestamp int64
	Topic     string
	Partition int
}

type MessageRow struct {
	Data      []byte
	Timestamp time.Time
	Topic     string
	Partition int
}
