package main

import (
	// "fmt"
	"github.com/leonyalin/todo-list/src"
)

func main() {
	io := src.NewIoManager()
	tm := src.NewTodoManager(io)
	tm.Init()
	// tm.AddTodo("This is a first todo")
	// tm.AddTodo("This is a second todo")
	// fmt.Println(tm.Print())
}
