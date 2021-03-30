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
