package main

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/ui"
)

func main() {
	screen := osui.NewScreen(App())
	screen.Run()
}

func App() *ui.PaginatorComponent {
	w, h := osui.GetTerminalSize()

	var keybindings ui.Class[*ui.TextComponent]
	keybindings.SetProperties(func(tc *ui.TextComponent) *ui.TextComponent {
		tc.Text = colors.Blue + tc.Text
		tc.Data.Y = h - 3
		return tc
	})

	return ui.Paginator(
		// Page 1
		ui.Div(
			ui.Text("Welcome to the example! The blue text on the bottom show the keybindings"),
			keybindings.Class(ui.Text("Next Page: Tab. Previous Page: shift + tab")),
		),

		// Page 2
		ui.Div(
			ui.WithPosition(0, 0, ui.Text("This blue square is a div. A div stores multiple components into one")),
			ui.WithPosition((w-30)/2, (h-4)/2,
				ui.WithSize(32, 7,
					ui.Div(ui.WithPosition(2, 2, ui.Text("This text is inside a div!"))).Params(ui.DivParams{Style: ui.DivStyle{Outline: colors.Blue, Foreground: colors.Red}}),
				)),
			keybindings.Class(ui.Text("Navigate: Ctrl+(W/A/S/D)")),
		),

		// Page 3
		ui.Div(
			ui.WithPosition(0, 0, ui.Text("This is a InputBox. It takes user input just like a GUI. "+colors.Red+colors.Bold+"NOTE: This InputBox is inside a div")),
			ui.WithPosition((w-32)/2, (h-3)/2, ui.InputBox(30)),
			keybindings.Class(ui.Text("Navigate: Ctrl+(W/A/S/D). Finish: Enter.")),
		),

		// Page 4
		ui.Div(
			ui.WithPosition(0, 0, ui.Text("This is a Button. To click it press Enter. "+colors.Red+colors.Bold+"NOTE: This Button is inside a div")),
			ui.WithPosition((w-20)/2, (h-3)/2, ui.Button("This is a button")),
			keybindings.Class(ui.Text("Navigate: Ctrl+(W/A/S/D). Click Button: Enter")),
		),

		// Page 5
		ui.Div(
			ui.WithPosition(0, 0, ui.Text("This is a Menu. It prompts you to select an item. "+colors.Red+colors.Bold+"NOTE: This Menu is inside a div")),
			ui.WithPosition(0, 2, ui.Menu("Item 1", "Item 2", "Item 3", "Item 4", "Item 5").Params(ui.MenuParams{OnSelected: func(mc *ui.MenuComponent, b bool) {}})),
			keybindings.Class(ui.Text("Navigate: Ctrl+(W/A/S/D). Navigate Menu: W/S. Select: Enter. Exit Menu: Q. ")),
		),
	)
}
