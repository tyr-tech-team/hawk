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
)
