package yaml

import (
	"github.com/tyr-tech-team/hawk/config/encoder"

	"github.com/ghodss/yaml"
)

//
const (
	YAML = "yaml"
)

type yamlEncoder struct{}

func (y yamlEncoder) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (y yamlEncoder) Decode(d []byte, v interface{}) error {
	return yaml.Unmarshal(d, v)
}

func (y yamlEncoder) String() string {
	return YAML
}

// NewEncoder -
func NewEncoder() encoder.Encoder {
	return yamlEncoder{}
}
