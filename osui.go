package osui

import (
	"fmt"
	"strings"

	"golang.org/x/term"
)

type ComponentData struct {
	X      uint16
	Y      uint16
	Screen *Screen
}

type Component interface {
	Render() string
	GetComponentData() *ComponentData
}

type Screen struct {
	components []Component
	Width      int
	Height     int
}

func NewScreen(c ...Component) *Screen {
	s := &Screen{components: c}
	for _, component := range s.components {
		componentData := component.GetComponentData()
		if componentData.Screen == nil {
			componentData.Screen = s
		} else {
			componentData.Screen.components = s.components
		}
	}
	return s
}

func (s *Screen) Render() error {
	Clear()
	width, height, err := term.GetSize(0)
	s.Width = width
	s.Height = height
	if err != nil {
		return fmt.Errorf("error getting the terminal size: %s", err)
	}
	frame := NewFrame(width, height)
	for _, component := range s.components {
		RenderOnFrame(component, &frame)
	}
	fmt.Print(strings.Join(frame, "\n"))
	return nil
}
