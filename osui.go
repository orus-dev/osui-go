package osui

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type Component interface {
	Render(*ComponentWrapper) string
	Run(*ComponentWrapper) error
}

type ComponentWrapper struct {
	Component Component
	X         uint16
	Y         uint16
	render    *Render
	Update    func(*ComponentWrapper, any) UpdateOutput_
}

func (c *ComponentWrapper) Run() {
	c.Component.Run(c)
}

func (cw *ComponentWrapper) Render(c Component) error {
	cw.Component = c
	return cw.render.Render()
}

func (c *ComponentWrapper) SetRender(r *Render) {
	c.render = r
}

func NewComponent(c Component) *ComponentWrapper {
	return &ComponentWrapper{
		Component: c,
		Update:    func(cw *ComponentWrapper, a any) UpdateOutput_ { return UpdateOutput.Continue },
	}
}

type Render struct {
	components []*ComponentWrapper
}

func NewRender() *Render {
	return &Render{components: []*ComponentWrapper{}}
}

func (r *Render) Add(c *ComponentWrapper) *ComponentWrapper {
	c.SetRender(r)
	r.components = append(r.components, c)
	return c
}

func (r *Render) Render() error {
	Clear()
	width, height, err := term.GetSize(0)
	if err != nil {
		return fmt.Errorf("error getting the terminal size: %s", err)
	}
	frame := make([]string, height)
	for i := 0; i < height; i++ {
		frame[i] = strings.Repeat(" ", width)
	}
	for _, c := range r.components {
		for i, line := range strings.Split(c.Component.Render(c), "\n") {
			if int(c.Y)+i < len(frame) {
				frame[int(c.Y)+i] = renderLine(frame[int(c.Y)+i], line, int(c.X))
			}
		}
	}
	fmt.Print(strings.Join(frame, "\n"))
	return nil
}

func Clear() {
	fmt.Print("\x1b[H\x1b[2J\x1b[3J")
}

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

func KeyUpdate(keys map[string]UpdateOutput_) func(c *ComponentWrapper, key any) UpdateOutput_ {
	return func(c *ComponentWrapper, k any) UpdateOutput_ {
		action, ok := keys[fmt.Sprint(k)]
		if ok {
			return action
		} else {
			return UpdateOutput.Continue
		}
	}
}

type UpdateOutput_ = int

type updateOutput struct {
	/*Continue the default check*/
	Continue UpdateOutput_
	/*Jump to the next check*/
	Jmp UpdateOutput_
	/*Leave the loop and stop the function*/
	Exit UpdateOutput_
}

var UpdateOutput = updateOutput{Continue: 0, Jmp: 1, Exit: 2}

/*
Keys
*/

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
