package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/keys"
)

type PaginatorComponent struct {
	Data            osui.ComponentData
	Components      []osui.Component
	ActiveComponent int
}

func (p *PaginatorComponent) GetComponentData() *osui.ComponentData {
	return &p.Data
}

func (p *PaginatorComponent) Render() string {
	width, _ := osui.GetTerminalSize()
	pgs := strings.Repeat(" ", (width-len(p.Components))/2)
	p.Data.Style.UseStyle()
	frame := osui.NewFrame(p.Data.Width, p.Data.Height)
	for i, c := range p.Components {
		data := c.GetComponentData()
		if i == p.ActiveComponent {
			pgs += p.Data.Style.SelectedForeground + data.Style.SelectedBackground + "•" + colors.Reset
			data.IsActive = true
		} else {
			pgs += colors.Reset + p.Data.Style.Foreground + data.Style.Background + "•"
			data.IsActive = false
		}
		if data.Width == 0 {
			data.Width = p.Data.Width
		}
		if data.Height == 0 {
			data.Height = p.Data.Height
		}
		data.DefaultColor = p.Data.DefaultColor
		data.Screen = p.Data.Screen
	}
	osui.RenderOnFrame(p.Components[p.ActiveComponent], &frame)
	for i, f := range frame {
		frame[i] = colors.Reset + p.Data.DefaultColor + f + colors.Reset
	}
	return fmt.Sprintf("%s\n%s", pgs+" ", colors.Reset+strings.Join(frame, "\n"))
}

func (p *PaginatorComponent) Update(ctx osui.UpdateContext) bool {
	if ctx.UpdateKind == osui.UpdateKindKey {
		switch p.Data.Keys[ctx.Key.Name] {
		case "previous":
			if p.ActiveComponent > 0 {
				p.updateActive(p.ActiveComponent - 1)
			} else {
				p.updateActive(len(p.Components) - 1)
			}
		case "next":
			if p.ActiveComponent < len(p.Components)-1 {
				p.updateActive(p.ActiveComponent + 1)
			} else {
				p.updateActive(0)
			}
		case "exit":
			p.Data.OnClick()
			fmt.Print("\n\n")
			return true
		default:
			if len(p.Components) > 0 {
				p.Components[p.ActiveComponent].GetComponentData().IsActive = p.Data.IsActive
				if p.Components[p.ActiveComponent].Update(ctx) {
					if p.ActiveComponent < len(p.Components)-1 {
						p.updateActive(p.ActiveComponent + 1)
					} else {
						p.updateActive(0)
					}
				}
			}
		}
	}
	return false
}

func (p *PaginatorComponent) updateActive(newIndex int) {
	if newIndex >= 0 && newIndex < len(p.Components) && len(p.Components) > 0 {
		p.ActiveComponent = newIndex
	}
}

func Paginator(param osui.Param, pages ...osui.Component) *PaginatorComponent {
	param.SetDefaultBindings(map[string]string{
		keys.ShiftTab: "previous",
		keys.Tab:      "next",
		keys.Escape:   "exit",
	})
	return param.UseParam(&PaginatorComponent{
		Components: pages,
	}).(*PaginatorComponent)
}
