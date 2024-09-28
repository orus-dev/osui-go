package ui

import (
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/isKey"
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
		data.DefaultColor = colors.Combine(d.Style.Background, d.Style.Foreground)
		data.Screen = d.Data.Screen
		osui.RenderOnFrame(c, &frame)
	}
	for i, f := range frame {
		frame[i] = colors.Combine(d.Style.Foreground, d.Style.Background) + f + colors.Reset
	}
	return strings.Join(frame, "\n")
}

func (d *DivComponent) Update(key string) bool {
	if isKey.Up(key) {
		d.updateActive(d.ActiveComponent - 1)
	} else if isKey.Down(key) {
		d.updateActive(d.ActiveComponent + 1)
	} else {
		if len(d.Components) > 0 {
			d.Components[d.ActiveComponent].GetComponentData().IsActive = d.Data.IsActive
			if d.Components[d.ActiveComponent].Update(key) {
				if d.ActiveComponent < len(d.Components)-1 {
					d.updateActive(d.ActiveComponent + 1)
				} else {
					return true
				}
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
	return osui.NewComponent(&DivComponent{
		Components: components,
		Style:      osui.SetDefaults(&DivStyle{}).(*DivStyle),
	}).(*DivComponent)
}
