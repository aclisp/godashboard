package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
)

// SignOutModal .
type SignOutModal struct {
	vecty.Core

	Open      bool `vecty:"prop"`
	OnConfirm func()
}

// Render .
func (c *SignOutModal) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("modal"), vecty.ClassMap{"is-active": c.Open}),
		elem.Div(
			vecty.Markup(vecty.Class("modal-background"),
				event.Click(func(e *vecty.Event) {
					c.close()
				}),
			),
		),
		elem.Div(
			vecty.Markup(vecty.Class("modal-content")),
			elem.Div(
				vecty.Markup(vecty.Class("box")),
				elem.Article(
					vecty.Markup(vecty.Class("media")),
					elem.Div(
						vecty.Markup(vecty.Class("media-left")),
						elem.Figure(
							vecty.Markup(vecty.Class("image", "is-64x64")),
							elem.Image(
								vecty.Markup(vecty.Property("src", "https://bulma.io/images/placeholders/128x128.png"),
									vecty.Property("alt", "Image"),
								),
							),
						),
					),
					elem.Div(
						vecty.Markup(vecty.Class("media-content")),
						elem.Div(
							vecty.Markup(vecty.Class("content")),
							elem.Paragraph(
								elem.Strong(
									vecty.Text("John Smith"),
								),
								elem.Small(
									vecty.Text("@johnsmith"),
								),
								elem.Small(
									vecty.Text("31m"),
								),
								elem.Break(),
								vecty.Text("Are you sure to sign out?"),
							),
						),
					),
				),
				elem.Button(
					vecty.Markup(vecty.Class("button", "is-success", "is-fullwidth", "mt-4"),
						event.Click(func(e *vecty.Event) {
							if c.OnConfirm != nil {
								c.OnConfirm()
							}
						}),
					),
					vecty.Text("Sign Out"),
				),
			),
		),
		elem.Button(
			vecty.Markup(vecty.Class("modal-close", "is-large"),
				vecty.Attribute("aria-label", "close"),
				event.Click(func(e *vecty.Event) {
					c.close()
				}),
			),
		),
	)
}

func (c *SignOutModal) close() {
	c.Open = false
	vecty.Rerender(c)
}
