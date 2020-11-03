package config

import (
	"fmt"
	"hawk/config/encoder/json"
	"hawk/config/source"
	"hawk/config/source/etcd"
	"hawk/config/source/file"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
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

// TestReadJson -
func TestReadJson(t *testing.T) {
	enc := json.NewEncoder()

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
func TestReadEtcd(t *testing.T) {
	etcdSource := etcd.NewSource(
		etcd.WithAddress("127.0.0.1:2379"),
		etcd.WithPrefix("user"),
		etcd.StripPrefix(true),
	)

	config := NewConfig(etcdSource)
	changeSet, err := config.Read()
	if err != nil {
		panic(err)
	}

	spew.Dump(changeSet)

	enc := json.NewEncoder()

	conf := CC{}
	enc.Decode(changeSet.Data, &conf)

	spew.Dump(conf.Hosts)
}
