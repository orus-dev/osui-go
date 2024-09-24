package main

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/styles"
	"github.com/orus-dev/osui/ui"
)

func main() {
	colors := ui.Text(
		styles.Red + "Red!\n" + styles.Green + "Green!\n" + styles.Blue + "Blue!\n",
	)
	input := ui.InputBox(30)
	paginator := ui.Paginator(colors, input)
	paginator.UpdatePages[1] = input.Update
	screen := osui.NewScreen(paginator)
	screen.Render()
	paginator.Run()
}
