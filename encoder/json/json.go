package json

import (
	"encoding/json"

	"github.com/tyr-tech-team/hawk/encoder"
)

//
const (
	JSON = "json"
)

type jsonEncoder struct{}

func (j jsonEncoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j jsonEncoder) Decode(d []byte, v interface{}) error {
	return json.Unmarshal(d, v)
}

func (j jsonEncoder) String() string {
	return JSON
}

// NewEncoder -
func NewEncoder() encoder.Encoder {
	return jsonEncoder{}
}
