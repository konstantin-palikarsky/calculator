package state

import (
	"errors"
)

type Stack []interface{}

func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("popping empty stack")
	}

	index := len(*stack) - 1
	element := (*stack)[index]
	*stack = (*stack)[:index]
	return element, nil
}

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Size() int {
	return len(*stack)
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("peeking empty stack")
	}
	return (*stack)[len(*stack)-1], nil
}
