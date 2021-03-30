package amount

import (
	"fmt"
	"testing"

	"github.com/mono83/currency"
	"github.com/stretchr/testify/assert"
)

var toStringData = []struct {
	Decimals int
	Given    float64
	Expected string
}{
	{Decimals: 0, Given: 0., Expected: "0 FOO"},
	{Decimals: 3, Given: 0., Expected: "0.000 FOO"},
	{Decimals: 5, Given: 0., Expected: "0.00000 FOO"},
	{Decimals: 1, Given: 2., Expected: "2.0 FOO"},
	{Decimals: 3, Given: -11.234, Expected: "-11.234 FOO"},
}

func TestToString(t *testing.T) {
	for _, data := range toStringData {
		t.Run(fmt.Sprint(data), func(t *testing.T) {
			assert.Equal(t, data.Expected, toString(currency.Short{ISO: "FOO", Decimals: data.Decimals}, data.Given))
		})
	}
}
