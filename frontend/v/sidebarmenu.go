package v

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	router "marwan.io/vecty-router"
)

// SidebarMenu is the menu inside the main page side bar
type SidebarMenu struct {
	vecty.Core

	parent string
	id     string
	icon   string
	text   string
	groups []sidebarGroup
}

type sidebarGroup struct {
	text  string
	items []sidebarEntry
}

type sidebarEntry struct {
	text  string
	route string
}

// Render a side bar item
func (m *SidebarMenu) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		vecty.Markup(vecty.Class("nav-item")),
		elem.Anchor(
			vecty.Markup(
				vecty.Class("nav-link", "collapsed"),
				prop.Href("#"),
				vecty.Data("toggle", "collapse"),
				vecty.Data("target", "#"+m.id),
				vecty.Attribute("aria-expanded", "true"),
				vecty.Attribute("aria-controls", m.id),
			),
			elem.Italic(vecty.Markup(vecty.Class("fas", "fa-fw", m.icon))),
			elem.Span(vecty.Text(m.text)),
		),
		elem.Div(
			vecty.Markup(
				prop.ID(m.id),
				vecty.Class("collapse"),
				vecty.Data("parent", "#"+m.parent),
			),
			elem.Div(
				vecty.Markup(vecty.Class("bg-white", "py-2", "collapse-inner", "rounded")),
				m.renderGroups(),
			),
		),
	)
}

func (m *SidebarMenu) renderGroups() vecty.List {
	list := make(vecty.List, 0, 10)
	for i, group := range m.groups {
		list = append(list, m.renderGroup(group, i))
	}
	return list
}

func (m *SidebarMenu) renderGroup(group sidebarGroup, index int) vecty.List {
	list := make(vecty.List, 0, 10)
	if index > 0 {
		list = append(list, elem.Div(vecty.Markup(vecty.Class("collapse-divider"))))
	}
	list = append(list, elem.Heading6(
		vecty.Markup(vecty.Class("collapse-header")),
		vecty.Text(group.text),
	))
	for _, item := range group.items {
		item := item
		list = append(list, elem.Anchor(
			vecty.Markup(
				vecty.Class("collapse-item"),
				prop.Href(item.route),
				event.Click(func(e *vecty.Event) {
					router.Redirect(item.route)
				}).PreventDefault(),
			),
			vecty.Text(item.text),
		))
	}
	return list
}
