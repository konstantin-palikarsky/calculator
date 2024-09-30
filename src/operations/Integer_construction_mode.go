package operations

import (
	"fmt"
	"tuwien.ac.at/calculator/v2/src/types"
)

// IntegerConstructionMode handles the construction of integers from digits
type IntegerConstructionMode struct {
	Calculator types.Calculator
}

// NewIntegerConstructionMode creates a new instance of IntegerConstructionMode
func NewIntegerConstructionMode(calc types.Calculator) *IntegerConstructionMode {
	return &IntegerConstructionMode{
		Calculator: calc,
	}
}

// Execute processes a single command in integer construction mode
func (i *IntegerConstructionMode) Execute(command rune) error {
	fmt.Printf("ExecutingInt command: %c\n", command)

	// Check for stack underflow
	if i.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow in integer construction mode")
	}

	// Get the top value from the stack
	top, err := i.Calculator.GetDataStack().Peek()
	if err != nil {
		return err
	}

	// Ensure the top value is an integer
	intValue, ok := top.(int)
	if !ok {
		return fmt.Errorf("top of stack is not an integer in integer construction mode")
	}

	switch {
	case command >= '0' && command <= '9':
		// Add the digit to the current integer value
		newValue := intValue*10 + int(command-'0')
		i.Calculator.GetDataStack().Pop()
		i.Calculator.GetDataStack().Push(newValue)
	case command == '.':
		// Convert to float and switch to decimal construction mode
		i.Calculator.GetDataStack().Pop()
		i.Calculator.GetDataStack().Push(float64(intValue))
		i.Calculator.SetOperationMode(-2)
	default:
		// Exit integer construction mode
		i.Calculator.SetOperationMode(0)
		i.Calculator.GetCommandStream().AddToFront(string(command))
	}

	return nil
}
