package currency

// Short contains short (brief) currency configuration
type Short struct {
	Numeric  string // ISO 4217 numeric currency code
	ISO      string // ISO currency code
	Decimals int    // Amount of decimals
}

func (s Short) String() string {
	return s.ISO
}
