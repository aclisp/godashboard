package s

import (
	"encoding/json"

	"google.golang.org/grpc/grpclog"

	"github.com/aclisp/godashboard/frontend/s/action"
	"github.com/aclisp/godashboard/frontend/s/dispatcher"
	"github.com/aclisp/godashboard/frontend/s/model"
	"github.com/aclisp/godashboard/frontend/s/storeutil"
)

var (
	// State is the app local state, which can be put in the browser localStorage.
	State = NewLocalState()

	// Backend is the data from backend
	Backend = NewBackendData()

	// Listeners is the listeners that will be invoked when the store changes.
	Listeners = storeutil.NewListenerRegistry()

	// updater is the backend data updater
	updater dynamicViewUpdater
)

// LocalState is the app local state, which can be put in the browser localStorage.
type LocalState struct {
	localState
	// SyncOngoing is true when there is an outstanding sync data request
	syncOngoing bool
}

// BackendData is the data from backend
type BackendData struct {
	data map[model.PackageEndpoint]model.PackageEndpointData
}

func init() {
	dispatcher.Register(actions)
}

// actions is the only place to mutate states
func actions(act interface{}) {
	grpclog.Infof("%#v", act)

	switch a := act.(type) {
	case *action.ChangeGateway:
		State.setGateway(a)

	case *action.ChangePackageEndpoint:
		if parsed, ok := a.ParseRoute(); ok {
			State.localState.CurrentPackageEndpoint = parsed
		}

	case *action.StartDynamicViewUpdating:
		updater.start()

	case *action.StopDynamicViewUpdating:
		updater.stop()

	case *action.SyncDynamicViewData:
		State.syncOngoing = true
		updater.sync(a.StartUpdater)

	case *action.SyncDynamicViewDataDone:
		a.SaveTo(Backend.data)
		State.syncOngoing = false

	case *action.ReplaceState:
		if err := json.Unmarshal([]byte(a.StateJSON), &State); err != nil {
			grpclog.Errorf("failed to unmarshal state: %v", err)
		}

	default:
		return // don't fire listeners
	}

	Listeners.Fire()
}

// Gateways is the avaiable gateways
func (s *LocalState) Gateways() []model.GatewayDefinition {
	return s.localState.Gateways
}

// CurrentPackageEndpoint is the current package-endpoint
func (s *LocalState) CurrentPackageEndpoint() model.PackageEndpoint {
	return s.localState.CurrentPackageEndpoint
}

// CurrentGatewayID is the current gateway identity
func (s *LocalState) CurrentGatewayID() [2]string {
	return s.localState.CurrentGatewayID
}

// SyncOngoing is true when there is an outstanding sync data request
func (s *LocalState) SyncOngoing() bool {
	return s.syncOngoing
}

// CurrentPackageEndpointData is the current package-endpoint data
func (b *BackendData) CurrentPackageEndpointData() model.PackageEndpointData {
	return b.data[State.CurrentPackageEndpoint()]
}

type localState struct {
	Gateways               []model.GatewayDefinition
	CurrentGatewayID       [2]string
	CurrentPackageEndpoint model.PackageEndpoint
}

// UnmarshalJSON implements json.Unmarshaler
func (s *LocalState) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.localState)
}

// MarshalJSON implements json.Marshaler
func (s *LocalState) MarshalJSON() ([]byte, error) {
	return json.Marshal(&s.localState)
}

func (s *LocalState) setGateway(a *action.ChangeGateway) {
	s.localState.CurrentGatewayID = a.GatewayID
	for i := range s.localState.Gateways {
		gateway := &s.localState.Gateways[i]
		if gateway.ID == a.GatewayID {
			gateway.Selected = true
		} else {
			gateway.Selected = false
		}
	}
}

// NewLocalState creates a local state with predefined values
func NewLocalState() *LocalState {
	s := new(LocalState)
	gateways := []model.GatewayDefinition{
		{Name: "Indonesia", ID: [2]string{"product", "ID"}},
		{Name: "Singapore", ID: [2]string{"product", "SG"}, Selected: true},
		{Name: "USA", ID: [2]string{"product", "US"}},
		{Name: "India", ID: [2]string{"product", "IN"}},
		{Name: "Arab", ID: [2]string{"product", "AE"}},
		{Name: "Brazil", ID: [2]string{"product", "BR"}},
		{Name: "Russia", ID: [2]string{"product", "RU"}},

		{Name: "Indonesia - testing", ID: [2]string{"testing", "ID"}},
		{Name: "Singapore - testing", ID: [2]string{"testing", "SG"}},
		{Name: "India - testing", ID: [2]string{"testing", "IN"}},
	}
	for _, gateway := range gateways {
		if gateway.Selected {
			s.localState.CurrentGatewayID = gateway.ID
		}
	}
	s.localState.Gateways = gateways
	return s
}

// NewBackendData creates an empty backend data
func NewBackendData() *BackendData {
	b := new(BackendData)
	b.data = make(map[model.PackageEndpoint]model.PackageEndpointData)
	return b
}
