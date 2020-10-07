package backend

import (
	"context"

	"google.golang.org/grpc/codes"

	dashboard "github.com/aclisp/godashboard/proto"
)

// Backend should be used to implement the server interface
// exposed by the generated server proto.
type Backend struct {
}

// Ensure struct implements interface
var _ dashboard.BackendServer = (*Backend)(nil)

// Ping test if the backend alive
func (b *Backend) Ping(ctx context.Context, req *dashboard.Hello) (*dashboard.Pong, error) {
	if req.GetMessage() == "Expect-Error" {
		return nil, newStatus(codes.Canceled, "operation canceled because client sent Expect-Error").err()
	}
	return &dashboard.Pong{
		Reply: req.GetMessage(),
	}, nil
}

// GetSidebarMenus gets all side-bar menus
func (b *Backend) GetSidebarMenus(ctx context.Context, req *dashboard.GetSidebarMenusReq) (*dashboard.GetSidebarMenusRes, error) {
	return &dashboard.GetSidebarMenusRes{
		Menus: []*dashboard.SidebarMenu{
			{
				Id:     "package1",
				FaIcon: "fa-cog",
				Text:   "Package 1",
				Groups: []*dashboard.SidebarGroup{
					{
						Text: "Service 1:",
						Items: []*dashboard.SidebarEntry{
							{
								Text:  "Method 1",
								Route: "/go/package1/Service1-Method1",
							},
							{
								Text:  "Method 2",
								Route: "/go/package1/Service1-Method2",
							},
						},
					},
				},
			},
			{
				Id:     "package2",
				FaIcon: "fa-wrench",
				Text:   "Package 2",
				Groups: []*dashboard.SidebarGroup{
					{
						Text: "Service 1:",
						Items: []*dashboard.SidebarEntry{
							{
								Text:  "Method 1",
								Route: "/go/package2/Service1-Method1",
							},
							{
								Text:  "Method 2",
								Route: "/go/package2/Service1-Method2",
							},
							{
								Text:  "Method 3",
								Route: "/go/package2/Service1-Method3",
							},
							{
								Text:  "Method 4",
								Route: "/go/package2/Service1-Method4",
							},
						},
					},
				},
			},
			{
				Id:     "package3",
				FaIcon: "fa-folder",
				Text:   "Package 3",
				Groups: []*dashboard.SidebarGroup{
					{
						Text: "Service 1:",
						Items: []*dashboard.SidebarEntry{
							{
								Text:  "Method 1",
								Route: "/go/package3/Service1-Method1",
							},
							{
								Text:  "Method 2",
								Route: "/go/package3/Service1-Method2",
							},
							{
								Text:  "Method 3",
								Route: "/go/package3/Service1-Method3",
							},
						},
					},
					{
						Text: "Service 2:",
						Items: []*dashboard.SidebarEntry{
							{
								Text:  "Method 1",
								Route: "/go/package3/Service2-Method1",
							},
							{
								Text:  "Method 2",
								Route: "/go/package3/Service2-Method2",
							},
						},
					},
				},
			},
		},
	}, nil
}
