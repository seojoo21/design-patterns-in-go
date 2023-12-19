package schema

type InvoiceSchema struct {
	Amount   float64
	Tax      float64
	Total    float64
	Currency string
}
