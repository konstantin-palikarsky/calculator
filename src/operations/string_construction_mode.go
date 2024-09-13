package operations

import (
	"fmt"
	"tuwien.ac.at/calculator/v2/src/calculator"
)

type StringConstructionMode struct {
	Calculator *calculator.Calculator
}

func NewStringConstructionMode(calc *calculator.Calculator) *StringConstructionMode {
	return &StringConstructionMode{
		Calculator: calc,
	}
}

func (s *StringConstructionMode) Execute(command rune) error {
	if s.Calculator.DataStack.IsEmpty() {
		return fmt.Errorf("stack underflow in string construction mode")
	}

	top, err := s.Calculator.DataStack.Peek()
	if err != nil {
		return err
	}

	str, ok := top.(string)
	if !ok {
		return fmt.Errorf("top of stack is not a string in string construction mode")
	}

	switch command {
	case '(':
		newStr := str + string(command)
		s.Calculator.DataStack.Pop()
		s.Calculator.DataStack.Push(newStr)
		s.Calculator.OperationMode++
	case ')':
		if s.Calculator.OperationMode > 1 {
			newStr := str + string(command)
			s.Calculator.DataStack.Pop()
			s.Calculator.DataStack.Push(newStr)
		}
		s.Calculator.OperationMode--
		if s.Calculator.OperationMode == 0 {
			return nil
		}
	default:
		newStr := str + string(command)
		s.Calculator.DataStack.Pop()
		s.Calculator.DataStack.Push(newStr)
	}

	return nil
}
