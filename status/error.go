package status

import "google.golang.org/grpc/codes"

// Errors -
type Errors interface {
	// Error 所使用
	Error() string
	// 預設字串顯示
	String() string
	// 取得錯誤代碼
	GetCode() int32
	// 取得 Grpc 錯誤代碼
	GetGrpcCode() codes.Code
	// 取得錯誤訊息
	GetMsg() string
	// 取得錯誤
	Err() error
	// 判斷是否相同
	Equal(err error) bool
}
