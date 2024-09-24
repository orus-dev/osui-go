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
	Data        osui.ComponentData
	Style       *PaginatorStyle
	Pages       []osui.Component
	UpdatePages map[int]func(string) bool
	ActivePage  int
}

func (p *PaginatorComponent) GetComponentData() *osui.ComponentData {
	return &p.Data
}

func (p PaginatorComponent) Render() string {
	pgs := strings.Repeat(" ", p.Data.Screen.Width/2)
	for page := range p.Pages {
		if page == p.ActivePage {
			pgs += p.Style.Active + "•" + styles.Reset
		} else {
			pgs += styles.Reset + "•"
		}
	}
	frame := osui.NewFrame(p.Data.Screen.Width, p.Data.Screen.Height-3)
	osui.RenderOnFrame(p.Pages[p.ActivePage], &frame)
	return fmt.Sprintf("%s\n%s", pgs, styles.Reset+strings.Join(frame, "\n"))
}

func (p *PaginatorComponent) Run() {
	for _, c := range p.Pages {
		d := c.GetComponentData()
		if d.Screen == nil {
			d.Screen = p.Data.Screen
		}
	}
	for {
		key, _ := osui.ReadKey()
		switch key {
		case "\x1b[Z":
			if p.ActivePage > 0 {
				p.ActivePage--
			} else {
				p.ActivePage = len(p.Pages) - 1
			}
		case osui.Key.Tab:
			if p.ActivePage < len(p.Pages)-1 {
				p.ActivePage++
			} else {
				p.ActivePage = 0
			}
		case osui.Key.Escape:
			fmt.Print("\n\n")
			return
		default:
			for i, c := range p.UpdatePages {
				if i == p.ActivePage {
					c(key)
				}
			}
		}
		p.Data.Screen.Render()
	}
}

func Paginator(pages ...osui.Component) *PaginatorComponent {
	return &PaginatorComponent{
		Pages:       pages,
		Style:       osui.SetDefaults(&PaginatorStyle{}).(*PaginatorStyle),
		UpdatePages: make(map[int]func(string) bool),
	}
}
