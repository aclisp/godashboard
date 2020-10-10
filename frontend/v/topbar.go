package v

import (
	"strings"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"

	"github.com/aclisp/godashboard/frontend/s"
	"github.com/aclisp/godashboard/frontend/s/action"
	"github.com/aclisp/godashboard/frontend/s/dispatcher"
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
		b.renderSelect("d-none", "d-sm-inline-block", "form-inline", "mr-auto", "ml-md-3", "my-2", "my-md-0", "mw-100", "navbar-search"),
		// Topbar Navbar
		elem.UnorderedList(
			vecty.Markup(vecty.Class("navbar-nav", "ml-auto")),
			// Nav Item - Search Dropdown (Visible Only XS)
			elem.ListItem(
				vecty.Markup(
					vecty.Class("nav-item", "dropdown", "no-arrow", "d-sm-none"),
				),
				elem.Anchor(
					vecty.Markup(
						vecty.Class("nav-link", "dropdown-toggle"),
						vecty.Property("href", "#"),
						vecty.Property("id", "searchDropdown"),
						vecty.Attribute("role", "button"),
						vecty.Data("toggle", "dropdown"),
						vecty.Attribute("aria-haspopup", "true"),
						vecty.Attribute("aria-expanded", "false"),
					),
					elem.Italic(
						vecty.Markup(
							vecty.Class("fas", "fa-cloud", "fa-fw"),
						),
					),
				),
				// Dropdown - Messages
				elem.Div(
					vecty.Markup(
						vecty.Class("dropdown-menu", "dropdown-menu-right", "p-3", "shadow", "animated--grow-in"),
						vecty.Attribute("aria-labelledby", "searchDropdown"),
					),
					b.renderSelect("form-inline", "mr-auto", "w-100", "navbar-search"),
				),
			),
			// Nav Item - Alerts
			b.renderAlerts(),
			// Nav Item - Messages
			b.renderMessages(),

			elem.Div(vecty.Markup(vecty.Class("topbar-divider", "d-none", "d-sm-block"))),
			// Nav Item - User Information
			b.renderUserInfo(),
		),
	)
}

