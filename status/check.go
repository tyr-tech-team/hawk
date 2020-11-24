package status

// CheckGRPCError -
func checkGRPCError(g GRPCCode) Status {
	switch g {
	case GRPCUnavailable:
		return RemoteHostNotFound
	default:
		return UnKnownError
	}
}
