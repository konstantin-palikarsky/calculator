package state

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func NewStack() *Stack {
	return &Stack{}
}

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

func (stack *Stack) Get(n int) (interface{}, error) {
	if n <= 0 || n > stack.Size() {
		return nil, fmt.Errorf("invalid index: %d", n)
	}
	index := len(*stack) - n
	return (*stack)[index], nil
}

func (stack *Stack) Remove(n int) error {
	if n <= 0 || n > stack.Size() {
		return fmt.Errorf("invalid index: %d", n)
	}
	index := len(*stack) - n
	*stack = append((*stack)[:index], (*stack)[index+1:]...)
	return nil
}
