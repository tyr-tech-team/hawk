package status

//
var (
	// NoError - 沒有錯誤
	NoError = NewStatus(LevelNONE, ServiceNONE, GRPCOK, ActionNono, "success")
	// UnknownError - 未知的錯誤
	UnKnownError = NewStatus(LevelERROR, ServiceNONE, GRPCUnknown, ActionNono, "unknown error")

	/*
	  一般常用錯誤
	*/
	// InvalidParameter - 錯誤的參數
	InvalidParameter = NewStatus(LevelWARNING, ServiceNONE, GRPCInvalidArgument, ActionCheck, "invalid Parameter")
	// ConnectFailed - 連線失敗
	ConnectFailed   = NewStatus(LevelFATAL, ServiceNONE, GRPCUnavailable, ActionConnect, "connect failed")
	ConnectTimeOut  = NewStatus(LevelWARNING, ServiceNONE, GRPCDeadlineExceeded, ActionConnect, "connect time out")
	TooManayConnect = NewStatus(LevelERROR, ServiceNONE, GRPCResourceExhausted, ActionConnect, "too manay connect")
	// Auth - 002

	// Card
	// EventLog
	// Item -
	// Brand -
	// Member -
	// NfcReader -
	// Order -
	// Storage -
	// Transaction -
	// User -
	// Website -
	// InitService -
	// SellOrder
	// BuyOrder
	CreateBuyOrderFailed = NewStatus(LevelERROR, ServiceBuyOrder, GRPCAborted, ActionCreate, "create buyOrder failed")
	// SMS
	// IDCard

)
