package status

//
var (
	NoError             = NewStatus(LevelNONE, ServiceNONE, GRPCOK, ActionNono, "成功", "success")
	UnKnownError        = NewStatus(LevelERROR, ServiceNONE, GRPCUnknown, ActionNono, "未知的錯誤", "unknown error")
	InvalidParameter    = NewStatus(LevelWARNING, ServiceNormal, GRPCInvalidArgument, ActionCheck, "錯誤的參數", "invalid parameter")
	ConnectTimeOut      = NewStatus(LevelWARNING, ServiceNormal, GRPCDeadlineExceeded, ActionConnect, "連線超時", "connect time out")
	NotFound            = NewStatus(LevelWARNING, ServiceNormal, GRPCNotFound, ActionFind, "資料找不到", "not found")
	DataWasExisted      = NewStatus(LevelERROR, ServiceNormal, GRPCAlreadyExists, ActionCheck, "資料已存在", "data was existed")
	CreatedFailed       = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionCreate, "建立失敗", "created failed")
	UpdatedFailed       = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionUpdate, "更新失敗", "updated failed")
	DeletedFailed       = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionDelete, "刪除失敗", "deleted failed")
	DecodedFailed       = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionDecode, "解碼失敗", "decoded failed")
	EncodedFailed       = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionEncode, "編碼失敗", "encoded failed")
	ExecutedFailed      = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionExecute, "執行失敗", "executed failed")
	OpenedFailed        = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, ActionOpen, "開啟失敗", "opened failed")
	DataLoss            = NewStatus(LevelERROR, ServiceNormal, GRPCDataLoss, ActionCheck, "資料已遺失", "data loss")
	ParameterOutOfRange = NewStatus(LevelERROR, ServiceNormal, GRPCOutOfRange, ActionCheck, "參數超過範圍", "parameter out of range")
	PermissionDenied    = NewStatus(LevelFATAL, ServiceNormal, GRPCPermissionDenied, ActionExecute, "權限被拒絕", "permission denied")
	ConnectFailed       = NewStatus(LevelFATAL, ServiceNormal, GRPCUnavailable, ActionConnect, "連線失敗", "connect failed")
	TooManayConnect     = NewStatus(LevelERROR, ServiceNormal, GRPCResourceExhausted, ActionConnect, "太多連線", "too manay connect")
	HealthCheckFailed   = NewStatus(LevelFATAL, ServiceNormal, GRPCAborted, ActionHealthCheck, "健康檢查失敗", "health check failed")
	RemoteHostNotFound  = NewStatus(LevelERROR, ServiceNormal, GRPCNotFound, ActionConnect, "找不到遠端呼叫地址", "remote host not found")

	// Auth -
	TokenGenerationFailed = NewStatus(LevelERROR, ServiceAuth, GRPCAborted, ActionCreate, "權杖產生失敗", "token generation failed")
	TokenWasExpired       = NewStatus(LevelFATAL, ServiceAuth, GRPCDeadlineExceeded, ActionCheck, "權杖已失效", "token was expired")
	TokenRevokedFailed    = NewStatus(LevelERROR, ServiceAuth, GRPCAborted, ActionDelete, "取消權杖失敗", "revoke token failed")
	InvalidToken          = NewStatus(LevelFATAL, ServiceAuth, GRPCUnauthenticated, ActionCheck, "錯誤的權杖", "invalid token")
	// Card
	CardWasNotEmpty = NewStatus(LevelERROR, ServiceCard, GRPCAlreadyExists, ActionCheck, "卡片不是空的", "card was not empty")
	// EventLog
	// Item -
	ItemParameterInvalid = NewStatus(LevelWARNING, ServiceItem, GRPCInvalidArgument, ActionCheck, "商品參數錯誤", "item parmeter invalid")
	ItemNotFound         = NewStatus(LevelWARNING, ServiceItem, GRPCNotFound, ActionFind, "找不到商品", "item not found")
	ItemCreatedFailed    = NewStatus(LevelERROR, ServiceItem, GRPCAborted, ActionCreate, "商品入庫失敗", "create item failed")
	ItemUpdatedFailed    = NewStatus(LevelERROR, ServiceItem, GRPCAborted, ActionUpdate, "商品更新資訊失敗", "update item failed")
	ItemDeletedFailed    = NewStatus(LevelERROR, ServiceItem, GRPCAborted, ActionDelete, "商品刪除失敗", "delete item failed")
	// Brand -
	BrandNotFound      = NewStatus(LevelWARNING, ServiceBrand, GRPCNotFound, ActionFind, "找不到品牌", "brand not found")
	BrandWasExists     = NewStatus(LevelERROR, ServiceBrand, GRPCAlreadyExists, ActionCheck, "品牌已存在", "brand is exists")
	BrandCreatedFailed = NewStatus(LevelERROR, ServiceBrand, GRPCAborted, ActionCreate, "品牌建立失敗", "create brand failed")
	BrandDeletedFailed = NewStatus(LevelERROR, ServiceBrand, GRPCAborted, ActionDelete, "刪除品牌失敗", "delete brand failed")
	// Member -
	MemberInvalidParameter = NewStatus(LevelWARNING, ServiceMember, GRPCInvalidArgument, ActionCheck, "錯誤的會員參數", "invalid member parameter ")
	MemberNotFound         = NewStatus(LevelWARNING, ServiceMember, GRPCNotFound, ActionFind, "找不到會員", "member not found")
	MemberCreateFailed     = NewStatus(LevelERROR, ServiceMember, GRPCAborted, ActionCreate, "建立會員失敗", "create member failed")
	MemberUpdatedFailed    = NewStatus(LevelERROR, ServiceMember, GRPCAborted, ActionUpdate, "更新會員資訊失敗", "update member failed")
	MemberDeletedFailed    = NewStatus(LevelERROR, ServiceMember, GRPCAborted, ActionDelete, "刪除會員失敗", "delete member failed")
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
	UserNotFound           = NewStatus(LevelWARNING, ServiceUser, GRPCNotFound, ActionFind, "找不到使用者", "user not found")
	UserWasDisabled        = NewStatus(LevelWARNING, ServiceUser, GRPCPermissionDenied, ActionCheck, "使用者已被禁止", "user was disabled")
	UserSignInFailed       = NewStatus(LevelWARNING, ServiceUser, GRPCPermissionDenied, ActionSignIn, "使用者登入失敗", "user signin failed")
	UserSignOutFailed      = NewStatus(LevelWARNING, ServiceUser, GRPCPermissionDenied, ActionSignOut, "使用者登出失敗", "user signout failed")
	UserDuplicateParameter = NewStatus(LevelWARNING, ServiceUser, GRPCAlreadyExists, ActionCheck, "使用者參數已重複", "user have duplicate parameter")
	// Website -
	// InitService -
	// SellOrder
	// BuyOrder
	CreatedBuyOrderFailed = NewStatus(LevelERROR, ServiceBuyOrder, GRPCAborted, ActionCreate, "建立收購訂單失敗", "create buyOrder failed")
	// SMS
	SmsFailedToSend = NewStatus(LevelERROR, ServiceSMS, GRPCUnknown, ActionCallAPI, "簡訊傳送失敗", "sms failed to send")
	// IDCard
)
