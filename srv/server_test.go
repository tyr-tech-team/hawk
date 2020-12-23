package srv

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tyr-tech-team/hawk/pkg/consul"
)

func Test_New(t *testing.T) {
	cli := consul.NewClient(context.TODO(), consul.DefaultConsulConfig())

	s := New(
		SetName("test-srv"),
		SetHost(""),
		SetPort("0"),
		SetGRPC(),
		SetEnableTraefik(),
		SetRegister(cli),
	)

	assert.NoError(t, s.Register())
	time.Sleep(1 * time.Second)
	assert.NoError(t, s.Deregister())
	s.Close()
}
