package page

import (
	"strconv"
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/hexops/vecty/style"
)

// TableView is the page showing a table
type TableView struct {
	vecty.Core
}

// Render TableView
func (table *TableView) Render() vecty.ComponentOrHTML {
	return elem.Div(
		// Page Heading
		elem.Heading1(
			vecty.Markup(vecty.Class("h3", "mb-2", "text-gray-800")),
			vecty.Text("Tables"),
		),
		elem.Paragraph(
			vecty.Markup(vecty.Class("mb-4")),
			vecty.Text("DataTables is a third party plugin that is used to generate the demo table below."),
		),
		// DataTales Example
		table.renderTableContainer(),
	)
}

// Mount is called when the table is mounted
func (table *TableView) Mount() {
	if js.Global().Get("jQuery").Truthy() {
		js.Global().Call("eval", `$('#dataTable').DataTable();`)
	}
}

func (table *TableView) renderTableContainer() *vecty.HTML {
	return elem.Div(
		vecty.Markup(vecty.Class("card", "shadow", "mb-4")),
		elem.Div(
			vecty.Markup(vecty.Class("card-header", "py-3")),
			elem.Heading6(
				vecty.Markup(vecty.Class("m-0", "font-weight-bold", "text-primary")),
				vecty.Text("DataTables Example"),
			),
		),
		elem.Div(
			vecty.Markup(vecty.Class("card-body")),
			elem.Div(
				vecty.Markup(vecty.Class("table-responsive")),
				table.renderTable(),
			),
		),
	)
}

func (table *TableView) renderTable() *vecty.HTML {
	list := make(vecty.List, 0, 100)
	for i := 0; i < 60; i++ {
		list = append(list, elem.TableRow(
			elem.TableData(vecty.Text("-")),
			elem.TableData(vecty.Text(strconv.Itoa(i))),
			elem.TableData(vecty.Text("-")),
			elem.TableData(vecty.Text("-")),
			elem.TableData(vecty.Text("-")),
			elem.TableData(vecty.Text("-")),
		))
	}
	return elem.Table(
		vecty.Markup(
			vecty.Class("table", "table-bordered"),
			prop.ID("dataTable"),
			style.Width("100%"),
			vecty.Property("cellspacing", "0"),
		),
		elem.TableHead(
			elem.TableRow(
				elem.TableHeader(vecty.Text("Name")),
				elem.TableHeader(vecty.Text("Position")),
				elem.TableHeader(vecty.Text("Office")),
				elem.TableHeader(vecty.Text("Age")),
				elem.TableHeader(vecty.Text("Start date")),
				elem.TableHeader(vecty.Text("Salary")),
			),
		),
		elem.TableFoot(
			elem.TableRow(
				elem.TableHeader(vecty.Text("Name")),
				elem.TableHeader(vecty.Text("Position")),
				elem.TableHeader(vecty.Text("Office")),
				elem.TableHeader(vecty.Text("Age")),
				elem.TableHeader(vecty.Text("Start date")),
				elem.TableHeader(vecty.Text("Salary")),
			),
		),
		elem.TableBody(
			list,
		),
	)
}
