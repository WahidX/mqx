package types

import "sync"

type Topic struct {
	Mu      sync.RWMutex
	Name    string
	Roffset int64
	Woffset int64
}

func (t *Topic) FileName() string {
	return "files/" + t.Name + ".msg"
}

func (t *Topic) String() string {
	return t.Name
}
