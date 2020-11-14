package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

// Header .
type Header struct {
	vecty.Core

	active bool
}

// Render .
func (c *Header) Render() vecty.ComponentOrHTML {
	return elem.Navigation(
		vecty.Markup(vecty.Class("navbar", "has-shadow")),
		c.renderBrand(),
		elem.Div(
			vecty.Markup(vecty.Class("navbar-menu"), vecty.ClassMap{"is-active": c.active}),
			elem.Div(
				vecty.Markup(vecty.Class("navbar-start")),
				elem.Div(
					vecty.Markup(vecty.Class("navbar-item")),
					elem.Small(vecty.Text("Publishing at the speed of technology")),
				),
			),
			c.renderUserControls(),
		),
	)
}

func (c *Header) renderBrand() *vecty.HTML {
	return elem.Div(
		vecty.Markup(vecty.Class("navbar-brand")),
		elem.Anchor(
			vecty.Markup(vecty.Class("navbar-item")),
			elem.Image(vecty.Markup(prop.Src("/images/logo.png"))),
		),
		elem.Div(
			vecty.Markup(vecty.Class("navbar-burger"), vecty.ClassMap{"is-active": c.active},
				event.Click(func(e *vecty.Event) {
					c.active = !c.active
					vecty.Rerender(c)
				})),
			elem.Span(),
			elem.Span(),
			elem.Span(),
		),
	)
}

func (c *Header) renderUserControls() *vecty.HTML {
	return elem.Div(
		vecty.Markup(vecty.Class("navbar-end")),
		elem.Div(
			vecty.Markup(vecty.Class("navbar-item", "has-dropdown", "is-hoverable")),
			elem.Div(
				vecty.Markup(vecty.Class("navbar-link")),
				vecty.Text("Alex Johnson"),
			),
			elem.Div(
				vecty.Markup(vecty.Class("navbar-dropdown")),
				elem.Anchor(
					vecty.Markup(vecty.Class("navbar-item")),
					elem.Span(
						vecty.Markup(vecty.Class("icon", "is-small", "mr-1")),
						elem.Italic(vecty.Markup(vecty.Class("fa", "fa-user-circle-o")))),
					vecty.Text("Profile"),
				),
				elem.Anchor(
					vecty.Markup(vecty.Class("navbar-item")),
					elem.Span(
						vecty.Markup(vecty.Class("icon", "is-small", "mr-1")),
						elem.Italic(vecty.Markup(vecty.Class("fa", "fa-bug")))),
					vecty.Text("Report bug"),
				),
				elem.Anchor(
					vecty.Markup(vecty.Class("navbar-item")),
					elem.Span(
						vecty.Markup(vecty.Class("icon", "is-small", "mr-1")),
						elem.Italic(vecty.Markup(vecty.Class("fa", "fa-sign-out")))),
					vecty.Text("Sign Out"),
				)),
		),
	)
}
