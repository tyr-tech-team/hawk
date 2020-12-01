package config

import (
	"github.com/tyr-tech-team/hawk/config/source"
	"github.com/tyr-tech-team/hawk/encoder"
)

// Reader -
type Reader interface {
	Read() ([]byte, error)
	ReadWithStruct(value interface{}) error
}

type reader struct {
	s  source.Source
	en encoder.Encoder
}

// NewReader -
func NewReader(s source.Source) Reader {
	return &reader{
		s: s,
	}
}

// Read - 讀取配置檔
func (r reader) Read() ([]byte, error) {

	return nil, nil
}

// ReadWithStruct - 讀取至結構
func (r reader) ReadWithStruct(value interface{}) error {
	return nil
}
