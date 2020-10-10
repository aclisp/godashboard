package action

import (
	"strings"

	"google.golang.org/grpc/status"

	"github.com/aclisp/godashboard/frontend/s/model"
	dashboard "github.com/aclisp/godashboard/proto"
)

// ChangeGateway changes the gateway
type ChangeGateway struct {
	GatewayID [2]string
}

// ChangePackageEndpoint changes the package-endpoint
type ChangePackageEndpoint struct {
	Route string
}

// ParseRoute converts route url to package and endpoint
func (a ChangePackageEndpoint) ParseRoute() (parsed model.PackageEndpoint, ok bool) {
	if !strings.HasPrefix(a.Route, "/go/") {
		return parsed, false
	}
	parts := strings.SplitN(a.Route, "/", 4)
	for i, part := range parts {
		switch i {
		case 2:
			parsed.PackageName = part
		case 3:
			parsed.EndpointName = part
		}
	}
	return parsed, true
}

// StartDynamicViewUpdating starts updating
type StartDynamicViewUpdating struct {
}

// StopDynamicViewUpdating stops updating
type StopDynamicViewUpdating struct {
}

// SyncDynamicViewData sync data with backend
type SyncDynamicViewData struct {
}

// SyncDynamicViewDataDone is dispatched on sync done
type SyncDynamicViewDataDone struct {
	Req *dashboard.QueryReq
	Res *dashboard.QueryRes
	Sta *status.Status
}

// SaveTo saves data to store
func (a SyncDynamicViewDataDone) SaveTo(store map[model.PackageEndpoint]model.PackageEndpointData) {
	key := model.PackageEndpoint{
		PackageName:  a.Req.Package,
		EndpointName: a.Req.Endpoint,
	}
	val := model.PackageEndpointData{
		Status:   a.Sta,
		QueryRes: a.Res,
	}
	store[key] = val
}
