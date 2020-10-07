package s

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	dashboard "github.com/aclisp/godashboard/proto"
)

var client dashboard.BackendClient

// Init the s (state) package
func Init() {
	cc, err := grpc.Dial("")
	if err != nil {
		grpclog.Errorf("state init: grpc dial: %v", err)
		return
	}

	client = dashboard.NewBackendClient(cc)

	// example backend communication
	pingBackend(client, "Hello-World")
	// example backend communication for error handling
	pingBackend(client, "Expect-Error")

	SidebarMenus = getSidebarMenus(client)

	grpclog.Infof("state init finished")
}

func pingBackend(c dashboard.BackendClient, message string) {
	resp, err := c.Ping(context.Background(), &dashboard.Hello{
		Message: message,
	})
	if err != nil {
		st := status.Convert(err)
		grpclog.Errorf("pingBackend: %v %q details: %v", st.Code(), st.Message(), st.Details())
		return
	}
	grpclog.Infof("pingBackend: %v", resp)
}
