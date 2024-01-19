package util

import "github.com/TwiN/go-color"

func YellowText(text string) string {
	return color.Colorize(color.Yellow, text)
}

func RedText(text string) string {
	return color.Colorize(color.Red, text)
}
