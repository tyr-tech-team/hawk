package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PublicIP(t *testing.T) {
	ip := PublishIP()
	assert.NotEmpty(t, ip)
	t.Log(ip)

}

func Test_LocalIP(t *testing.T) {
	ip := LocalIP()
	assert.NotEmpty(t, ip)
	t.Log(ip)

}

func Test_HostIP(t *testing.T) {
	ip := HostIP()
	assert.NotEmpty(t, ip)
	t.Log(ip)
}

func Test_GetFreePort(t *testing.T) {
	p, err := GetFreePort()
	assert.NoError(t, err)
	t.Log(p)
}
