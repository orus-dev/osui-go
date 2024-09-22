# osui
A Component-based TUI library written in go!


```go
package main

import (
	"github.com/orus-dev/osui/osui"
	"github.com/orus-dev/osui/osui/ui"
)

func main() {
	components := osui.NewRender()
	components.Add(ui.Text("Hello, World!"))
	components.Render()
}
```