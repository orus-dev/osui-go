<h1 style="text-align:center;"><img src="assets/osui.png" width="170px"></img></h1>
<h1 style="text-align:center;"></img>A Component-based TUI library written in go</h1>
<p style="text-align:center;"><img src="assets/osui.gif"></img></p>


# Hello World
```go
package main

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/ui"
)

func main() {
    app := ui.Div(
        ui.Text("Hello, World!")
    )

    screen := osui.NewScreen(app)
    screen.Run()
}
```

# [Documentation](https://github.com/orus-dev/osui/wiki)
