package operations

import (
	"fmt"
	"math"
	"strconv"
	"tuwien.ac.at/calculator/v2/src/types"
)

type ExecutionMode struct {
	Calculator types.Calculator
}

func NewExecutionMode(calc types.Calculator) *ExecutionMode {
	return &ExecutionMode{
		Calculator: calc,
	}
}

// Execute processes a single command in execution mode
func (e *ExecutionMode) Execute(command rune) error {
	fmt.Printf("Executing command: %c\n", command)
	switch {
	case command >= '0' && command <= '9':
		// Push digit as integer and enter integer construction mode
		e.Calculator.GetDataStack().Push(int(command - '0'))
		e.Calculator.SetOperationMode(-1)
	case command == '.':
		// Push 0.0 and enter decimal place construction mode
		e.Calculator.GetDataStack().Push(0.0)
		e.Calculator.SetOperationMode(-2)
	case command == '(':
		// Push empty string and enter string construction mode
		e.Calculator.GetDataStack().Push("")
		e.Calculator.SetOperationMode(1)
	case command >= 'A' && command <= 'Z' || command >= 'a' && command <= 'z':
		// Push register content
		return e.pushRegisterContent(command)
	case command == '=', command == '<', command == '>':
		// Execute comparison operation
		return e.executeComparison(command)
	case command == '+', command == '-', command == '*', command == '/', command == '%':
		// Execute arithmetic operation
		return e.executeArithmetic(command)
	case command == '&', command == '|':
		// Execute logic operation
		return e.executeLogic(command)
	case command == '_':
		// Execute null check
		return e.executeNullCheck()
	case command == '~':
		// Execute negation
		return e.executeNegation()
	case command == '?':
		// Execute integer conversion
		return e.executeIntegerConversion()
	case command == '!':
		// Execute copy operation
		return e.executeCopy()
	case command == '$':
		// Execute delete operation
		return e.executeDelete()
	case command == '@':
		// Execute apply immediately
		return e.executeApplyImmediately()
	case command == '\\':
		// Execute apply later
		return e.executeApplyLater()
	case command == '#':
		// Push stack size
		e.Calculator.GetDataStack().Push(e.Calculator.GetDataStack().Size())
	case command == '\'':
		// Execute read input
		return e.executeReadInput()
	case command == '"':
		// Execute write output
		return e.executeWriteOutput()
	default:
		// Ignore other characters
		return nil
	}
	return nil
}

// pushRegisterContent pushes the content of a register onto the stack
func (e *ExecutionMode) pushRegisterContent(register rune) error {
	value, exists := e.Calculator.GetRegisters()[register]
	if !exists {
		return fmt.Errorf("register %c not found", register)
	}
	e.Calculator.GetDataStack().Push(value)
	return nil
}

// executeComparison performs a comparison operation on the top two stack items
func (e *ExecutionMode) executeComparison(op rune) error {
	// Check if there are at least two items on the stack
	if e.Calculator.GetDataStack().Size() < 2 {
		return fmt.Errorf("stack underflow: not enough operands for comparison")
	}

	// Pop the top two items from the stack
	b, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}
	a, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	// Perform the comparison
	result, err := e.compare(a, b, op)
	if err != nil {
		return err
	}

	// Push the result back onto the stack
	e.Calculator.GetDataStack().Push(result)
	return nil
}

// compare compares two values based on their types and the comparison operator
func (e *ExecutionMode) compare(a, b interface{}, op rune) (int, error) {
	epsilon := 1e-9

	// Type assertions
	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	// Convert integers to floats if necessary
	if aIsInt && bIsFloat {
		aFloat, aIsFloat = float64(aInt), true
	} else if bIsInt && aIsFloat {
		bFloat, bIsFloat = float64(bInt), true
	}

	// Compare floats
	if aIsFloat && bIsFloat {
		diff := aFloat - bFloat
		absA, absB := math.Abs(aFloat), math.Abs(bFloat)
		maxAbs := math.Max(absA, absB)
		if maxAbs > 1.0 {
			epsilon *= maxAbs
		}
		switch op {
		case '=':
			return boolToInt(math.Abs(diff) <= epsilon), nil
		case '<':
			return boolToInt(diff < -epsilon), nil
		case '>':
			return boolToInt(diff > epsilon), nil
		}
	}

	// Compare integers
	if aIsInt && bIsInt {
		switch op {
		case '=':
			return boolToInt(aInt == bInt), nil
		case '<':
			return boolToInt(aInt < bInt), nil
		case '>':
			return boolToInt(aInt > bInt), nil
		}
	}

	// Compare strings
	if aIsString && bIsString {
		switch op {
		case '=':
			return boolToInt(aString == bString), nil
		case '<':
			return boolToInt(aString < bString), nil
		case '>':
			return boolToInt(aString > bString), nil
		}
	}

	// Compare numbers with strings
	if (aIsInt || aIsFloat) && bIsString {
		return boolToInt(op == '<'), nil
	} else if aIsString && (bIsInt || bIsFloat) {
		return boolToInt(op == '>'), nil
	}

	return 0, fmt.Errorf("incomparable types")
}

