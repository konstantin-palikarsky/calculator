package operations

import (
	"fmt"
	"tuwien.ac.at/calculator/v2/src/types"
)

type StringConstructionMode struct {
	Calculator types.Calculator
}

func NewStringConstructionMode(calc types.Calculator) *StringConstructionMode {
	return &StringConstructionMode{
		Calculator: calc,
	}
}

func (s *StringConstructionMode) Execute(command rune) error {
	fmt.Printf("ExecutingString command: %c\n", command)
	if s.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow in string construction mode")
	}

	top, err := s.Calculator.GetDataStack().Peek()
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
		s.Calculator.GetDataStack().Pop()
		s.Calculator.GetDataStack().Push(newStr)
		s.Calculator.SetOperationMode(s.Calculator.GetOperationMode() + 1)
	case ')':
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
		newStr := str + string(command)
		s.Calculator.GetDataStack().Pop()
		s.Calculator.GetDataStack().Push(newStr)
	}

	return nil
}
