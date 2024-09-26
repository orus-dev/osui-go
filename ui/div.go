package ui

import (
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
)

type DivStyle struct {
	Background string `defaults:"" type:"bg"`
	Foreground string `defaults:"" type:"fg"`
}

type DivComponent struct {
	Data            osui.ComponentData
	Style           *DivStyle
	Components      []osui.Component
	ActiveComponent int
}

func (d *DivComponent) GetComponentData() *osui.ComponentData {
	return &d.Data
}

func (d *DivComponent) Render() string {
	osui.UseStyle(d.Style)
	frame := osui.NewFrame(d.Data.Width, d.Data.Height)
	for i, c := range d.Components {
		data := c.GetComponentData()
		if i == d.ActiveComponent {
			data.IsActive = true
		} else {
			data.IsActive = false
		}
		if data.Width == 0 {
			data.Width = d.Data.Width
		}
		if data.Height == 0 {
			data.Height = d.Data.Height
		}
		data.DefaultColor = colors.Combine(d.Style.Foreground, d.Style.Background)
		data.Screen = d.Data.Screen
		osui.RenderOnFrame(c, &frame)
	}
	for i, f := range frame {
		frame[i] = colors.Reset + d.Data.DefaultColor + colors.Combine(d.Style.Foreground, d.Style.Background) + f + colors.Reset + d.Data.DefaultColor
	}
	return strings.Join(frame, "\n")
}

func (d *DivComponent) Run() {
	for _, c := range d.Components {
		data := c.GetComponentData()
		data.Screen = d.Data.Screen
	}
	for {
		key, _ := osui.ReadKey()
		if d.Update(key) {
			return
		}
	}
}

func (d *DivComponent) Update(key string) bool {
	switch key {

	case osui.Key.Up:
		d.updateActive(d.ActiveComponent - 1)
	case osui.Key.Down:
		d.updateActive(d.ActiveComponent + 1)
	default:
		if len(d.Components) > 0 {
			d.Components[d.ActiveComponent].GetComponentData().IsActive = d.Data.IsActive
			if d.Components[d.ActiveComponent].Update(key) {
				return true
				// if d.ActiveComponent < len(d.Components)-1 {
				// 	d.updateActive(d.ActiveComponent + 1)
				// } else {
				// 	return true
				// }
			}
		}
	}
	return false
}

func (d *DivComponent) SetStyle(c interface{}) {
	d.Style = osui.SetDefaults(c.(*DivStyle)).(*DivStyle)
}

func (d *DivComponent) updateActive(newIndex int) {
	if newIndex >= 0 && newIndex < len(d.Components) && len(d.Components) > 0 {
		d.ActiveComponent = newIndex
	}
}

func Div(components ...osui.Component) *DivComponent {
	return &DivComponent{
		Components: components,
		Style:      osui.SetDefaults(&DivStyle{}).(*DivStyle),
	}
}
