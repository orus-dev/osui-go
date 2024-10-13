package ui

import (
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/keys"
)

type DivComponent struct {
	Data            osui.ComponentData
	Components      []osui.Component
	ActiveComponent int
}

func (d *DivComponent) GetComponentData() *osui.ComponentData {
	return &d.Data
}

func (d *DivComponent) Render() string {
	d.Data.Style.UseStyle()
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
		data.DefaultColor = colors.Combine(d.Data.Style.Background, d.Data.Style.Foreground)
		data.Screen = d.Data.Screen
		osui.RenderOnFrame(c, &frame)
	}
	if d.Data.Style.Outline == "" {
		for i, f := range frame {
			frame[i] = colors.Combine(d.Data.Style.Foreground, d.Data.Style.Background) + f + colors.Reset
		}
		return strings.Join(frame, "\n")
	} else {
		for i, f := range frame {
			frame[i] = d.Data.Style.Outline + "│" + colors.Reset + colors.Combine(d.Data.Style.Foreground, d.Data.Style.Background) + f + colors.Reset + d.Data.Style.Outline + "│" + colors.Reset
		}
	}
	return " " + d.Data.Style.Outline + strings.Repeat("_", d.Data.Width-2) + colors.Reset + "\n" + strings.Join(frame, "\n") + "\n " + d.Data.Style.Outline + strings.Repeat("‾", d.Data.Width-2) + colors.Reset
}

func (d *DivComponent) Update(ctx osui.UpdateContext) bool {
	if ctx.UpdateKind == osui.UpdateKindKey {
		switch d.Data.Keys[ctx.Key.Name] {
		case "up":
			d.updateActive(findClosestComponent(d.Components, d.ActiveComponent, "up"))
		case "down":
			d.updateActive(findClosestComponent(d.Components, d.ActiveComponent, "down"))
		case "left":
			d.updateActive(findClosestComponent(d.Components, d.ActiveComponent, "left"))
		case "right":
			d.updateActive(findClosestComponent(d.Components, d.ActiveComponent, "right"))
		default:
			d.Components[d.ActiveComponent].GetComponentData().IsActive = d.Data.IsActive
			if d.Components[d.ActiveComponent].Update(ctx) {
				d.updateActive(findClosestComponent(d.Components, d.ActiveComponent, "down"))
			}
		}
	} else if ctx.UpdateKind == osui.UpdateKindTick {
		d.Components[d.ActiveComponent].GetComponentData().IsActive = d.Data.IsActive
		if d.Components[d.ActiveComponent].Update(ctx) {
			d.updateActive(findClosestComponent(d.Components, d.ActiveComponent, "down"))
		}
	}

	return false
}

func (d *DivComponent) updateActive(newIndex int) {
	if newIndex >= 0 && newIndex < len(d.Components) && len(d.Components) > 0 {
		d.ActiveComponent = newIndex
	}
}

func Div(param osui.Param, components ...osui.Component) *DivComponent {
	param.SetDefaultBindings(map[string]string{
		keys.CtrlW: "up",
		keys.CtrlS: "down",
		keys.CtrlA: "left",
		keys.CtrlD: "right",
	})
	return param.UseParam(&DivComponent{
		Components: components,
	}).(*DivComponent)
}
