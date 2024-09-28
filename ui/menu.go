package ui

import (
	"strings"

	"github.com/orus-dev/osui"
	"github.com/orus-dev/osui/colors"
	"github.com/orus-dev/osui/isKey"
)

type MenuStyle struct {
	Fg             string `default:"" type:"fg"`
	Bg             string `default:"" type:"bg"`
	SelectedFg     string `default:"\x1b[34m" type:"fg"`
	SelectedBg     string `default:"" type:"bg"`
	Cursor         string `default:"> " type:"fg"`
	CursorInactive string `default:"  " type:"fg"`
}

type MenuComponent struct {
	Data         osui.ComponentData
	Style        *MenuStyle
	Items        []string
	SelectedItem int
	onSelected   func(*MenuComponent, bool)
}

func (m *MenuComponent) GetComponentData() *osui.ComponentData {
	return &m.Data
}

func (m *MenuComponent) Render() string {
	osui.UseStyle(m.Style)
	res := []string{}

	cursor := m.Style.Cursor
	if !m.Data.IsActive {
		cursor = m.Style.CursorInactive
	}
	d, _ := osui.CompressString(m.Style.Cursor, "")
	empty := strings.Repeat(" ", len(d))

	for i, item := range m.Items {
		if i == m.SelectedItem {
			res = append(res, cursor+colors.Combine(m.Style.SelectedFg, m.Style.SelectedBg)+item+colors.Reset+m.Data.DefaultColor)
		} else {
			res = append(res, empty+colors.Combine(m.Style.Fg, m.Style.Bg)+item+colors.Reset+m.Data.DefaultColor)
		}
	}

	return strings.Join(res, "\n")
}

func (m *MenuComponent) Update(key string) bool {
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
		m.onSelected(m, true)
		return true
	} else if isKey.Escape(key) {
		m.onSelected(m, false)
		return true
	}
	return false
}

func (m *MenuComponent) SetStyle(c interface{}) {
	m.Style = osui.SetDefaults(c.(*MenuStyle)).(*MenuStyle)
}

func (m *MenuComponent) OnSelected(f func(*MenuComponent, bool)) *MenuComponent {
	m.onSelected = f
	return m
}

func Menu(items ...string) *MenuComponent {
	return &MenuComponent{
		Items: items,
		Style: osui.SetDefaults(&MenuStyle{}).(*MenuStyle),
	}
}
