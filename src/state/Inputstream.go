package state

import (
	"bufio"
	"os"
	"strings"
)

type InputStream struct {
	reader *bufio.Reader
}

func NewInputStream() *InputStream {
	return &InputStream{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (is *InputStream) ReadLine() (string, error) {
	line, err := is.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimRight(line, "\r\n"), nil
}
