package model

import (
	"google.golang.org/grpc/status"

	dashboard "github.com/aclisp/godashboard/proto"
)

// GatewayDefinition defines a gateway
type GatewayDefinition struct {
	Name     string
	ID       [2]string
	Selected bool
}

// PackageEndpoint is the key type
type PackageEndpoint struct {
	PackageName  string
	EndpointName string
}

// PackageEndpointData is the value type
type PackageEndpointData struct {
	Status   *status.Status
	QueryRes *dashboard.QueryRes
}
