package entities

type Message struct {
	Header    map[string]any
	Data      []byte
	Timestamp int64
	Topic     string
	Partition int
}
