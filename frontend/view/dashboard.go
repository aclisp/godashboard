package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Dashboard .
type Dashboard struct {
	vecty.Core
}

// Render .
func (c *Dashboard) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&Header{},
		elem.Section(
			vecty.Markup(vecty.Class("section")),
			elem.Div(
				vecty.Markup(vecty.Class("columns")),
				&Sidebar{},
				&Main{},
			),
		),
	)
}
