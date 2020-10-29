package status

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

// GRPCCode -
type GRPCCode int64

func (g GRPCCode) String() string {
	return fmt.Sprintf("%02d", g)
}

//
const (
	OK                 GRPCCode = GRPCCode(codes.OK)
	Canceled           GRPCCode = GRPCCode(codes.Canceled)
	Unknown            GRPCCode = GRPCCode(codes.Unknown)
	InvalidArgument    GRPCCode = GRPCCode(codes.InvalidArgument)
	DeadlineExceeded   GRPCCode = GRPCCode(codes.DeadlineExceeded)
	NotFound           GRPCCode = GRPCCode(codes.NotFound)
	AlreadyExists      GRPCCode = GRPCCode(codes.AlreadyExists)
	PermissionDenied   GRPCCode = GRPCCode(codes.PermissionDenied)
	ResourceExhausted  GRPCCode = GRPCCode(codes.ResourceExhausted)
	FailedPrecondition GRPCCode = GRPCCode(codes.FailedPrecondition)
	Aborted            GRPCCode = GRPCCode(codes.Aborted)
	OutOfRange         GRPCCode = GRPCCode(codes.OutOfRange)
	Unimplemented      GRPCCode = GRPCCode(codes.Unimplemented)
	Internal           GRPCCode = GRPCCode(codes.Internal)
	Unavailable        GRPCCode = GRPCCode(codes.Unavailable)
	DataLoss           GRPCCode = GRPCCode(codes.DataLoss)
	Unauthenticated    GRPCCode = GRPCCode(codes.Unauthenticated)
)
