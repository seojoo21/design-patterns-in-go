package factory

import (
	"github.com/seojoo21/design-patterns-in-go/creational-patterns/factory-method/global-invoice/country"
	"github.com/seojoo21/design-patterns-in-go/creational-patterns/factory-method/global-invoice/schema"
)

func InvoiceFactory(countryName string, amount float64) *schema.InvoiceSchema {
	switch countryName {
	case "korea":
		inv := &country.KoreaInvoice{}
		return inv.GenerateInvoice(amount)
	case "china":
		inv := &country.ChinaInvoice{}
		return inv.GenerateInvoice(amount)
	case "japan":
		inv := &country.JapanInvoice{}
		return inv.GenerateInvoice(amount)
	default:
		inv := &country.KoreaInvoice{}
		return inv.GenerateInvoice(amount)
	}
}
