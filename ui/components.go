package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
)

type ComponentText struct {
	Text string
}

func (t ComponentText) Render(cw *osui.ComponentWrapper) string {
	return t.Text
}

func (t ComponentText) Run(cw *osui.ComponentWrapper) error {
	return nil
}

type ComponentInputBox struct {
	max_size uint
	data     string
}

func (s ComponentInputBox) Render(cw *osui.ComponentWrapper) string {
	return fmt.Sprintf(
		" %s\n|%s%s|\n %s",
		strings.Repeat("_", int(s.max_size)),
		s.data,
		strings.Repeat(" ", int(s.max_size)-len(s.data)),
		strings.Repeat("‾", int(s.max_size)),
	)
}

func (s ComponentInputBox) Run(cw *osui.ComponentWrapper) error {
	for {
		fmt.Printf("\x1B[%d;%dH", cw.Y+2, int(cw.X+2)+len(s.data))
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
			return nil

		}
		switch key {

		case osui.Key.Enter:
			fmt.Print("\n\n")
			return nil

		case osui.Key.Backspace:
			if len(s.data) > 0 {
				s.data = s.data[:len(s.data)-1]
			}

		default:
			if int(s.max_size) > len(s.data) {
				s.data += key
			}
		}
		cw.Render(s)
	}
	return nil
}
