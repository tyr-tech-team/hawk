package status

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ServiceCode(t *testing.T) {
	a := ServiceAuth
	assert.Equal(t, a.String(), "002")
}

func Test_NewStatus(t *testing.T) {
	x := NewStatus(LevelERROR, ServiceNONE, GRPCAlreadyExists, ActionCreate, "123888")

	s := ConvertStatus(x.Err())
	fmt.Println(s)

}

func Test_WithDetail(t *testing.T) {
	x := NewStatus(LevelWARNING, ServiceNONE, GRPCAlreadyExists, ActionCreate, "123888")

	s := ConvertStatus(x)
	assert.NotEqual(t, s, UnKnownError)
	fmt.Println(s)

	x = x.WithDetail([]string{"1", "2", "3"}...)

	fmt.Println(x.Err())
	x = x.WithDetail([]string{"4", "5"}...)
	fmt.Println(x.Err())
}

func Test_WithServiceCode(t *testing.T) {
	x := NewStatus(LevelWARNING, ServiceNONE, GRPCAlreadyExists, ActionCreate, "123888")
	fmt.Println(x)
	s := ConvertStatus(x)
	assert.NotEqual(t, s, UnKnownError)
	fmt.Println(s)

	s = s.SetServiceCode(ServiceBrand)
	fmt.Println(s)
	s = s.WithDetail([]string{"4", "5"}...)
	fmt.Println("s:", s)

	fmt.Println("x:", x)
}

func Test_Equal(t *testing.T) {
	x1 := UserNotFound
	x2 := UserNotFound
	assert.Equal(t, Equal(x1, x2), true)

	x1 = x1.WithDetail("123")
	assert.Equal(t, Equal(x1, x2), true)

	t.Log("x1", x1)
	t.Log("x2", x2)

	x2 = x2.SetServiceCode(ServiceItem)
	assert.Equal(t, Equal(x1, x2), false)
	t.Log("x1", x1)
	t.Log("x2", x2)

}
