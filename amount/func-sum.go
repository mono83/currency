package amount

import (
	"errors"
	"github.com/mono83/currency"
)

// Sum performs summing of given amounts
func Sum(amounts ...currency.Amount) (currency.Amount, error) {
	l := len(amounts)
	if l == 0 {
		return nil, errors.New("no amounts given")
	} else if l == 1 {
		return amounts[0], nil
	}

	int := true
	sumInt := int64(0)
	sumF64 := 0.
	curr := amounts[0].Currency()

	for _, a := range amounts {
		if a.Currency() != curr {
			return nil, errors.New("cannot sum amounts with different currency")
		}
		if a.Amount() != 0. {
			if int {
				if i, ok := a.(integer); ok {
					x, ok := iadd(sumInt, i.amount)
					if !ok {
						return nil, errors.New("int64 overflow")
					}
					sumInt += x
				} else {
					int = false
				}
			}

			x, ok := fadd(sumF64, a.Amount())
			if !ok {
				return nil, errors.New("float64 overflow")
			}
			sumF64 += x
		}
	}

	if int {
		return Integer(curr, sumInt), nil
	}

	return New(curr, sumF64), nil
}
