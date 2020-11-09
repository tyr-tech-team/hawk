package config

import (
	"fmt"
	j "hawk/config/encoder/json"
	y "hawk/config/encoder/yaml"
	"hawk/config/source"
	"hawk/config/source/consul"
	"hawk/config/source/file"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

type C struct {
	Hosts string `json:"hosts" yaml:"hosts"`
}

// Config -
type CC struct {
	Hosts Hosts `json:"hosts"`
}

type Hosts struct {
	Database string `json:"database"`
	Test     string `json:"test"`
}

type Service struct {
	Website string `yaml:"website"`
}

type Conf struct {
	Channel map[string]string `json:"channel" yaml:"channel"`
	Service Services          `json:"service" yaml:"service"`
}

// Service -
type Services struct {
	Card   string `json:"card" yaml:"card"`
	Member string `json:"member" yaml:"member"`
}

// TestReadJson -
func TestReadJson(t *testing.T) {
	enc := j.NewEncoder()

	opts := file.WithPath("../config.json")

	fileSource := file.NewSource(opts, source.WithEncoder(enc))
	cc := NewConfig(fileSource)

	changeSet, err := cc.Read()

	assert.NoError(t, err)

	conf := C{}
	enc.Decode(changeSet.Data, &conf)

	assert.NoError(t, err)

	fmt.Println(conf.Hosts)
}

// TestReadEtcd -

// TestReadConsul
func TestReadConsul(t *testing.T) {
	consulSource := consul.NewSource(
		consul.SetKey("gateway"),
		consul.SetAddrs("127.0.0.1:8500"),
		consul.SetConfigType("yaml"),
	)
	c := NewConfig(consulSource)

	chSet, err := c.Read()

	if err != nil {
		fmt.Println(err.Error())
	}

	enct := y.NewEncoder()

	conft := bson.M{}

	enct.Decode(chSet.Data, &conft)

	spew.Dump(conft)

}
