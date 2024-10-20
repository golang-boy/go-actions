package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	testCases := []struct {
		name string
		req  *Request
	}{
		{
			name: "Test Request",
			req:  &Request{Data: []byte("test data")},
		},
	}

	for _, tc := range testCases {
		// Encode
		encoded := EncodeReq(tc.req)
		// Decode
		req := DecodeReq(encoded)

		assert.Equal(t, tc.req, req)
	}

}
