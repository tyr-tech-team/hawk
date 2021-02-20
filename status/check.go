package status

import "fmt"

// CheckGRPCError -
func checkGRPCError(g GRPCCode) Status {
	switch g {
	case GRPCUnavailable:
		return ConnectFailed
	case GRPCUnimplemented:
		return RemoteHostNotFound
	default:
		fmt.Println(g.String())
		return UnKnownError
	}
}
