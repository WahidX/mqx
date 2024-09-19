package handler

type Command byte

const (
	Ping Command = iota
	Publish
	Listen
	// More commands can be added here
)
