package util

import (
	"fmt"
	"testing"

	"github.com/leonyalin/todo-list/src/models"
	"github.com/stretchr/testify/assert"
)

func TestColors(t *testing.T) {
	// setup
	text := "sample text"
	todo1 := models.NewTodo("test todo 1")
	todo2 := models.NewTodo("test todo 2")
	todo2.Completed = true
	tt := []struct {
		title    string
		expected string
		actual   string
	}{
		{title: "Yellow text", expected: YellowText(text), actual: fmt.Sprintf("\x1b[33m%s\x1b[0m", text)},
		{title: "Red text", expected: RedText(text), actual: fmt.Sprintf("\x1b[31m%s\x1b[0m", text)},
		{title: "Green text", expected: GreenText(text), actual: fmt.Sprintf("\x1b[32m%s\x1b[0m", text)},
		{title: "Options text", expected: OptionText("1", text), actual: fmt.Sprintf("[\x1b[32m%s\x1b[0m]: %s", "1", text)},
		{title: "Todo text 1", expected: TodoText(1, todo1), actual: fmt.Sprintf("[\x1b[31m%s\x1b[0m]: %s - \x1b[31mActive\x1b[0m", "1", todo1.Description)},
		{title: "Todo text 2", expected: TodoText(2, todo2), actual: fmt.Sprintf("[\x1b[32m%s\x1b[0m]: %s - \x1b[32mCompleted\x1b[0m", "2", todo2.Description)},
		{title: "Todo text 3", expected: TodoText(-1, todo1), actual: fmt.Sprintf("%s - \x1b[31mActive\x1b[0m", todo1.Description)},
	}

	for _, test := range tt {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(t, test.expected, test.actual)
		})
	}

	// teardown
}
