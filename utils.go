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

func renderLine(frameLine, line string, x int) string {
	result := []rune(frameLine) // Convert frameLine to a slice of runes for mutability
	visibleCount := 0           // Track visible characters in frameLine

	for _, c := range line {
		// Skip over non-visible characters when counting the frameLine length
		for visibleCount < len(result) {
			if result[visibleCount] == '\t' || result[visibleCount] == '\b' {
				visibleCount++ // Skip these characters without counting them
			} else {
				break
			}
		}

		// If visibleCount exceeds the length of frameLine, stop
		if x+visibleCount >= len(result) {
			break
		}

		// Overlay the character from the line onto frameLine at the correct position
		result[x+visibleCount] = c
		visibleCount++ // Increment visible character count
	}

	return string(result) // Convert the result back to a string
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

func DebugRender(c Component) {
	s := NewScreen(c)
	s.CustomRender = func() {

	}
	// data := c.GetComponentData()
	fmt.Println(colors.Bold + "---OBJ RENDER---" + colors.Reset)
	for _, line := range strings.Split(c.Render(), "\n") {
		fmt.Printf("%#v\n", line)
	}
	s.Run()
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
