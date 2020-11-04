package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

// Main .
type Main struct {
	vecty.Core

	navMenus []NavMenu
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
	for _, m := range c.navMenus {
		vl = append(vl, router.NewRoute(m.Link, m.Component, router.NewRouteOpts{ExactMatch: true}))
	}
	return vl
}
