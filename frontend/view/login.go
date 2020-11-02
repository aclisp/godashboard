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
						elem.Form(
							vecty.Markup(
								vecty.Class("box"),
							),
							vecty.Text("Login"),
						),
					),
				),
			),
		),
	)
}
