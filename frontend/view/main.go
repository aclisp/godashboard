package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

// Main .
type Main struct {
	vecty.Core

	NavMenus []NavMenu
}

// Render .
func (c *Main) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("column")),
		c.renderRoutes(),
	)
}

func (c *Main) renderRoutes() vecty.List {
	var vl vecty.List
	for _, m := range c.NavMenus {
		vl = append(vl, router.NewRoute(m.Link, m.Component, router.NewRouteOpts{ExactMatch: true}))
	}
	return vl
}
