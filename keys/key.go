package keys

func Enter(s string) bool {
	return s == "\r" || s == "\n"
}

func Tab(s string) bool {
	return s == "\t"
}

func ShiftTab(s string) bool {
	return s == "\x1b[Z"
}

func Backspace(s string) bool {
	return s == "\x7f" || s == "\b"
}

func Escape(s string) bool {
	return s == "\x1b"
}

func Up(s string) bool {
	return s == "\x1b[A"
}

func Down(s string) bool {
	return s == "\x1b[B"
}

func Right(s string) bool {
	return s == "\x1b[C"
}

func Left(s string) bool {
	return s == "\x1b[D"
}

func Char(s string, s1 string) bool {
	return s == s1
}

func CtrlW(s string) bool {
	return s == "\x17" 
}

func CtrlS(s string) bool {
	return s == "\x13" 
}

func CtrlA(s string) bool {
	return s == "\x01" 
}

func CtrlD(s string) bool {
	return s == "\x04" 
}