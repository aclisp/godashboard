package v

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	router "marwan.io/vecty-router"

	dashboard "github.com/aclisp/godashboard/proto"
)

// SidebarMenu is the menu inside the main page side bar
type SidebarMenu struct {
	vecty.Core

	parent string
	data   *dashboard.SidebarMenu
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
				vecty.Data("target", "#"+m.data.Id),
				vecty.Attribute("aria-expanded", "true"),
				vecty.Attribute("aria-controls", m.data.Id),
			),
			elem.Italic(vecty.Markup(vecty.Class("fas", "fa-fw", m.data.FaIcon))),
			elem.Span(vecty.Text(m.data.Text)),
		),
		elem.Div(
			vecty.Markup(
				prop.ID(m.data.Id),
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
	for i, group := range m.data.Groups {
		list = append(list, m.renderGroup(group, i))
	}
	return list
}

func (m *SidebarMenu) renderGroup(group *dashboard.SidebarGroup, index int) vecty.List {
	list := make(vecty.List, 0, 10)
	if index > 0 {
		list = append(list, elem.Div(vecty.Markup(vecty.Class("collapse-divider"))))
	}
	if group.Text != "" {
		list = append(list, elem.Heading6(
			vecty.Markup(vecty.Class("collapse-header")),
			vecty.Text(group.Text),
		))
	}
	for _, item := range group.Items {
		item := item
		list = append(list, elem.Anchor(
			vecty.Markup(
				vecty.Class("collapse-item"),
				prop.Href(item.Route),
				event.Click(func(e *vecty.Event) {
					router.Redirect(item.Route)
				}).PreventDefault(),
			),
			vecty.Text(item.Text),
		))
	}
	return list
}
