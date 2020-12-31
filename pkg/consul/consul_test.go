package consul

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tyr-tech-team/hawk/config"
)

var testCli Client

func Test_NewClient(t *testing.T) {
	testCli = NewClient(context.TODO(), DefaultConsulConfig())
	_, err := testCli.Client().Status().Leader()
	assert.NoError(t, err)
}

func Test_Register(t *testing.T) {
	testCli.SetRegisterConfig(config.ServiceRegister{
		Name:    "test",
		Tags:    []string{"dev-test"},
		Port:    10800,
		Address: "localhost",
	})

	err := testCli.Register()
	assert.NoError(t, err)
}

func Test_Deregister(t *testing.T) {
	time.Sleep(30 * time.Second)
	assert.NoError(t, testCli.Deregister())
}
