package country

import "github.com/seojoo21/design-patterns-in-go/creational-patterns/abstract-factory/global-invoice-v2/schema"

type Japan struct{}

func (j *Japan) GenerateInvoice(amount float64) *schema.InvoiceSchema {
	return &schema.InvoiceSchema{
		Amount:   amount,
		Tax:      j.getTax(amount),
		Total:    amount + j.getTax(amount),
		Currency: "JPY",
	}
}

func (j *Japan) getTax(amount float64) float64 {
	// Japan has a 8% VAT
	vatRate := 0.08
	return amount * vatRate
}

func (j *Japan) CalculateCorporateTax(amount float64) float64 {
	rate := 0.15
	return amount * (1 + rate)
}
