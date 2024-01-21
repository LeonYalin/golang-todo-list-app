package src

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/leonyalin/todo-list/src/util"
)

type IOManageable interface {
	Scan() string
	Write(msg string)
	Question(msg string, testAns func(answer string) bool) string
}

type IOManager struct {
	name string
}

func (iom *IOManager) Scan() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading user input:", err)
	}
	line = strings.Trim(line, "\n")
	return line
}

func (iom *IOManager) Question(msg string, testAns func(ans string) bool) string {
	var answer string
	iom.Write(msg)
	answer = iom.Scan()
	for {
		if testAns(answer) {
			break
		}
		iom.Write(util.INVALID_OPTION_TEXT)
		iom.Write(msg)
		answer = iom.Scan()
	}
	return answer
}

func (iom *IOManager) Write(text string) {
	io.WriteString(os.Stdout, text)
}

func NewIoManager() *IOManager {
	return &IOManager{name: ""}
}
