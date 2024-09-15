package calculator

import (
	"fmt"
	"tuwien.ac.at/calculator/v2/src/operations"
	"tuwien.ac.at/calculator/v2/src/types"
)

type Calculator struct {
	CommandStream           types.CommandStream
	DataStack               types.Stack
	Registers               map[rune]interface{}
	OperationMode           int64
	InputStream             types.InputStream
	OutputStream            types.OutputStream
	ExecutionMode           types.ExecutionMode
	IntegerConstructionMode types.ConstructionMode
	DecimalConstructionMode types.ConstructionMode
	StringConstructionMode  types.ConstructionMode
}

func NewCalculator(cs types.CommandStream, ds types.Stack, is types.InputStream, os types.OutputStream) *Calculator {
	calc := &Calculator{
		CommandStream: cs,
		DataStack:     ds,
		Registers:     make(map[rune]interface{}),
		OperationMode: 0,
		InputStream:   is,
		OutputStream:  os,
	}
	calc.ExecutionMode = operations.NewExecutionMode(calc)
	calc.IntegerConstructionMode = operations.NewIntegerConstructionMode(calc)
	calc.DecimalConstructionMode = operations.NewDecimalConstructionMode(calc)
	calc.StringConstructionMode = operations.NewStringConstructionMode(calc)
	calc.InitializeRegisters()
	return calc
}

func (c *Calculator) InitializeRegisters() {
	for ch := 'A'; ch <= 'Z'; ch++ {
		c.Registers[ch] = ""
	}
	for ch := 'a'; ch <= 'z'; ch++ {
		c.Registers[ch] = ""
	}
	c.Registers['a'] = "Welcome to the Postfix Calculator!"
	c.GetCommandStream().AddToBack("a\"")
}

func (c *Calculator) Run() {
	for !c.CommandStream.IsEmpty() {
		command, err := c.CommandStream.Pop()
		if err != nil {
			c.OutputStream.WriteLine(fmt.Sprintf("Error: %v", err))
			break
		}

		err = c.ExecuteCommand(command.(rune))
		if err != nil {
			c.OutputStream.WriteLine(fmt.Sprintf("Error: %v", err))
			break
		}
	}
}

func (c *Calculator) ExecuteCommand(command rune) error {
	switch {
	case c.OperationMode == 0:
		return c.ExecutionMode.Execute(command)
	case c.OperationMode == -1:
		return c.IntegerConstructionMode.Execute(command)
	case c.OperationMode < -1:
		return c.DecimalConstructionMode.Execute(command)
	case c.OperationMode > 0:
		return c.StringConstructionMode.Execute(command)
	default:
		return fmt.Errorf("invalid operation mode: %d", c.OperationMode)
	}
}

func (c *Calculator) GetDataStack() types.Stack {
	return c.DataStack
}

func (c *Calculator) GetRegisters() map[rune]interface{} {
	return c.Registers
}

func (c *Calculator) GetOperationMode() int64 {
	return c.OperationMode
}

func (c *Calculator) SetOperationMode(mode int64) {
	c.OperationMode = mode
}

func (c *Calculator) GetInputStream() types.InputStream {
	return c.InputStream
}

func (c *Calculator) GetOutputStream() types.OutputStream {
	return c.OutputStream
}

func (c *Calculator) GetCommandStream() types.CommandStream {
	return c.CommandStream
}
