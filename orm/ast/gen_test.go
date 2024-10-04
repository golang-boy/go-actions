package ast

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGen(t *testing.T) {
	buffer := &bytes.Buffer{}

	err := Gen(buffer, "testdata/user.go")
	require.NoError(t, err)
	assert.Equal(t, `
package testdata	
	`, buffer.String())
}
