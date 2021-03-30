package amount

import "github.com/mono83/currency"

// Must function returns amount if error is nil
// or panics if error present
func Must(amount currency.Amount, err error) currency.Amount {
	if err != nil {
		panic(err)
	}

	return amount
}
