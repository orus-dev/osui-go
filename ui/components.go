package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
)

type ComponentText struct{ Text string }

func (t ComponentText) Render(cw *osui.ComponentWrapper) string {
	return t.Text
}

func (t ComponentText) Run(cw *osui.ComponentWrapper, a ...any) any {
	return nil
}

type ComponentInputBox struct {
	max_size uint
}

func (s ComponentInputBox) Render(cw *osui.ComponentWrapper) string {
	return fmt.Sprintf(
		" %s\n|%s%s|\n %s",
		strings.Repeat("_", int(s.max_size)),
		cw.Data,
		strings.Repeat(" ", int(s.max_size)-len(cw.Data)),
		strings.Repeat("‾", int(s.max_size)),
	)
}

func (s ComponentInputBox) Run(cw *osui.ComponentWrapper, a ...any) any {
	for {
		fmt.Printf("\x1B[%d;%dH", cw.Y+2, int(cw.X+2)+len(cw.Data))
		key, err := osui.ReadKey()
		if err != nil {
			fmt.Println(err)
			break
		}
		switch cw.Update(cw, key) {

		case osui.UpdateOutput.Jmp:
			continue
		case osui.UpdateOutput.Exit:
			fmt.Print("\n\n")
			return cw.Data

		}
		switch key {

		case osui.Key.Enter:
			fmt.Print("\n\n")
			return cw.Data

		case osui.Key.Backspace:
			if len(cw.Data) > 0 {
				cw.Data = cw.Data[:len(cw.Data)-1]
			}

		default:
			if int(s.max_size) > len(cw.Data) {
				cw.Data += key
			}
		}
		cw.Render.Render()
	}
	return cw.Data
}
