// Package status provides status ﳑ
package status

import "fmt"

// ActionCode -
type ActionCode int64

func (a ActionCode) String() string {
	return fmt.Sprintf("%03d", a)
}

const (
	// ActionNone 0
	ActionNone ActionCode = iota
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
	//ActionDecode 解析
	ActionDecode
	//ActionEncode 加密
	ActionEncode
	//ActionExecute 執行
	ActionExecute
	//ActionSignIn 登入
	ActionSignIn
	// ActionSignOut 登出
	ActionSignOut
	// ActionOpen 開啟
	ActionOpen
	// ActionUpload 上傳
	ActionUpload
	// ActionDownload 載
	ActionDownload
	// ActionPay 付
	ActionPay
	//ActionRefund 退款
	ActionRefund
	//ActionHealthCheck 健康檢查
	ActionHealthCheck
	//ActionCancel 取消
	ActionCancel
	//ActionUse 使用
	ActionUse
	//ActionEnable 啟用
	ActionEnable
	//ActionDisable 禁用
	ActionDisable
	//ActionReturn 返回
	ActionReturn
	//ActionExist 已存在
	ActionExist
)
