package calculator

import (
	"fmt"
	"tuwien.ac.at/calculator/v2/src/operations"
	"tuwien.ac.at/calculator/v2/src/state"
)

type Calculator struct {
	CommandStream           *state.CommandStream
	DataStack               *state.Stack
	Registers               map[rune]interface{}
	OperationMode           int64
	InputStream             *state.InputStream
	OutputStream            *state.OutputStream
	ExecutionMode           *operations.ExecutionMode
	IntegerConstructionMode *operations.IntegerConstructionMode
	DecimalConstructionMode *operations.DecimalConstructionMode
	StringConstructionMode  *operations.StringConstructionMode
}

func NewCalculator(cs *state.CommandStream, ds *state.Stack, is *state.InputStream, os *state.OutputStream) *Calculator {
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

func main() {
	cs := state.NewCommandStream()
	ds := state.NewStack()
	is := state.NewInputStream()
	os := state.NewOutputStream()

	calc := NewCalculator(cs, ds, is, os)

	initialCommand, ok := calc.Registers['a'].(string)
	if !ok {
		calc.OutputStream.WriteLine("Error: Invalid content in register 'a'")
		return
	}
	calc.CommandStream.AddToBack(initialCommand)

	calc.Run()

	for {
		calc.OutputStream.WriteLine("Enter a command (or 'quit' to exit):")
		input, err := calc.InputStream.ReadLine()
		if err != nil {
			calc.OutputStream.WriteLine(fmt.Sprintf("Error reading input: %v", err))
			continue
		}

		if input == "quit" {
			break
		}

		calc.CommandStream.AddToBack(input)
		calc.Run()
	}
}
