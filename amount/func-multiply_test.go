package amount

import (
	"testing"

	"github.com/mono83/currency"
	"github.com/stretchr/testify/assert"
)

func TestMultiplyCeilZero(t *testing.T) {
	a := Must(MultiplyCeil(New(currency.Testing().Short(), 0.), 10))
	_, ok := a.(zero)
	assert.True(t, ok)
}

func TestMultiplyFloorZero(t *testing.T) {
	a := Must(MultiplyFloor(New(currency.Testing().Short(), 0.), 10))
	_, ok := a.(zero)
	assert.True(t, ok)
}

func TestMultiplyCeilAmount64(t *testing.T) {
	a := Must(MultiplyCeil(New(currency.Testing().Short(), 11.77), 0.5))
	assert.Equal(t, New(currency.Testing().Short(), 5.89), a)
}

func TestMultiplyFloorAmount64(t *testing.T) {
	a := Must(MultiplyFloor(New(currency.Testing().Short(), 11.77), 0.5))
	assert.Equal(t, New(currency.Testing().Short(), 5.88), a)
}

func TestMultiplyCeilInteger(t *testing.T) {
	a := Must(MultiplyCeil(Integer(currency.Testing().Short(), 1177), 0.5))
	if assert.Equal(t, Integer(currency.Testing().Short(), 589), a) {
		_, ok := a.(integer)
		assert.True(t, ok)
	}
}

func TestMultiplyFloorInteger(t *testing.T) {
	a := Must(MultiplyFloor(Integer(currency.Testing().Short(), 1177), 0.5))
	if assert.Equal(t, Integer(currency.Testing().Short(), 588), a) {
		_, ok := a.(integer)
		assert.True(t, ok)
	}
}
