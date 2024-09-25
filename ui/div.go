package ui

import (
	"strings"

	"github.com/orus-dev/osui"
)

type DivComponent struct {
	Data            osui.ComponentData
	Components      []osui.Component
	ActiveComponent int
}

func (p *DivComponent) GetComponentData() *osui.ComponentData {
	return &p.Data
}

func (p DivComponent) Render() string {
	frame := osui.NewFrame(p.Data.Width, p.Data.Height)
	for _, c := range p.Components {
		data := c.GetComponentData()
		data.Width = p.Data.Width
		data.Height = p.Data.Height
		data.Screen = p.Data.Screen
		osui.RenderOnFrame(c, &frame)
	}
	return strings.Join(frame, "\n")
}

func (p *DivComponent) Run() {
	for _, c := range p.Components {
		d := c.GetComponentData()
		d.Screen = p.Data.Screen
	}
	for {
		key, _ := osui.ReadKey()
		if p.Update(key) {
			return
		}
	}
}

func (p *DivComponent) Update(key string) bool {
	switch key {

	case "":
		p.updateActive(p.ActiveComponent)
	case osui.Key.Up:
		p.updateActive(p.ActiveComponent - 1)
	case osui.Key.Down:
		p.updateActive(p.ActiveComponent + 1)
	default:
		if len(p.Components) > 0 {
			p.Components[p.ActiveComponent].GetComponentData().IsActive = p.Data.IsActive
			if p.Components[p.ActiveComponent].Update(key) {
				if p.ActiveComponent < len(p.Components)-1 {
					p.updateActive(p.ActiveComponent + 1)
				} else {
					return true
				}
			}
		}
	}
	return false
}

func (p *DivComponent) updateActive(newIndex int) {
	if newIndex >= 0 && newIndex < len(p.Components) && len(p.Components) > 0 {
		p.Components[p.ActiveComponent].GetComponentData().IsActive = false
		p.ActiveComponent = newIndex
		p.Components[p.ActiveComponent].GetComponentData().IsActive = p.Data.IsActive
		p.Components[p.ActiveComponent].Update("")
	}
}

func Div(components ...osui.Component) *DivComponent {
	return &DivComponent{
		Components: components,
	}
}
