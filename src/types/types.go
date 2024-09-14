package types

type Calculator interface {
	GetDataStack() Stack
	GetRegisters() map[rune]interface{}
	GetOperationMode() int64
	SetOperationMode(mode int64)
	GetInputStream() InputStream
	GetOutputStream() OutputStream
	GetCommandStream() CommandStream
}

type Stack interface {
	Push(value interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	Size() int
	Get(n int) (interface{}, error)
	Remove(n int) error
}

type InputStream interface {
	ReadLine() (string, error)
}

type OutputStream interface {
	WriteLine(value interface{})
}

type CommandStream interface {
	AddToFront(str string)
	AddToBack(str string)
	Pop() (interface{}, error)
	IsEmpty() bool
	PrintContents() string
}

type ExecutionMode interface {
	Execute(command rune) error
}

type ConstructionMode interface {
	Execute(command rune) error
}
