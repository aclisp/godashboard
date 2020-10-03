package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	router "marwan.io/vecty-router"
)

// View404 is a 404 page
type View404 struct {
	vecty.Core
}

// Render View404
func (v *View404) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("text-center")),
		elem.Div(
			vecty.Markup(
				vecty.Class("error", "mx-auto"),
				vecty.Data("text", "404"),
			),
			vecty.Text("404"),
		),
		elem.Paragraph(
			vecty.Markup(vecty.Class("lead", "text-gray-800", "mb-5")),
			vecty.Text("Page Not Found"),
		),
		elem.Paragraph(
			vecty.Markup(vecty.Class("text-gray-500", "mb-0")),
			vecty.Text("It looks like you found a glitch in the matrix..."),
		),
		elem.Anchor(
			vecty.Markup(
				prop.Href("/"),
				event.Click(func(e *vecty.Event) {
					router.Redirect("/")
				}).PreventDefault(),
			),
			vecty.Text("← Back to Dashboard"),
		),
	)
}
