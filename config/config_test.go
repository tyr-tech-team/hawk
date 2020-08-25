package config

import (
	"fmt"
	"hawk/config/encoder/json"
	"hawk/config/encoder/yaml"
	"hawk/config/source"
	"hawk/config/source/file"
	"testing"

	"github.com/stretchr/testify/assert"
)

type C struct {
	Hosts string `json:"hosts" yaml:"hosts"`
}

// TestReadFromYaml -
func TestReadYaml(t *testing.T) {
	enc := yaml.NewEncoder()

	opts := file.WithPath("../config.yaml")

	fileSource := file.NewSource(opts, source.WithEncoder(enc))

	changeSet, err := fileSource.Read()

	assert.NoError(t, err)

	conf := C{}
	enc.Decode(changeSet.Data, &conf)

	assert.NoError(t, err)

	fmt.Println(conf.Hosts)
}

// TestReadJson -
func TestReadJson(t *testing.T) {
	enc := json.NewEncoder()

	opts := file.WithPath("../config.json")

	fileSource := file.NewSource(opts, source.WithEncoder(enc))

	changeSet, err := fileSource.Read()

	assert.NoError(t, err)

	conf := C{}
	enc.Decode(changeSet.Data, &conf)

	assert.NoError(t, err)

	fmt.Println(conf.Hosts)
}
