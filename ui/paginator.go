package ui

import (
	"fmt"
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/isKey"
)

type PaginatorParams struct {
	Style  DivStyle
	Width  int
	Height int
}

type PaginatorStyle struct {
	Active   string `default:"\033[32m"`
	Inactive string `default:""`
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

func (p *PaginatorComponent) Render() string {
	width, _ := osui.GetTerminalSize()
	pgs := strings.Repeat(" ", (width-len(p.Components))/2)
	osui.UseStyle(p.Style)
	frame := osui.NewFrame(p.Data.Width, p.Data.Height)
	for i, c := range p.Components {
		data := c.GetComponentData()
		if i == p.ActiveComponent {
			pgs += p.Style.Active + "•" + colors.Reset
			data.IsActive = true
		} else {
			pgs += colors.Reset + p.Style.Inactive + "•"
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
	return fmt.Sprintf("%s\n%s", pgs, colors.Reset+strings.Join(frame, "\n"))
}

func (p *PaginatorComponent) Update(key string) bool {
	if key == "\x1b[Z" {
		if p.ActiveComponent > 0 {
			p.updateActive(p.ActiveComponent - 1)
		} else {
			p.updateActive(len(p.Components) - 1)
		}
	} else if isKey.Tab(key) {
		if p.ActiveComponent < len(p.Components)-1 {
			p.updateActive(p.ActiveComponent + 1)
		} else {
			p.updateActive(0)
		}
	} else if isKey.Escape(key) {
		fmt.Print("\n\n")
		return true
	} else {
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
		p.ActiveComponent = newIndex
	}
}

func (b *PaginatorComponent) Params(p interface{}) *PaginatorComponent {
	param := p.(PaginatorParams)
	b.Style = osui.SetDefaults(&param.Style).(*PaginatorStyle)
	return b
}

func Paginator(pages ...osui.Component) *PaginatorComponent {
	return &PaginatorComponent{
		Components: pages,
		Style:      osui.SetDefaults(&PaginatorStyle{}).(*PaginatorStyle),
	}
}
