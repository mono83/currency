package currency

// Currency contains full information about currency
type Currency struct {
	Numeric  string // ISO 4217 numeric currency code
	ISO      string // ISO currency code
	Name     string // Currency name
	Decimals int    // Amount of decimals
}

func (c Currency) String() string {
	return c.ISO
}

// Short returns short (brief) representation of currency
func (c Currency) Short() Short {
	return Short{
		Numeric: c.Numeric,
		ISO:     c.ISO,
	}
}
