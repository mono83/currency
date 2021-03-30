package amount

import (
	"errors"
	"math"

	"github.com/mono83/currency"
)

// MultiplyCeil applies multiplier on given amount and rounds up result
// to decimal digits
func MultiplyCeil(a currency.Amount, multiplier float64) (currency.Amount, error) {
	return multiply(a, multiplier, math.Ceil)
}

// MultiplyFloor applies multiplier on given amount and rounds down result
// to decimal digits
func MultiplyFloor(a currency.Amount, multiplier float64) (currency.Amount, error) {
	return multiply(a, multiplier, math.Floor)
}

func multiply(a currency.Amount, multiplier float64, rounding func(float64) float64) (currency.Amount, error) {
	if a == nil {
		return nil, errors.New("nil amount")
	}
	if multiplier == 0. || a.Amount() == 0. {
		return Zero(a.Currency()), nil
	}

	if i, ok := a.(integer); ok {
		c, ok := imultiply(i.amount, multiplier, rounding)
		if !ok {
			return nil, errors.New("int64 overflow")
		}
		return Integer(a.Currency(), c), nil
	}

	c, ok := fmultiply(a.Amount()*divider(a.Currency().Decimals), multiplier, rounding)
	if !ok {
		return nil, errors.New("float64 overflow")
	}
	return New(a.Currency(), c/divider(a.Currency().Decimals)), nil
}
