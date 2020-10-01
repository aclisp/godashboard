package v

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// Sidebar is the main page side bar
type Sidebar struct {
	vecty.Core
}

// Render a side bar
func (b *Sidebar) Render() vecty.ComponentOrHTML {
	id := "accordionSidebar"
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
		// Nav Item - Pages Collapse Menu
		&SidebarMenu{
			parent: id,
			id:     "collapseTwo",
			icon:   "fa-cog",
			text:   "Components",
			groups: []sidebarGroup{
				{
					text: "Custom Components:",
					items: []sidebarEntry{
						{text: "Buttons", href: "buttons.html"},
						{text: "Cards", href: "cards.html"},
					},
				},
			},
		},
		// Nav Item - Utilities Collapse Menu
		&SidebarMenu{
			parent: id,
			id:     "collapseUtilities",
			icon:   "fa-wrench",
			text:   "Utilities",
			groups: []sidebarGroup{
				{
					text: "Custom Utilities:",
					items: []sidebarEntry{
						{text: "Colors", href: "utilities-color.html"},
						{text: "Borders", href: "utilities-border.html"},
						{text: "Animations", href: "utilities-animation.html"},
						{text: "Other", href: "utilities-other.html"},
					},
				},
			},
		},
		// Divider
		b.renderDivider(),
		// Heading
		b.renderHeading("Addons"),
		// Nav Item - Pages Collapse Menu
		&SidebarMenu{
			parent: id,
			id:     "collapsePages",
			icon:   "fa-folder",
			text:   "Pages",
			groups: []sidebarGroup{
				{
					text: "Login Screens:",
					items: []sidebarEntry{
						{text: "Login", href: "login.html"},
						{text: "Register", href: "register.html"},
						{text: "Forgot Password", href: "forgot-password.html"},
					},
				},
				{
					text: "Other Pages:",
					items: []sidebarEntry{
						{text: "404 Page", href: "404.html"},
						{text: "Blank Page", href: "blank.html"},
					},
				},
			},
		},
		// Nav Item - Charts
		b.renderItem("fa-chart-area", "Charts", "charts.html"),
		// Nav Item - Tables
		b.renderItem("fa-table", "Tables", "tables.html"),
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
			prop.Href("index.html"),
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
		vecty.Markup(vecty.Class("nav-item", "active")),
		elem.Anchor(
			vecty.Markup(
				vecty.Class("nav-link"),
				prop.Href("index.html"),
			),
			elem.Italic(vecty.Markup(vecty.Class("fas", "fa-fw", "fa-tachometer-alt"))),
			elem.Span(vecty.Text("Dashboard")),
		),
	)
}

func (b *Sidebar) renderHeading(text string) *vecty.HTML {
	return elem.Div(
		vecty.Markup(vecty.Class("sidebar-heading")),
		vecty.Text(text),
	)
}

func (b *Sidebar) renderItem(icon, text, href string) *vecty.HTML {
	return elem.ListItem(
		vecty.Markup(vecty.Class("nav-item")),
		elem.Anchor(
			vecty.Markup(vecty.Class("nav-link"), prop.Href(href)),
			elem.Italic(vecty.Markup(vecty.Class("fas", "fa-fw", icon))),
			elem.Span(vecty.Text(text)),
		),
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
