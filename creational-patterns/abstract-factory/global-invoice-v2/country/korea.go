package country

import "github.com/seojoo21/design-patterns-in-go/creational-patterns/abstract-factory/global-invoice-v2/schema"

type Korea struct{}

func (k *Korea) GenerateInvoice(amount float64) *schema.InvoiceSchema {
	return &schema.InvoiceSchema{
		Amount:   amount,
		Tax:      k.getTax(amount),
		Total:    amount + k.getTax(amount),
		Currency: "KRW",
	}
}

func (k *Korea) getTax(amount float64) float64 {
	// Korea has a 10% VAT
	vatRate := 0.1
	return amount * vatRate
}

func (k *Korea) CalculateCorporateTax(amount float64) float64 {
	rate := 0.2
	return amount * (1 + rate)
}
