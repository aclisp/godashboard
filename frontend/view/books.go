package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Books .
type Books struct {
	vecty.Core
}

// Render .
func (c *Books) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading1(
			vecty.Markup(vecty.Class("title")),
			vecty.Text("Books"),
		),
	)
}
