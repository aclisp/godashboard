package backend

import (
	"context"

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
	return &dashboard.Pong{
		Reply: req.GetMessage(),
	}, nil
}
