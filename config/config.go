package config

import (
	"github.com/tyr-tech-team/hawk/config/source"
	"github.com/tyr-tech-team/hawk/encoder"
)

// Reader -
type Reader interface {
	Read() ([]byte, error)
	ReadWith(value interface{}) error
}

type reader struct {
	s  source.Source
	en encoder.Encoder
}

// NewReader -
func NewReader(s source.Source, en encoder.Encoder) Reader {
	return &reader{
		s:  s,
		en: en,
	}
}

// Read - 讀取配置檔
func (r reader) Read() ([]byte, error) {
	data, err := r.s.Read()
	if err != nil {
		return []byte{}, err
	}
	return data.Data, nil
}

// ReadWithStruct - 讀取至結構
func (r reader) ReadWith(value interface{}) error {
	data, err := r.s.Read()
	if err != nil {
		return err
	}

	return r.en.Decode(data.Data, value)
}
