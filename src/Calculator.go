package main

import (
	"tuwien.ac.at/calculator/v2/src/state"
)

var registerSet = make(map[rune]interface{})
var dataStack state.Stack
var commandStream state.CommandStream
var operationMode int64 = 0

var inputStream *InputStream
var outputWriter *OutputStream

func init() {
	registerSet['a'] = "Hello! Try our Postix Calculator!"
	inputStream = InitializeInputStream()
	outputWriter = InitializeOutputStream()
}

func main() {
	commandStream.AddToFront("front")
	commandStream.AddToFront("front")
	commandStream.AddToFront("firstCommand")
	commandStream.AddToBack("lastCommand")
	commandStream.PrintValues()

}
