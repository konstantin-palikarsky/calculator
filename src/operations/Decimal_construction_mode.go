package operations

import (
	"fmt"
	"math"
	"tuwien.ac.at/calculator/v2/src/types"
)

// DecimalConstructionMode handles the construction of decimal places for floating-point numbers
type DecimalConstructionMode struct {
	Calculator types.Calculator
}

// NewDecimalConstructionMode creates a new instance of DecimalConstructionMode
func NewDecimalConstructionMode(calc types.Calculator) *DecimalConstructionMode {
	return &DecimalConstructionMode{
		Calculator: calc,
	}
}

// Execute processes a single command in decimal construction mode
func (d *DecimalConstructionMode) Execute(command rune) error {
	fmt.Printf("ExecutingDec command: %c\n", command)

	// Check for stack underflow
	if d.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow in decimal construction mode")
	}

	// Get the top value from the stack
	top, err := d.Calculator.GetDataStack().Peek()
	if err != nil {
		return err
	}

	// Ensure the top value is a float
	floatValue, ok := top.(float64)
	if !ok {
		return fmt.Errorf("top of stack is not a float in decimal construction mode")
	}

	switch {
	case command >= '0' && command <= '9':
		// Add the digit to the current float value
		digit := float64(command - '0')
		multiplier := math.Pow(10, float64(d.Calculator.GetOperationMode()+1))
		newValue := floatValue + digit*multiplier
		d.Calculator.GetDataStack().Pop()
		d.Calculator.GetDataStack().Push(newValue)
		d.Calculator.SetOperationMode(d.Calculator.GetOperationMode() - 1)
	case command == '.':
		// Start a new float number
		d.Calculator.GetDataStack().Push(0.0)
		d.Calculator.SetOperationMode(-2)
	default:
		// Exit decimal construction mode
		d.Calculator.SetOperationMode(0)
		d.Calculator.GetCommandStream().AddToFront(string(command))
	}

	return nil
}
