package keys

const (
	Enter     = "Enter"
	Tab       = "Tab"
	ShiftTab  = "ShiftTab"
	Backspace = "Backspace"
	Escape    = "Escape"
	Up        = "Up"
	Down      = "Down"
	Left      = "Left"
	Right     = "Right"
	CtrlW     = "CtrlW"
	CtrlS     = "CtrlS"
	CtrlA     = "CtrlA"
	CtrlD     = "CtrlD"
	Error     = "Error"
)

func Char(key string) string {
	return "Char(" + key + ")"
}

func GetKeyName(key string) string {
	switch key {
	case "\r", "\n":
		return Enter
	case "\x7f", "\b":
		return Backspace
	case "\t":
		return Tab
	case "\x1b[Z":
		return ShiftTab
	case "\x1b":
		return Escape
	case "\x1b[A":
		return Up
	case "\x1b[B":
		return Down
	case "\x1b[C":
		return Right
	case "\x1b[D":
		return Left
	case "\x17":
		return CtrlW
	case "\x13":
		return CtrlS
	case "\x01":
		return CtrlA
	case "\x04":
		return CtrlD
	default:
		return Char(key)
	}
}
