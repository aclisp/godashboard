package v

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// Topbar is the main page top bar
type Topbar struct {
	vecty.Core
}

// Render a top bar
func (b *Topbar) Render() vecty.ComponentOrHTML {
	return elem.Navigation(
		vecty.Markup(vecty.Class("navbar", "navbar-expand", "navbar-light", "bg-white", "topbar", "mb-4", "static-top", "shadow")),
		// Sidebar Toggle (Topbar)
		elem.Button(
			vecty.Markup(
				prop.ID("sidebarToggleTop"),
				vecty.Class("btn", "btn-link", "d-md-none", "rounded-circle", "mr-3"),
			),
			elem.Italic(vecty.Markup(vecty.Class("fa", "fa-bars"))),
		),
		// Topbar Search
		b.renderSearch(),
		// Topbar Navbar
		elem.UnorderedList(
			vecty.Markup(vecty.Class("navbar-nav", "ml-auto")),
			elem.Div(vecty.Markup(vecty.Class("topbar-divider", "d-none", "d-sm-block"))),
			// Nav Item - User Information
			b.renderUserInfo(),
		),
	)
}

func (b *Topbar) renderSearch() *vecty.HTML {
	return elem.Form(
		vecty.Markup(vecty.Class("d-none", "d-sm-inline-block", "form-inline", "mr-auto", "ml-md-3", "my-2", "my-md-0", "mw-100", "navbar-search")),
		elem.Div(
			vecty.Markup(vecty.Class("input-group")),
			elem.Input(
				vecty.Markup(
					prop.Type(prop.TypeText),
					vecty.Class("form-control", "bg-light", "border-0", "small"),
					prop.Placeholder("Search for..."),
					vecty.Attribute("aria-label", "Search"),
					vecty.Attribute("aria-describedby", "basic-addon2"),
				),
			),
			elem.Div(
				vecty.Markup(vecty.Class("input-group-append")),
				elem.Button(
					vecty.Markup(
						prop.Type(prop.TypeButton),
						vecty.Class("btn", "btn-primary"),
					),
					elem.Italic(vecty.Markup(vecty.Class("fas", "fa-search", "fa-sm"))),
				),
			),
		),
	)
}

func (b *Topbar) renderUserInfo() *vecty.HTML {
	id := "userDropdown"
	return elem.ListItem(
		vecty.Markup(vecty.Class("nav-item", "dropdown", "no-arrow")),
		elem.Anchor(
			vecty.Markup(
				vecty.Class("nav-link", "dropdown-toggle"),
				prop.Href("#"),
				prop.ID(id),
				vecty.Attribute("role", "button"),
				vecty.Data("toggle", "dropdown"),
				vecty.Attribute("aria-haspopup", "true"),
				vecty.Attribute("aria-expanded", "false"),
			),
			elem.Span(
				vecty.Markup(vecty.Class("mr-2", "d-none", "d-lg-inline", "text-gray-600", "small")),
				vecty.Text("Alice"),
			),
			elem.Image(
				vecty.Markup(
					vecty.Class("img-profile", "rounded-circle"),
					prop.Src("https://source.unsplash.com/QAB-WJcbgJk/60x60"),
				),
			),
		),
		// Dropdown - User Information
		elem.Div(
			vecty.Markup(
				vecty.Class("dropdown-menu", "dropdown-menu-right", "shadow", "animated--grow-in"),
				vecty.Attribute("aria-labelledby", id),
			),
			elem.Anchor(
				vecty.Markup(vecty.Class("dropdown-item"), prop.Href("#")),
				elem.Italic(vecty.Markup(vecty.Class("fas", "fa-user", "fa-sm", "fa-fw", "mr-2", "text-gray-400"))),
				vecty.Text("Profile"),
			),
			elem.Anchor(
				vecty.Markup(vecty.Class("dropdown-item"), prop.Href("#")),
				elem.Italic(vecty.Markup(vecty.Class("fas", "fa-cogs", "fa-sm", "fa-fw", "mr-2", "text-gray-400"))),
				vecty.Text("Settings"),
			),
			elem.Anchor(
				vecty.Markup(vecty.Class("dropdown-item"), prop.Href("#")),
				elem.Italic(vecty.Markup(vecty.Class("fas", "fa-list", "fa-sm", "fa-fw", "mr-2", "text-gray-400"))),
				vecty.Text("Activity Log"),
			),
			elem.Div(vecty.Markup(vecty.Class("dropdown-divider"))),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("dropdown-item"),
					prop.Href("#"),
					vecty.Data("toggle", "modal"),
					vecty.Data("target", "#logoutModal"),
				),
				elem.Italic(vecty.Markup(vecty.Class("fas", "fa-sign-out-alt", "fa-sm", "fa-fw", "mr-2", "text-gray-400"))),
				vecty.Text("Logout"),
			),
		),
	)
}