package main

import (
	"tuwien.ac.at/calculator/v2/src/calculator"
	"tuwien.ac.at/calculator/v2/src/state"
)

func main() {
	// Initialize the calculator components
	cs := state.NewCommandStream()
	ds := state.NewStack()
	is := state.NewInputStream()
	os := state.NewOutputStream()

	// Create a new calculator instance
	calc := calculator.NewCalculator(cs, ds, is, os)

	// Retrieve and set the initial command from register 'a'
	initialCommand, ok := calc.GetRegisters()['a'].(string)
	if !ok {
		calc.GetOutputStream().WriteLine("Error: Invalid content in register 'a'")
		return
	}
	initialCommand = ""
	calc.GetCommandStream().AddToBack(initialCommand)
	calc.GetCommandStream().AddToBack("(\n  # 2 >\n  (\n    +\n    3!@\n  )\n  (2$)\n  ?\n)@\n")
	// Execute the initial command
	calc.Run()

	// Main input loop
	/*for {
		// Prompt for user input
		calc.GetOutputStream().WriteLine("Enter a command (or 'quit' to exit):")
		input, err := calc.GetInputStream().ReadLine()
		if err != nil {
			calc.GetOutputStream().WriteLine(fmt.Sprintf("Error reading input: %v", err))
			continue
		}

		// Check for quit command
		if input == "quit" {
			break
		}

		// Add user input to command stream
		calc.GetCommandStream().AddToBack(input)

		// Execute the command
		calc.Run()
	}*/
}
