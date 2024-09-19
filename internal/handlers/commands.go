package handlers

type Command byte

// TODO: Need to add content structure documentation
const (
	Ping Command = iota
	Publish
	Listen
	// More commands can be added here
)
