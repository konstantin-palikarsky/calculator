package operations

import (
	"fmt"
	"tuwien.ac.at/calculator/v2/src/types"
)

// StringConstructionMode handles the construction of strings from characters
type StringConstructionMode struct {
	Calculator types.Calculator
}

// NewStringConstructionMode creates a new instance of StringConstructionMode
func NewStringConstructionMode(calc types.Calculator) *StringConstructionMode {
	return &StringConstructionMode{
		Calculator: calc,
	}
}

// Execute processes a single command in string construction mode
func (s *StringConstructionMode) Execute(command rune) error {
	fmt.Printf("ExecutingString command: %c\n", command)

	// Check for stack underflow
	if s.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow in string construction mode")
	}

	// Get the top value from the stack
	top, err := s.Calculator.GetDataStack().Peek()
	if err != nil {
		return err
	}

	// Ensure the top value is a string
	str, ok := top.(string)
	if !ok {
		return fmt.Errorf("top of stack is not a string in string construction mode")
	}

	switch command {
	case '(':
		// Add '(' to the string and increment operation mode
		newStr := str + string(command)
		s.Calculator.GetDataStack().Pop()
		s.Calculator.GetDataStack().Push(newStr)
		s.Calculator.SetOperationMode(s.Calculator.GetOperationMode() + 1)
	case ')':
		// Add ')' if operation mode > 1, then decrement operation mode
		if s.Calculator.GetOperationMode() > 1 {
			newStr := str + string(command)
			s.Calculator.GetDataStack().Pop()
			s.Calculator.GetDataStack().Push(newStr)
		}
		s.Calculator.SetOperationMode(s.Calculator.GetOperationMode() - 1)
		if s.Calculator.GetOperationMode() == 0 {
			return nil
		}
	default:
		// Add the character to the string
		newStr := str + string(command)
		s.Calculator.GetDataStack().Pop()
		s.Calculator.GetDataStack().Push(newStr)
	}

	return nil
}
