package osui

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/nathan-fiscaletti/consolesize-go"
	"github.com/orus-dev/osui/colors"
	"golang.org/x/term"
)

var re = regexp.MustCompile(`(\x1b\[([0-9;]*)[a-zA-Z])+`)

func RenderLine(frame_, line_ string, x int, fm, lm map[int]string) string {
	var res strings.Builder
	frame := []rune(frame_)
	line := []rune(line_)

	flen, llen := len(frame), len(line)
	for i := 0; i < flen; i++ {
		if i >= x && i-x < llen {
			if v, ok := lm[i-x]; ok {
				res.WriteString(v)
			}
			res.WriteRune(line[i-x])
		} else {
			if v, ok := fm[i]; ok {
				res.WriteString(v)
			}
			res.WriteRune(frame[i])
		}
	}
	return res.String()
}

func CompressString(input string) (string, map[int]string) {
	matchesMap := make(map[int]string)
	res := []rune{}

	i := 0
	for i < len([]rune(input)) {
		loc := re.FindStringIndex(string([]rune(input)[i:]))

		if loc != nil && loc[0] == 0 {
			ansiSeq := re.FindString(string([]rune(input)[i:]))
			matchesMap[len(res)] = ansiSeq
			i += len([]rune(ansiSeq))
		} else {
			res = append(res, []rune(input)[i])
			i++
		}
	}

	return string(res), matchesMap
}

func RenderOnFrame(c Component, frame *[]string) {
	componentData := c.GetComponentData()
	for i, line := range strings.Split(c.Render(), "\n") {
		if int(componentData.Y)+i < len(*frame) {
			fo, fm := CompressString((*frame)[int(componentData.Y)+i])
			lo, lm := CompressString(line)
			(*frame)[int(componentData.Y)+i] = RenderLine(fo, lo, componentData.X, fm, lm)
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

func NewFrame(width, height int) []string {
	frame := make([]string, height)
	for i := 0; i < height; i++ {
		frame[i] = strings.Repeat(" ", width)
	}
	return frame
}

func Clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func SetDefaults(p interface{}) interface{} {
	v := reflect.ValueOf(p)
	if v.Kind() != reflect.Ptr {
		panic("SetDefaults: expected a pointer to a struct")
	}
	val := v.Elem()
	if val.Kind() != reflect.Struct {
		panic("SetDefaults: expected a pointer to a struct")
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		structField := typ.Field(i)
		if defaultValue, ok := structField.Tag.Lookup("default"); ok {
			if field.Kind() == reflect.String && field.String() == "" {
				field.SetString(defaultValue)
			}
		}
	}
	return p
}

func UseStyle(p interface{}) {
	SetDefaults(p)
	v := reflect.ValueOf(p)
	if v.Kind() != reflect.Ptr {
		panic("SetStyle: expected a pointer to a struct")
	}
	val := v.Elem()
	if val.Kind() != reflect.Struct {
		panic("SetStyle: expected a pointer to a struct")
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		structField := typ.Field(i)
		if styleType, ok := structField.Tag.Lookup("type"); ok {
			if field.Kind() == reflect.String {
				if styleType == "bg" || styleType == "background" {
					field.SetString(colors.AsBg(field.String()))
				} else {
					field.SetString(colors.AsFg(field.String()))
				}
			}
		}
	}

}

func ShowCursor() {
	fmt.Print("\033[?25h")
}

func HideCursor() {
	fmt.Print("\033[?25l")
}

func LogicValue(b bool, _if, _else string) string {
	if b {
		return _if
	}
	return _else
}

func LogicValueInt(b bool, _if, _else int) int {
	if b {
		return _if
	}
	return _else
}

// Get the terminal size
func GetTerminalSize() (int, int) {
	return consolesize.GetConsoleSize()
}
