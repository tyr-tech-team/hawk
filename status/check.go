package status

// CheckGRPCError -
func checkGRPCError(g GRPCCode) Status {
	switch g {
	case GRPCUnavailable:
		return ConnectFailed
	case GRPCUnimplemented:
		return RemoteHostNotFound
	default:
		return UnKnownError
	}
}
