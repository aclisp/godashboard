package s

import (
	"context"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	dashboard "github.com/aclisp/godashboard/proto"
)

// SidebarMenus is retrieved at init
var SidebarMenus []*dashboard.SidebarMenu

func getSidebarMenus(c dashboard.BackendClient) []*dashboard.SidebarMenu {
	resp, err := c.GetSidebarMenus(context.Background(), &dashboard.GetSidebarMenusReq{})
	if err != nil {
		st := status.Convert(err)
		grpclog.Errorf("getSidebarMenus: %v %q details: %v", st.Code(), st.Message(), st.Details())
		return nil
	}
	return resp.Menus
}
