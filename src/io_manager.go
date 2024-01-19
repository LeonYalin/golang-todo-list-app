package src

import (
	"fmt"
	"io"
	"os"
)

type IOManageable interface {
	Scan() string
	Write(msg string)
}

type IOManager struct {
	name string
}

func (iom *IOManager) Scan() (text string) {
	fmt.Scanln(text)
	return text
}

func (iom *IOManager) Write(text string) {
	io.WriteString(os.Stdout, text)
}

func NewIoManager() *IOManager {
	return &IOManager{name: ""}
}
