package osui

import (
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
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
	pattern := `\x1b\[([0-9;]*)m`
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
		panic("setDefaults: expected a pointer to a struct")
	}
	val := v.Elem()
	if val.Kind() != reflect.Struct {
		panic("setDefaults: expected a pointer to a struct")
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
