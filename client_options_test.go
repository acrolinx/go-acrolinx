package acrolinx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithAPIToken(t *testing.T) {
	client, err := NewClient("signature", "https://example.com", WithAPIToken("sOmEaPiToKeN"))
	assert.NoError(t, err)

	assert.Equal(t, "sOmEaPiToKeN", client.accessToken)
}
