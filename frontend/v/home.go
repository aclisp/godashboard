package v

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// HomeView is the home
type HomeView struct {
	vecty.Core
}

// Render HomeView
func (home *HomeView) Render() vecty.ComponentOrHTML {
	return elem.Div(
		// Page Heading
		elem.Heading1(
			vecty.Markup(vecty.Class("h3", "mb-2", "text-gray-800")),
			vecty.Text("Home"),
		),
		elem.Paragraph(
			vecty.Markup(vecty.Class("mb-4")),
			vecty.Text("Home is a start page."),
		),
	)
}
