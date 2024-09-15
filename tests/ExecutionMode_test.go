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
		{"Modulo", "17 5 %", 2},
		{"Complex Expression", "15 2 3 4 + * -", 1},
		{"Division by Zero", "10 0 /", ""},
		{"Modulo by Zero", "10 0 %", ""},
		{"Float Addition", "3.5 2.7 +", 6.2},
		{"Mixed Type Addition", "5 3.5 +", 8.5},

		{"Equal Integers", "5 5 =", 1},
		{"Not Equal Integers", "5 6 =", 0},
		{"Less Than", "3 5 <", 1},
		{"Greater Than", "5 3 >", 1},
		{"Equal Floats", "3.14 3.14 =", 1},
		{"Not Equal Floats", "3.14 3.15 =", 0},
		{"Less Than Floats", "3.14 3.15 <", 1},
		{"Greater Than Floats", "3.15 3.14 >", 1},
		{"String Comparison", "(abc) (abd) <", 1},

		{"AND True", "1 1 &", 1},
		{"AND False", "1 0 &", 0},
		{"OR True", "1 0 |", 1},
		{"OR False", "0 0 |", 0},

		{"Null Check Empty String", "() _", 1},
		{"Null Check Non-Empty String", "(a) _", 0},
		{"Null Check Zero", "0 _", 1},
		{"Null Check Non-Zero", "1 _", 0},
		{"Negation Integer", "5 ~", -5},
		{"Negation Float", "3.14 ~", -3.14},
		{"Integer Conversion", "5.7 ?", 5},
		{"Copy", "1 2 3 2 !", 2},
		{"Delete", "1 2 3 2 $", 3},

		{"Push Multiple", "1 2 3 4 5", 5},
		{"Size", "1 2 3 #", 3},

		{"Single Element", "42", 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := state.NewCommandStream()
			ds := state.NewStack()
			is := state.NewInputStream()
			os := state.NewOutputStream()

			calc := calculator.NewCalculator(cs, ds, is, os)

			for _, char := range tt.input {
				calc.GetCommandStream().AddToBack(string(char))
			}

			calc.Run()

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
