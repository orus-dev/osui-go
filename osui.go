package osui

import (
	"fmt"
	"os"
	"strings"

	"github.com/orus-dev/osui/colors"
	"golang.org/x/term"
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
	Component Component
}

func NewScreen(c Component) *Screen {
	HideCursor()
	s := &Screen{Component: c}
	return s
}

func (s *Screen) Render() {
	width, height := GetTerminalSize()
	frame := NewFrame(width, height)
	data := s.Component.GetComponentData()
	if data.Height == 0 {
		data.Height = height
	}
	if data.Width == 0 {
		data.Width = width
	}
	data.Screen = s
	data.IsActive = true
	data.DefaultColor = colors.Reset
	RenderOnFrame(s.Component, &frame)
	Clear()
	fmt.Print(strings.Join(frame, ""))
}

func (s *Screen) Run() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	data := s.Component.GetComponentData()
	data.Screen = s
	for {
		s.Render()
		k, _ := ReadKey()
		if s.Component.Update(k) {
			ShowCursor()
			return
		}
	}
}
