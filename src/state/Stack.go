package state

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []interface{}

func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack *Stack) PushFromInputStream() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Calculator input: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")

	stack.Push(determineType(text))
}

func determineType(token string) interface{} {
	if val, err := strconv.Atoi(token); err == nil {
		return val
	}
	if val, err := strconv.ParseFloat(token, 64); err == nil {
		return val
	}
	return token
}

func (stack *Stack) PopToOutputStream() {
	if val, err := stack.Pop(); err == nil {
		fmt.Print(val)
	}
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.isEmpty() {
		return nil, errors.New("popping empty stack")
	}

	index := len(*stack) - 1
	element := (*stack)[index]
	*stack = (*stack)[:index]
	return element, nil
}

func (stack *Stack) isEmpty() bool {
	return len(*stack) == 0
}
