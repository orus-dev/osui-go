package cutils

import "github.com/orus-dev/osui"

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
