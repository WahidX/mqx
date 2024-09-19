package handler

type Command byte

const (
	Publish Command = iota + 1
	Listen
	// More commands can be added here
)
