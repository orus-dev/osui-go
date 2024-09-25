package osui

import (
	"fmt"
	"strings"

	"golang.org/x/term"
)

type ComponentData struct {
	X        uint16
	Y        uint16
	Width    int
	Height   int
	IsActive bool
	Screen   *Screen
}

type Component interface {
	Render() string
	GetComponentData() *ComponentData
	Update(string) bool
}

type Screen struct {
	component Component
}

func NewScreen(c Component) *Screen {
	HideCursor()
	s := &Screen{component: c}
	return s
}

func (s *Screen) Render() error {
	Clear()
	width, height, err := term.GetSize(0)
	if err != nil {
		return fmt.Errorf("error getting the terminal size: %s", err)
	}
	frame := NewFrame(width, height)
	data := s.component.GetComponentData()
	data.Width = width
	data.Height = height
	data.Screen = s
	data.IsActive = true
	RenderOnFrame(s.component, &frame)
	fmt.Print(strings.Join(frame, "\n"))
	return nil
}

func (s *Screen) Run() {
	s.component.GetComponentData().Screen = s
	s.component.Update("OSUI:RENDER")
	for {
		s.Render()
		k, _ := ReadKey()
		if s.component.Update(k) {
			ShowCursor()
			return
		}
	}
}
