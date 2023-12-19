package interfaces

type Invoice interface {
	GenerateInvoice(amount float64) float64
}
