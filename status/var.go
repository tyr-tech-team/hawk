package status

//
var (
	NoError             = NewStatus(LevelNONE, ServiceNONE, GRPCOK, NewDescCode(0), "成功", "success")
	UnKnownError        = NewStatus(LevelERROR, ServiceNONE, GRPCUnknown, NewDescCode(0), "未知的錯誤", "unknown error")
	InternalError       = NewStatus(LevelFATAL, ServiceNONE, GRPCInternal, NewDescCode(0), "內部錯誤")
	InvalidParameter    = NewStatus(LevelWARNING, ServiceNormal, GRPCInvalidArgument, NewDescCode(5), "錯誤的參數", "invalid parameter")
	ConnectTimeOut      = NewStatus(LevelWARNING, ServiceNormal, GRPCDeadlineExceeded, NewDescCode(7), "連線超時", "connect time out")
	NotFound            = NewStatus(LevelWARNING, ServiceNormal, GRPCNotFound, NewDescCode(2), "資料找不到", "not found")
	DataWasExisted      = NewStatus(LevelERROR, ServiceNormal, GRPCAlreadyExists, NewDescCode(5), "資料已存在", "data was existed")
	CreatedFailed       = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, NewDescCode(1), "建立失敗", "created failed")
	UpdatedFailed       = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, NewDescCode(3), "更新失敗", "updated failed")
	DeletedFailed       = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, NewDescCode(4), "刪除失敗", "deleted failed")
	DecodedFailed       = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, NewDescCode(8), "解碼失敗", "decoded failed")
	EncodedFailed       = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, NewDescCode(9), "編碼失敗", "encoded failed")
	ExecutedFailed      = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, NewDescCode(10), "執行失敗", "executed failed")
	OpenedFailed        = NewStatus(LevelERROR, ServiceNormal, GRPCAborted, NewDescCode(13), "開啟失敗", "opened failed")
	DataLoss            = NewStatus(LevelERROR, ServiceNormal, GRPCDataLoss, NewDescCode(5), "資料已遺失", "data loss")
	ParameterOutOfRange = NewStatus(LevelERROR, ServiceNormal, GRPCOutOfRange, NewDescCode(5), "參數超過範圍", "parameter out of range")
	PermissionDenied    = NewStatus(LevelFATAL, ServiceNormal, GRPCPermissionDenied, NewDescCode(10), "權限被拒絕", "permission denied")
	ConnectFailed       = NewStatus(LevelFATAL, ServiceNormal, GRPCUnavailable, NewDescCode(7), "連線失敗", "connect failed")
	TooManayConnect     = NewStatus(LevelERROR, ServiceNormal, GRPCResourceExhausted, NewDescCode(7), "太多連線", "too manay connect")
	HealthCheckFailed   = NewStatus(LevelFATAL, ServiceNormal, GRPCAborted, NewDescCode(18), "健康檢查失敗", "health check failed")
	RemoteHostNotFound  = NewStatus(LevelERROR, ServiceNormal, GRPCNotFound, NewDescCode(7), "找不到遠端呼叫地址", "remote host not found")

	// Auth -
	TokenGenerationFailed = NewStatus(LevelERROR, ServiceAuth, GRPCAborted, NewDescCode(1), "權杖產生失敗", "token generation failed")
	TokenWasExpired       = NewStatus(LevelFATAL, ServiceAuth, GRPCDeadlineExceeded, NewDescCode(5), "權杖已失效", "token was expired")
	TokenRevokedFailed    = NewStatus(LevelERROR, ServiceAuth, GRPCAborted, NewDescCode(4), "取消權杖失敗", "revoke token failed")
	InvalidToken          = NewStatus(LevelFATAL, ServiceAuth, GRPCUnauthenticated, NewDescCode(5), "錯誤的權杖", "invalid token")
	// Card
	CardInvalidParemeter = NewStatus(LevelWARNING, ServiceCard, GRPCInvalidArgument, NewDescCode(5), "卡片參數錯誤", "invalid card parameter")
	CardWasNotEmpty      = NewStatus(LevelERROR, ServiceCard, GRPCAlreadyExists, NewDescCode(5), "卡片不是空的", "card was not empty")
	CardNotFound         = NewStatus(LevelWARNING, ServiceCard, GRPCNotFound, NewDescCode(5), "找不到卡片紀錄", "card not found")
	CardCreatedFailed    = NewStatus(LevelERROR, ServiceCard, GRPCAborted, NewDescCode(1), "建立卡片失敗", "create card failed")
	CardUpdatedFailed    = NewStatus(LevelERROR, ServiceCard, GRPCAborted, NewDescCode(3), "更新卡片失敗", "update card failed")
	CardRevokeFailed     = NewStatus(LevelERROR, ServiceCard, GRPCAborted, NewDescCode(4), "卡片註銷失敗", "Revoke card failed")
	CardCheckFailed      = NewStatus(LevelWARNING, ServiceCard, GRPCFailedPrecondition, NewDescCode(5), "卡片檢查失敗", "check card failed")
	// EventLog
	// Item -
	ItemParameterInvalid = NewStatus(LevelWARNING, ServiceItem, GRPCInvalidArgument, NewDescCode(5), "商品參數錯誤", "item parmeter invalid")
	ItemNotFound         = NewStatus(LevelWARNING, ServiceItem, GRPCNotFound, NewDescCode(2), "找不到商品", "item not found")
	ItemCreatedFailed    = NewStatus(LevelERROR, ServiceItem, GRPCAborted, NewDescCode(1), "商品入庫失敗", "create item failed")
	ItemUpdatedFailed    = NewStatus(LevelERROR, ServiceItem, GRPCAborted, NewDescCode(3), "商品更新資訊失敗", "update item failed")
	ItemDeletedFailed    = NewStatus(LevelERROR, ServiceItem, GRPCAborted, NewDescCode(4), "商品刪除失敗", "delete item failed")
	// Brand -
	BrandNotFound      = NewStatus(LevelWARNING, ServiceBrand, GRPCNotFound, NewDescCode(2), "找不到品牌", "brand not found")
	BrandWasExists     = NewStatus(LevelERROR, ServiceBrand, GRPCAlreadyExists, NewDescCode(5), "品牌已存在", "brand is exists")
	BrandCreatedFailed = NewStatus(LevelERROR, ServiceBrand, GRPCAborted, NewDescCode(1), "品牌建立失敗", "create brand failed")
	BrandDeletedFailed = NewStatus(LevelERROR, ServiceBrand, GRPCAborted, NewDescCode(4), "刪除品牌失敗", "delete brand failed")
	// Member -
	MemberInvalidParameter = NewStatus(LevelWARNING, ServiceMember, GRPCInvalidArgument, NewDescCode(5), "錯誤的會員參數", "invalid member parameter ")
	MemberNotFound         = NewStatus(LevelWARNING, ServiceMember, GRPCNotFound, NewDescCode(2), "找不到會員", "member not found")
	MemberCreatedFailed    = NewStatus(LevelERROR, ServiceMember, GRPCAborted, NewDescCode(1), "建立會員失敗", "create member failed")
	MemberUpdatedFailed    = NewStatus(LevelERROR, ServiceMember, GRPCAborted, NewDescCode(3), "更新會員資訊失敗", "update member failed")
	MemberDeletedFailed    = NewStatus(LevelERROR, ServiceMember, GRPCAborted, NewDescCode(4), "刪除會員失敗", "delete member failed")
	MemberPhoneUsed        = NewStatus(LevelWARNING, ServiceMember, GRPCAborted, NewDescCode(1), "手機號碼已被使用", "phone has been used")
	MemberSiginFailed      = NewStatus(LevelWARNING, ServiceMember, GRPCPermissionDenied, NewDescCode(1), "會員登入失敗", "member singIn failed")
	// NfcReader -
	NotFoundTheNFCCard       = NewStatus(LevelERROR, ServiceNFCReader, GRPCNotFound, NewDescCode(2), "讀卡機讀取不到卡片", "nfc reader not find the card")
	NotFoundTheNFCCardReader = NewStatus(LevelFATAL, ServiceNFCReader, GRPCFailedPrecondition, NewDescCode(5), "找不到讀卡機", "not found the card reader")
	// Order -
	// Storage -
	UploadFileNotFound   = NewStatus(LevelERROR, ServiceStorage, GRPCNotFound, NewDescCode(2), "找不到上傳的檔案", "upload file not found")
	UploadFileFailed     = NewStatus(LevelERROR, ServiceStorage, GRPCAborted, NewDescCode(14), "上傳檔案失敗", "upload file failed")
	DownloadFileFailed   = NewStatus(LevelERROR, ServiceStorage, GRPCAborted, NewDescCode(15), "下載檔案失敗", "download file failed")
	UploadFileOutOfRange = NewStatus(LevelERROR, ServiceStorage, GRPCOutOfRange, NewDescCode(5), "上傳檔案超出範圍", "upload file out of range")
	// Transaction -
	PaymentMethodCreatedCheckFailed = NewStatus(LevelERROR, ServiceTransaction, GRPCInvalidArgument, NewDescCode(1), "建立付款方式檢查失敗", "create payment method check failed")
	PaymentMethodCreatedFailed      = NewStatus(LevelERROR, ServiceTransaction, GRPCAborted, NewDescCode(2), "付款方式建立失敗", "create payment method failed")
	PaymentMethodNotFound           = NewStatus(LevelWARNING, ServiceTransaction, GRPCNotFound, NewDescCode(1), "找不到此付款方式", "can't found this payment method")
	PaymentMethodWasExists          = NewStatus(LevelWARNING, ServiceTransaction, GRPCAlreadyExists, NewDescCode(1), "付款方式已存在", "payment method was exists")
	PaymentSlipUnMatchRules         = NewStatus(LevelWARNING, ServiceTransaction, GRPCInvalidArgument, NewDescCode(6), "此付款單不符合付款方式規則", "payment slip unmatch payment method")
	PaymentSlipCheckFailed          = NewStatus(LevelWARNING, ServiceTransaction, GRPCInvalidArgument, NewDescCode(2), "付款單參數錯誤", "invalid payment slip parameter")
	PaymentSlipNotFound             = NewStatus(LevelWARNING, ServiceTransaction, GRPCNotFound, NewDescCode(2), "找不到此交易單", "payment slip not found")
	PaymentSlipCreatedFailed        = NewStatus(LevelERROR, ServiceTransaction, GRPCAborted, NewDescCode(1), "交易建立失敗", "create payment slip failed")
	PaymentSlipUpdatedFailed        = NewStatus(LevelERROR, ServiceTransaction, GRPCAborted, NewDescCode(3), "交易更新失敗", "update payment slip failed")
	PaymentSlipDeleteFailed         = NewStatus(LevelERROR, ServiceTransaction, GRPCAborted, NewDescCode(4), "交易刪除失敗", "delete payment slip failed")
	PaymentSlipCancelFailed         = NewStatus(LevelERROR, ServiceTransaction, GRPCAborted, NewDescCode(15), "付款單取消失敗", "payment slip cancel failed")
	PaymentSlipPayFailed            = NewStatus(LevelERROR, ServiceTransaction, GRPCAborted, NewDescCode(16), "付款單支付失敗", "payment slip pay failed")
	PaymentSlipRefunedFailed        = NewStatus(LevelERROR, ServiceTransaction, GRPCAborted, NewDescCode(17), "付款單退款失敗", "payment slip refund failed")
	// User -
	UserNotFound           = NewStatus(LevelWARNING, ServiceUser, GRPCNotFound, NewDescCode(2), "找不到使用者", "user not found")
	UserWasDisabled        = NewStatus(LevelWARNING, ServiceUser, GRPCPermissionDenied, NewDescCode(5), "使用者已被禁止", "user was disabled")
	UserSignInFailed       = NewStatus(LevelWARNING, ServiceUser, GRPCPermissionDenied, NewDescCode(11), "使用者登入失敗", "user signin failed")
	UserSignOutFailed      = NewStatus(LevelWARNING, ServiceUser, GRPCPermissionDenied, NewDescCode(12), "使用者登出失敗", "user signout failed")
	UserDuplicateParameter = NewStatus(LevelWARNING, ServiceUser, GRPCAlreadyExists, NewDescCode(5), "使用者參數已重複", "user have duplicate parameter")
	// Website -
	// InitService -
	// SellOrder
	SellOrderCreatedFailed = NewStatus(LevelERROR, ServiceSellOrder, GRPCAborted, NewDescCode(1), "銷售訂單建立失敗", "create sellorder failed")
	SellOrderCheckedFailed = NewStatus(LevelERROR, ServiceSellOrder, GRPCFailedPrecondition, NewDescCode(5), "銷售訂單檢查錯誤", "check sellorder failed")

	// BuyOrder
	CreatedBuyOrderFailed = NewStatus(LevelERROR, ServiceBuyOrder, GRPCAborted, NewDescCode(1), "建立收購訂單失敗", "create buyOrder failed")
	// SMS
	SmsFailedToSend       = NewStatus(LevelERROR, ServiceSMS, GRPCUnknown, NewDescCode(6), "簡訊傳送失敗", "sms failed to send")
	SmsVerificationFailed = NewStatus(LevelERROR, ServiceSMS, GRPCInvalidArgument, NewDescCode(5), "驗證碼錯誤", "wrong sms code")
	// Category
	CategorySortCheckFailed = NewStatus(LevelWARNING, ServiceCategory, GRPCInvalidArgument, NewDescCode(5), "分類更新排序檢查失敗", "category sort check failed")
	CategoryWasUsed         = NewStatus(LevelWARNING, ServiceCategory, GRPCFailedPrecondition, NewDescCode(20), "分類已被使用", "category was used ")
	CategoryHasQuickFilter  = NewStatus(LevelWARNING, ServiceCategory, GRPCFailedPrecondition, NewDescCode(24), "分類下已有快篩", "Category has quickfilter")
	CategoryDeletedFailed   = NewStatus(LevelERROR, ServiceCategory, GRPCAborted, NewDescCode(4), "分類刪除失敗", "Category delete failed")
	CategoryUpdatedFailed   = NewStatus(LevelERROR, ServiceCategory, GRPCAborted, NewDescCode(3), "分類更新失敗", "Category update failed")
	CategoryCreatedFailed   = NewStatus(LevelERROR, ServiceCategory, GRPCAborted, NewDescCode(1), "分類建立失敗", "Category create failed")
	CategoryNotFound        = NewStatus(LevelWARNING, ServiceCategory, GRPCNotFound, NewDescCode(2), "找不到此分類", "Category not found")
	CategoryWasExisted      = NewStatus(LevelERROR, ServiceCategory, GRPCAlreadyExists, NewDescCode(24), "分類已存在", "Category was existed")

	// QuickFilterBeUsed -
	QuickFilterBeUsed            = NewStatus(LevelWARNING, ServiceQuickFilter, GRPCAlreadyExists, NewDescCode(20), "快篩商品已被使用", "quickfilter be used")
	QuickFilterDeletedFailed     = NewStatus(LevelERROR, ServiceQuickFilter, GRPCAborted, NewDescCode(4), "快篩刪除失敗", "quickfilter delete filed")
	QuickFilterUpdatedFailed     = NewStatus(LevelERROR, ServiceQuickFilter, GRPCAborted, NewDescCode(3), "快篩更新失敗", "quickfilter delete filed")
	QuickFilterCreatedFailed     = NewStatus(LevelERROR, ServiceQuickFilter, GRPCAborted, NewDescCode(1), "快篩建立失敗", "quickfilter created filed")
	QuickFilterNotFound          = NewStatus(LevelWARNING, ServiceQuickFilter, GRPCNotFound, NewDescCode(1), "找不到此快篩", "quickfilter not found")
	QuickFilterWasExisted        = NewStatus(LevelERROR, ServiceQuickFilter, GRPCAlreadyExists, NewDescCode(24), "快篩已存在", "quickfilter was existed")
	QuickFilterCreatedOutOfRange = NewStatus(LevelWARNING, ServiceQuickFilter, GRPCOutOfRange, NewDescCode(1), "建立快篩數量超出範圍", "created quickfilter out of range")
	QuickFilterInvalidParameter  = NewStatus(LevelWARNING, ServiceQuickFilter, GRPCInvalidArgument, NewDescCode(5), "快篩參數錯誤", "quickfilter have invalid parameter")

	// Receipt
	ReceiptIssuedFailed    = NewStatus(LevelERROR, ServiceReceipt, GRPCAborted, NewDescCode(1), "核發發票失敗", "issue receipt failed")
	ReceiptInvalidedFailed = NewStatus(LevelERROR, ServiceReceipt, GRPCAborted, NewDescCode(2), "作廢發票失敗", "invalid receipt failed")
	ReceiptNotFound        = NewStatus(LevelWARNING, ServiceReceipt, GRPCNotFound, NewDescCode(0), "找不到此付款單發票", "this receipt not found")
	ReceiptInvalidBarcode  = NewStatus(LevelWARNING, ServiceReceipt, GRPCInvalidArgument, NewDescCode(1), "無效的手機條碼", "invalid barcode")
	ReceiptInvalidLovecode = NewStatus(LevelWARNING, ServiceReceipt, GRPCInvalidArgument, NewDescCode(2), "無效的捐獻碼", "invalid lovecode")
)
