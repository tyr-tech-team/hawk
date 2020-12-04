package config

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tyr-tech-team/hawk/config/source"
	"github.com/tyr-tech-team/hawk/pkg/consul"
)

func Test_ConsulConfig(t *testing.T) {

	cli := consul.NewClient(context.TODO(), consul.Config{
		Address: "localhost:8500",
	})

	s := source.NewConsul(cli, "auth")

	r := NewReader(s, YAML)
	d, err := r.Read()
	assert.NoError(t, err)
	assert.NotEmpty(t, d)
	t.Log(string(d))

	x := &struct {
		Redis Redis `yaml:"redis"`
	}{}
	err = r.ReadWith(x)
	assert.NoError(t, err)
	assert.NotEmpty(t, x)

	t.Log(x)
}
