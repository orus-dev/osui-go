package ui

import (
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/isKey"
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
	ctx := osui.Context{Key: key}
	m.Data.OnUpdate(&ctx)
	if ctx.Response != 0 {
		return false
	}

	if isKey.Down(key) {
		if m.SelectedItem+1 < len(m.Items) {
			m.SelectedItem++
		} else {
			m.SelectedItem = 0
		}
	} else if isKey.Up(key) {
		if m.SelectedItem > 0 {
			m.SelectedItem--
		} else {
			m.SelectedItem = len(m.Items) - 1
		}
	} else if isKey.Enter(key) {
		m.Data.OnClick()
		return true
	}
	return false
}

func Menu(param osui.Param, items ...string) *MenuComponent {
	return param.UseParam(&MenuComponent{
		Items: items,
	}).(*MenuComponent)
}
