package amount

import (
	"math"
)

var infPositive = math.Inf(1)
var infNegative = math.Inf(-1)

func iadd(a, b int64) (int64, bool) {
	if a == 0 || b == 0 || (a < 0 && b > 0) || (a > 0 && b < 0) {
		// No overflow possible
	} else {
		if a > 0 {
			// Both numbers are positive
			if math.MaxInt64-b < a {
				// Positive integer overflow
				return 0, false
			}
		} else {
			// Both numbers are negative
			if math.MinInt64-b > a {
				// Negative integer overflow
				return 0, false
			}
		}
	}

	return a + b, true
}

func imultiply(a int64, b float64, f func(float64) float64) (int64, bool) {
	if b == infPositive || b == infNegative {
		// Multiplier is infinity
		return 0, false
	}

	c := f(float64(a) * b)
	if c > float64(math.MaxInt64) || c < float64(math.MinInt64) {
		return 0, false
	}

	return int64(c), true
}

func fadd(a, b float64) (float64, bool) {
	if a == infNegative || b == infNegative || a == infPositive || b == infPositive {
		// At least one of operands is infinity
		return 0, false
	}
	c := a + b
	if c == infNegative || c == infPositive {
		// Result is infinity
		return 0, false
	}

	return c, true
}

func fmultiply(a, b float64, f func(float64) float64) (float64, bool) {
	if a == infNegative || b == infNegative || a == infPositive || b == infPositive {
		// Something is infinity
		return 0, false
	}

	c := f(a * b)
	if c == infPositive || c == infNegative {
		return 0, false
	}
	return c, true
}
