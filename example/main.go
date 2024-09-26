package main

// IMPORTANT NOTE: Needs bugfix. "Some text here" is distorted

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/ui"
)

func main() {
	// innerDiv := ui.Div(ui.Text("ABC"))

	// innerDiv.Data.X = 4
	// innerDiv.Style.BackgroundColor = colors.Red
	// innerDiv.Data.Width = 5
	// innerDiv.Data.Height = 3

	input := ui.InputBox(30)
	input.Style.Cursor = colors.Green
	input.Style.Background = colors.Red
	input.Style.Outline = colors.Blue

	someText := ui.Text("Some text here")
	someText.Data.X = 40
	someText.Data.Y = 1

	div := ui.Div(input, someText)

	screen := osui.NewScreen(div)
	screen.Run()
}
