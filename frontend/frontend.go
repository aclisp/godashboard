package main

import (
	"context"
	"io/ioutil"
	"syscall/js"

	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	dashboard "github.com/aclisp/godashboard/proto"
)

// Build with Go WASM fork

//go:generate bash -c "GOOS=js GOARCH=wasm go build -o ./html/main.wasm frontend.go"

//go:generate bash -c "cp $DOLLAR(go env GOROOT)/misc/wasm/wasm_exec.js ./html/wasm_exec.js"

// Integrate generated JS into a Go file for static loading.
//go:generate bash -c "go run -mod=mod assets_generate.go"

var document js.Value

// DivWriter is an io.Writer
type DivWriter js.Value

func (d DivWriter) Write(p []byte) (n int, err error) {
	node := document.Call("createElement", "div")
	node.Set("innerHTML", string(p))
	js.Value(d).Call("appendChild", node)
	return len(p), nil
}

func init() {
	document = js.Global().Get("document")
	div := document.Call("getElementById", "target")
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(DivWriter(div), ioutil.Discard, ioutil.Discard))
}

func removeContentLoadingIndicator() {
	document = js.Global().Get("document")
	div := document.Call("getElementById", "contents-loader")
	div.Call("remove")
}

func main() {
	removeContentLoadingIndicator()

	cc, err := grpc.Dial("")
	if err != nil {
		grpclog.Println(err)
		return
	}
	client := dashboard.NewBackendClient(cc)

	// example backend communication
	pingBackend(client, "Hello-World")
	// example backend communication for error handling
	pingBackend(client, "Expect-Error")

	grpclog.Println("finished")
}

func pingBackend(c dashboard.BackendClient, message string) {
	resp, err := c.Ping(context.Background(), &dashboard.Hello{
		Message: message,
	})
	if err != nil {
		st := status.Convert(err)
		grpclog.Println(st.Code(), st.Message(), st.Details())
		return
	}
	grpclog.Println(resp)
}
