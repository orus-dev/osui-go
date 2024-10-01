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

func RenderLine(frameLine, line_ string, x int, fm *map[int]string, lm *map[int]string) string {
	frame := []rune(frameLine)
	line := []rune(line_)

	for i := range frame {
		if i >= x && i-x < len(line) {
			frame[i] = line[i-x]
			delete(*fm, i)
		}
		if v, ok := (*lm)[i-x]; ok {
			if i-x > 0 {
				delete(*lm, i-x)
				(*lm)[i] = v
			}
		}
	}

	return string(frame)
}

func RenderOnFrame(c Component, frame *[]string) {
	componentData := c.GetComponentData()
	for i, line := range strings.Split(c.Render(), "\n") {
		if int(componentData.Y)+i < len(*frame) {
			fo, fm := CompressString((*frame)[int(componentData.Y)+i], "")
			lo, lm := CompressString(line, "")
			r := RenderLine(fo, lo, int(componentData.X), &fm, &lm)
			(*frame)[int(componentData.Y)+i] = DecompressString(r, fm, lm)
		}
	}
}

func CompressString(input, repl string) (string, map[int]string) {
	pattern := `\x1b\[([0-9;]*)[a-zA-Z]`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringIndex(input, -1)
	matchesMap := make(map[int]string)
	n := 0
	for _, match := range matches {
		start := match[0]
		end := match[1]
		matchesMap[start-n] = input[start:end]
		n += len(input[start:end])
	}
	return re.ReplaceAllString(input, repl), matchesMap
}

func DecompressString(modified string, fm map[int]string, lm map[int]string) string {
	res := ""
	for i, c := range modified {
		if v, ok := fm[i]; ok {
			res += v
		}
		if v, ok := lm[i]; ok {
			res += v
		}
		res += string(c)
	}
	return res
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
