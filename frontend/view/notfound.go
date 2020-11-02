package view

import "github.com/hexops/vecty"

// NotFound .
type NotFound struct {
	vecty.Core
}

// Render .
func (c *NotFound) Render() vecty.ComponentOrHTML {
	return vecty.Text("404 page not found")
}
