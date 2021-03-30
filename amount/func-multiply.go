package amount

import (
	"math"

	"github.com/mono83/currency"
)

// MultiplyCeil applies multiplier on given amount and rounds up result
// to decimal digits
func MultiplyCeil(a currency.Amount, multiplier float64) currency.Amount {
	if multiplier == 0. || a.Amount() == 0. {
		return Zero(a.Currency())
	}

	if i, ok := a.(integer); ok {
		return Integer(
			a.Currency(),
			int(math.Ceil(float64(i.amount)*multiplier)),
		)
	}

	return New(
		a.Currency(),
		math.Ceil(a.Amount()*divider(a.Currency().Decimals)*multiplier)/divider(a.Currency().Decimals),
	)
}

// MultiplyFloor applies multiplier on given amount and rounds down result
// to decimal digits
func MultiplyFloor(a currency.Amount, multiplier float64) currency.Amount {
	if multiplier == 0. || a.Amount() == 0. {
		return Zero(a.Currency())
	}

	if i, ok := a.(integer); ok {
		return Integer(
			a.Currency(),
			int(math.Floor(float64(i.amount)*multiplier)),
		)
	}

	return New(
		a.Currency(),
		math.Floor(a.Amount()*divider(a.Currency().Decimals)*multiplier)/divider(a.Currency().Decimals),
	)
}
