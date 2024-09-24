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
	if err != nil {
		return fmt.Errorf("error getting the terminal size: %s", err)
	}
	frame := make([]string, height)
	for i := 0; i < height; i++ {
		frame[i] = strings.Repeat(" ", width)
	}
	for _, component := range s.components {
		RenderOnFrame(component, &frame)
	}
	fmt.Print(strings.Join(frame, "\n"))
	return nil
}
