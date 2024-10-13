package ui

import (
	"github.com/orus-dev/osui"
)

type TextComponent struct {
	Data osui.ComponentData
	Text string
}

func (t *TextComponent) Update(ctx osui.UpdateContext) bool {
	return false
}

func (t *TextComponent) GetComponentData() *osui.ComponentData {
	return &t.Data
}

func (t TextComponent) Render() string {
	return t.Text
}

func Text(param osui.Param, text string) *TextComponent {
	param.SetDefaultBindings(map[string]string{})
	return param.UseParam(&TextComponent{Text: text}).(*TextComponent)
}
