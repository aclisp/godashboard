package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	router "marwan.io/vecty-router"
)

// Sidebar .
type Sidebar struct {
	vecty.Core

	navMenus []NavMenu
}

// Render .
func (c *Sidebar) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("column", "is-4-tablet", "is-3-desktop", "is-2-widescreen")),
		elem.Navigation(
			vecty.Markup(vecty.Class("menu")),
			elem.Paragraph(
				vecty.Markup(vecty.Class("menu-label")),
				vecty.Text("Menu"),
			),
			elem.UnorderedList(
				vecty.Markup(vecty.Class("menu-list")),
				c.renderMenuItems(),
			),
		),
	)
}

func (c *Sidebar) renderMenuItems() vecty.List {
	var vl vecty.List
	for i, m := range c.navMenus {
		i := i
		m := m // https://golang.org/doc/faq#closures_and_goroutines
		vl = append(vl, elem.ListItem(
			elem.Anchor(
				vecty.Markup(vecty.ClassMap{"is-active": m.active}, prop.Href(m.Link),
					event.Click(func(e *vecty.Event) {
						c.setActiveMenuItem(i)
						router.Redirect(m.Link)
					}).PreventDefault()),
				elem.Span(
					vecty.Markup(vecty.Class("icon")),
					elem.Italic(vecty.Markup(vecty.Class("fa", m.Icon))),
				),
				vecty.Text(m.Text),
			),
		))
	}
	return vl
}

func (c *Sidebar) setActiveMenuItem(index int) {
	for i := range c.navMenus {
		if i == index {
			c.navMenus[i].active = true
		} else {
			c.navMenus[i].active = false
		}
	}
}
