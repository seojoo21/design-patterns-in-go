package country

import "github.com/seojoo21/design-patterns-in-go/creational-patterns/factory-method/global-invoice/schema"

type KoreaInvoice struct{}

func (ki *KoreaInvoice) GenerateInvoice(amount float64) *schema.InvoiceSchema {
	return &schema.InvoiceSchema{
		Amount:   amount,
		Tax:      ki.getTax(amount),
		Total:    amount + ki.getTax(amount),
		Currency: "KRW",
	}
}

func (ki *KoreaInvoice) getTax(amount float64) float64 {
	// Korea has a 10% VAT
	vatRate := 0.1
	return amount * vatRate
}
