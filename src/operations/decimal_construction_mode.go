package operations

import (
	"fmt"
	"math"
	"tuwien.ac.at/calculator/v2/src/calculator"
)

type DecimalConstructionMode struct {
	Calculator *calculator.Calculator
}

func NewDecimalConstructionMode(calc *calculator.Calculator) *DecimalConstructionMode {
	return &DecimalConstructionMode{
		Calculator: calc,
	}
}

func (d *DecimalConstructionMode) Execute(command rune) error {
	if d.Calculator.DataStack.IsEmpty() {
		return fmt.Errorf("stack underflow in decimal construction mode")
	}

	top, err := d.Calculator.DataStack.Peek()
	if err != nil {
		return err
	}

	floatValue, ok := top.(float64)
	if !ok {
		return fmt.Errorf("top of stack is not a float in decimal construction mode")
	}

	switch {
	case command >= '0' && command <= '9':
		digit := float64(command - '0')
		multiplier := math.Pow(10, float64(d.Calculator.OperationMode+1))
		newValue := floatValue + digit*multiplier
		d.Calculator.DataStack.Pop()
		d.Calculator.DataStack.Push(newValue)
		d.Calculator.OperationMode--
	case command == '.':
		d.Calculator.DataStack.Push(0.0)
		d.Calculator.OperationMode = -2
	default:
		d.Calculator.OperationMode = 0
		d.Calculator.CommandStream.AddToFront(string(command))
	}

	return nil
}
