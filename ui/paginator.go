package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/styles"
)

type PaginatorStyle struct {
	Active string `default:"\033[32m"`
}

type PaginatorComponent struct {
	Data            osui.ComponentData
	Style           *PaginatorStyle
	Components      []osui.Component
	ActiveComponent int
}

func (p *PaginatorComponent) GetComponentData() *osui.ComponentData {
	return &p.Data
}

func (p PaginatorComponent) Render() string {
	pgs := strings.Repeat(" ", p.Data.Width/2)
	for page, c := range p.Components {
		if page == p.ActiveComponent {
			pgs += p.Style.Active + "•" + styles.Reset
		} else {
			pgs += styles.Reset + "•"
		}
		d := c.GetComponentData()
		if d.Screen == nil {
			d.Screen = p.Data.Screen
		}
		d.Height = p.Data.Height - 1
		d.Width = p.Data.Width
	}
	frame := osui.NewFrame(p.Data.Width, p.Data.Height-3)
	osui.RenderOnFrame(p.Components[p.ActiveComponent], &frame)
	return fmt.Sprintf("%s\n%s", pgs, styles.Reset+strings.Join(frame, "\n"))
}

func (p *PaginatorComponent) Update(key string) bool {
	switch key {
	case "":
		p.updateActive(p.ActiveComponent)
	case "\x1b[Z":
		if p.ActiveComponent > 0 {
			p.updateActive(p.ActiveComponent - 1)
		} else {
			p.updateActive(len(p.Components) - 1)
		}
	case osui.Key.Tab:
		if p.ActiveComponent < len(p.Components)-1 {
			p.updateActive(p.ActiveComponent + 1)
		} else {
			p.updateActive(0)
		}
	case osui.Key.Escape:
		fmt.Print("\n\n")
		return true
	default:
		if len(p.Components) > 0 {
			p.Components[p.ActiveComponent].GetComponentData().IsActive = p.Data.IsActive
			if p.Components[p.ActiveComponent].Update(key) {
				if p.ActiveComponent < len(p.Components)-1 {
					p.updateActive(p.ActiveComponent + 1)
				} else {
					p.updateActive(0)
				}
			}
		}
	}
	return false
}

func (p *PaginatorComponent) updateActive(newIndex int) {
	if newIndex >= 0 && newIndex < len(p.Components) && len(p.Components) > 0 {
		p.Components[p.ActiveComponent].GetComponentData().IsActive = false
		p.ActiveComponent = newIndex
		p.Components[p.ActiveComponent].GetComponentData().IsActive = p.Data.IsActive
		p.Components[p.ActiveComponent].Update("")
	}
}

func Paginator(pages ...osui.Component) *PaginatorComponent {
	return &PaginatorComponent{
		Components: pages,
		Style:      osui.SetDefaults(&PaginatorStyle{}).(*PaginatorStyle),
	}
}
