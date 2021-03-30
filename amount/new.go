package amount

import "github.com/mono83/currency"

func New(currency currency.Short, amount float64) currency.Amount {
	if amount == 0. {
		return Zero(currency)
	}

	return amount64{
		currency: currency,
		amount:   amount,
	}
}

type amount64 struct {
	currency currency.Short
	amount   float64
}

func (a amount64) Currency() currency.Short {
	return a.currency
}

func (a amount64) Amount() float64 {
	return a.amount
}

func (a amount64) String() string {
	return toString(a.currency, a.amount)
}
