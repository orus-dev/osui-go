package main

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/styles"
	"github.com/orus-dev/osui/ui"
)

func main() {
	colors := ui.Text(
		styles.Red + "Red!\n" + styles.Green + "Green!\n" + styles.Blue + "Blue!" + styles.Reset,
	)
	paginator := ui.Paginator(ui.Div(colors), ui.Div(ui.InputBox(30)))
	screen := osui.NewScreen(paginator)
	screen.Run()
}
