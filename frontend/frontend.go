package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"syscall/js"

	"github.com/hexops/vecty"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/grpclog"

	"github.com/aclisp/godashboard/frontend/s"
	"github.com/aclisp/godashboard/frontend/s/action"
	"github.com/aclisp/godashboard/frontend/s/dispatcher"
	"github.com/aclisp/godashboard/frontend/v"
)

//go:generate bash -c "./jsbundle.sh"

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
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard))
}

func removeContentLoadingIndicator() {
	document = js.Global().Get("document")
	div := document.Call("getElementById", "contents-loader")
	div.Call("remove")
}

func main() {
	removeContentLoadingIndicator()
	s.Init()
	attachLocalStorage()
	if err := vecty.RenderInto("body", &v.Body{}); err != nil {
		panic(err)
	}
	select {} // run Go forever
}

func attachLocalStorage() {
	s.Listeners.Add(nil, func() {
		data, err := json.Marshal(s.State)
		if err != nil {
			grpclog.Errorf("failed to marshal state: %v", err)
			return
		}
		js.Global().Get("localStorage").Call("setItem", "state", string(data))
	})

	if data := js.Global().Get("localStorage").Call("getItem", "state"); !data.IsNull() {
		dispatcher.Dispatch(&action.ReplaceState{StateJSON: data.String()})
	}
}
