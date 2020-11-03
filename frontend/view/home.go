package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Home .
type Home struct {
	vecty.Core
}

// Render .
func (c *Home) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading1(
			vecty.Markup(vecty.Class("title")),
			vecty.Text("Home"),
		),
	)
}
