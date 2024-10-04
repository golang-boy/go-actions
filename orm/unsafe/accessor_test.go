package unsafe

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_unsafeAccessor_GetField(t *testing.T) {

	type User struct {
		Name string
		Age  int
	}

	a := NewUnsafeAccessor(&User{Name: "Tom", Age: 18})

	val, err := a.GetField("Age")
	require.NoError(t, err)
	assert.Equal(t, 18, val)

	err = a.SetField("Age", 19)
	assert.NoError(t, err)
}
