package state

type CommandStream []interface{}

func (stream *CommandStream) AddToFront(value interface{}) {
	*stream = append([]interface{}{value}, *stream...)
}

func (stream *CommandStream) AddToBack(value interface{}) {
	*stream = append(*stream, value)
}
