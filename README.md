# osui
A Component-based TUI library written in go!

![OSUI](osui.gif)


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
