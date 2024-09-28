package main

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	. "github.com/orus-dev/osui/ui"
)

func main() {
	screen := osui.NewScreen(App())
	screen.Run()
}

func App() *PaginatorComponent {
	w, h := osui.GetTerminalSize()
	return Paginator(
		// Page 1
		Div(
			Text("Welcome to the example! Press tab to go to the next page or press shift + tab to go to a previous page"),
		),

		// Page 2
		Div(
			WithPosition(0, 0, Text("This blue square is a div. A div stores multiple components into one. To navigate to each of those components you can use arrow keys ↑↓")),
			WithPosition((w-30)/2, (h-4)/2,
				WithSize(30, 4,
					WithStyle(&DivStyle{Background: colors.Blue, Foreground: colors.Red},
						Div(WithPosition(0, 0, Text("Hello, World!"))),
					))),
		),

		// Page 3
		Div(
			WithPosition(0, 0, Text("This is a InputBox. It takes user input just like a GUI. "+colors.Red+colors.Bold+"NOTE: This InputBox is inside a div")),
			WithPosition((w-32)/2, (h-3)/2, InputBox(30)),
		),

		// Page 4
		Div(
			WithPosition(0, 0, Text("This is a Button. To click it press Enter. "+colors.Red+colors.Bold+"NOTE: This Button is inside a div")),
			WithPosition((w-20)/2, (h-3)/2, Button("This is a button", BtnParams{OnClick: func(bc *ButtonComponent) bool { return false }})),
		),
	)
}