// boolToInt converts a boolean value to an integer (1 for true, 0 for false)
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// executeArithmetic performs arithmetic operations on the top two stack items
func (e *ExecutionMode) executeArithmetic(op rune) error {
	fmt.Printf("Executing arithmetic operation: %c\n", op)

	// Check if there are at least two items on the stack
	if e.Calculator.GetDataStack().Size() < 2 {
		return fmt.Errorf("stack underflow: not enough operands for arithmetic operation")
	}

	// Pop the top two items from the stack
	b, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}
	fmt.Printf("Popped b: %v\n", b)

	a, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}
	fmt.Printf("Popped a: %v\n", a)

	// Perform the arithmetic operation
	result, err := e.performArithmetic(a, b, op)
	if err != nil {
		return err
	}

	// Push the result back onto the stack
	e.Calculator.GetDataStack().Push(result)
	fmt.Printf("Result of %v %c %v = %v\n", a, op, b, result)
	fmt.Printf("Pushed result to stack: %v\n", result)
	return nil
}

// performArithmetic executes the actual arithmetic operation based on operand types
func (e *ExecutionMode) performArithmetic(a, b interface{}, op rune) (interface{}, error) {
	// Type assertions for operands
	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	// Convert integers to floats if necessary
	if aIsInt && bIsFloat {
		aFloat, aIsFloat = float64(aInt), true
	} else if bIsInt && aIsFloat {
		bFloat, bIsFloat = float64(bInt), true
	}

	// Handle float operations
	if aIsFloat && bIsFloat {
		switch op {
		case '+':
			return aFloat + bFloat, nil
		case '-':
			return aFloat - bFloat, nil
		case '*':
			return aFloat * bFloat, nil
		case '/':
			if math.Abs(bFloat) < 1e-9 {
				return "", nil
			}
			return aFloat / bFloat, nil
		case '%':
			return "", nil
		}
	}

	// Handle integer operations
	if aIsInt && bIsInt {
		switch op {
		case '+':
			return aInt + bInt, nil
		case '-':
			return aInt - bInt, nil
		case '*':
			return aInt * bInt, nil
		case '/':
			if bInt == 0 {
				return "", nil
			}
			return aInt / bInt, nil
		case '%':
			if bInt == 0 {
				return "", nil
			}
			return aInt % bInt, nil
		}
	}

	// Handle string operations
	if aIsString && bIsString && op == '+' {
		return aString + bString, nil
	}
	if aIsString && bIsInt && op == '+' {
		return aString + string(rune(bInt)), nil
	}

	return "", fmt.Errorf("unsupported arithmetic operation")
}

// executeLogic performs logical operations on the top two stack items
func (e *ExecutionMode) executeLogic(op rune) error {
	// Check if there are at least two items on the stack
	if e.Calculator.GetDataStack().Size() < 2 {
		return fmt.Errorf("stack underflow: not enough operands for logic operation")
	}

	// Pop the top two items from the stack
	b, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}
	a, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	// Ensure both operands are integers
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if !aIsInt || !bIsInt {
		e.Calculator.GetDataStack().Push("")
		return nil
	}

	// Perform the logical operation
	var result int
	switch op {
	case '&':
		result = boolToInt(aInt != 0 && bInt != 0)
	case '|':
		result = boolToInt(aInt != 0 || bInt != 0)
	default:
		return fmt.Errorf("unknown logic operation: %c", op)
	}

	// Push the result back onto the stack
	e.Calculator.GetDataStack().Push(result)
	return nil
}

// executeNullCheck checks if the top stack value is "null-like" and pushes the result
func (e *ExecutionMode) executeNullCheck() error {
	// Check for stack underflow
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	// Pop the top value from the stack
	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	// Initialize result to 0 (false)
	result := 0

	// Check if the value is "null-like" based on its type
	switch v := value.(type) {
	case string:
		if v == "" {
			result = 1
		}
	case int:
		if v == 0 {
			result = 1
		}
	case float64:
		if math.Abs(v) < 1e-9 {
			result = 1
		}
	}

	// Push the result onto the stack
	e.Calculator.GetDataStack().Push(result)
	return nil
}

// executeNegation negates the top stack value or pushes an empty string
func (e *ExecutionMode) executeNegation() error {
	// Check for stack underflow
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	// Pop the top value from the stack
	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	// Negate the value based on its type
	switch v := value.(type) {
	case int:
		e.Calculator.GetDataStack().Push(-v)
	case float64:
		e.Calculator.GetDataStack().Push(-v)
	default:
		e.Calculator.GetDataStack().Push("")
	}

	return nil
}

