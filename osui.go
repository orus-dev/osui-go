package osui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui/colors"
)

type ComponentData struct {
	X            int
	Y            int
	Width        int
	Height       int
	DefaultColor string
	IsActive     bool
	Screen       *Screen
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

func (s *Screen) Render() {
	Clear()
	width, height := GetTerminalSize()
	frame := NewFrame(width, height)
	data := s.component.GetComponentData()
	if data.Height == 0 {
		data.Height = height
	}
	if data.Width == 0 {
		data.Width = width
	}
	data.Screen = s
	data.IsActive = true
	data.DefaultColor = colors.Reset
	RenderOnFrame(s.component, &frame)
	fmt.Print(strings.Join(frame, "\n"))
}

func (s *Screen) Run() {
	data := s.component.GetComponentData()
	data.Screen = s
	for {
		s.Render()
		k, _ := ReadKey()
		if s.component.Update(k) {
			ShowCursor()
			return
		}
	}
}
