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

	// State is the app local state, which can be put in the browser localStorage.
	State LocalState

	// Backend is the data from backend
	Backend BackendData

	// Listeners is the listeners that will be invoked when the store changes.
	Listeners = storeutil.NewListenerRegistry()

	// updater is the backend data updater
	updater dynamicViewUpdater
)

// LocalState is the app local state, which can be put in the browser localStorage.
type LocalState struct {
	// CurrentGatewayID is the current gateway identity
	currentGatewayID [2]string
	// CurrentPackageEndpoint is the current package-endpoint
	currentPackageEndpoint model.PackageEndpoint
	// SyncOngoing is true when there is an outstanding sync data request
	syncOngoing bool
}

// BackendData is the data from backend
type BackendData struct {
	data map[model.PackageEndpoint]model.PackageEndpointData
}

func init() {
	dispatcher.Register(actions)

	for _, gateway := range Gateways {
		if gateway.Selected {
			State.currentGatewayID = gateway.ID
		}
	}

	Backend.data = make(map[model.PackageEndpoint]model.PackageEndpointData)
}

// actions is the only place to mutate states
func actions(act interface{}) {
	grpclog.Infof("%#v", act)

	switch a := act.(type) {
	case *action.ChangeGateway:
		State.currentGatewayID = a.GatewayID

	case *action.ChangePackageEndpoint:
		if parsed, ok := a.ParseRoute(); ok {
			State.currentPackageEndpoint = parsed
		}

	case *action.StartDynamicViewUpdating:
		updater.start()

	case *action.StopDynamicViewUpdating:
		updater.stop()

	case *action.SyncDynamicViewData:
		State.syncOngoing = true
		updater.sync()

	case *action.SyncDynamicViewDataDone:
		a.SaveTo(Backend.data)
		State.syncOngoing = false

	default:
		return // don't fire listeners
	}

	Listeners.Fire()
}

// CurrentPackageEndpoint is the current package-endpoint
func (s *LocalState) CurrentPackageEndpoint() model.PackageEndpoint {
	return s.currentPackageEndpoint
}

// CurrentGatewayID is the current gateway identity
func (s *LocalState) CurrentGatewayID() [2]string {
	return s.currentGatewayID
}

// SyncOngoing is true when there is an outstanding sync data request
func (s *LocalState) SyncOngoing() bool {
	return s.syncOngoing
}

// CurrentPackageEndpointData is the current package-endpoint data
func (b *BackendData) CurrentPackageEndpointData() model.PackageEndpointData {
	return b.data[State.CurrentPackageEndpoint()]
}
