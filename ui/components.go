package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
)

type TextComponent struct {
	Data osui.ComponentData
	Text string
}

func (t *TextComponent) GetComponentData() *osui.ComponentData {
	return &t.Data
}

func (t TextComponent) Render() string {
	return t.Text
}

type InputBoxComponent struct {
	Data      osui.ComponentData
	max_size  uint
	InputData string
}

func (s *InputBoxComponent) GetComponentData() *osui.ComponentData {
	return &s.Data
}

func (s InputBoxComponent) Render() string {
	return fmt.Sprintf(
		" %s\n|%s%s|\n %s",
		strings.Repeat("_", int(s.max_size)),
		s.InputData,
		strings.Repeat(" ", int(s.max_size)-len(s.InputData)),
		strings.Repeat("‾", int(s.max_size)),
	)
}

func (s *InputBoxComponent) Read() error {
	for {
		fmt.Printf("\x1B[%d;%dH", s.Data.Y+2, int(s.Data.X+2)+len(s.InputData))
		key, err := osui.ReadKey()
		if err != nil {
			fmt.Print("\n\n")
			fmt.Println(err)
			break
		}
		switch key {
		case osui.Key.Enter:
			fmt.Print("\n\n")
			return nil

		case osui.Key.Backspace:
			if len(s.InputData) > 0 {
				s.InputData = s.InputData[:len(s.InputData)-1]
			}

		default:
			if int(s.max_size) > len(s.InputData) {
				s.InputData += key
			}
		}
		s.Data.Screen.Render()
	}
	return nil
}
