package status

//
var (
	// NoError - 沒有錯誤
	NoError = NewStatus(LevelNONE, ServiceNONE, GRPCOK, ActionNono, "success")
	// UnknownError - 未知的錯誤
	UnknownError = NewStatus(LevelERROR, ServiceNONE, GRPCUnknown, ActionNono, "unknown error")
)
