package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

// Body renders the <body> tag
type Body struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (c *Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		router.NewRoute("/go/login", &Login{}, router.NewRouteOpts{ExactMatch: true}),
		router.NotFoundHandler(&NotFound{}),
	)
}
