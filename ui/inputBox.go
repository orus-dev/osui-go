package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/keys"
)

type InputBoxComponent struct {
	Data      osui.ComponentData
	max_size  uint
	cursor    uint
	InputData string
}

func (s *InputBoxComponent) GetComponentData() *osui.ComponentData {
	return &s.Data
}

func (s InputBoxComponent) Render() string {
	s.Data.Style.UseStyle()
	if s.max_size > uint(len(s.InputData)) {
		return fmt.Sprintf(
			" %s\n%s│%s%s│%s\n %s",
			colors.Reset+s.Data.Style.Outline+strings.Repeat("_", int(s.max_size))+colors.Reset,
			colors.Reset+s.Data.Style.Outline,
			colors.Combine(s.Data.Style.Foreground, s.Data.Style.Background)+s.InputData+osui.LogicValue(s.Data.IsActive, s.Data.Style.CursorColor+"█"+colors.Combine(s.Data.Style.Foreground, s.Data.Style.Background), ""),
			strings.Repeat(" ", int(s.max_size)-len(s.InputData)-osui.LogicValueInt(s.Data.IsActive, 1, 0))+colors.Reset+s.Data.Style.Outline,
			colors.Reset+s.Data.DefaultColor,
			s.Data.Style.Outline+strings.Repeat("‾", int(s.max_size))+colors.Reset+s.Data.DefaultColor,
		)
	}

	return fmt.Sprintf(
		" %s\n%s│%s%s\n %s",
		colors.Reset+s.Data.Style.Outline+strings.Repeat("_", int(s.max_size))+colors.Reset,
		colors.Reset+s.Data.Style.Outline,
		colors.Combine(s.Data.Style.Foreground, s.Data.Style.Background)+s.InputData+osui.LogicValue(s.Data.IsActive, s.Data.Style.CursorColor+"█"+colors.Reset, s.Data.Style.Outline+"|"+colors.Reset),
		colors.Reset+s.Data.DefaultColor,
		colors.Reset+s.Data.Style.Outline+strings.Repeat("‾", int(s.max_size))+colors.Reset+s.Data.DefaultColor,
	)
}

func (s *InputBoxComponent) Update(ctx osui.UpdateContext) bool {
	if ctx.UpdateKind == osui.UpdateKindKey {
		switch s.Data.Keys[ctx.Key.Name] {
		case "done":
			s.Data.OnClick()
			return true
		case "remove":
			if len(s.InputData) > 0 {
				s.InputData = s.InputData[:len(s.InputData)-1]
			}
		case "left":
			if s.cursor > 0 {
				s.cursor--
			}
		case "right":
			if s.cursor < s.max_size {
				s.cursor++
			}
		default:
			if ctx.Key.Chars[1] == '\x00' {
				if int(s.max_size) > len(s.InputData) {
					s.InputData += string(ctx.Key.Chars[0])
				}
			}
		}
	}

	return false
}

func InputBox(param osui.Param, max_size uint) *InputBoxComponent {
	param.SetDefaultBindings(map[string]string{
		keys.Enter:     "done",
		keys.Backspace: "remove",
		keys.Left:      "left",
		keys.Right:     "right",
	})
	return param.UseParam(&InputBoxComponent{max_size: max_size, Data: osui.ComponentData{}}).(*InputBoxComponent)
}
