package osui

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func renderLine(frameLine, line string, x int) string {
	result := ""
	if x >= len(frameLine) {
		return result
	}

	lChars := []rune(line)
	for i, c := range frameLine {
		if i >= x && len(lChars) > i-x {
			result += string(lChars[i-x])
		} else {
			result += string(c)
		}
	}
	return result
}

func RenderOnFrame(c Component, frame *[]string) {
	componentData := c.GetComponentData()
	for i, line := range strings.Split(c.Render(), "\n") {
		if int(componentData.Y)+i < len(*frame) {
			(*frame)[int(componentData.Y)+i] = renderLine((*frame)[int(componentData.Y)+i], line, int(componentData.X))
		}
	}
}

func ReadKey() (string, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	var b [3]byte
	n, err := os.Stdin.Read(b[:])
	if err != nil {
		return "", err
	}
	return string(b[:n]), nil
}

func Clear() {
	fmt.Print("\x1b[H\x1b[2J\x1b[3J")
}

type Key_ = string

type key struct {
	Enter     Key_
	Tab       Key_
	Backspace Key_
	Escape    Key_
	Up        Key_
	Down      Key_
	Right     Key_
	Left      Key_
}

var Key = key{
	Enter:     "\r",
	Tab:       "\t",
	Backspace: "\x7f",
	Escape:    "\x1b",
	Up:        "\x1b[A",
	Down:      "\x1b[B",
	Right:     "\x1b[C",
	Left:      "\x1b[D",
}
