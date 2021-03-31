package currency

import "strings"

// Add registers new currency in global currencies registry
func Add(currencies ...Currency) {
	for _, c := range currencies {
		add(c)
	}
}

// ByNumeric searches for currency by numeric code
func ByNumeric(s string) (Currency, bool) {
	c, ok := byNumeric[s]
	if ok {
		return *c, true
	}

	return Currency{}, false
}

// ByISO searches for currency by ISO code
func ByISO(s string) (Currency, bool) {
	c, ok := byISO[strings.ToUpper(s)]
	if ok {
		return *c, ok
	}

	return Currency{}, false
}

// EUR returns EUR currency
func EUR() Currency {
	c, ok := ByISO("EUR")
	if !ok {
		panic("Euro currency missing")
	}
	return c
}

// USD returns USD currency
func USD() Currency {
	c, ok := ByISO("USD")
	if !ok {
		panic("US dollar currency missing")
	}
	return c
}

// Testing returns testing currency (XTS)
func Testing() Currency {
	return xts
}
