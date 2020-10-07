package v

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

// PageContent is the page content
type PageContent struct {
	vecty.Core
}

// Render the page content
func (c *PageContent) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("container-fluid")),
		router.NewRoute("/", &HomeView{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/404", &View404{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/blank", &Blank{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/tables", &TableView{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/go/{package}/{endpoint}", &DynamicView{}, router.NewRouteOpts{ExactMatch: true}),
		// Note that this handler only works for router.Link and router.Redirect accesses.
		// Directly accessing a non-existant route won't be handled by this.
		router.NotFoundHandler(&notFound{}),
	)
}

type notFound struct {
	vecty.Core
}

func (nf *notFound) Render() vecty.ComponentOrHTML {
	return elem.Div(
		// Page Heading
		elem.Heading1(
			vecty.Markup(vecty.Class("h3", "mb-4", "text-gray-800")),
			vecty.Text("Under construction..."),
		),
	)
}
