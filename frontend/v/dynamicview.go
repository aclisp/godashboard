package v

import (
	"fmt"
	"strings"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"google.golang.org/grpc/status"
	router "marwan.io/vecty-router"

	"github.com/aclisp/godashboard/frontend/s"
	"github.com/aclisp/godashboard/frontend/s/action"
	"github.com/aclisp/godashboard/frontend/s/dispatcher"
	dashboard "github.com/aclisp/godashboard/proto"
)

// DynamicView is rendered by backend-returned data
type DynamicView struct {
	vecty.Core
}

func (d *DynamicView) getPackageAndEndpoint() (packageName, endpointName string) {
	// vars := router.GetNamedVar(d)
	// packageName = vars["package"]
	// endpointName = vars["endpoint"]
	// ...here we use state instead
	packageName = s.State.CurrentPackageEndpoint().PackageName
	endpointName = s.State.CurrentPackageEndpoint().EndpointName
	return
}

// Render DynamicView
func (d *DynamicView) Render() vecty.ComponentOrHTML {
	packageName, endpointName := d.getPackageAndEndpoint()

	if !s.Listeners.Has(d) {
		s.Listeners.Add(d, func() {
			vecty.Rerender(d)
		})
	}

	data := s.Backend.CurrentPackageEndpointData()
	var progress, status, queryRes vecty.ComponentOrHTML
	if s.State.SyncOngoing() {
		progress = d.renderProgress()
	} else {
		if data.Status != nil {
			status = d.renderStatus(data.Status)
		}
		if data.QueryRes != nil {
			queryRes = d.renderQueryRes(data.QueryRes)
		}
	}

	return elem.Div(
		// Page Heading
		elem.Div(
			vecty.Markup(vecty.Class("d-sm-flex", "align-items-center", "justify-content-between", "mb-4")),
			elem.Heading1(
				vecty.Markup(vecty.Class("h3", "mb-0", "text-gray-800")),
				vecty.Text("Dynamic View"),
			),
			elem.Anchor(
				vecty.Markup(
					prop.Href("#"),
					vecty.Class("d-none", "d-sm-inline-block", "btn", "btn-sm", "btn-primary", "shadow-sm"),
					event.Click(func(e *vecty.Event) {
						dispatcher.Dispatch(&action.SyncDynamicViewData{})
					}).PreventDefault(),
				),
				elem.Italic(
					vecty.Markup(
						vecty.Class("fas", "fa-sync-alt", "fa-sm", "text-white-50"),
					),
				),
				vecty.Text(" Sync Data"),
			),
		),
		elem.Paragraph(
			vecty.Markup(vecty.Class("mb-4")),
			vecty.Text(fmt.Sprintf("package is %q, endpoint is %q, gateway is %v",
				"net.ihago."+strings.ReplaceAll(packageName, "-", "."),
				endpointName, s.State.CurrentGatewayID())),
		),
		progress,
		status,
		queryRes,
	)
}

// Mount is called when the view is mounted
func (d *DynamicView) Mount() {
	vars := router.GetNamedVar(d)
	packageName := vars["package"]
	endpointName := vars["endpoint"]

	dispatcher.Dispatch(&action.StartDynamicViewUpdating{})
	if !s.State.SyncOngoing() {
		dispatcher.Dispatch(&action.ChangePackageEndpoint{Route: fmt.Sprintf("/go/%s/%s", packageName, endpointName)})
		dispatcher.Dispatch(&action.SyncDynamicViewData{})
	}
}

// Unmount is called when the view is unmounted
func (d *DynamicView) Unmount() {
	dispatcher.Dispatch(&action.StopDynamicViewUpdating{})
}

func (d *DynamicView) renderProgress() *vecty.HTML {
	return elem.Paragraph(
		vecty.Markup(vecty.Class("mb-4")),
		vecty.Text("Syncing data, please wait..."),
	)
}

func (d *DynamicView) renderStatus(st *status.Status) *vecty.HTML {
	return elem.Paragraph(
		vecty.Markup(vecty.Class("mb-4")),
		vecty.Text(fmt.Sprintf("error: %v %q details: %v", st.Code(), st.Message(), st.Details())),
	)
}

func (d *DynamicView) renderQueryRes(res *dashboard.QueryRes) vecty.ComponentOrHTML {
	tables := make(vecty.List, len(res.Tables))
	for i := range res.Tables {
		tables[i] = &TableContainer{
			id:   fmt.Sprintf("dataTable-%d", i),
			Data: res.Tables[i],
		}
	}
	return tables
}
