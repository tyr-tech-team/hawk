package status

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_ServiceCode(t *testing.T) {
	a := ServiceAuth

	assert.Equal(t, a.String(), "001")

}
