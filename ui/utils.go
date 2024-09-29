package ui

import (
	"fmt"
	"math"
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

func findClosestComponent(a []osui.Component, i int, d string) int {
	if len(a) == 0 || i < 0 || i >= len(a) {
		return -1 // Invalid input
	}

	current := a[i].GetComponentData()
	closestIndex := -1
	minDistance := math.MaxFloat64

	for j, c := range a {
		component := c.GetComponentData()
		if i == j {
			continue // Skip the current component
		}

		// Check if the component is in the correct direction
		isValidDirection := false
		switch d {
		case "up":
			if component.Y < current.Y {
				isValidDirection = true
			}
		case "down":
			if component.Y > current.Y {
				isValidDirection = true
			}
		case "left":
			if component.X < current.X {
				isValidDirection = true
			}
		case "right":
			if component.X > current.X {
				isValidDirection = true
			}
		}

		// If the direction is valid, calculate distance
		if isValidDirection {
			distance := math.Hypot(float64(component.X-current.X), float64(component.Y-current.Y))
			if distance < minDistance {
				minDistance = distance
				closestIndex = j
			}
		}
	}

	return closestIndex
}
