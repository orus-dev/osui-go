package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
)

type InputBoxComponent struct {
	Data      osui.ComponentData
	max_size  uint
	InputData string
}

func (s *InputBoxComponent) GetComponentData() *osui.ComponentData {
	return &s.Data
}

func (s InputBoxComponent) Render() string {
	if s.max_size > uint(len(s.InputData)) {
		return fmt.Sprintf(
			" %s\n|%s%s|\n %s",
			strings.Repeat("_", int(s.max_size)),
			s.InputData+osui.LogicValue(s.Data.IsActive, "█", ""),
			strings.Repeat(" ", int(s.max_size)-len(s.InputData)-osui.LogicValueInt(s.Data.IsActive, 1, 0)),
			strings.Repeat("‾", int(s.max_size)),
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
	case "":
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
	return &InputBoxComponent{max_size: max_size}
}
