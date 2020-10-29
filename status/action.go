package status

import "fmt"

// ActionCode -
type ActionCode int64

func (a ActionCode) String() string {
	return fmt.Sprintf("%03d", a)
}

//
const (
	ActionNono ActionCode = iota
	ActionCreate
	ActionFind
	ActionUpdate
	ActionDelete
	ActionCheck
	ActionCallAPI
)
