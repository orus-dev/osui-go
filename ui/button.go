package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/isKey"
)

type ButtonStyle struct {
	Background string `default:"" type:"bg"`
	Foreground string `default:"" type:"fg"`
	Outline    string `default:"" type:"fg"`

	ActiveBackground string `default:"" type:"bg"`
	ActiveForeground string `default:"\x1b[34m" type:"fg"`
	ActiveOutline    string `default:"" type:"fg"`

	ClickedBackground string `default:"" type:"bg"`
	ClickedForeground string `default:"\x1b[32m" type:"fg"`
	ClickedOutline    string `default:"" type:"fg"`
}

type ButtonComponent struct {
	Data    osui.ComponentData
	Style   *ButtonStyle
	Text    string
	Toggle  bool
	Clicked bool
	OnClick func(*ButtonComponent) bool
}

func (b *ButtonComponent) Render() string {
	osui.UseStyle(b.Style)

	if b.Clicked {
		return fmt.Sprintf(" %s\n%s│%s│%s\n %s",
			colors.Reset+b.Style.ClickedOutline+strings.Repeat("_", b.Data.Width)+colors.Reset+b.Data.DefaultColor,
			colors.Reset+b.Style.ClickedOutline,
			colors.Reset+colors.Combine(b.Style.ClickedBackground, b.Style.ClickedForeground)+centerText(b.Text, b.Data.Width)+colors.Reset+b.Style.ClickedOutline,
			colors.Reset+b.Data.DefaultColor,
			colors.Reset+b.Style.ClickedOutline+strings.Repeat("‾", b.Data.Width)+colors.Reset+b.Data.DefaultColor,
		)
	}

	if b.Data.IsActive {
		return fmt.Sprintf(" %s\n%s│%s│%s\n %s",
			colors.Reset+b.Style.ActiveOutline+strings.Repeat("_", b.Data.Width)+colors.Reset+b.Data.DefaultColor,
			colors.Reset+b.Style.ActiveOutline,
			colors.Reset+colors.Combine(b.Style.ActiveBackground, b.Style.ActiveForeground)+centerText(b.Text, b.Data.Width)+colors.Reset+b.Style.ActiveOutline,
			colors.Reset+b.Data.DefaultColor,
			colors.Reset+b.Style.ActiveOutline+strings.Repeat("‾", b.Data.Width)+colors.Reset+b.Data.DefaultColor,
		)
	}

	return fmt.Sprintf(" %s\n%s│%s│%s\n %s",
		colors.Reset+b.Style.Outline+strings.Repeat("_", b.Data.Width)+colors.Reset+b.Data.DefaultColor,
		colors.Reset+b.Style.Outline,
		colors.Reset+colors.Combine(b.Style.Background, b.Style.Foreground)+centerText(b.Text, b.Data.Width)+colors.Reset+b.Style.Outline,
		colors.Reset+b.Data.DefaultColor,
		colors.Reset+b.Style.Outline+strings.Repeat("‾", b.Data.Width)+colors.Reset+b.Data.DefaultColor,
	)
}

func (b *ButtonComponent) Update(key string) bool {
	if isKey.Enter(key) {
		if !b.OnClick(b) {
			if b.Toggle {
				b.Clicked = !b.Clicked
				return false
			}
			b.Clicked = true
			b.Data.Screen.Render()
			time.Sleep(time.Millisecond * 120)
			b.Clicked = false
		}
	}
	return false
}

func (b *ButtonComponent) GetComponentData() *osui.ComponentData {
	return &b.Data
}

func (b *ButtonComponent) SetStyle(style interface{}) {
	b.Style = osui.SetDefaults(style).(*ButtonStyle)
}

type BtnParams struct {
	Style   ButtonStyle
	Toggle  bool
	OnClick func(*ButtonComponent) bool
	Width   int
}

func Button(text string, p ...BtnParams) *ButtonComponent {
	if len(p) > 0 {
		param := p[0]
		return &ButtonComponent{Text: text, Style: osui.SetDefaults(&param.Style).(*ButtonStyle),
			Data: osui.ComponentData{
				Width:  osui.LogicValueInt(param.Width == 0, 20, param.Width),
				Height: 1,
			},
			OnClick: fnLogicValue(param.OnClick),
		}

	}
	return &ButtonComponent{Text: text,
		Style: osui.SetDefaults(&ButtonStyle{}).(*ButtonStyle),
		Data: osui.ComponentData{
			Width:  20,
			Height: 1,
		},
		OnClick: func(bc *ButtonComponent) bool { return false },
	}
}
