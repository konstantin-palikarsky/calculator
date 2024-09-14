package operations

import (
	"fmt"
	"math"
	"tuwien.ac.at/calculator/v2/src/types"
)

type DecimalConstructionMode struct {
	Calculator types.Calculator
}

func NewDecimalConstructionMode(calc types.Calculator) *DecimalConstructionMode {
	return &DecimalConstructionMode{
		Calculator: calc,
	}
}
func (d *DecimalConstructionMode) Execute(command rune) error {
	if d.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow in decimal construction mode")
	}

	top, err := d.Calculator.GetDataStack().Peek()
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
		multiplier := math.Pow(10, float64(d.Calculator.GetOperationMode()+1))
		newValue := floatValue + digit*multiplier
		d.Calculator.GetDataStack().Pop()
		d.Calculator.GetDataStack().Push(newValue)
		d.Calculator.SetOperationMode(d.Calculator.GetOperationMode() - 1)
	case command == '.':
		d.Calculator.GetDataStack().Push(0.0)
		d.Calculator.SetOperationMode(-2)
	default:
		d.Calculator.SetOperationMode(0)
		d.Calculator.GetCommandStream().AddToFront(string(command))
	}

	return nil
}
