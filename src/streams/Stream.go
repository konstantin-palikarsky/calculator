package streams

type Stream struct {
	queue []interface{}
}

func NewStream() Stream {
	return Stream{make([]interface{}, 0, 16)}
}

func (stream *Stream) AddToFront(value interface{}) {
	stream.queue = append([]interface{}{value}, stream.queue...)
}

func (stream *Stream) AddToBack(value interface{}) {
	stream.queue = append(stream.queue, value)
}
