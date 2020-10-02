generate:
	protoc -Iproto proto/dashboard.proto --go_out=plugins=grpc:proto --go_opt=paths=source_relative
	go generate -x ./frontend/
	goimports -local github.com/aclisp/godashboard -format-only -l -w frontend backend main.go
	gofmt -l -w -s frontend backend main.go
	golint ./...
	GOOS=js GOARCH=wasm go vet ./...
	go build

serve:
	GOOS=js GOARCH=wasm go build -o frontend/html/main.wasm frontend/frontend.go
	go run -tags=dev main.go
