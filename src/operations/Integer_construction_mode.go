package operations

import (
	"fmt"
	"tuwien.ac.at/calculator/v2/src/types"
)

type IntegerConstructionMode struct {
	Calculator types.Calculator
}

func NewIntegerConstructionMode(calc types.Calculator) *IntegerConstructionMode {
	return &IntegerConstructionMode{
		Calculator: calc,
	}
}

func (i *IntegerConstructionMode) Execute(command rune) error {
	fmt.Printf("ExecutingInt command: %c\n", command)
	if i.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow in integer construction mode")
	}

	top, err := i.Calculator.GetDataStack().Peek()
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
		i.Calculator.GetDataStack().Pop()
		i.Calculator.GetDataStack().Push(newValue)
	case command == '.':
		i.Calculator.GetDataStack().Pop()
		i.Calculator.GetDataStack().Push(float64(intValue))
		i.Calculator.SetOperationMode(-2)
	default:
		i.Calculator.SetOperationMode(0)
		i.Calculator.GetCommandStream().AddToFront(string(command))
	}

	return nil
}
