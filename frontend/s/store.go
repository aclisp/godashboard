package s

import (
	"google.golang.org/grpc/grpclog"

	"github.com/aclisp/godashboard/frontend/s/action"
	"github.com/aclisp/godashboard/frontend/s/dispatcher"
	"github.com/aclisp/godashboard/frontend/s/model"
	"github.com/aclisp/godashboard/frontend/s/storeutil"
)

var (
	// Gateways is the avaiable gateways
	Gateways = []model.GatewayDefinition{
		{Name: "Indonesia", ID: [2]string{"product", "ID"}, Selected: false},
		{Name: "Singapore", ID: [2]string{"product", "SG"}, Selected: true},
		{Name: "USA", ID: [2]string{"product", "US"}, Selected: false},
	}
	// CurrentGatewayID is the current gateway identity
	CurrentGatewayID = [2]string{"product", "ID"}

	// PackageEndpointData is the data
	PackageEndpointData = make(map[model.PackageEndpoint]model.PackageEndpointData)

	// CurrentPackageEndpoint is the current package-endpoint
	CurrentPackageEndpoint model.PackageEndpoint

	// Listeners is the listeners that will be invoked when the store changes.
	Listeners = storeutil.NewListenerRegistry()

	updater dynamicViewUpdater
)

func init() {
	dispatcher.Register(actions)

	for _, gateway := range Gateways {
		if gateway.Selected {
			CurrentGatewayID = gateway.ID
		}
	}
}

// actions is the only place to mutate states
func actions(act interface{}) {
	grpclog.Infof("%#v", act)

	switch a := act.(type) {
	case *action.ChangeGateway:
		CurrentGatewayID = a.GatewayID

	case *action.ChangePackageEndpoint:
		if parsed, ok := a.ParseRoute(); ok {
			CurrentPackageEndpoint = parsed
		}

	case *action.StartDynamicViewUpdating:
		updater.start()
	case *action.StopDynamicViewUpdating:
		updater.stop()
	case *action.SyncDynamicViewData:
		updater.sync()
	case *action.SyncDynamicViewDataDone:
		a.SaveTo(PackageEndpointData)

	default:
		return // don't fire listeners
	}

	Listeners.Fire()
}
