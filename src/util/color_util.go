package util

import (
	"fmt"
	"strconv"

	"github.com/TwiN/go-color"
	"github.com/leonyalin/todo-list/src/models"
)

func YellowText(text string) string {
	return color.Colorize(color.Yellow, text)
}

func RedText(text string) string {
	return color.Colorize(color.Red, text)
}

func GreenText(text string) string {
	return color.Colorize(color.Green, text)
}

func OptionText(index string, text string) string {
	return fmt.Sprintf("[%s]: %s", GreenText(index), text)
}

func TodoText(index int, todo *models.Todo) string {
	indexText := ""
	if index == -1 {
		indexText = ""
	} else {
		if todo.Completed {
			indexText = fmt.Sprintf("[%s]: ", GreenText(strconv.Itoa(index)))
		} else {
			indexText = fmt.Sprintf("[%s]: ", RedText(strconv.Itoa(index)))
		}
	}

	text := ""
	if todo.Completed {
		text = fmt.Sprintf("%s - %s", todo.Description, GreenText("Completed"))
	} else {
		text = fmt.Sprintf("%s - %s", todo.Description, RedText("Active"))
	}
	return fmt.Sprintf("%s%s", indexText, text)
}
