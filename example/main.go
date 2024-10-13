package main

import (
	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/ui"
)

func main() {
	// screen := osui.NewScreen()
	// my_btn := ui.Button(osui.Param{},
	// 	"Button",
	// )
	// screen.Run(my_btn)

	// var menu ui.Id[*ui.MenuComponent]
	// screen := osui.NewScreen()
	// screen.Run(ui.Div(osui.Param{},

	// 	menu.Id(ui.Menu(osui.Param{OnClick: func() {
	// 		screen.SetComponent(ui.Text(osui.Param{}, menu.Component.Items[menu.Component.SelectedItem]))
	// 		screen.Exit()
	// 	}},
	// 		"Item 1", "Item 24324322", "Item 366",
	// 	)),

	// ))

	// screen := osui.NewScreen()
	// my_btn := ui.InputBox(osui.Param{},
	// 	24,
	// )
	// screen.Run(my_btn)

	screen := osui.NewScreen()

	my_btn := ui.Paginator(osui.Param{OnClick: func() {screen.Exit()}},
		ui.Div(osui.Param{},
			ui.Text(osui.Param{}, "Page 1"),
		),
		ui.Div(osui.Param{},
			ui.Text(osui.Param{}, "Page 2"),
		),
	)

	screen.Run(my_btn)
}
