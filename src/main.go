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

	initialCommand, ok := calc.GetRegisters()['a'].(string)
	if !ok {
		calc.GetOutputStream().WriteLine("Error: Invalid content in register 'a'")
		return
	}
	calc.GetCommandStream().AddToBack(initialCommand)

	calc.Run()

	for {
		calc.GetOutputStream().WriteLine("Enter a command (or 'quit' to exit):")
		input, err := calc.GetInputStream().ReadLine()
		if err != nil {
			calc.GetOutputStream().WriteLine(fmt.Sprintf("Error reading input: %v", err))
			continue
		}

		if input == "quit" {
			break
		}

		calc.GetCommandStream().AddToBack(input)
		fmt.Printf("Command stream after input: %s\n")
		calc.GetCommandStream().PrintContents()
		calc.Run()
	}
}
