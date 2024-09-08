package main

import (
	"bufio"
	"os"
)

type InputStream struct {
	reader *bufio.Reader
}

func InitializeInputStream() *InputStream {
	return &InputStream{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (is *InputStream) ReadLine() (string, error) {
	return is.reader.ReadString('\n')
}
