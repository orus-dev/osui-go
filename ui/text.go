package ui

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
)

type TextComponent struct {
	Data osui.ComponentData
	Text string
}

func (t *TextComponent) Update(string) bool {
	return false
}

func (t *TextComponent) GetComponentData() *osui.ComponentData {
	return &t.Data
}

func (t TextComponent) Render() string {
	return colors.Reset + t.Data.DefaultColor + t.Text
}

func (t *TextComponent) SetStyle(interface{}) {}

func Text(text string) *TextComponent {
	return &TextComponent{Text: text}
}
