package v

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// Body renders the <body> tag
type Body struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (b *Body) Render() vecty.ComponentOrHTML {
	id := "page-top"
	return elem.Body(
		vecty.Markup(prop.ID(id)),
		// Page Wrapper
		elem.Div(
			vecty.Markup(prop.ID("wrapper")),
			// Sidebar
			&Sidebar{},
			// Content Wrapper
			elem.Div(
				vecty.Markup(
					prop.ID("content-wrapper"),
					vecty.Class("d-flex", "flex-column"),
				),
				// Main Content
				&MainArea{},
				b.renderFooter(),
			),
		),
		// Scroll to Top Button
		elem.Anchor(
			vecty.Markup(
				vecty.Class("scroll-to-top", "rounded"),
				vecty.Markup(prop.Href("#"+id)),
			),
			elem.Italic(vecty.Markup(vecty.Class("fas", "fa-angle-up"))),
		),
		// Logout Modal
		b.renderLogoutModal(),
		// Run JavaScript
		elem.Script(vecty.Markup(prop.Src("/jsbundle/godashboard.bundle.js"))),
	)
}

func (b *Body) renderFooter() *vecty.HTML {
	return elem.Footer(
		vecty.Markup(vecty.Class("sticky-footer", "bg-white")),
		elem.Div(
			vecty.Markup(vecty.Class("container", "my-auto")),
			elem.Div(
				vecty.Markup(vecty.Class("copyright", "text-center", "my-auto")),
				elem.Span(
					vecty.Text("Copyright © Your Website 2020"),
				),
			),
		),
	)
}

func (b *Body) renderLogoutModal() *vecty.HTML {
	id := "exampleModalLabel"
	return elem.Div(
		vecty.Markup(
			vecty.Class("modal", "fade"),
			prop.ID("logoutModal"),
			vecty.Attribute("tabindex", "-1"),
			vecty.Attribute("role", "dialog"),
			vecty.Attribute("aria-labelledby", id),
			vecty.Attribute("aria-hidden", "true"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("modal-dialog"),
				vecty.Attribute("role", "document"),
			),
			elem.Div(
				vecty.Markup(vecty.Class("modal-content")),
				elem.Div(
					vecty.Markup(vecty.Class("modal-header")),
					elem.Heading5(
						vecty.Markup(
							prop.ID(id),
							vecty.Class("modal-title"),
						),
						vecty.Text("Ready to Leave?"),
					),
					elem.Button(
						vecty.Markup(
							vecty.Class("close"),
							prop.Type(prop.TypeButton),
							vecty.Data("dismiss", "modal"),
							vecty.Attribute("aria-label", "Close"),
						),
						elem.Span(
							vecty.Markup(vecty.Attribute("aria-hidden", "true")),
							vecty.Text("×"),
						),
					),
				),
				elem.Div(
					vecty.Markup(vecty.Class("modal-body")),
					vecty.Text("Select \"Logout\" below if you are ready to end your current session."),
				),
				elem.Div(
					vecty.Markup(vecty.Class("modal-footer")),
					elem.Button(
						vecty.Markup(
							vecty.Class("btn", "btn-secondary"),
							prop.Type(prop.TypeButton),
							vecty.Data("dismiss", "modal"),
						),
						vecty.Text("Cancel"),
					),
					elem.Anchor(
						vecty.Markup(
							vecty.Class("btn", "btn-primary"),
							prop.Href("/login.html"),
						),
						vecty.Text("Logout"),
					),
				),
			),
		),
	)
}
