package currency

// Amount contains amount
type Amount interface {
	Currency() Short
	Amount() float64
}
