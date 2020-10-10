package v

import (
	"fmt"
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/hexops/vecty/style"

	dashboard "github.com/aclisp/godashboard/proto"
)

// TableContainer wraps a data table
type TableContainer struct {
	vecty.Core

	id   string
	data *dashboard.TableInfo
}

// Render a data table
func (t *TableContainer) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("card", "shadow", "mb-4")),
		elem.Div(
			vecty.Markup(vecty.Class("card-header", "py-3")),
			elem.Heading6(
				vecty.Markup(vecty.Class("m-0", "font-weight-bold", "text-primary")),
				vecty.Text(t.data.Name),
			),
		),
		elem.Div(
			vecty.Markup(vecty.Class("card-body")),
			elem.Div(
				vecty.Markup(vecty.Class("table-responsive")),
				t.renderTable(),
			),
		),
	)
}

func (t *TableContainer) renderTable() *vecty.HTML {
	rows := make(vecty.List, 0, len(t.data.Rows))
	for _, row := range t.data.Rows {
		cells := make(vecty.List, 0, len(row.Infos))
		for _, cell := range row.Infos {
			cells = append(cells, elem.TableData(vecty.Text(cell.Content)))
		}
		rows = append(rows, elem.TableRow(cells))
	}
	heads := make(vecty.List, 0, len(t.data.Ths))
	for _, head := range t.data.Ths {
		heads = append(heads, elem.TableHeader(vecty.Text(head)))
	}
	return elem.Table(
		vecty.Markup(
			vecty.Class("table", "table-bordered"),
			prop.ID(t.id),
			style.Width("100%"),
			vecty.Property("cellspacing", "0"),
		),
		elem.TableHead(elem.TableRow(heads)),
		elem.TableFoot(elem.TableRow(heads)),
		elem.TableBody(rows),
	)
}

// Mount is called when the table is mounted
func (t *TableContainer) Mount() {
	script := fmt.Sprintf(`$('#%s').DataTable();`, t.id)
	if js.Global().Get("jQuery").Truthy() {
		js.Global().Call("eval", script)
	}
}