func (b *Topbar) renderSearch(formClass ...string) *vecty.HTML {
	return elem.Form(
		vecty.Markup(vecty.Class(formClass...)),
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

func (b *Topbar) renderSelect(formClass ...string) *vecty.HTML {
	options := make(vecty.List, len(s.Gateways))
	for i, gateway := range s.Gateways {
		options[i] = elem.Option(
			vecty.Markup(
				prop.Value(strings.Join(gateway.ID[:], ",")),
				vecty.MarkupIf(gateway.Selected, vecty.Property("selected", "selected")),
			),
			vecty.Text(gateway.Name),
		)
	}

	return elem.Form(
		vecty.Markup(vecty.Class(formClass...)),
		elem.Div(
			vecty.Markup(vecty.Class("input-group")),
			elem.Div(
				vecty.Markup(vecty.Class("input-group-prepend")),
				elem.Label(
					vecty.Markup(prop.For("selectGateway"), vecty.Class("input-group-text", "bg-light")),
					vecty.Text("Select Gateway:"),
				),
			),
			elem.Select(
				vecty.Markup(
					prop.ID("selectGateway"),
					vecty.Class("form-control", "bg-light", "small", "custom-select"),
					event.Change(func(e *vecty.Event) {
						var changeGateway action.ChangeGateway
						for i, s := range strings.SplitN(e.Target.Get("value").String(), ",", 2) {
							changeGateway.GatewayID[i] = s
						}
						dispatcher.Dispatch(&changeGateway)
						dispatcher.Dispatch(&action.SyncDynamicViewData{})
					}).PreventDefault(),
				),
				options,
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

func (b *Topbar) renderAlerts() *vecty.HTML {
	return elem.ListItem(
		vecty.Markup(
			vecty.Class("nav-item", "dropdown", "no-arrow", "mx-1"),
		),
		elem.Anchor(
			vecty.Markup(
				vecty.Class("nav-link", "dropdown-toggle"),
				vecty.Property("href", "#"),
				vecty.Property("id", "alertsDropdown"),
				vecty.Attribute("role", "button"),
				vecty.Data("toggle", "dropdown"),
				vecty.Attribute("aria-haspopup", "true"),
				vecty.Attribute("aria-expanded", "false"),
			),
			elem.Italic(
				vecty.Markup(
					vecty.Class("fas", "fa-bell", "fa-fw"),
				),
			),
			// Counter - Alerts
			elem.Span(
				vecty.Markup(
					vecty.Class("badge", "badge-danger", "badge-counter"),
				),
				vecty.Text("3+"),
			),
		),
		// Dropdown - Alerts
		elem.Div(
			vecty.Markup(
				vecty.Class("dropdown-list", "dropdown-menu", "dropdown-menu-right", "shadow", "animated--grow-in"),
				vecty.Attribute("aria-labelledby", "alertsDropdown"),
			),
			elem.Heading6(
				vecty.Markup(
					vecty.Class("dropdown-header"),
				),
				vecty.Text("Alerts Center"),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("dropdown-item", "d-flex", "align-items-center"),
					vecty.Property("href", "#"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("mr-3"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("icon-circle", "bg-primary"),
						),
						elem.Italic(
							vecty.Markup(
								vecty.Class("fas", "fa-file-alt", "text-white"),
							),
						),
					),
				),
				elem.Div(
					elem.Div(
						vecty.Markup(
							vecty.Class("small", "text-gray-500"),
						),
						vecty.Text("December 12, 2019"),
					),
					elem.Span(
						vecty.Markup(
							vecty.Class("font-weight-bold"),
						),
						vecty.Text("A new monthly report is ready to download!"),
					),
				),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("dropdown-item", "d-flex", "align-items-center"),
					vecty.Property("href", "#"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("mr-3"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("icon-circle", "bg-success"),
						),
						elem.Italic(
							vecty.Markup(
								vecty.Class("fas", "fa-donate", "text-white"),
							),
						),
					),
				),
				elem.Div(
					elem.Div(
						vecty.Markup(
							vecty.Class("small", "text-gray-500"),
						),
						vecty.Text("December 7, 2019"),
					),
					vecty.Text("$290.29 has been deposited into your account!"),
				),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("dropdown-item", "d-flex", "align-items-center"),
					vecty.Property("href", "#"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("mr-3"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("icon-circle", "bg-warning"),
						),
						elem.Italic(
							vecty.Markup(
								vecty.Class("fas", "fa-exclamation-triangle", "text-white"),
							),
						),
					),
				),
				elem.Div(
					elem.Div(
						vecty.Markup(
							vecty.Class("small", "text-gray-500"),
						),
						vecty.Text("December 2, 2019"),
					),
					vecty.Text("Spending Alert: We've noticed unusually high spending for your account."),
				),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("dropdown-item", "text-center", "small", "text-gray-500"),
					vecty.Property("href", "#"),
				),
				vecty.Text("Show All Alerts"),
			),
		),
	)
}

func (b *Topbar) renderMessages() *vecty.HTML {
	return elem.ListItem(
		vecty.Markup(
			vecty.Class("nav-item", "dropdown", "no-arrow", "mx-1"),
		),
		elem.Anchor(
			vecty.Markup(
				vecty.Class("nav-link", "dropdown-toggle"),
				vecty.Property("href", "#"),
				vecty.Property("id", "messagesDropdown"),
				vecty.Attribute("role", "button"),
				vecty.Data("toggle", "dropdown"),
				vecty.Attribute("aria-haspopup", "true"),
				vecty.Attribute("aria-expanded", "false"),
			),
			elem.Italic(
				vecty.Markup(
					vecty.Class("fas", "fa-envelope", "fa-fw"),
				),
			),
			// Counter - Messages
			elem.Span(
				vecty.Markup(
					vecty.Class("badge", "badge-danger", "badge-counter"),
				),
				vecty.Text("7"),
			),
		),
		// Dropdown - Messages
		elem.Div(
			vecty.Markup(
				vecty.Class("dropdown-list", "dropdown-menu", "dropdown-menu-right", "shadow", "animated--grow-in"),
				vecty.Attribute("aria-labelledby", "messagesDropdown"),
			),
			elem.Heading6(
				vecty.Markup(
					vecty.Class("dropdown-header"),
				),
				vecty.Text("Message Center"),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("dropdown-item", "d-flex", "align-items-center"),
					vecty.Property("href", "#"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("dropdown-list-image", "mr-3"),
					),
					elem.Image(
						vecty.Markup(
							vecty.Class("rounded-circle"),
							vecty.Property("src", "https://source.unsplash.com/fn_BT9fwg_E/60x60"),
							vecty.Property("alt", ""),
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("status-indicator", "bg-success"),
						),
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("font-weight-bold"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("text-truncate"),
						),
						vecty.Text("Hi there! I am wondering if you can help me with a problem I've been having."),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("small", "text-gray-500"),
						),
						vecty.Text("Emily Fowler · 58m"),
					),
				),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("dropdown-item", "d-flex", "align-items-center"),
					vecty.Property("href", "#"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("dropdown-list-image", "mr-3"),
					),
					elem.Image(
						vecty.Markup(
							vecty.Class("rounded-circle"),
							vecty.Property("src", "https://source.unsplash.com/AU4VPcFN4LE/60x60"),
							vecty.Property("alt", ""),
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("status-indicator"),
						),
					),
				),
				elem.Div(
					elem.Div(
						vecty.Markup(
							vecty.Class("text-truncate"),
						),
						vecty.Text("I have the photos that you ordered last month, how would you like them sent to you?"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("small", "text-gray-500"),
						),
						vecty.Text("Jae Chun · 1d"),
					),
				),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("dropdown-item", "d-flex", "align-items-center"),
					vecty.Property("href", "#"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("dropdown-list-image", "mr-3"),
					),
					elem.Image(
						vecty.Markup(
							vecty.Class("rounded-circle"),
							vecty.Property("src", "https://source.unsplash.com/CS2uCrpNzJY/60x60"),
							vecty.Property("alt", ""),
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("status-indicator", "bg-warning"),
						),
					),
				),
				elem.Div(
					elem.Div(
						vecty.Markup(
							vecty.Class("text-truncate"),
						),
						vecty.Text("Last month's report looks great, I am very happy with the progress so far, keep up the good work!"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("small", "text-gray-500"),
						),
						vecty.Text("Morgan Alvarez · 2d"),
					),
				),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("dropdown-item", "d-flex", "align-items-center"),
					vecty.Property("href", "#"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("dropdown-list-image", "mr-3"),
					),
					elem.Image(
						vecty.Markup(
							vecty.Class("rounded-circle"),
							vecty.Property("src", "https://source.unsplash.com/Mv9hjnEUHR4/60x60"),
							vecty.Property("alt", ""),
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("status-indicator", "bg-success"),
						),
					),
				),
				elem.Div(
					elem.Div(
						vecty.Markup(
							vecty.Class("text-truncate"),
						),
						vecty.Text("Am I a good boy? The reason I ask is because someone told me that people say this to all dogs, even if they aren't good..."),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("small", "text-gray-500"),
						),
						vecty.Text("Chicken the Dog · 2w"),
					),
				),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("dropdown-item", "text-center", "small", "text-gray-500"),
					vecty.Property("href", "#"),
				),
				vecty.Text("Read More Messages"),
			),
		),
	)
}
