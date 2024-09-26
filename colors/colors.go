package colors

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	Reset     = "\x1b[0m"
	Bold      = "\x1b[1m"
	Underline = "\x1b[4m"
	Italic    = "\x1b[3m"
	Reverse   = "\x1b[7m"
	Strike    = "\x1b[9m"

	Black   = "\x1b[30m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan    = "\x1b[36m"
	White   = "\x1b[37m"
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
