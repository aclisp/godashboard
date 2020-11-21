package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
)

// ReportBugModal .
type ReportBugModal struct {
	vecty.Core

	Open    bool `vecty:"prop"`
	Success bool `vecty:"prop"`
}

// Render .
func (c *ReportBugModal) Render() vecty.ComponentOrHTML {
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
			vecty.Markup(vecty.Class("modal-card")),
			elem.Header(
				vecty.Markup(vecty.Class("modal-card-head")),
				elem.Paragraph(
					vecty.Markup(vecty.Class("modal-card-title")),
					vecty.Text("Report a Bug"),
				),
				elem.Button(
					vecty.Markup(vecty.Class("delete"),
						vecty.Attribute("aria-label", "close"),
						event.Click(func(e *vecty.Event) {
							c.close()
						}),
					),
				),
			),
			elem.Section(
				vecty.Markup(vecty.Class("modal-card-body")),
				vecty.If(c.Success, c.renderNotification()),
				elem.TextArea(
					vecty.Markup(vecty.Class("textarea"),
						vecty.Property("placeholder", "Let us know what problems you faced."),
					),
				),
			),
			elem.Footer(
				vecty.Markup(vecty.Class("modal-card-foot")),
				elem.Button(
					vecty.Markup(vecty.Class("button", "is-success"),
						event.Click(func(e *vecty.Event) {
							c.submit()
						}),
					),
					vecty.Text("Send"),
				),
				elem.Button(
					vecty.Markup(vecty.Class("button"),
						event.Click(func(e *vecty.Event) {
							c.close()
						}),
					),
					vecty.Text("Close"),
				),
			),
		),
	)
}

func (c *ReportBugModal) renderNotification() *vecty.HTML {
	return elem.Div(
		vecty.Markup(vecty.Class("notification", "is-success")),
		elem.Button(
			vecty.Markup(vecty.Class("delete"),
				event.Click(func(e *vecty.Event) {
					c.hideNotification()
				}),
			),
		),
		elem.Span(
			vecty.Markup(vecty.Class("fa", "fa-bug", "mr-1")),
		),
		vecty.Text("Thanks. Your bug has been reported."),
	)
}

func (c *ReportBugModal) close() {
	c.Open = false
	vecty.Rerender(c)
}

func (c *ReportBugModal) submit() {
	c.Success = true
	vecty.Rerender(c)
}

func (c *ReportBugModal) hideNotification() {
	c.Success = false
	vecty.Rerender(c)
}
