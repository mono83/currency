package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByNumeric(t *testing.T) {
	c, ok := ByNumeric("963")
	if assert.True(t, ok) {
		assert.Equal(t, "XTS", c.ISO)
	}
	c, ok = ByNumeric("000")
	assert.False(t, ok)
}

func TestByISO(t *testing.T) {
	c, ok := ByISO("XTS")
	if assert.True(t, ok) {
		assert.Equal(t, "963", c.Numeric)
	}
	c, ok = ByISO("000")
	assert.False(t, ok)
}
