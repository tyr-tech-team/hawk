package status

import "fmt"

// ActionCode -
type ActionCode int64

func (a ActionCode) String() string {
	return fmt.Sprintf("%03d", a)
}

//
const (
	//
	ActionNono ActionCode = iota
	// ActionCreate - 建立動作
	ActionCreate
	// ActionFind - 查詢動作
	ActionFind
	// ActionUpdate - 更新動作
	ActionUpdate
	// ActionDelete - 刪除動作
	ActionDelete
	// ActionCheck -  檢查動作
	ActionCheck
	// ActionCallAPI - 呼叫API
	ActionCallAPI
	// ActionConnect - 建立鏈結
	ActionConnect
	// 解析
	ActionDecode
	// 加密
	ActionEncode
	// 執行
	ActionExecute
	// 登入
	ActionSignIn
	// 登出
	ActionSignOut
	// 開啟
	ActionOpen
	// 上傳
	ActionUpload
	// 下載
	ActionDownload
	// 支付
	ActionPay
	// 退款
	ActionRefund
)
