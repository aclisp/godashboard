package s

import (
	"fmt"

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
	// Listeners is the listeners that will be invoked when the store changes.
	Listeners = storeutil.NewListenerRegistry()
)

func init() {
	dispatcher.Register(actions)

	for _, gateway := range Gateways {
		if gateway.Selected {
			CurrentGatewayID = gateway.ID
		}
	}
}

func actions(act interface{}) {
	fmt.Printf("%#v\n", act)

	switch a := act.(type) {
	case *action.ChangeGateway:
		CurrentGatewayID = a.GatewayID

	default:
		return // don't fire listeners
	}

	Listeners.Fire()
}
