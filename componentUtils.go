package osui

import (
	"github.com/orus-dev/osui/colors"
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
	CursorColor    string `json:"cursorColor"`
	InactiveCursor string `json:"inactiveCursor"`
}

func (s *Style) SetDefaults() {
	defaultValue(&s.Cursor, "> ")
	defaultValue(&s.SelectedForeground, colors.Black)
	defaultValue(&s.SelectedBackground, colors.Magenta)
	defaultValue(&s.ActiveOutline, colors.Blue)
	defaultValue(&s.ClickedForeground, colors.Blue)
	defaultValue(&s.ClickedOutline, colors.Blue)
	s.UseStyle()
}

func (s *Style) UseStyle() {
	s.Foreground = colors.AsFg(s.Foreground)
	s.Background = colors.AsBg(s.Background)
	s.Outline = colors.AsFg(s.Outline)

	s.ActiveForeground = colors.AsFg(s.ActiveForeground)
	s.ActiveBackground = colors.AsBg(s.ActiveBackground)
	s.ActiveOutline = colors.AsFg(s.ActiveOutline)

	s.ClickedForeground = colors.AsFg(s.ClickedForeground)
	s.ClickedBackground = colors.AsBg(s.ClickedBackground)
	s.ClickedOutline = colors.AsFg(s.ClickedOutline)

	s.SelectedForeground = colors.AsFg(s.SelectedForeground)
	s.SelectedBackground = colors.AsBg(s.SelectedBackground)
	s.SelectedOutline = colors.AsFg(s.SelectedOutline)
}

func defaultValue(sd *string, s string) {
	if *sd == "" {
		*sd = s
	}
}

type Param struct {
	Style   Style `json:"style"`
	X       int   `json:"y"`
	Y       int   `json:"x"`
	Width   int   `json:"width"`
	Height  int   `json:"height"`
	OnClick func()
	Keys    map[string]string
	Toggle  bool `json:"useToggle"`
}

func (p *Param) SetDefaultBindings(keys map[string]string) {
	p.Style.SetDefaults()
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
	Keys         map[string]string
	Toggle       bool
	Screen       *Screen
}
