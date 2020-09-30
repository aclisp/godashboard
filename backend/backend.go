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
