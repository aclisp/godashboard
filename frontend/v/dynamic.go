package v

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

// DynamicView is rendered by backend-returned data
type DynamicView struct {
	vecty.Core
}

// Render DynamicView
func (d *DynamicView) Render() vecty.ComponentOrHTML {
	vars := router.GetNamedVar(d)
	packageName := vars["package"]
	endpointName := vars["endpoint"]

	return elem.Div(
		// Page Heading
		elem.Heading1(
			vecty.Markup(vecty.Class("h3", "mb-2", "text-gray-800")),
			vecty.Text("Dynamic View"),
		),
		elem.Paragraph(
			vecty.Markup(vecty.Class("mb-4")),
			vecty.Text(fmt.Sprintf("package is %q, endpoint is %q", packageName, endpointName)),
		),
	)
}
