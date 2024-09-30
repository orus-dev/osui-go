package ui

import (
	"fmt"
	"math"
	"strings"

	"github.com/orus-dev/osui"
)

func WithPosition(x, y int, c osui.Component) osui.Component {
	data := c.GetComponentData()
	data.X = x
	data.Y = y
	return c
}

func WithSize(width, height int, c osui.Component) osui.Component {
	data := c.GetComponentData()
	data.Width = width
	data.Height = height
	return c
}

func centerText(text string, width int) string {
	if len(text) > width {
		return text[:width]
	}
	padding := (width - len(text)) / 2
	leftPadding := strings.Repeat(" ", padding)
	rightPadding := strings.Repeat(" ", padding)

	return fmt.Sprintf("%s%s%s", leftPadding, text, rightPadding)
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

type Id[T osui.Component] struct {
	changer   func(T) T
	Component T
}

func (i *Id[T]) Id(c T) T {
	i.Component = c
	if i.changer != nil {
		return i.changer(i.Component)
	}
	return i.Component
}

func (i *Id[T]) SetProperties(c func(T) T) {
	i.changer = c
}

type Class[T osui.Component] struct {
	changer func(T) T
}

func (c *Class[T]) Class(co T) T {
	if c.changer != nil {
		return c.changer(co)
	}
	return co
}

func (c *Class[T]) SetProperties(ch func(T) T) {
	c.changer = ch
}
