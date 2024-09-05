package entities

type ListenerRequest struct {
	Topic     string
	Partition int
	Offset    int64

	// more configs
}
