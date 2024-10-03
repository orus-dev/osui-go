package main

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/ui"
)

var students []string = []string{}

func main() {
	screen := osui.NewScreen(App())
	screen.Run()
}

func App() *ui.DivComponent {
	return renderRoot()
}

func renderRoot() *ui.DivComponent {
	var root ui.Id[*ui.DivComponent]

	return root.Id(ui.Div(
		ui.WithPosition(0, 0, ui.Text("student managment system (terminal edition permium) made by fefek")),
		ui.WithPosition(0, 2, ui.Button("add student")),
		ui.WithPosition(25, 2, ui.Button("delete student")),
		ui.WithPosition(50, 2, ui.Button("info")),
		ui.WithPosition(0, 5, ui.Text("students")),
		ui.WithPosition(0, 7, ui.Menu(students...)),
	))
}
