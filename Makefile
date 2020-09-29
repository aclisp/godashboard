generate:
	protoc -Iproto proto/dashboard.proto --go_out=plugins=grpc:proto --go_opt=paths=source_relative
	go generate -x ./frontend/
	go build

serve:
	GOOS=js GOARCH=wasm go build -o frontend/html/main.wasm frontend/frontend.go
	go run -tags=dev main.go
