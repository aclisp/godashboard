# Go dashboard

A dashboard application built in Go and gRPC using WebAssembly.

## Features

* Pure Go in both frontend and backend. No need to learn another frontend JavaScript framework.
* Type-safe Web APIs defined in gRPC+Protobuf. Move away from REST+JSON.
* Standalone executable with all the assets embedded for easy deployment.

## Development

* Run `make generate` to regenerate the protofiles and the frontend with all static resources embedded in bundle.go. For more information about embedding, check:
  - [vfsgen](https://github.com/shurcooL/vfsgen)
  - and the reason [why it is chosen](https://tech.townsourced.com/post/embedding-static-files-in-go/)
* Run `make serve` to start the web server. In this case, the static resources embedding is not used. They are served from disk as usual.
  - Reload index.html could trigger a rebuild of main.wasm, just like wasmserve. Happy hacking!
* Run `go build` (after running `make generate`) to get the standalone executable for release deployment

## Design

```
+----------------------+
|   Admin dashboard    |
+----------------------+
| frontend Golang code |
+----------+-----------+
|  golang  |  protobuf |
|  stdlib  |  message  |
|          |  objects  |
+----------+-----------+
|  browser Fetch API   |
+----------------------+
|        HTTP2         |
+----------------------+
| backend gRPC server  |
+----------------------+
|   Business System    |
+----------------------+
```

## Limitations

* To support gRPC-Web in the browser, an older grpc-go with a customized [patch](https://github.com/grpc/grpc-go/pull/2174) applied must be used. As the grpc-go and protobuf-go changed frequently, an older tool chain must be kept to build the project.

## Dependencies

* Golang 1.14+
* The Google protobuf compiler, `protoc`.
* $ GO111MODULE=on go get github.com/golang/protobuf/protoc-gen-go@v1.3.2 (do not use cmd/protoc-gen-go-grpc)
* github.com/golang/protobuf (do not use google.golang.org/protobuf)
* github.com/aclisp/grpc-go @branch add-grpc-web-client (forked from johanbrandhorst/grpc-go in case the upstream might be destroyed)
* github.com/improbable-eng/grpc-web @v0.9.6 (this is the latest version work with grpc-go above, since which CloseNotifier is [removed](https://github.com/improbable-eng/grpc-web/pull/478))

## References

* [dotGo 2019 - Johan Brandhorst - Get Going With WebAssembly](https://youtu.be/osVHH7rjpxs?t=773)
* [Live Go module index](https://dmitri.shuralyov.com/projects/live-module-index/)
* [Go WebAssembly Wiki](https://github.com/golang/go/wiki/WebAssembly)
* [The state of gRPC in the browser](https://grpc.io/blog/state-of-grpc-web/)
