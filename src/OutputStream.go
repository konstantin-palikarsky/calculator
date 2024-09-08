package main

import (
	"fmt"
)

type OutputStream struct{}

func InitializeOutputStream() *OutputStream {
	return &OutputStream{}
}

func (os *OutputStream) WriteLine(value interface{}) {
	fmt.Println(value)
}
