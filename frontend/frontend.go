package main

import (
	"io/ioutil"
	"os"
	"syscall/js"

	"github.com/hexops/vecty"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/grpclog"

	"github.com/aclisp/godashboard/frontend/view"
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

func grpcSetDivLogger() {
	document = js.Global().Get("document")
	div := document.Call("getElementById", "target")
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(DivWriter(div), ioutil.Discard, ioutil.Discard))
}

func grpcSetConsoleLogger() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard))
}

func init() {
	grpcSetConsoleLogger()
}

func removeContentLoadingIndicator() {
	document = js.Global().Get("document")
	div := document.Call("getElementById", "contents-loader")
	div.Call("remove")
}

func main() {
	removeContentLoadingIndicator()

	vecty.SetTitle("Go Dashboard")
	vecty.RenderBody(&view.Body{})
}
