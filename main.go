package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/aclisp/godashboard/backend"
	"github.com/aclisp/godashboard/frontend/bundle"
	dashboard "github.com/aclisp/godashboard/proto"
)

var logger *logrus.Logger

func init() {
	logger = backend.Logger
	// Should only be done from init functions
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(logger.Out, ioutil.Discard, ioutil.Discard))
	// for pprof and trace
	grpc.EnableTracing = true
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
}

func main() {
	gs := grpc.NewServer()
	dashboard.RegisterBackendServer(gs, &backend.Backend{})
	wrappedServer := grpcweb.WrapServer(gs)

	handler := func(resp http.ResponseWriter, req *http.Request) {
		// Redirect gRPC and gRPC-Web requests to the gRPC-Web Websocket Proxy server
		if req.ProtoMajor == 2 && strings.Contains(req.Header.Get("Content-Type"), "application/grpc") {
			logger.Infof("got gRPC request: %v", req.URL.String())
			wrappedServer.ServeHTTP(resp, req)
			return
		}

		// Serve the WASM client
		wasmContentTypeSetter(http.FileServer(bundle.Assets)).ServeHTTP(resp, req)
	}
	router := mux.NewRouter()
	router.PathPrefix("/debug/").Handler(http.DefaultServeMux)
	router.PathPrefix("/").Handler(http.HandlerFunc(handler))

	addr := "localhost:10000"
	httpsSrv := &http.Server{
		Addr:    addr,
		Handler: router,
		// Some security settings
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
		},
	}

	logger.Info("Serving on https://" + addr)
	logger.Fatal(httpsSrv.ListenAndServeTLS("./insecure/cert.pem", "./insecure/key.pem"))
}

func wasmContentTypeSetter(fn http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if strings.Contains(req.URL.Path, ".wasm") {
			w.Header().Set("content-type", "application/wasm")
		}
		fn.ServeHTTP(w, req)
	}
}
