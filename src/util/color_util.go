package util

import (
	"fmt"

	"github.com/TwiN/go-color"
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
