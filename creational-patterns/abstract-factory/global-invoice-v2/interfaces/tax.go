package interfaces

type Tax interface {
	CalculateCorporateTax(amount float64) float64
}