// executeIntegerConversion converts the top stack value to an integer or pushes an empty string
func (e *ExecutionMode) executeIntegerConversion() error {
	// Check for stack underflow
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	// Pop the top value from the stack
	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	// Convert or replace the value based on its type
	switch v := value.(type) {
	case float64:
		e.Calculator.GetDataStack().Push(int(v))
	case int, string:
		e.Calculator.GetDataStack().Push("")
	default:
		return fmt.Errorf("unexpected type for integer conversion")
	}

	return nil
}

// executeCopy copies the nth item from the stack to the top
func (e *ExecutionMode) executeCopy() error {
	// Check for stack underflow
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	// Pop the top value (n) from the stack
	n, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	// Ensure n is a valid integer within the stack range
	nInt, ok := n.(int)
	if !ok || nInt <= 0 || nInt > e.Calculator.GetDataStack().Size() {
		// If n is not a valid integer or out of range, do nothing
		return nil
	}

	// Get the nth item from the stack
	value, err := e.Calculator.GetDataStack().Get(nInt)
	if err != nil {
		return err
	}

	// Push the copied value onto the stack
	e.Calculator.GetDataStack().Push(value)
	return nil
}

/*func (e *ExecutionMode) executeCopy() error {
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}
	n, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}
	nInt, ok := n.(int)
	if !ok || nInt <= 0 || nInt > e.Calculator.GetDataStack().Size() {
		e.Calculator.GetDataStack().Push(n)
		return nil
	}
	value, err := e.Calculator.GetDataStack().Get(e.Calculator.GetDataStack().Size() - nInt + 1)
	if err != nil {
		e.Calculator.GetDataStack().Push(n)
		return err
	}
	e.Calculator.GetDataStack().Push(n)
	e.Calculator.GetDataStack().Push(value)
	return nil
}*/

// executeDelete removes the nth item from the stack
func (e *ExecutionMode) executeDelete() error {
	// Check for stack underflow
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	// Pop the top value (n) from the stack
	n, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	// Ensure n is a valid integer within the stack range
	nInt, ok := n.(int)
	if !ok || nInt <= 0 || nInt > e.Calculator.GetDataStack().Size() {
		// If n is invalid, push it back and do nothing
		e.Calculator.GetDataStack().Push(n)
		return nil
	}

	// Remove the nth item from the stack
	return e.Calculator.GetDataStack().Remove(nInt)
}

// executeApplyImmediately executes the string at the top of the stack immediately
func (e *ExecutionMode) executeApplyImmediately() error {
	// Check for stack underflow
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	// Pop the top value from the stack
	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	// Ensure the value is a string
	str, ok := value.(string)
	if !ok {
		return nil // Do nothing if not a string
	}

	// Add the string to the front of the command stream for immediate execution
	e.Calculator.GetCommandStream().AddToFront(str)
	return nil
}

// executeApplyLater adds the string at the top of the stack to the end of the command stream
func (e *ExecutionMode) executeApplyLater() error {
	// Check for stack underflow
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	// Pop the top value from the stack
	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	// Ensure the value is a string
	str, ok := value.(string)
	if !ok {
		return nil // Do nothing if not a string
	}

	// Add the string to the back of the command stream for later execution
	e.Calculator.GetCommandStream().AddToBack(str)
	return nil
}

// executeReadInput reads a line from the input stream and pushes it onto the data stack
func (e *ExecutionMode) executeReadInput() error {
	// Read a line from the input stream
	input, err := e.Calculator.GetInputStream().ReadLine()
	if err != nil {
		return err
	}

	// Try to parse the input as an integer
	if intValue, err := strconv.Atoi(input); err == nil {
		e.Calculator.GetDataStack().Push(intValue)
		return nil
	}

	// If not an integer, try to parse as a float
	if floatValue, err := strconv.ParseFloat(input, 64); err == nil {
		e.Calculator.GetDataStack().Push(floatValue)
		return nil
	}

	// If neither integer nor float, treat as a string
	e.Calculator.GetDataStack().Push(input)
	return nil
}

// executeWriteOutput pops a value from the data stack and writes it to the output stream
func (e *ExecutionMode) executeWriteOutput() error {
	// Check if the stack is empty
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	// Pop a value from the stack
	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	// Convert the value to a string based on its type
	var output string
	switch v := value.(type) {
	case string:
		output = v
	case int:
		output = strconv.Itoa(v)
	case float64:
		output = strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Errorf("unexpected type for output")
	}

	// Write the output to the output stream
	e.Calculator.GetOutputStream().WriteLine(output)
	return nil
}
