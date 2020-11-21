package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

// Dashboard .
type Dashboard struct {
	vecty.Core

	navMenus      []NavMenu
	openBugReport bool
	openSignOut   bool
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
		&Header{OnReportBug: c.onReportBug, OnSignOut: c.onSignOut},
		elem.Section(
			vecty.Markup(vecty.Class("section")),
			elem.Div(
				vecty.Markup(vecty.Class("columns")),
				&Sidebar{NavMenus: c.navMenus},
				&Main{NavMenus: c.navMenus},
			),
		),
		&ReportBugModal{Open: c.openBugReport, Success: false},
		&SignOutModal{Open: c.openSignOut, OnConfirm: c.onConfirmSignOut},
	)
}

func (c *Dashboard) onReportBug() {
	c.openBugReport = true
	c.openSignOut = false
	vecty.Rerender(c)
}

func (c *Dashboard) onSignOut() {
	c.openBugReport = false
	c.openSignOut = true
	vecty.Rerender(c)
}

func (c *Dashboard) onConfirmSignOut() {
	c.openBugReport = false
	c.openSignOut = false
	router.Redirect("/go/login")
}
