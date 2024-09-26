package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
)

type InputBoxStyle struct {
	Background string `defaults:"" type:"bg"`
	Foreground string `defaults:"" type:"fg"`
	Outline    string `defaults:"" type:"fg"`
	Cursor     string `defaults:"" type:"fg"`
}

type InputBoxComponent struct {
	Data      osui.ComponentData
	Style     *InputBoxStyle
	max_size  uint
	InputData string
}

func (s *InputBoxComponent) GetComponentData() *osui.ComponentData {
	return &s.Data
}

func (s InputBoxComponent) Render() string {
	osui.UseStyle(s.Style)
	if s.max_size > uint(len(s.InputData)) {
		return fmt.Sprintf(
			" %s\n%s|%s%s|%s\n %s",
			colors.Reset+s.Style.Outline+strings.Repeat("_", int(s.max_size))+colors.Reset,
			colors.Reset+s.Style.Outline,
			colors.Combine(s.Style.Foreground, s.Style.Background)+s.InputData+osui.LogicValue(s.Data.IsActive, s.Style.Cursor+"█"+colors.Combine(s.Style.Foreground, s.Style.Background), ""),
			strings.Repeat(" ", int(s.max_size)-len(s.InputData)-osui.LogicValueInt(s.Data.IsActive, 1, 0))+colors.Reset+s.Style.Outline,
			colors.Reset+s.Data.DefaultColor,
			s.Style.Outline+strings.Repeat("‾", int(s.max_size))+colors.Reset+s.Data.DefaultColor,
		)
	}
	return fmt.Sprintf(
		" %s\n|%s\n %s",
		strings.Repeat("_", int(s.max_size)),
		s.InputData+osui.LogicValue(s.Data.IsActive, "█", "|"),
		strings.Repeat("‾", int(s.max_size)),
	)
}

func (s *InputBoxComponent) Read() error {
	for {
		key, err := osui.ReadKey()
		if err != nil {
			fmt.Print("\n\n")
			fmt.Println(err)
			break
		}
		if s.Update(key) {
			fmt.Print("\n\n")
			return nil
		}
	}
	return nil
}

func (s *InputBoxComponent) Update(key string) bool {
	switch key {
	case osui.Key.Enter:
		return true
	case osui.Key.Backspace:
		if len(s.InputData) > 0 {
			s.InputData = s.InputData[:len(s.InputData)-1]
		}

	default:
		if len(key) == 1 {
			if int(s.max_size) > len(s.InputData) {
				s.InputData += key
			}
		}
	}
	return false
}

func InputBox(max_size uint) *InputBoxComponent {
	return &InputBoxComponent{max_size: max_size, Style: osui.SetDefaults(&InputBoxStyle{}).(*InputBoxStyle)}
}
