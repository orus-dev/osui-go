package ui

import (
	"github.com/orus-dev/osui"
)

func InputBox(max_size uint) *osui.ComponentWrapper {
	return osui.NewComponent(ComponentInputBox{max_size: max_size})
}

func Text(text string) *osui.ComponentWrapper {
	return osui.NewComponent(ComponentText{Text: text})
}
