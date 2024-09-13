package main

import (
	"fmt"
	"tuwien.ac.at/calculator/v2/src/calculator"
	"tuwien.ac.at/calculator/v2/src/state"
)

func main() {
	cs := state.NewCommandStream()
	ds := state.NewStack()
	is := state.NewInputStream()
	os := state.NewOutputStream()

	calc := calculator.NewCalculator(cs, ds, is, os)

	initialCommand, ok := calc.Registers['a'].(string)
	if !ok {
		calc.OutputStream.WriteLine("Error: Invalid content in register 'a'")
		return
	}
	calc.CommandStream.AddToBack(initialCommand)

	calc.Run()

	for {
		calc.OutputStream.WriteLine("Enter a command (or 'quit' to exit):")
		input, err := calc.InputStream.ReadLine()
		if err != nil {
			calc.OutputStream.WriteLine(fmt.Sprintf("Error reading input: %v", err))
			continue
		}

		if input == "quit" {
			break
		}

		calc.CommandStream.AddToBack(input)
		calc.Run()
	}
}
