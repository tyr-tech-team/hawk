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
	GRPCOK                 GRPCCode = GRPCCode(codes.OK)
	GRPCCanceled           GRPCCode = GRPCCode(codes.Canceled)
	GRPCUnknown            GRPCCode = GRPCCode(codes.Unknown)
	GRPCInvalidArgument    GRPCCode = GRPCCode(codes.InvalidArgument)
	GRPCDeadlineExceeded   GRPCCode = GRPCCode(codes.DeadlineExceeded)
	GRPCNotFound           GRPCCode = GRPCCode(codes.NotFound)
	GRPCAlreadyExists      GRPCCode = GRPCCode(codes.AlreadyExists)
	GRPCPermissionDenied   GRPCCode = GRPCCode(codes.PermissionDenied)
	GRPCResourceExhausted  GRPCCode = GRPCCode(codes.ResourceExhausted)
	GRPCFailedPrecondition GRPCCode = GRPCCode(codes.FailedPrecondition)
	GRPCAborted            GRPCCode = GRPCCode(codes.Aborted)
	GRPCOutOfRange         GRPCCode = GRPCCode(codes.OutOfRange)
	GRPCUnimplemented      GRPCCode = GRPCCode(codes.Unimplemented)
	GRPCInternal           GRPCCode = GRPCCode(codes.Internal)
	GRPCUnavailable        GRPCCode = GRPCCode(codes.Unavailable)
	GRPCDataLoss           GRPCCode = GRPCCode(codes.DataLoss)
	GRPCUnauthenticated    GRPCCode = GRPCCode(codes.Unauthenticated)
)
