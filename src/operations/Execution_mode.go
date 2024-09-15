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

func (e *ExecutionMode) Execute(command rune) error {
	switch {
	case command >= '0' && command <= '9':
		e.Calculator.GetDataStack().Push(int(command - '0'))
		e.Calculator.SetOperationMode(-1)
	case command == '.':
		e.Calculator.GetDataStack().Push(0.0)
		e.Calculator.SetOperationMode(-2)
	case command == '(':
		e.Calculator.GetDataStack().Push("")
		e.Calculator.SetOperationMode(1)
	case command >= 'A' && command <= 'Z' || command >= 'a' && command <= 'z':
		return e.pushRegisterContent(command)
	case command == '=', command == '<', command == '>':
		return e.executeComparison(command)
	case command == '+', command == '-', command == '*', command == '/', command == '%':
		return e.executeArithmetic(command)
	case command == '&', command == '|':
		return e.executeLogic(command)
	case command == '_':
		return e.executeNullCheck()
	case command == '~':
		return e.executeNegation()
	case command == '?':
		return e.executeIntegerConversion()
	case command == '!':
		return e.executeCopy()
	case command == '$':
		return e.executeDelete()
	case command == '@':
		return e.executeApplyImmediately()
	case command == '\\':
		return e.executeApplyLater()
	case command == '#':
		e.Calculator.GetDataStack().Push(e.Calculator.GetDataStack().Size())
	case command == '\'':
		return e.executeReadInput()
	case command == '"':
		return e.executeWriteOutput()
	default:
		return nil
	}
	return nil
}

func (e *ExecutionMode) pushRegisterContent(register rune) error {
	value, exists := e.Calculator.GetRegisters()[register]
	if !exists {
		return fmt.Errorf("register %c not found", register)
	}
	e.Calculator.GetDataStack().Push(value)
	return nil
}

func (e *ExecutionMode) executeComparison(op rune) error {
	if e.Calculator.GetDataStack().Size() < 2 {
		return fmt.Errorf("stack underflow: not enough operands for comparison")
	}

	b, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}
	a, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	result, err := e.compare(a, b, op)
	if err != nil {
		return err
	}

	e.Calculator.GetDataStack().Push(result)
	return nil
}

func (e *ExecutionMode) compare(a, b interface{}, op rune) (int, error) {
	epsilon := 1e-9

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsInt && bIsFloat {
		aFloat, aIsFloat = float64(aInt), true
	} else if bIsInt && aIsFloat {
		bFloat, bIsFloat = float64(bInt), true
	}

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

	if (aIsInt || aIsFloat) && bIsString {
		return boolToInt(op == '<'), nil
	} else if aIsString && (bIsInt || bIsFloat) {
		return boolToInt(op == '>'), nil
	}

	return 0, fmt.Errorf("incomparable types")
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (e *ExecutionMode) executeArithmetic(op rune) error {
	fmt.Printf("Executing arithmetic operation: %c\n", op)

	if e.Calculator.GetDataStack().Size() < 2 {
		return fmt.Errorf("stack underflow: not enough operands for arithmetic operation")
	}

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

	result, err := e.performArithmetic(a, b, op)
	if err != nil {
		return err
	}

	e.Calculator.GetDataStack().Push(result)
	fmt.Printf("Result of %v %c %v = %v\n", a, op, b, result)
	fmt.Printf("Pushed result to stack: %v\n", result)
	return nil
}

func (e *ExecutionMode) performArithmetic(a, b interface{}, op rune) (interface{}, error) {
	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsInt && bIsFloat {
		aFloat, aIsFloat = float64(aInt), true
	} else if bIsInt && aIsFloat {
		bFloat, bIsFloat = float64(bInt), true
	}

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

	if aIsString && bIsString && op == '+' {
		return aString + bString, nil
	}
	if aIsString && bIsInt && op == '+' {
		return aString + string(rune(bInt)), nil
	}

	return "", fmt.Errorf("unsupported arithmetic operation")
}

func (e *ExecutionMode) executeLogic(op rune) error {
	if e.Calculator.GetDataStack().Size() < 2 {
		return fmt.Errorf("stack underflow: not enough operands for logic operation")
	}

	b, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}
	a, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if !aIsInt || !bIsInt {
		e.Calculator.GetDataStack().Push("")
		return nil
	}

	var result int
	switch op {
	case '&':
		result = boolToInt(aInt != 0 && bInt != 0)
	case '|':
		result = boolToInt(aInt != 0 || bInt != 0)
	default:
		return fmt.Errorf("unknown logic operation: %c", op)
	}

	e.Calculator.GetDataStack().Push(result)
	return nil
}

func (e *ExecutionMode) executeNullCheck() error {
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	result := 0
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

	e.Calculator.GetDataStack().Push(result)
	return nil
}

func (e *ExecutionMode) executeNegation() error {
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

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

func (e *ExecutionMode) executeIntegerConversion() error {
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

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

func (e *ExecutionMode) executeCopy() error {
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	n, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	nInt, ok := n.(int)
	if !ok || nInt <= 0 || nInt > e.Calculator.GetDataStack().Size() {
		return nil
	}

	value, err := e.Calculator.GetDataStack().Get(nInt - 1)
	if err != nil {
		return err
	}

	e.Calculator.GetDataStack().Push(value)
	return nil
}

func (e *ExecutionMode) executeDelete() error {
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	n, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	nInt, ok := n.(int)
	if !ok || nInt <= 0 || nInt > e.Calculator.GetDataStack().Size() {
		return nil
	}

	return e.Calculator.GetDataStack().Remove(nInt - 1)
}

func (e *ExecutionMode) executeApplyImmediately() error {
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	str, ok := value.(string)
	if !ok {
		return nil
	}

	e.Calculator.GetCommandStream().AddToFront(str)
	return nil
}

func (e *ExecutionMode) executeApplyLater() error {
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

	str, ok := value.(string)
	if !ok {
		return nil
	}

	e.Calculator.GetCommandStream().AddToBack(str)
	return nil
}

func (e *ExecutionMode) executeReadInput() error {
	input, err := e.Calculator.GetInputStream().ReadLine()
	if err != nil {
		return err
	}

	if intValue, err := strconv.Atoi(input); err == nil {
		e.Calculator.GetDataStack().Push(intValue)
		return nil
	}

	if floatValue, err := strconv.ParseFloat(input, 64); err == nil {
		e.Calculator.GetDataStack().Push(floatValue)
		return nil
	}

	e.Calculator.GetDataStack().Push(input)
	return nil
}

func (e *ExecutionMode) executeWriteOutput() error {
	if e.Calculator.GetDataStack().IsEmpty() {
		return fmt.Errorf("stack underflow")
	}

	value, err := e.Calculator.GetDataStack().Pop()
	if err != nil {
		return err
	}

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
	e.Calculator.GetOutputStream().WriteLine(output)
	return nil
}
