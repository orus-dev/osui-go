package main

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/cutils"
	"github.com/orus-dev/osui/ui"
)

func main() {
	text := colors.Red + "This is a div!"
	width, h := osui.GetTerminalSize()
	paginator := ui.Paginator(
		ui.Div(
			ui.Text("Welcome to the example! Press tab to go to the next page or press tab to go to a previous page"),
		),
		ui.Div(cutils.WithPosition((width-30)/2, (h-4)/2, cutils.WithSize(30, 4, cutils.WithStyle(&ui.DivStyle{Background: colors.Blue}, ui.Div(ui.Text(text)))))),
		// ui.Div(
		// 	ui.Text("This is a InputBox. It takes user input and"),
		// 	cutils.WithPosition(0, 2, ui.InputBox(20)),
		// ),
	)
	screen := osui.NewScreen(paginator)
	screen.Run()
}
