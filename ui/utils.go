package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
)

func centerText(text string, width int) string {
	if len(text) > width {
		return text[:width]
	}
	padding := (width - len(text)) / 2
	leftPadding := strings.Repeat(" ", padding)
	rightPadding := strings.Repeat(" ", padding)

	return fmt.Sprintf("%s%s%s", leftPadding, text, rightPadding)
}

func fnLogicValue(f func(*ButtonComponent) bool) func(*ButtonComponent) bool {
	if f == nil {
		return func(*ButtonComponent) bool { return false }
	} else {
		return f
	}
}

func WithPosition(x, y int, c osui.Component) osui.Component {
	data := c.GetComponentData()
	data.X = x
	data.Y = y
	return c
}

func WithStyle(style interface{}, c osui.Component) osui.Component {
	c.SetStyle(style)
	return c
}

func WithSize(width, height int, c osui.Component) osui.Component {
	data := c.GetComponentData()
	data.Width = width
	data.Height = height
	return c
}
