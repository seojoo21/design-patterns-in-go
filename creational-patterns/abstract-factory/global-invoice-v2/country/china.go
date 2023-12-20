package country

import "github.com/seojoo21/design-patterns-in-go/creational-patterns/abstract-factory/global-invoice-v2/schema"

type China struct{}

func (c *China) GenerateInvoice(amount float64) *schema.InvoiceSchema {
	return &schema.InvoiceSchema{
		Amount:   amount,
		Tax:      c.getTax(amount),
		Total:    amount + c.getTax(amount),
		Currency: "CNY",
	}
}

func (c *China) getTax(amount float64) float64 {
	// China has a 13% VAT
	vatRate := 0.13
	return amount * vatRate
}

func (c *China) CalculateCorporateTax(amount float64) float64 {
	rate := 0.25
	return amount * (1 + rate)
}
