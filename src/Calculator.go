package main

import "fmt"
import "tuwien.ac.at/calculator/v2/src/streams"

var registerSet = make(map[rune]interface{})
var operationMode int64
var dataStack = make([]interface{}, 8)

var commandStream = streams.NewStream()
var inputStream = streams.NewStream()
var outputStream = streams.NewStream()

func init() {
	operationMode = 0

	//runs once before main, should use to set the registers
	registerSet['a'] = "firstCommand"
}

func main() {
	fmt.Println("hi world", operationMode)
	commandStream.AddToBack("lastCommand")
	commandStream.AddToFront("firstCommand")

	fmt.Println(commandStream)
}
