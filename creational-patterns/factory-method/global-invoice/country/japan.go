package country

import "github.com/seojoo21/design-patterns-in-go/creational-patterns/factory-method/global-invoice/schema"

type JapanInvoice struct{}

func (ji *JapanInvoice) GenerateInvoice(amount float64) *schema.InvoiceSchema {
	return &schema.InvoiceSchema{
		Amount:   amount,
		Tax:      ji.getTax(amount),
		Total:    amount + ji.getTax(amount),
		Currency: "JPY",
	}
}

func (ji *JapanInvoice) getTax(amount float64) float64 {
	// Japan has a 8% VAT
	vatRate := 0.08
	return amount * vatRate
}
