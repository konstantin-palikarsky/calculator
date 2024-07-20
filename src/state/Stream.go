package state

import "fmt"

type CommandStream []interface{}

func (stream *CommandStream) AddToFront(value interface{}) {
	*stream = append([]interface{}{value}, *stream...)
}

func (stream *CommandStream) AddToBack(value interface{}) {
	*stream = append(*stream, value)
}

func (stream *CommandStream) PrintValues() {
	fmt.Println("Command Stream")
	for i, elem := range *stream {
		fmt.Println(i, ":", elem)
	}
}
