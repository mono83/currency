package amount

import (
	"math"

	"github.com/mono83/currency"
)

func Integer(c currency.Short, amount int) currency.Amount {
	if amount == 0 {
		return Zero(c)
	}

	return integer{
		currency: c,
		amount:   amount,
	}
}

type integer struct {
	currency currency.Short
	amount   int
}

func (i integer) Currency() currency.Short {
	return i.currency
}

func (i integer) Amount() float64 {
	return float64(i.amount) / divider(i.currency.Decimals)
}

func (i integer) String() string {
	return toString(i.currency, i.Amount())
}

func divider(decimals int) float64 {
	if decimals == 0 {
		return 1.
	} else if decimals == 1 {
		return 10.
	} else if decimals == 2 {
		return 100.
	} else if decimals == 3 {
		return 1000.
	}

	return math.Pow10(decimals)
}
