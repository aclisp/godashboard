package view

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Body renders the <body> tag
type Body struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (b *Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Text("Hello Vecty!"),
	)
}
