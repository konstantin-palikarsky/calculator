package state

import "fmt"

type CommandStream struct {
	commands []rune
}

func NewCommandStream() *CommandStream {
	return &CommandStream{
		commands: make([]rune, 0),
	}
}

func (cs *CommandStream) AddToFront(str string) {
	for i := len(str) - 1; i >= 0; i-- {
		cs.commands = append([]rune{rune(str[i])}, cs.commands...)
	}
}

func (cs *CommandStream) AddToBack(str string) {
	cs.commands = append(cs.commands, []rune(str)...)
}

func (cs *CommandStream) Pop() (interface{}, error) {
	if len(cs.commands) == 0 {
		return nil, fmt.Errorf("command stream is empty")
	}
	command := cs.commands[0]
	cs.commands = cs.commands[1:]
	return command, nil
}

func (cs *CommandStream) IsEmpty() bool {
	return len(cs.commands) == 0
}

func (cs *CommandStream) PrintContents() string {
	return fmt.Sprintf("%v", cs.commands)
}
