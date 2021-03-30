package currency

var (
	byNumeric = map[string]*Currency{}
	byISO     = map[string]*Currency{}

	// Testing currency
	// For tesing purposes this currency should have 2 decimal digits
	// just like other major currencies
	xts = Currency{
		Numeric:  "963",
		ISO:      "XTS",
		Decimals: 2,
		Name:     "Testing currency",
	}
)

func add(c Currency) {
	ref := &c

	byNumeric[c.Numeric] = ref
	byISO[c.ISO] = ref
}

func init() {
	Add(
		Currency{Numeric: "840", ISO: "USD", Decimals: 2, Name: "US Dollar"},
		Currency{Numeric: "978", ISO: "EUR", Decimals: 2, Name: "Euro"},

		// Registering testing currency
		xts,
	)
}
