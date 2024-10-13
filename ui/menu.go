package ui

import (
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/keys"
)

type MenuComponent struct {
	Data         osui.ComponentData
	Items        []string
	SelectedItem int
}

func (m *MenuComponent) GetComponentData() *osui.ComponentData {
	return &m.Data
}

func (m *MenuComponent) Render() string {
	m.Data.Style.UseStyle()
	res := []string{}

	cursor := m.Data.Style.Cursor
	if !m.Data.IsActive {
		cursor = m.Data.Style.InactiveCursor
	}
	d, _ := osui.CompressString(m.Data.Style.Cursor)
	empty := strings.Repeat(" ", len(d))

	for i, item := range m.Items {
		if i == m.SelectedItem {
			res = append(res, colors.Reset+cursor+m.Data.Style.SelectedBackground+m.Data.Style.SelectedForeground+item+"\x1b[0m ")
		} else {
			res = append(res, colors.Reset+empty+m.Data.Style.Foreground+m.Data.Style.Background+item)
		}
	}

	return strings.Join(res, "\n")
}

func (m *MenuComponent) Update(ctx osui.UpdateContext) bool {
	if ctx.UpdateKind == osui.UpdateKindKey {
		switch m.Data.Keys[ctx.Key.Name] {
		case "up":
			if m.SelectedItem > 0 {
				m.SelectedItem--
			} else {
				m.SelectedItem = len(m.Items) - 1
			}
		case "down":
			if m.SelectedItem+1 < len(m.Items) {
				m.SelectedItem++
			} else {
				m.SelectedItem = 0
			}
		case "select":
			m.Data.OnClick()
			return true
		case "exit":
			return true
		}
	}

	return false
}

func Menu(param osui.Param, items ...string) *MenuComponent {
	param.SetDefaultBindings(map[string]string{
		keys.Up:    "up",
		keys.Down:  "down",
		keys.Enter: "select",
	})
	return param.UseParam(&MenuComponent{
		Items: items,
	}).(*MenuComponent)
}
