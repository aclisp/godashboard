package backend

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// errorStatus provides error handling of the backend
type errorStatus struct {
	code    codes.Code
	message string
	details []proto.Message
}

// newStatus creates a default error status with the same reason as msg
// Note:
// details must be filled because `grpc-status-details-bin` is mandatory for frontend
// to correctly decode an error using `status.Convert`.
func newStatus(c codes.Code, msg string) errorStatus {
	return errorStatus{
		code:    c,
		message: msg,
		details: []proto.Message{&errdetails.ErrorInfo{
			Reason: msg,
		}},
	}
}

func (e errorStatus) err() error {
	st := status.New(e.code, e.message)
	detSt, err := st.WithDetails(e.details...)
	if err == nil {
		return detSt.Err()
	}
	return st.Err()
}
