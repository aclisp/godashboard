THIS_MODULE = github.com/aclisp/godashboard
CHECK_FILES = frontend backend tools main.go
GOIMPORTS = goimports -local $(THIS_MODULE) -format-only -l -w $(CHECK_FILES)

generate:
	protoc -Iproto proto/dashboard.proto --go_out=plugins=grpc:proto --go_opt=paths=source_relative
	go generate -x ./frontend/
	$(GOIMPORTS)
	gofmt -l -w -s $(CHECK_FILES)
	go list ./... | xargs -L1 golint
	GOOS=js GOARCH=wasm go vet ./...
	go build

serve:
	GOOS=js GOARCH=wasm go build -o frontend/html/main.wasm frontend/frontend.go
	go run -tags=dev main.go

check:
	$(GOIMPORTS)
	go list ./... | xargs -L1 golint
