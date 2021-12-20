package status

import "fmt"

// DescCode -
type DescCode int64

// String -
func (d DescCode) String() string {
	return fmt.Sprintf("%03d", d)
}

// NewDescCode -
func NewDescCode(code int64) DescCode {
	return DescCode(code)
}
