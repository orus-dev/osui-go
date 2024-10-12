package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/isKey"
)

type ButtonComponent struct {
	Data    osui.ComponentData
	Text    string
	clicked bool
}

func (b *ButtonComponent) Render() string {
	b.Data.Style.UseStyle()

	if b.clicked {
		return fmt.Sprintf(" %s\n%s│%s│%s\n %s",
			colors.Reset+b.Data.Style.ClickedOutline+strings.Repeat("_", b.Data.Width-2)+colors.Reset+b.Data.DefaultColor,
			colors.Reset+b.Data.Style.ClickedOutline,
			colors.Reset+colors.Combine(b.Data.Style.ClickedBackground, b.Data.Style.ClickedForeground)+centerText(b.Text, b.Data.Width-2)+colors.Reset+b.Data.Style.ClickedOutline,
			colors.Reset+b.Data.DefaultColor,
			colors.Reset+b.Data.Style.ClickedOutline+strings.Repeat("‾", b.Data.Width-2)+colors.Reset+b.Data.DefaultColor,
		)
	}

	if b.Data.IsActive {
		return fmt.Sprintf(" %s\n%s│%s│%s\n %s",
			colors.Reset+b.Data.Style.ActiveOutline+strings.Repeat("_", b.Data.Width-2)+colors.Reset+b.Data.DefaultColor,
			colors.Reset+b.Data.Style.ActiveOutline,
			colors.Reset+colors.Combine(b.Data.Style.ActiveBackground, b.Data.Style.ActiveForeground)+centerText(b.Text, b.Data.Width-2)+colors.Reset+b.Data.Style.ActiveOutline,
			colors.Reset+b.Data.DefaultColor,
			colors.Reset+b.Data.Style.ActiveOutline+strings.Repeat("‾", b.Data.Width-2)+colors.Reset+b.Data.DefaultColor,
		)
	}

	return fmt.Sprintf(" %s\n%s│%s│%s\n %s",
		colors.Reset+b.Data.Style.Outline+strings.Repeat("_", b.Data.Width-2)+colors.Reset+b.Data.DefaultColor,
		colors.Reset+b.Data.Style.Outline,
		colors.Reset+colors.Combine(b.Data.Style.Background, b.Data.Style.Foreground)+centerText(b.Text, b.Data.Width-2)+colors.Reset+b.Data.Style.Outline,
		colors.Reset+b.Data.DefaultColor,
		colors.Reset+b.Data.Style.Outline+strings.Repeat("‾", b.Data.Width-2)+colors.Reset+b.Data.DefaultColor,
	)
}

func (b *ButtonComponent) Update(key string) bool {
	if isKey.Enter(key) {
		b.Data.OnClick()
		if b.Data.Toggle {
			b.clicked = !b.clicked
			return false
		}
		b.clicked = true
		b.Data.Screen.Render()
		time.Sleep(time.Millisecond * 120)
		b.clicked = false
	}
	return false
}

func (b *ButtonComponent) GetComponentData() *osui.ComponentData {
	return &b.Data
}

func Button(param osui.Param, text string) *ButtonComponent {
	return param.UseParam(&ButtonComponent{Text: text,
		Data: osui.ComponentData{
			Width:  20,
			Height: 1,
		},
		clicked: false,
	}).(*ButtonComponent)
}
