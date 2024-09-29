package osui

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/orus-dev/osui/colors"
	"golang.org/x/term"
)

func renderLine(frameLine, line_ string, x int) string {
	frame := []rune(frameLine)
	line := []rune(line_)

	i := 0
	v := 0
	for j, c := range frame {
		if i+j >= x && v < len(line) {
			if c != '\b' && line[v] == '\t' {
				frame = append(frame[:i+j], append([]rune{' '}, frame[i+j:]...)...)
			}
			frame[i+j] = line[v]
			v++
		}
		if c == '\b' {
			i++
		}
	}

	return string(frame)
}

func RenderOnFrame(c Component, frame *[]string) {
	componentData := c.GetComponentData()
	for i, lo := range strings.Split(c.Render(), "\n") {
		if int(componentData.Y)+i < len(*frame) {
			fo := (*frame)[int(componentData.Y)+i]
			f, fa := CompressString(fo, "\b")
			line, la := CompressString(lo, "\t")
			r := renderLine(f, line, int(componentData.X))
			(*frame)[int(componentData.Y)+i] = DecompressString(DecompressString(r, la, "\t"), fa, "\b")
		}
	}
}

func CompressString(input, repl string) (string, []string) {
	pattern := `\x1b\[([0-9;]*)[a-zA-Z]`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(input, -1)
	return re.ReplaceAllString(input, repl), matches
}

func DecompressString(modified string, outputArray []string, c string) string {
	parts := strings.Split(modified, c)
	reconstructed := ""
	for i, part := range parts {
		reconstructed += part
		if i < len(outputArray) {
			reconstructed += outputArray[i]
		}
	}
	return reconstructed
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

func GetTerminalSize() (int, int) {
	w, h, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	return w, h
}
