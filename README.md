# osui
A Component-based TUI library written in go!


```go
package main

import (
	"github.com/orus-dev/osui/osui"
	"github.com/orus-dev/osui/osui/ui"
)

func main() {

	app := ui.Div(
		ui.Text("Hello, World!")
	)

	screen := osui.NewScreen(app)
	screen.Run()
}
```