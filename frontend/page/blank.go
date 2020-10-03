package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Blank is a blank page
type Blank struct {
	vecty.Core
}

// Render Blank
func (blank *Blank) Render() vecty.ComponentOrHTML {
	return elem.Div(
		// Page Heading
		elem.Heading1(
			vecty.Markup(vecty.Class("h3", "mb-4", "text-gray-800")),
			vecty.Text("Blank Page"),
		),
	)
}
