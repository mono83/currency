package amount

import (
	"testing"

	"github.com/mono83/currency"
	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	a := Zero(currency.Testing().Short())

	assert.Equal(t, 0., a.Amount())
	assert.Equal(t, currency.Testing().Short(), a.Currency())
}

func TestAutoZero(t *testing.T) {
	z, ok := New(currency.Testing().Short(), 0.).(zero)
	if assert.True(t, ok) {
		assert.Equal(t, Zero(currency.Testing().Short()), z)
	}

	z, ok = Integer(currency.Testing().Short(), 0).(zero)
	if assert.True(t, ok) {
		assert.Equal(t, Zero(currency.Testing().Short()), z)
	}
}
