package currency

var (
	byNumeric = map[string]*Currency{}
	byISO     = map[string]*Currency{}
)

func add(c Currency) {
	ref := &c

	byNumeric[c.Numeric] = ref
	byISO[c.ISO] = ref
}

func init() {
	Add(
		Currency{Numeric: "840", ISO: "USD", Decimals: 2, Name: "US Dollar"},
		Currency{Numeric: "963", ISO: "XTS", Decimals: 2, Name: "Testing currency"},
		Currency{Numeric: "978", ISO: "EUR", Decimals: 2, Name: "Euro"},
	)
}
