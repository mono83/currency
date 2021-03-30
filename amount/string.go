package amount

import (
	"fmt"
	"strconv"

	"github.com/mono83/currency"
)

func toString(c currency.Short, a float64) string {
	format := "%.0f %s"
	if c.Decimals > 0 {
		format = "%." + strconv.Itoa(c.Decimals) + "f %s"
	}

	return fmt.Sprintf(format, a, c.ISO)
}
