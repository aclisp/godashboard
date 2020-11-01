package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Body renders the <body> tag
type Body struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (b *Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Section(
			vecty.Markup(vecty.Class("hero", "is-primary", "is-fullheight")),
			elem.Div(
				vecty.Markup(vecty.Class("hero-body")),
				vecty.Text("Login"),
			),
		),
	)
}
