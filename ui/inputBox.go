package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
)

type InputBoxComponent struct {
	Data      osui.ComponentData
	cursor    string
	max_size  uint
	InputData string
}

func (s *InputBoxComponent) GetComponentData() *osui.ComponentData {
	return &s.Data
}

func (s InputBoxComponent) Render() string {
	return fmt.Sprintf(
		" %s\n|%s%s|\n %s%s",
		strings.Repeat("_", int(s.max_size)),
		s.InputData,
		strings.Repeat(" ", int(s.max_size)-len(s.InputData)),
		strings.Repeat("‾", int(s.max_size)),
		s.cursor,
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
	// s.cursor = fmt.Sprintf("\x1B[%d;%dH", s.Data.Y+2, int(s.Data.X+2)+len(s.InputData))
	// s.Render()
	switch key {
	case osui.Key.Enter:
		return true

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
	return false
}

func InputBox(max_size uint) *InputBoxComponent {
	return &InputBoxComponent{max_size: max_size}
}
