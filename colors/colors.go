package colors

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
	Italic    = "\033[3m"
	Reverse   = "\033[7m"
	Strike    = "\033[9m"

	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

func Rgb(r, g, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func AsBg(s string) string {
	re, _ := regexp.Compile(`\x1b\[3([0-9]+)`)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		groups := re.FindStringSubmatch(match)
		if len(groups) >= 2 {
			return fmt.Sprintf("\x1b[4%s", groups[1])
		}
		return match
	})
}


func AsFg(s string) string {
	re, _ := regexp.Compile(`\x1b\[4([0-9]+)`)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		groups := re.FindStringSubmatch(match)
		if len(groups) >= 2 {
			return fmt.Sprintf("\x1b[3%s", groups[1])
		}
		return match
	})
}

func Combine(s, s1 string) string {
	if strings.Contains(s, Reset) {
		return s + s1
	} else {
		return s1 + s
	}
}
