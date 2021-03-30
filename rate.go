package currency

import (
	"fmt"
	"time"
)

// Rate contains currency exchange rate
type Rate struct {
	From Short
	To   Short
	Rate float32
	Time time.Time
}

func (r Rate) String() string {
	return fmt.Sprintf(
		"[1 %s = %.3f %s @ %s]",
		r.From.String(),
		r.Rate,
		r.To.String(),
		r.Time.Truncate(time.Second).String(),
	)
}
