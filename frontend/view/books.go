package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/hexops/vecty/style"
)

// Books .
type Books struct {
	vecty.Core
}

// Render .
func (c *Books) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading1(
			vecty.Markup(vecty.Class("title")),
			vecty.Text("Books"),
		),
		c.renderNavbar(),
		c.renderBooks(),
		c.renderPagination(),
	)
}

func (c *Books) renderNavbar() *vecty.HTML {
	return elem.Navigation(
		vecty.Markup(vecty.Class("level")),
		elem.Div(
			vecty.Markup(vecty.Class("level-left")),
			elem.Div(
				vecty.Markup(vecty.Class("level-item")),
				elem.Paragraph(
					vecty.Markup(vecty.Class("subtitle", "is-5")),
					elem.Strong(
						vecty.Markup(vecty.Class("mr-1")),
						vecty.Text("6")),
					vecty.Text("books"),
				),
			),
			elem.Paragraph(
				vecty.Markup(vecty.Class("level-item")),
				elem.Anchor(
					vecty.Markup(vecty.Class("button", "is-success")),
					vecty.Text("New"),
				),
			),
			elem.Div(
				vecty.Markup(vecty.Class("level-item", "is-hidden-tablet-only")),
				elem.Div(
					vecty.Markup(vecty.Class("field", "has-addons")),
					elem.Paragraph(
						vecty.Markup(vecty.Class("control")),
						elem.Input(
							vecty.Markup(vecty.Class("input"),
								prop.Type(prop.TypeText),
								prop.Placeholder("Book name, ISBN…"))),
					),
					elem.Paragraph(
						vecty.Markup(vecty.Class("control")),
						elem.Button(
							vecty.Markup(vecty.Class("button")),
							vecty.Text("Search")),
					),
				),
			),
		),
		elem.Div(
			vecty.Markup(vecty.Class("level-right")),
			elem.Div(
				vecty.Markup(vecty.Class("level-item")),
				vecty.Text("Order by"),
			),
			elem.Div(
				vecty.Markup(vecty.Class("level-item")),
				elem.Div(
					vecty.Markup(vecty.Class("select")),
					elem.Select(
						elem.Option(vecty.Text("Publish date")),
						elem.Option(vecty.Text("Price")),
						elem.Option(vecty.Text("Page count")),
					),
				),
			),
		),
	)
}

func (c *Books) renderBooks() *vecty.HTML {
	var books vecty.List
	for i := 0; i < 6; i++ {
		books = append(books, c.renderBook())
	}
	return elem.Div(
		vecty.Markup(vecty.Class("columns", "is-multiline")),
		books,
	)
}

func (c *Books) renderBook() *vecty.HTML {
	return elem.Div(
		vecty.Markup(vecty.Class("column", "is-12-tablet", "is-6-desktop", "is-4-widescreen")),
		elem.Article(
			vecty.Markup(vecty.Class("box")),
			elem.Div(
				vecty.Markup(vecty.Class("media")),
				elem.Aside(
					vecty.Markup(vecty.Class("media-left")),
					elem.Image(
						vecty.Markup(
							prop.Src("/images/tensorflow.jpg"),
							style.Width(style.Px(80)),
						),
					),
				),
				elem.Div(
					vecty.Markup(vecty.Class("media-content")),
					elem.Paragraph(
						vecty.Markup(vecty.Class("title", "is-5", "is-spaced", "is-marginless")),
						elem.Anchor(
							vecty.Text("TensorFlow For Machine Intelligence"),
						),
					),
					elem.Paragraph(
						vecty.Markup(vecty.Class("subtitle", "is-marginless")),
						vecty.Text("$22.99"),
					),
					elem.Div(
						vecty.Markup(vecty.Class("content", "is-small")),
						vecty.Text("270 pages"),
						elem.Break(),
						vecty.Text("ISBN: 9781939902351"),
						elem.Break(),
						elem.Anchor(
							vecty.Text("Edit"),
						),
						elem.Span(vecty.Text("·")),
						elem.Anchor(
							vecty.Text("Delete"),
						),
						elem.Paragraph(),
					),
				),
			),
		),
	)
}

func (c *Books) renderPagination() *vecty.HTML {
	return elem.Navigation(
		vecty.Markup(vecty.Class("pagination")),
		elem.Anchor(
			vecty.Markup(vecty.Class("pagination-previous")),
			vecty.Text("Previous"),
		),
		elem.Anchor(
			vecty.Markup(vecty.Class("pagination-next")),
			vecty.Text("Next page"),
		),
		elem.UnorderedList(
			vecty.Markup(vecty.Class("pagination-list")),
			elem.ListItem(
				elem.Anchor(
					vecty.Markup(vecty.Class("pagination-link")),
					vecty.Text("1"),
				),
			),
			elem.ListItem(
				elem.Span(
					vecty.Markup(vecty.Class("pagination-ellipsis")),
					vecty.Text("…"),
				),
			),
			elem.ListItem(
				elem.Anchor(
					vecty.Markup(vecty.Class("pagination-link")),
					vecty.Text("45"),
				),
			),
			elem.ListItem(
				elem.Anchor(
					vecty.Markup(vecty.Class("pagination-link", "is-current")),
					vecty.Text("46"),
				),
			),
			elem.ListItem(
				elem.Anchor(
					vecty.Markup(vecty.Class("pagination-link")),
					vecty.Text("47"),
				),
			),
			elem.ListItem(
				elem.Span(
					vecty.Markup(vecty.Class("pagination-ellipsis")),
					vecty.Text("…"),
				),
			),
			elem.ListItem(
				elem.Anchor(
					vecty.Markup(vecty.Class("pagination-link")),
					vecty.Text("86"),
				),
			),
		),
	)
}
