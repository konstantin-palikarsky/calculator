package main

import (
	"tuwien.ac.at/calculator/v2/src/state"
)

var registerSet = make(map[rune]interface{})
var dataStack state.Stack
var commandStream state.CommandStream
var operationMode int64 = 0

func init() {
	registerSet['a'] = "'"
}

func main() {
	commandStream.AddToBack("lastCommand")
	commandStream.AddToFront("firstCommand")

	dataStack.PushFromInputStream()

	dataStack.PopToOutputStream()
}
