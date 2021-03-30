package amount

import "github.com/mono83/currency"

func Zero(c currency.Short) currency.Amount {
	return zero(c)
}

type zero currency.Short

func (z zero) Currency() currency.Short {
	return currency.Short(z)
}

func (zero) Amount() float64 {
	return 0.
}

func (z zero) String() string {
	return toString(currency.Short(z), 0.)
}
