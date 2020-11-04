package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Dashboard .
type Dashboard struct {
	vecty.Core

	navMenus []NavMenu
}

// NavMenu .
type NavMenu struct {
	Icon      string
	Text      string
	Link      string
	Component vecty.Component
	active    bool
}

// NewDashboard .
func NewDashboard(routePrefix string) *Dashboard {
	return &Dashboard{
		navMenus: []NavMenu{
			{Icon: "fa-tachometer", Text: "Dashboard", Link: routePrefix + "/home", Component: &Home{}},
			{Icon: "fa-book", Text: "Books", Link: routePrefix + "/books", Component: &Books{}},
			//{Icon: "fa-address-book", Text: "Customers", Link: routePrefix + "/customers"},
			//{Icon: "fa-file-text-o", Text: "Orders", Link: routePrefix + "/orders"},
		},
	}
}

// Render .
func (c *Dashboard) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&Header{},
		elem.Section(
			vecty.Markup(vecty.Class("section")),
			elem.Div(
				vecty.Markup(vecty.Class("columns")),
				&Sidebar{navMenus: c.navMenus},
				&Main{navMenus: c.navMenus},
			),
		),
	)
}
