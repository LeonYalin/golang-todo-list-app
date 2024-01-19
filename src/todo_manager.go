package src

import (
	"fmt"
	"github.com/leonyalin/todo-list/src/util"
	"strings"
)

type TodoManager struct {
	io    IOManageable
	todos []*Todo
}

func (tm *TodoManager) AddTodo(description string) {
	tm.todos = append(tm.todos, NewTodo(description))
}

func (tm *TodoManager) Print() string {
	var sb strings.Builder
	for _, t := range tm.todos {
		sb.WriteString(t.Print() + "\n")
	}
	return sb.String()
}

func (tm *TodoManager) Init() {
	tm.greetingDialog()
}

func (tm *TodoManager) greetingDialog() {
	tm.io.Write(fmt.Sprintf("\n%s and welcome to my %s app!\n\n", util.YellowText("Hello"), util.RedText("TODO List")))
}

func (tm *TodoManager) mainMenuDialog() {
	// TODO
}

func NewTodoManager(io IOManageable) *TodoManager {
	return &TodoManager{io: io, todos: []*Todo{}}
}
