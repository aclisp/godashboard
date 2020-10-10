package v

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	router "marwan.io/vecty-router"

	"github.com/aclisp/godashboard/frontend/s"
	dashboard "github.com/aclisp/godashboard/proto"
)

// Sidebar is the main page side bar
type Sidebar struct {
	vecty.Core
}

// Render a side bar
func (b *Sidebar) Render() vecty.ComponentOrHTML {
	id := "accordionSidebar"

	ifaces := make(vecty.List, len(s.SidebarMenus))
	for i := range ifaces {
		ifaces[i] = &SidebarMenu{
			parent: id,
			data:   s.SidebarMenus[i],
		}
	}

	return elem.UnorderedList(
		vecty.Markup(
			prop.ID(id),
			vecty.Class("navbar-nav", "bg-gradient-primary", "sidebar", "sidebar-dark", "accordion"),
		),
		// Sidebar - Brand
		b.renderBrand(),
		// Divider
		b.renderDivider("my-0"),
		// Nav Item - Dashboard
		b.renderNavItemDashboard(),
		// Divider
		b.renderDivider(),
		// Heading
		b.renderHeading("Interface"),
		// Nav Item - Menu
		ifaces,
		// Divider
		b.renderDivider(),
		// Heading
		b.renderHeading("Addons"),
		&SidebarMenu{
			parent: id,
			data: &dashboard.SidebarMenu{
				Id:     "site-settings",
				FaIcon: "fa-cog",
				Text:   "Settings",
				Groups: []*dashboard.SidebarGroup{
					{
						Items: []*dashboard.SidebarEntry{
							{Text: "Sidebar Menus", Route: "/sidebar-menus"},
						},
					},
				},
			},
		},
		// Nav Item - Charts
		b.renderItem("fa-chart-area", "Charts", "/charts"),
		// Nav Item - Tables
		b.renderItem("fa-table", "Tables", "/tables"),
		// Divider
		b.renderDivider("d-none", "d-md-block"),
		// Sidebar Toggler (Sidebar)
		b.renderToggler(),
	)
}

func (b *Sidebar) renderBrand() *vecty.HTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.Class("sidebar-brand", "d-flex", "align-items-center", "justify-content-center"),
			prop.Href("/index.html"),
		),
		elem.Div(
			vecty.Markup(vecty.Class("sidebar-brand-icon", "rotate-n-15")),
			elem.Italic(vecty.Markup(vecty.Class("fas", "fa-laugh-wink"))),
		),
		elem.Div(
			vecty.Markup(vecty.Class("sidebar-brand-text", "mx-3")),
			vecty.Text("YS Admin "),
			elem.Superscript(vecty.Text("2")),
		),
	)
}

func (b *Sidebar) renderDivider(class ...string) *vecty.HTML {
	c := make([]string, 0, 10)
	c = append(c, "sidebar-divider")
	c = append(c, class...)
	return elem.HorizontalRule(vecty.Markup(vecty.Class(c...)))
}

func (b *Sidebar) renderNavItemDashboard() *vecty.HTML {
	return elem.ListItem(
		vecty.Markup(vecty.Class("nav-item")),
		b.navLink("fa-tachometer-alt", "Dashboard", "/"),
	)
}

func (b *Sidebar) renderHeading(text string) *vecty.HTML {
	return elem.Div(
		vecty.Markup(vecty.Class("sidebar-heading")),
		vecty.Text(text),
	)
}

func (b *Sidebar) renderItem(icon, text, route string) *vecty.HTML {
	return elem.ListItem(
		vecty.Markup(vecty.Class("nav-item")),
		b.navLink(icon, text, route),
	)
}

func (b *Sidebar) renderToggler() *vecty.HTML {
	return elem.Div(
		vecty.Markup(vecty.Class("text-center", "d-none", "d-md-inline")),
		elem.Button(
			vecty.Markup(
				vecty.Class("rounded-circle", "border-0"),
				prop.ID("sidebarToggle"),
			),
		),
	)
}

func (b *Sidebar) navLink(icon, text, route string) *vecty.HTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.Class("nav-link"),
			prop.Href(route),
			event.Click(func(e *vecty.Event) {
				router.Redirect(route)
			}).PreventDefault(),
		),
		elem.Italic(vecty.Markup(vecty.Class("fas", "fa-fw", icon))),
		elem.Span(vecty.Text(text)),
	)
}
