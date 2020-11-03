package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

// Main .
type Main struct {
	vecty.Core
}

// Render .
func (c *Main) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("column")),
		router.NewRoute("/go/dashboard/home", &Home{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/go/dashboard/books", &Books{}, router.NewRouteOpts{ExactMatch: true}),
	)
}
