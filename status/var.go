package status

//
var (
	NoError          = NewStatus(LevelNONE, ServiceNONE, GRPCOK, ActionNono, "成功", "success")
	UnKnownError     = NewStatus(LevelERROR, ServiceNONE, GRPCUnknown, ActionNono, "未知的錯誤", "unknown error")
	InvalidParameter = NewStatus(LevelWARNING, ServiceNormal, GRPCInvalidArgument, ActionCheck, "錯誤的參數", "invalid parameter")
	ConnectTimeOut   = NewStatus(LevelWARNING, ServiceNormal, GRPCDeadlineExceeded, ActionConnect, "連線超時", "connect time out")
	NotFound         = NewStatus(LevelWARNING, ServiceNormal, GRPCNotFound, ActionFind, "資料找不到", "not found")
	DataWasExisted   = NewStatus(LevelERROR, ServiceNormal, GRPCAlreadyExists, ActionCheck, "資料已存在", "data was existed")
	CreatedFailed    = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionCreate, "建立失敗", "created failed")
	UpdatedFailed    = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionUpdate, "更新失敗", "updated failed")
	DeletedFailed    = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionDelete, "刪除失敗", "deleted failed")
	DecodedFailed    = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionDecode, "解碼失敗", "decoded failed")
	EncodedFailed    = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionEncode, "編碼失敗", "encoded failed")
	ExecutedFailed   = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionExecute, "執行失敗", "executed failed")
	OpenedFailed     = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionOpen, "開啟失敗", "opened failed")
	DataLoss         = NewStatus(LevelERROR, ServiceNormal, GRPCDataLoss, ActionCheck, "資料已遺失", "data loss")
	PermissionDenied = NewStatus(LevelFATAL, ServiceNormal, GRPCPermissionDenied, ActionExecute, "權限被拒絕", "permission denied")
	ConnectFailed    = NewStatus(LevelFATAL, ServiceNormal, GRPCUnavailable, ActionConnect, "連線失敗", "connect failed")
	TooManayConnect  = NewStatus(LevelERROR, ServiceNormal, GRPCResourceExhausted, ActionConnect, "太多連線", "too manay connect")

	// Auth - 002
	TokenGenerationFailed = NewStatus(LevelERROR, ServiceAuth, GRPCAborted, ActionCreate, "權杖產生失敗", "token generation failed")
	TokenWasExpired       = NewStatus(LevelFATAL, ServiceAuth, GRPCDeadlineExceeded, ActionCheck, "權杖已失效", "token was expired")
	InvalidToken          = NewStatus(LevelFATAL, ServiceAuth, GRPCUnauthenticated, ActionCheck, "錯誤的權杖", "invalid token")
	// Card
	CardWasNotEmpty = NewStatus(LevelERROR, ServiceCard, GRPCAlreadyExists, ActionCheck, "卡片不是空的", "card was not empty")
	// EventLog
	// Item -
	ItemParameterInvalid = NewStatus(LevelWARNING, ServiceItem, GRPCInvalidArgument, ActionCheck, "商品參數錯誤", "item parmeter invalid")
	// Brand -
	BrandNotFound  = NewStatus(LevelWARNING, ServiceBrand, GRPCNotFound, ActionFind, "找不到品牌", "brand not found")
	BrandWasExists = NewStatus(LevelERROR, ServiceBrand, GRPCAlreadyExists, ActionCheck, "品牌已存在", "brand is exists")
	// Member -
	// NfcReader -
	NotFoundTheNFCCard       = NewStatus(LevelERROR, ServiceNFCReader, GRPCNotFound, ActionFind, "讀卡機讀取不到卡片", "nfc reader not find the card")
	NotFoundTheNFCCardReader = NewStatus(LevelFATAL, ServiceNFCReader, GRPCFailedPrecondition, ActionCheck, "找不到讀卡機", "not found the card reader")
	// Order -
	// Storage -
	UploadFileNotFound   = NewStatus(LevelERROR, ServiceStorage, GRPCNotFound, ActionFind, "找不到上傳的檔案", "upload file not found")
	UploadFileFailed     = NewStatus(LevelERROR, ServiceStorage, GRPCAborted, ActionUpload, "上傳檔案失敗", "upload file failed")
	DownloadFileFailed   = NewStatus(LevelERROR, ServiceStorage, GRPCAborted, ActionDownload, "下載檔案失敗", "download file failed")
	UploadFileOutOfRange = NewStatus(LevelERROR, ServiceStorage, GRPCOutOfRange, ActionCheck, "上傳檔案超出範圍", "upload file out of range")
	// Transaction -
	// User -
	UserNotFound      = NewStatus(LevelWARNING, ServiceUser, GRPCNotFound, ActionFind, "找不到使用者", "user not found")
	UserWasDisabled   = NewStatus(LevelWARNING, ServiceUser, GRPCPermissionDenied, ActionCheck, "使用者已被禁止", "user was disabled")
	UserSignInFailed  = NewStatus(LevelWARNING, ServiceUser, GRPCPermissionDenied, ActionSignIn, "使用者登入失敗", "user signin failed")
	UserSignOutFailed = NewStatus(LevelWARNING, ServiceUser, GRPCPermissionDenied, ActionSignOut, "使用者登出失敗", "user signout failed")
	// Website -
	// InitService -
	// SellOrder
	// BuyOrder
	CreateBuyOrderFailed = NewStatus(LevelERROR, ServiceBuyOrder, GRPCAborted, ActionCreate, "建立收購訂單失敗", "create buyOrder failed")
	// SMS
	InvalidMobileNumber       = NewStatus(LevelERROR, ServiceSMS, GRPCUnknown, ActionCallAPI, "無效的手機號碼", "invalid mobile number")
	UndeliverableAfterTimeout = NewStatus(LevelERROR, ServiceSMS, GRPCUnknown, ActionCallAPI, "逾時無法送達", "undeliverable after timeout")
	HasExpired                = NewStatus(LevelERROR, ServiceSMS, GRPCUnknown, ActionCallAPI, "簡訊已過期", "has expired")
	SuspensionOfService       = NewStatus(LevelERROR, ServiceSMS, GRPCUnknown, ActionCallAPI, "系統暫停服務", "suspend of service")
	// IDCard
)
