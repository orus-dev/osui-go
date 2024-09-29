package ui

import (
	"github.com/orus-dev/osui"
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
	return t.Text
}

func Text(text string) *TextComponent {
	return &TextComponent{Text: text}
}
