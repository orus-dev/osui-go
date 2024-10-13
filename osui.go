package osui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/orus-dev/osui/colors"
	"golang.org/x/term"
)

type Key struct {
	Chars [3]rune
	Name  string
}

type UpdateContext struct {
	UpdateKind uint8
	Tick       uint8
	Key        Key
}

type Component interface {
	Render() string
	GetComponentData() *ComponentData
	Update(UpdateContext) bool
}

type Screen struct {
	Component Component
	tickRate  uint8
	running   bool
}

func NewScreen() *Screen {
	s := Screen{}
	s.tickRate = 1
	return &s
}

func (s *Screen) Render() {
	width, height := GetTerminalSize()
	data := s.Component.GetComponentData()
	if data.Height == 0 {
		data.Height = height
	}
	if data.Width == 0 {
		data.Width = width
	}
	frame := NewFrame(width, height)
	RenderOnFrame(s.Component, &frame)
	Clear()
	fmt.Print(strings.Join(frame, ""))
}

func (s *Screen) Run(c Component) {
	s.SetComponent(c)
	HideCursor()
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	s.running = true
	for s.running {
		s.Render()
		s.Component.Update(UpdateContext{UpdateKind: UpdateKindKey, Key: ReadKey()})
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
}

func (s *Screen) _() {
	var tick uint8 = 0
	for s.running {
		s.Component.Update(UpdateContext{UpdateKind: UpdateKindTick, Tick: tick})
		if tick == 255 {
			tick = 0
		}
		time.Sleep(time.Duration(1000/int(s.tickRate)) * time.Millisecond)
	}
}

func (s *Screen) SetComponent(c Component) {
	s.Component = c
	width, height := GetTerminalSize()
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
}

func (s *Screen) Exit() {
	s.running = false
	ShowCursor()
}

const (
	UpdateKindKey = uint8(iota)
	UpdateKindTick
)
