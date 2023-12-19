package country

import "github.com/seojoo21/design-patterns-in-go/creational-patterns/factory-method/global-invoice/schema"

type ChinaInvoice struct{}

func (ci *ChinaInvoice) GenerateInvoice(amount float64) *schema.InvoiceSchema {
	return &schema.InvoiceSchema{
		Amount:   amount,
		Tax:      ci.getTax(amount),
		Total:    amount + ci.getTax(amount),
		Currency: "CNY",
	}
}

func (ci *ChinaInvoice) getTax(amount float64) float64 {
	// China has a 13% VAT
	vatRate := 0.13
	return amount * vatRate
}
