package v

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// MainArea is the page main area
type MainArea struct {
	vecty.Core
}

// Render a side bar
func (m *MainArea) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(prop.ID("content")),
		// Topbar
		&Topbar{},
		// Begin Page Content
		&PageContent{},
	)
}
