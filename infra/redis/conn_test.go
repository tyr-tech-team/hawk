package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tyr-tech-team/hawk/config"
)

func Test_NewDial(t *testing.T) {
	c := config.Redis{
		Host:     "localhost:6379",
		Database: 0,
	}

	cli, err := NewDial(c)
	assert.NoError(t, err)
	assert.NotEmpty(t, cli)

}
