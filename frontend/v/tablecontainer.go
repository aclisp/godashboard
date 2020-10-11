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
	Data *dashboard.TableInfo `vecty:"prop"`
}

// Render a data table
func (t *TableContainer) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("card", "shadow", "mb-4")),
		elem.Div(
			vecty.Markup(vecty.Class("card-header", "py-3")),
			elem.Heading6(
				vecty.Markup(vecty.Class("m-0", "font-weight-bold", "text-primary")),
				vecty.Text(t.Data.Name),
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
	rows := make(vecty.List, 0, len(t.Data.Rows))
	for _, row := range t.Data.Rows {
		cells := make(vecty.List, 0, len(row.Infos))
		for _, cell := range row.Infos {
			cells = append(cells, t.renderCell(cell))
		}
		rows = append(rows, elem.TableRow(cells))
	}
	heads := make(vecty.List, 0, len(t.Data.Ths))
	for _, head := range t.Data.Ths {
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
	//fmt.Printf("TableContainer.Mount: id=%s\n", t.id)
	script := fmt.Sprintf(`$('#%s').DataTable();`, t.id)
	if js.Global().Get("jQuery").Truthy() {
		js.Global().Call("eval", script)
	}
}

func (t *TableContainer) renderCell(cell *dashboard.TdInfo) *vecty.HTML {
	var td *vecty.HTML
	switch cell.TdType {
	case dashboard.TdType_TDTYPE_IMG, dashboard.TdType_TDTYPE_AVATAR:
		td = elem.Image(vecty.Markup(prop.Src(cell.Content),
			style.Height(style.Px(64)),
			style.Width(style.Px(64))))
	default:
		td = vecty.Text(cell.Content)
	}
	return elem.TableData(td)
}
