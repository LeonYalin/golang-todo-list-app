package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var table []Todo

func TestMain(m *testing.M) {
	table = []Todo{
		{ID: "1", Description: "1 Todo", Completed: false},
		{ID: "2", Description: "2 Todo", Completed: true},
	}
	os.Exit(m.Run())
}

func TestTodoCreation(t *testing.T) {
	for _, item := range table {
		todo := NewTodo(item.Description)
		assert.NotNil(t, todo.ID)
		assert.Equal(t, todo.Completed, false)
		assert.Equal(t, todo.Description, item.Description)
		assert.Contains(t, todo.Print(), item.Description)
	}
}
