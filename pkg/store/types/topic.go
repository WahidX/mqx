package types

type Topic struct {
	Name   string
	Offset int64
}

func (t *Topic) FileName() string {
	return "files/" + t.Name + ".msg"
}

func (t *Topic) String() string {
	return t.Name
}
