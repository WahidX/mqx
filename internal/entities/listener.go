package entities

type ListenerRequest struct {
	Topic  string
	Offset int64

	// more configs
}
