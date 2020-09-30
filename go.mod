module github.com/aclisp/godashboard

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hexops/vecty v0.5.1-0.20200925084631-56dda858f26b
	github.com/improbable-eng/grpc-web v0.9.6
	github.com/rs/cors v1.7.0 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546 // indirect
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/net v0.0.0-20200925080053-05aa5d4ee321
	google.golang.org/genproto v0.0.0-20200925023002-c2d885f95484
	google.golang.org/grpc v1.27.0
	marwan.io/vecty-router v0.0.0-20200914150808-f30c81f0deb5 // indirect
)

replace google.golang.org/grpc => github.com/aclisp/grpc-go v1.2.1-0.20180625151142-1f109e898476
