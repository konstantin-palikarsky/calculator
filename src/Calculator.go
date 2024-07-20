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
	commandStream.AddToFront("front")
	commandStream.AddToFront("front")
	commandStream.AddToFront("firstCommand")
	commandStream.AddToBack("lastCommand")
	commandStream.PrintValues()

	dataStack.PushFromInputStream()
	dataStack.PopToOutputStream()
}
