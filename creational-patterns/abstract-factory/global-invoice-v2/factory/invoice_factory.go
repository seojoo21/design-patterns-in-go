package factory

import "github.com/seojoo21/design-patterns-in-go/creational-patterns/abstract-factory/global-invoice-v2/interfaces"

type InvoiceFactory interface {
	GenerateInvoice(amount float64) interfaces.Invoice
	CaculateCorporateTax(amount float64) interfaces.Tax
}
