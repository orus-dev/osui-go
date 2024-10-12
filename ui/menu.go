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
			res = append(res, cursor+colors.Combine(m.Data.Style.SelectedBackground, m.Data.Style.SelectedBackground)+item+colors.Reset+m.Data.DefaultColor)
		} else {
			res = append(res, empty+colors.Combine(m.Data.Style.Foreground, m.Data.Style.Background)+item+colors.Reset+m.Data.DefaultColor)
		}
	}

	return strings.Join(res, "\n")
}

func (m *MenuComponent) Update(key string) bool {
	if f, ok := m.Data.Keys["up"]; ok && f(key) {
		if m.SelectedItem > 0 {
			m.SelectedItem--
		} else {
			m.SelectedItem = len(m.Items) - 1
		}
	} else if f, ok := m.Data.Keys["down"]; ok && f(key) {
		if m.SelectedItem+1 < len(m.Items) {
			m.SelectedItem++
		} else {
			m.SelectedItem = 0
		}
	} else if f, ok := m.Data.Keys["select"]; ok && f(key) {
		m.Data.OnClick()
		return true
	}

	return false
}

func Menu(param osui.Param, items ...string) *MenuComponent {
	param.SetDefaultBindings(map[string]func(string) bool{
		"up":     keys.Up,
		"down":   keys.Down,
		"select": keys.Enter,
	})
	return param.UseParam(&MenuComponent{
		Items: items,
	}).(*MenuComponent)
}
