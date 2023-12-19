package main

import (
	"fmt"

	"github.com/seojoo21/design-patterns-in-go/creational-patterns/factory-method/global-invoice/factory"
)

func main() {
	inv := factory.CreateInvoice("korea", 1000)
	fmt.Printf("Invoice Amount :: %s %.2f\n", inv.Currency, inv.Amount)
	fmt.Printf("Invoice Tax :: %s %.2f\n", inv.Currency, inv.Tax)
	fmt.Printf("Invoice Total :: %s %.2f\n", inv.Currency, inv.Total)
}
