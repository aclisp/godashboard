package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Sidebar .
type Sidebar struct {
	vecty.Core
}

// Render .
func (c *Sidebar) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("column", "is-4-tablet", "is-3-desktop", "is-2-widescreen")),
		elem.Navigation(
			vecty.Markup(vecty.Class("menu")),
			elem.Paragraph(
				vecty.Markup(vecty.Class("menu-label")),
				vecty.Text("Menu"),
			),
			elem.UnorderedList(
				vecty.Markup(vecty.Class("menu-list")),
				elem.ListItem(
					elem.Anchor(
						elem.Span(
							vecty.Markup(vecty.Class("icon")),
							elem.Italic(vecty.Markup(vecty.Class("fa", "fa-tachometer"))),
						),
						vecty.Text("Dashboard"),
					),
				),
				elem.ListItem(
					elem.Anchor(
						elem.Span(
							vecty.Markup(vecty.Class("icon")),
							elem.Italic(vecty.Markup(vecty.Class("fa", "fa-book"))),
						),
						vecty.Text("Books"),
					),
				),
				elem.ListItem(
					elem.Anchor(
						elem.Span(
							vecty.Markup(vecty.Class("icon")),
							elem.Italic(vecty.Markup(vecty.Class("fa", "fa-address-book"))),
						),
						vecty.Text("Customers"),
					),
				),
				elem.ListItem(
					elem.Anchor(
						elem.Span(
							vecty.Markup(vecty.Class("icon")),
							elem.Italic(vecty.Markup(vecty.Class("fa", "fa-file-text-o"))),
						),
						vecty.Text("Orders"),
					),
				),
			),
		),
	)
}
