package operations

import (
	"fmt"
	"tuwien.ac.at/calculator/v2/src/calculator"
)

type IntegerConstructionMode struct {
	Calculator *calculator.Calculator
}

func NewIntegerConstructionMode(calc *calculator.Calculator) *IntegerConstructionMode {
	return &IntegerConstructionMode{
		Calculator: calc,
	}
}

func (i *IntegerConstructionMode) Execute(command rune) error {
	if i.Calculator.DataStack.IsEmpty() {
		return fmt.Errorf("stack underflow in integer construction mode")
	}

	top, err := i.Calculator.DataStack.Peek()
	if err != nil {
		return err
	}

	intValue, ok := top.(int)
	if !ok {
		return fmt.Errorf("top of stack is not an integer in integer construction mode")
	}

	switch {
	case command >= '0' && command <= '9':
		newValue := intValue*10 + int(command-'0')
		i.Calculator.DataStack.Pop()
		i.Calculator.DataStack.Push(newValue)
	case command == '.':
		i.Calculator.DataStack.Pop()
		i.Calculator.DataStack.Push(float64(intValue))
		i.Calculator.OperationMode = -2
	default:
		i.Calculator.OperationMode = 0
	}

	return nil
}
