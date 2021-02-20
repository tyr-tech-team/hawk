package status

import "fmt"

type DescCode int64

func (d DescCode) String() string {
	return fmt.Sprintf("%03d", d)
}

func NewDescCode(code int64) DescCode {
	return DescCode(code)
}
