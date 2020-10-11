package backend

import (
	"context"
	"strings"

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
				Id:     "money-api-redpacket",
				FaIcon: "fa-envelope",
				Text:   "红包",
				Groups: []*dashboard.SidebarGroup{
					{
						Items: []*dashboard.SidebarEntry{
							{
								Text:  "档位配置",
								Route: "/go/money-api-redpacket/Config",
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

// Query data
func (b *Backend) Query(ctx context.Context, req *dashboard.QueryReq) (*dashboard.QueryRes, error) {
	//logger.Debugf("Backend.Query: req: %v", req)
	res := new(dashboard.QueryRes)
	err := RoundTrip(req.YsadminReq, Link{
		FromService: "ysadminv2",
		URL:         RegionalURL[req.Environment][req.GatewayCode],
		ToService:   "net.ihago." + strings.ReplaceAll(req.Package, "-", "."),
		Method:      "YsadminRPCService.Query",
	}, res)
	if err != nil {
		return nil, newStatus(codes.Unavailable, err.Error()).err()
	}
	return res, nil
	/*
		return &dashboard.QueryRes{
			Tables: []*dashboard.TableInfo{
				{
					Name: "Table One",
					Ths:  []string{"Column One", "Column Two", "Column Three"},
					Rows: []*dashboard.TdRow{
						{
							Infos: []*dashboard.TdInfo{
								{Content: "aaa"},
								{Content: "bbb"},
								{Content: "ccc"},
							},
						},
						{
							Infos: []*dashboard.TdInfo{
								{Content: "xxx"},
								{Content: "yyy"},
								{Content: "zzz"},
							},
						},
						{
							Infos: []*dashboard.TdInfo{
								{Content: "111"},
								{Content: "222"},
								{Content: "333"},
							},
						},
					},
				},
			},
		}, nil
	*/
}

// Commit data
func (b *Backend) Commit(ctx context.Context, req *dashboard.CommitReq) (*dashboard.CommitRes, error) {
	return nil, nil
}
