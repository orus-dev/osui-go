package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/isKey"
)

type InputBoxStyle struct {
	Background string `default:"" type:"bg"`
	Foreground string `default:"" type:"fg"`
	Outline    string `default:"" type:"fg"`
	Cursor     string `default:"" type:"fg"`
}

type InputBoxComponent struct {
	Data      osui.ComponentData
	Style     *InputBoxStyle
	max_size  uint
	cursor    uint
	InputData string
}

func (s *InputBoxComponent) GetComponentData() *osui.ComponentData {
	return &s.Data
}

func (s InputBoxComponent) Render() string {
	osui.UseStyle(s.Style)
	if s.max_size > uint(len(s.InputData)) {
		return fmt.Sprintf(
			" %s\n%s│%s%s│%s\n %s",
			colors.Reset+s.Style.Outline+strings.Repeat("_", int(s.max_size))+colors.Reset,
			colors.Reset+s.Style.Outline,
			colors.Combine(s.Style.Foreground, s.Style.Background)+s.InputData+osui.LogicValue(s.Data.IsActive, s.Style.Cursor+"█"+colors.Combine(s.Style.Foreground, s.Style.Background), ""),
			strings.Repeat(" ", int(s.max_size)-len(s.InputData)-osui.LogicValueInt(s.Data.IsActive, 1, 0))+colors.Reset+s.Style.Outline,
			colors.Reset+s.Data.DefaultColor,
			s.Style.Outline+strings.Repeat("‾", int(s.max_size))+colors.Reset+s.Data.DefaultColor,
		)
	}

	return fmt.Sprintf(
		" %s\n%s│%s%s\n %s",
		colors.Reset+s.Style.Outline+strings.Repeat("_", int(s.max_size))+colors.Reset,
		colors.Reset+s.Style.Outline,
		colors.Combine(s.Style.Foreground, s.Style.Background)+s.InputData+osui.LogicValue(s.Data.IsActive, s.Style.Cursor+"█"+colors.Reset, s.Style.Outline+"|"+colors.Reset),
		colors.Reset+s.Data.DefaultColor,
		colors.Reset+s.Style.Outline+strings.Repeat("‾", int(s.max_size))+colors.Reset+s.Data.DefaultColor,
	)
}

func (s *InputBoxComponent) Update(key string) bool {
	if isKey.Enter(key) {
		return true
	} else if isKey.Backspace(key) {
		if len(s.InputData) > 0 {
			s.InputData = s.InputData[:len(s.InputData)-1]
		}
	} else if isKey.Left(key) {
		if s.cursor > 0 {
			s.cursor--
		}
	} else if isKey.Right(key) {
		if s.cursor < s.max_size {
			s.cursor++
		}
	} else {
		if len(key) == 1 {
			if int(s.max_size) > len(s.InputData) {
				s.InputData += key
			}
		}
	}
	return false
}

func (d *InputBoxComponent) SetStyle(c interface{}) {
	d.Style = osui.SetDefaults(c.(*InputBoxStyle)).(*InputBoxStyle)
}

func InputBox(max_size uint) *InputBoxComponent {
	return &InputBoxComponent{max_size: max_size, Style: osui.SetDefaults(&InputBoxStyle{}).(*InputBoxStyle)}
}
