package ui

import (
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/isKey"
)

type DivParams struct {
	Style  DivStyle
	Width  int
	Height int
}

type DivStyle struct {
	Background string `default:"" type:"bg"`
	Foreground string `default:"" type:"fg"`
	Outline    string `default:"" type:"fg"`
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
	frame := osui.NewFrame(d.Data.Width-2, d.Data.Height-2)
	for i, c := range d.Components {
		data := c.GetComponentData()
		if i == d.ActiveComponent {
			data.IsActive = d.Data.IsActive
		} else {
			data.IsActive = false
		}
		if data.Width == 0 {
			data.Width = d.Data.Width - 2
		}
		if data.Height == 0 {
			data.Height = d.Data.Height - 2
		}
		data.DefaultColor = colors.Combine(d.Style.Background, d.Style.Foreground)
		data.Screen = d.Data.Screen
		osui.RenderOnFrame(c, &frame)
	}
	if d.Style.Outline == "" {
		for i, f := range frame {
			frame[i] = colors.Combine(d.Style.Foreground, d.Style.Background) + f + colors.Reset
		}
		return strings.Join(frame, "\n")
	} else {
		for i, f := range frame {
			frame[i] = d.Style.Outline + "│" + colors.Combine(d.Style.Foreground, d.Style.Background) + f + colors.Reset + d.Style.Outline + "│" + colors.Reset
		}
	}
	return " " + d.Style.Outline + strings.Repeat("_", d.Data.Width-2) + colors.Reset + "\n" + strings.Join(frame, "\n") + "\n " + d.Style.Outline + strings.Repeat("‾", d.Data.Width-2) + colors.Reset
}

func (d *DivComponent) Update(key string) bool {
	if isKey.CtrlW(key) {
		d.updateActive(findClosestComponent(d.Components, d.ActiveComponent, "up"))
	} else if isKey.CtrlS(key) {
		d.updateActive(findClosestComponent(d.Components, d.ActiveComponent, "down"))
	} else if isKey.CtrlA(key) {
		d.updateActive(findClosestComponent(d.Components, d.ActiveComponent, "left"))
	} else if isKey.CtrlD(key) {
		d.updateActive(findClosestComponent(d.Components, d.ActiveComponent, "right"))
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

func (b *DivComponent) Params(param DivParams) *DivComponent {
	b.Style = osui.SetDefaults(&param.Style).(*DivStyle)
	b.Data.Width = osui.LogicValueInt(param.Width == 0, 20, param.Width)
	return b
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
