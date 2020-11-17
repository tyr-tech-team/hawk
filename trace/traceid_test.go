package trace

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewTraceID(t *testing.T) {
	id := NewTraceID()
	assert.NotEmpty(t, id)
	t.Log(id)
}

func Test_SetTraceID(t *testing.T) {
	id := NewTraceID()
	assert.NotEmpty(t, id)
	t.Log(id)
	ctx := context.TODO()
	nctx := SetTraceID(ctx, id)
	tid := GetTraceID(nctx)
	assert.Equal(t, id, tid)
}
