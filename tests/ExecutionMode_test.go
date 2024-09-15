package operations

import (
	"testing"
	"tuwien.ac.at/calculator/v2/src/calculator"
	"tuwien.ac.at/calculator/v2/src/state"
)

func TestExecutionMode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{"Addition", "5 3 +", 8},
		{"Subtraction", "10 4 -", 6},
		{"Multiplication", "3 4 *", 12},
		{"Division", "20 5 /", 4},
		{"Complex Expression", "15 2 3 4 + * -", 1},
		{"Division by Zero", "10 0 /", ""},
		{"Comparison Equal", "5 5 =", 1},
		{"Comparison Less Than", "3 5 <", 1},
		{"Comparison Greater Than", "5 3 >", 1},
		{"Logic AND", "1 1 &", 1},
		{"Logic OR", "1 0 |", 1},
		{"Null Check", "0 _", 1},
		{"Negation", "5 ~", -5},
		{"Integer Conversion", "5.5 ?", 5},
		{"Copy", "1 2 3 2 !", 2},
		{"Delete", "1 2 3 2 $", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := state.NewCommandStream()
			ds := state.NewStack()
			is := state.NewInputStream()
			os := state.NewOutputStream()

			calc := calculator.NewCalculator(cs, ds, is, os)

			// Add commands to the command stream
			for _, char := range tt.input {
				calc.GetCommandStream().AddToBack(string(char))
			}

			// Run the calculator
			calc.Run()

			// Check the result
			if !calc.GetDataStack().IsEmpty() {
				result, err := calc.GetDataStack().Pop()
				if err != nil {
					t.Errorf("Error popping result: %v", err)
					return
				}

				if result != tt.expected {
					t.Errorf("Expected %v, but got %v", tt.expected, result)
					t.Errorf("Stack contents: %s", calc.GetDataStack().String())
				}
			} else if tt.expected != "" {
				t.Errorf("Expected %v, but stack is empty", tt.expected)
			}
		})
	}
}
