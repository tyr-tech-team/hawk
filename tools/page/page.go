// Package page provides page ﳑ
package page

// DefaultSize - 預設每一頁初始顯示筆數
var DefaultSize int64

// ConvertMongoSizeAndLimit -
func ConvertMongoSizeAndLimit(page, size int64) (skip int64, limit int64) {
	d := &Data{
		Page: page,
		Size: size,
	}

	d.ToSkipLimit()

	return d.Skip, d.Limit
}

// Data -
type Data struct {
	Page  int64 `json:"page" yaml:"page"`
	Size  int64 `json:"size" yaml:"size"`
	Skip  int64 `json:"skip" yaml:"skip"`
	Limit int64 `json:"limit" yaml:"limit"`
}

// ToSkipLimit  - 轉換至 SKIP LIMIT
func (d *Data) ToSkipLimit() {
	if d.Page < 1 {
		d.Page = 1
	}

	if d.Size == 0 {
		d.Size = DefaultSize
	}

	d.Skip = d.Size * (d.Page - 1)
	d.Limit = d.Size
}
