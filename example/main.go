package main

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/cutils"
	"github.com/orus-dev/osui/ui"
)

func main() {
	w, h := osui.GetTerminalSize()
	paginator := ui.Paginator(
		ui.Div(
			ui.Text("Welcome to the example! Press tab to go to the next page or press shift + tab to go to a previous page"),
		),
		ui.Div(
			cutils.WithPosition(0, 0, ui.Text("This blue square is a div. A div stores multiple components into one. To navigate to each of those components you can use arrow keys ↑↓")),
			cutils.WithPosition((w-30)/2, (h-4)/2,
				cutils.WithSize(30, 4,
					cutils.WithStyle(&ui.DivStyle{Background: colors.Blue, Foreground: colors.Red},
						ui.Div(cutils.WithPosition(0, 0, ui.Text("Hello, World!"))),
					)))),
		ui.Div(
			cutils.WithPosition(0, 0, ui.Text("This is a InputBox. It takes user input just like a GUI. "+colors.Red+colors.Bold+"NOTE: This InputBox is inside a div")),
			cutils.WithPosition((w-32)/2, (h-3)/2, ui.InputBox(30)),
		),
	)
	screen := osui.NewScreen(paginator)
	screen.Run()
}
