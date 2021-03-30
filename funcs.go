package currency

import "strings"

// Add registers new currency in global currencies registry
func Add(c Currency) {
	ref := &c

	byNumeric[c.Numeric] = ref
	byISO[c.ISO] = ref
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
