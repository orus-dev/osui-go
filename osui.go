package osui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui/colors"
	"golang.org/x/term"
)

type ComponentData struct {
	X            uint16
	Y            uint16
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
	component    Component
	CustomRender func()
}

func NewScreen(c Component) *Screen {
	HideCursor()
	s := &Screen{component: c}
	return s
}

func (s *Screen) Render() error {
	if s.CustomRender != nil {
		s.CustomRender()
		return nil
	}
	Clear()
	width, height, err := term.GetSize(0)
	if err != nil {
		return fmt.Errorf("error getting the terminal size: %s", err)
	}
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
	return nil
}

func (s *Screen) Run() {
	s.component.GetComponentData().Screen = s
	for {
		s.Render()
		k, _ := ReadKey()
		if s.component.Update(k) {
			ShowCursor()
			return
		}
	}
}
