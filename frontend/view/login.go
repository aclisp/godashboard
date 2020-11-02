package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Login .
type Login struct {
	vecty.Core
}

// Render .
func (c *Login) Render() vecty.ComponentOrHTML {
	return elem.Section(
		vecty.Markup(
			vecty.Class("hero", "is-primary", "is-fullheight"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("hero-body"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("container"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("columns", "is-centered"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("column", "is-5-tablet", "is-4-desktop", "is-3-widescreen"),
						),
						c.renderForm(),
					),
				),
			),
		),
	)
}

func (c *Login) renderForm() *vecty.HTML {
	return elem.Form(
		vecty.Markup(
			vecty.Class("box"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("field", "has-text-centered"),
			),
			elem.Image(
				vecty.Markup(
					vecty.Property("src", "/images/logo-bis.png"),
					vecty.Style("width", "167"),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("field"),
			),
			elem.Label(
				vecty.Markup(
					vecty.Class("label"),
				),
				vecty.Text("Email"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("control", "has-icons-left"),
				),
				elem.Input(
					vecty.Markup(
						vecty.Class("input"),
						vecty.Property("type", "email"),
						vecty.Property("placeholder", "e.g. alexjohnson@gmail.com"),
						vecty.Attribute("required", ""),
					),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("icon", "is-small", "is-left"),
					),
					elem.Italic(
						vecty.Markup(
							vecty.Class("fa", "fa-envelope"),
						),
					),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("field"),
			),
			elem.Label(
				vecty.Markup(
					vecty.Class("label"),
				),
				vecty.Text("Password"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("control", "has-icons-left"),
				),
				elem.Input(
					vecty.Markup(
						vecty.Class("input"),
						vecty.Property("type", "password"),
						vecty.Property("placeholder", "********"),
						vecty.Attribute("required", ""),
					),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("icon", "is-small", "is-left"),
					),
					elem.Italic(
						vecty.Markup(
							vecty.Class("fa", "fa-lock"),
						),
					),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("field"),
			),
			elem.Label(
				vecty.Markup(
					vecty.Class("checkbox"),
				),
				elem.Input(
					vecty.Markup(
						vecty.Property("type", "checkbox"),
						vecty.Attribute("required", ""),
					),
				),
				vecty.Text(" Remember me"),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("field"),
			),
			elem.Button(
				vecty.Markup(
					vecty.Class("button", "is-success"),
				),
				vecty.Text("Login"),
			),
		),
	)
}
