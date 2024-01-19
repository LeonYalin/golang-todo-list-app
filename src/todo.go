package src

import (
	"fmt"

	"github.com/google/uuid"
)

type Todo struct {
	ID          string
	Description string
	Completed   bool
}

func (t *Todo) Print() string {
	return fmt.Sprintf("{ID: %s, Description: %s, Completed: %t}", t.ID, t.Description, t.Completed)
}

func NewTodo(desc string) *Todo {
	return &Todo{ID: uuid.New().String(), Description: desc, Completed: false}
}
