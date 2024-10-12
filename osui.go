package osui

import (
	"fmt"
	"os"
	"strings"

	"github.com/orus-dev/osui/colors"
	"golang.org/x/term"
)

type Style struct {
	Background string `json:"background"`
	Foreground string `json:"foreground"`
	Outline    string `json:"outline"`

	ActiveBackground string `json:"activeBackground"`
	ActiveForeground string `json:"activeForeground"`
	ActiveOutline    string `json:"activeOutline"`

	ClickedBackground string `json:"clickedBackground"`
	ClickedForeground string `json:"clickedForeground"`
	ClickedOutline    string `json:"clickedOutline"`

	SelectedBackground string `json:"selectedBackground"`
	SelectedForeground string `json:"selectedForeground"`
	SelectedOutline    string `json:"selectedOutline"`

	Cursor         string `json:"cursor"`
	InactiveCursor string `json:"inactiveCursor"`
}

func (s *Style) SetDefaults() {
	if s.Cursor != "" {
		s.Cursor = "> "
	}
}
func (s *Style) UseStyle() {
	s.Background = colors.AsBg(s.Background)
	s.Foreground = colors.AsBg(s.Foreground)
	s.Outline = colors.AsBg(s.Outline)

	s.ActiveBackground = colors.AsBg(s.ActiveBackground)
	s.ActiveForeground = colors.AsBg(s.ActiveForeground)
	s.ActiveOutline = colors.AsBg(s.ActiveOutline)

	s.ClickedBackground = colors.AsBg(s.ClickedBackground)
	s.ClickedForeground = colors.AsFg(s.ClickedForeground)
	s.ClickedOutline = colors.AsFg(s.ClickedOutline)
}

type Param struct {
	Style   Style `json:"style"`
	X       int   `json:"y"`
	Y       int   `json:"x"`
	Width   int   `json:"width"`
	Height  int   `json:"height"`
	OnClick func()
	Keys    map[string]func(string) bool
	Toggle  bool `json:"useToggle"`
}

func (p *Param) SetDefaultBindings(keys map[string]func(string) bool) {
	if p.Keys == nil {
		p.Keys = keys
		return
	}
	for k, f := range keys {
		if _, ok := p.Keys[k]; !ok {
			p.Keys[k] = f
		}
	}
}

func (p *Param) UseParam(c Component) Component {
	data := c.GetComponentData()

	p.Style.SetDefaults()
	data.Style = p.Style
	data.X = p.X
	data.Y = p.Y
	if p.Height != 0 {
		data.Height = p.Height
	}
	if p.Width != 0 {
		data.Width = p.Width
	}

	if p.OnClick != nil {
		data.OnClick = p.OnClick
	} else {
		data.OnClick = func() {}
	}

	data.Toggle = p.Toggle

	data.Keys = p.Keys

	return c
}

type ComponentData struct {
	X            int
	Y            int
	Width        int
	Height       int
	DefaultColor string
	IsActive     bool
	Style        Style
	OnClick      func()
	Keys         map[string]func(string) bool
	Toggle       bool
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

func NewScreen() *Screen {
	return &Screen{}
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
	for {
		s.Render()
		k, _ := ReadKey()
		if s.Component.Update(k) {
			ShowCursor()
			return
		}
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
