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

func (b *MenuComponent) Params(param MenuParams) *MenuComponent {
	b.Style = osui.SetDefaults(&param.Style).(*MenuStyle)
	return b
}

func Hello(name string) *HelloComponent {
	return osui.NewComponent(&HelloComponent{
		Name:  name,
		Style: osui.SetDefaults(&HelloStyle{}).(*HelloStyle),
	}).(*HelloComponent)
}
```