package state

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	dashboard "github.com/aclisp/godashboard/proto"
)

var client dashboard.BackendClient

func init() {
	cc, err := grpc.Dial("")
	if err != nil {
		log.Panicf("state init: grpc dial: %v", err)
	}
	client = dashboard.NewBackendClient(cc)
}

// PingBackend .
func PingBackend(message string) (reply string) {
	resp, err := client.Ping(context.Background(), &dashboard.Hello{
		Message: message,
	})
	if err != nil {
		st := status.Convert(err)
		grpclog.Println(st.Code(), st.Message(), st.Details())
		return st.Message()
	}
	grpclog.Println(resp)
	return resp.Reply
}
