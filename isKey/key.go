package isKey

func Enter(s string) bool {
	return s == "\r"
}

func Tab(s string) bool {
	return s == "\t"
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