package amount

import (
	"errors"
	"github.com/mono83/currency"
)

// ConvertCeil performs currency conversion for given currency using given rate
// and rounds up result to decimal digits
func ConvertCeil(a currency.Amount, rate currency.Rate) (currency.Amount, error) {
	return convert(a, rate, MultiplyCeil)
}

// ConvertFloor performs currency conversion for given currency using given rate
// and rounds down result to decimal digits
func ConvertFloor(a currency.Amount, rate currency.Rate) (currency.Amount, error) {
	return convert(a, rate, MultiplyCeil)
}

func convert(a currency.Amount, rate currency.Rate, multiplier func(currency.Amount, float64) (currency.Amount, error)) (currency.Amount, error) {
	if a.Currency() != rate.From {
		return nil, errors.New("amount and rate from currencies not match")
	}
	if rate.From == rate.To {
		return a, nil
	}
	if a.Amount() == 0. {
		return Zero(rate.To), nil
	}

	c, err := multiplier(a, float64(rate.Rate))
	if err != nil {
		return nil, err
	}

	if i, ok := c.(integer); ok {
		return integer{
			currency: rate.To,
			amount:   i.amount,
		}, nil
	}

	return New(rate.To, c.Amount()), nil
}
