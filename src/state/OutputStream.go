package state

import (
	"fmt"
)

type OutputStream struct{}

func NewOutputStream() *OutputStream {
	return &OutputStream{}
}

func (os *OutputStream) WriteLine(value interface{}) {
	fmt.Println(value)
}
