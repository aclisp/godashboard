package view

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"golang.org/x/sync/errgroup"

	"github.com/aclisp/godashboard/frontend/state"
)

// Home .
type Home struct {
	vecty.Core

	ping1 [2]string
	ping2 [2]string
}

// Render .
func (c *Home) Render() vecty.ComponentOrHTML {
	state.AddListener(c, func() { vecty.Rerender(c) })

	return elem.Div(
		elem.Heading1(
			vecty.Markup(vecty.Class("title")),
			vecty.Text("Home"),
		),
		elem.Paragraph(vecty.Text(fmt.Sprintf("%q", c.ping1))),
		elem.Paragraph(vecty.Text(fmt.Sprintf("%q", c.ping2))),
	)
}

// Mount .
func (c *Home) Mount() {
	c.ping1 = [2]string{"Hello-World", ""}
	c.ping2 = [2]string{"Expect-Error", ""}

	go func() {
		g := new(errgroup.Group)
		g.Go(func() error {
			c.ping1[1] = state.PingBackend(c.ping1[0])
			return nil
		})
		g.Go(func() error {
			c.ping2[1] = state.PingBackend(c.ping2[0])
			return nil
		})
		g.Wait()
		vecty.Rerender(c)
	}()
}
