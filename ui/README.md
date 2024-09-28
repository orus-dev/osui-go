# Component template (use rename symbol)

```go
type HelloStyle struct {
	Fg string `defaults:"" type:"fg"`
}

type HelloComponent struct {
	Data  osui.ComponentData
	Style *HelloStyle
	Name  string
}

func (h *HelloComponent) GetComponentData() *osui.ComponentData {
	return &h.Data
}

func (h *HelloComponent) Render() string {
	osui.UseStyle(h.Style)
	return "Hello " + h.Style.Fg + h.Name + colors.Reset + "!"
}

func (h *HelloComponent) Update(key string) bool {
	return false
}

func (h *HelloComponent) SetStyle(c interface{}) {
	h.Style = osui.SetDefaults(c.(*HelloStyle)).(*HelloStyle)
}

func Hello(name string) *HelloComponent {
	return osui.NewComponent(&HelloComponent{
		Name:  name,
		Style: osui.SetDefaults(&HelloStyle{}).(*HelloStyle),
	}).(*HelloComponent)
}
```