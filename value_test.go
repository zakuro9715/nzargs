package nzargs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlag(t *testing.T) {
	v := NewFlag("name", "a", "b")
	assert.Equal(t, v.Type(), TypeFlag)
	assert.Equal(t, v.Text(), "--name=a,b")
	assert.Equal(t, v.Flag(), v)
	assert.Nil(t, v.Arg())
	assert.Equal(t, v.Values, []string{"a", "b"})
}

func TestArg(t *testing.T) {
	v := NewArg("value")
	assert.Equal(t, v.Type(), TypeArg)
	assert.Equal(t, v.Text(), "value")
	assert.Nil(t, v.Flag())
	assert.Equal(t, v.Arg(), v)
}
