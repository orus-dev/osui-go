package main

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/ui"
)

func main() {
	screen := osui.NewScreen()
	my_btn := ui.Menu(osui.Param{Style: osui.Style{Cursor: "> "}},
		"item 1", "item 2", "item 3", "item 4",
	)
	screen.Run(my_btn)
}
