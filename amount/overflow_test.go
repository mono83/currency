package amount

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var iaddData = []struct {
	Expected bool
	A, B     int64
}{
	{true, 1, 2},
	{true, 1, -2},
	{true, -1, -2},
	{true, math.MaxInt64, 0},
	{true, 0, math.MaxInt64},
	{true, math.MaxInt64, math.MinInt64},
	{true, math.MinInt64, math.MaxInt64},
	{false, math.MaxInt64, 1},
	{false, 1, math.MaxInt64},
	{false, math.MinInt64, -1},
	{false, -1, math.MinInt64},
	{false, math.MaxInt64, math.MaxInt64},
	{false, math.MinInt64, math.MinInt64},
}

func TestIAdd(t *testing.T) {
	for _, data := range iaddData {
		t.Run(fmt.Sprint(data), func(t *testing.T) {
			_, ok := iadd(data.A, data.B)
			assert.Equal(t, data.Expected, ok)
		})
	}
}

var faddData = []struct {
	Expected bool
	A, B     float64
}{
	{true, 1., 2.},
	{true, 1., -2.},
	{true, -1., -2.},
	{true, math.MaxFloat64, 0},
	{true, -math.MaxFloat64, 0},
	{true, math.SmallestNonzeroFloat64, 0},
	{false, math.Inf(1), 2.},
	{false, math.Inf(-1), 2.},
	{false, 1, math.Inf(1)},
	{false, 1, math.Inf(-1)},
	{false, math.Inf(1), math.Inf(1)},
	{false, math.Inf(-1), math.Inf(-1)},
	{false, math.Inf(1), math.Inf(-1)},
	{false, math.Inf(-1), math.Inf(1)},
	{false, math.MaxFloat64, math.MaxFloat64},
	{false, -math.MaxFloat64, -math.MaxFloat64},
}

func TestFAdd(t *testing.T) {
	for _, data := range faddData {
		t.Run(fmt.Sprint(data), func(t *testing.T) {
			_, ok := fadd(data.A, data.B)
			assert.Equal(t, data.Expected, ok)
		})
	}
}
