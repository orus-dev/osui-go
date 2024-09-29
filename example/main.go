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
			Text("Welcome to the example! The blue text on the bottom show the keybindings"),
			WithPosition(0, h-2, Text(colors.Blue+"Next Page: Tab. Previous Page: shift + tab")),
		),

		// Page 2
		Div(
			WithPosition(0, 0, Text("This blue square is a div. A div stores multiple components into one")),
			WithPosition((w-30)/2, (h-4)/2,
				WithSize(30, 5,
					Div(WithPosition(4, 1, Button("Hello, World!").Params(BtnParams{Width: 19}))).Params(DivParams{Style: DivStyle{Outline: colors.Blue}}),
				)),
			WithPosition(0, h-2, Text(colors.Blue+"Navigate: Ctrl+(W/A/S/D)")),
		),

		// Page 3
		Div(
			WithPosition(0, 0, Text("This is a InputBox. It takes user input just like a GUI. "+colors.Red+colors.Bold+"NOTE: This InputBox is inside a div")),
			WithPosition((w-32)/2, (h-3)/2, InputBox(30)),
			WithPosition(0, h-2, Text(colors.Blue+"Navigate: W/S. Select: Enter. Exit Menu: Q. ")),
		),

		// Page 4
		Div(
			WithPosition(0, 0, Text("This is a Button. To click it press Enter. "+colors.Red+colors.Bold+"NOTE: This Button is inside a div")),
			WithPosition((w-20)/2, (h-3)/2, Button("This is a button")),
			WithPosition(0, h-2, Text(colors.Blue+"Navigate: Ctrl+(W/A/S/D). Click Button: Enter")),
		),

		// Page 5
		Div(
			WithPosition(0, 0, Text("This is a Menu. It prompts you to select a item. "+colors.Red+colors.Bold+"NOTE: This Menu is inside a div")),
			WithPosition(0, 2, Menu("Item 1", "Item 2", "Item 3", "Item 4", "Item 5").Params(MenuParams{})),
			WithPosition(0, h-2, Text(colors.Blue+"Navigate: Ctrl+(W/A/S/D). Navigate Menu: W/S. Select: Enter. Exit Menu: Q. ")),
		),
	)
}
