package main

import (
	"fmt"

	"github.com/seojoo21/design-patterns-in-go/creational-patterns/abstract-factory/global-invoice-v2/country"
)

func main() {
	// Create a factory for China
	chinaFactory := &country.China{}
	chinaInvoice := chinaFactory.GenerateInvoice(1000)
	chinaCorporateTax := chinaFactory.CalculateCorporateTax(1000)

	fmt.Printf("China Invoice Amount :: %s %.2f\n", chinaInvoice.Currency, chinaInvoice.Amount)
	fmt.Printf("China Invoice Tax :: %s %.2f\n", chinaInvoice.Currency, chinaInvoice.Tax)
	fmt.Printf("China Invoice Total :: %s %.2f\n", chinaInvoice.Currency, chinaInvoice.Total)
	fmt.Printf("China Corporate Tax :: CNY %.2f\n", chinaCorporateTax)

	// Create a factory for Japan
	japanFactory := &country.Japan{}
	japanInvoice := japanFactory.GenerateInvoice(1000)
	japanCorporateTax := japanFactory.CalculateCorporateTax(1000)

	fmt.Printf("\nJapan Invoice Amount :: %s %.2f\n", japanInvoice.Currency, japanInvoice.Amount)
	fmt.Printf("Japan Invoice Tax :: %s %.2f\n", japanInvoice.Currency, japanInvoice.Tax)
	fmt.Printf("Japan Invoice Total :: %s %.2f\n", japanInvoice.Currency, japanInvoice.Total)
	fmt.Printf("Japan Corporate Tax :: JPY %.2f\n", japanCorporateTax)

	// Create a factory for Korea
	koreaFactory := &country.Korea{}
	koreaInvoice := koreaFactory.GenerateInvoice(1000)
	koreaCorporateTax := koreaFactory.CalculateCorporateTax(1000)

	fmt.Printf("\nKorea Invoice Amount :: %s %.2f\n", koreaInvoice.Currency, koreaInvoice.Amount)
	fmt.Printf("Korea Invoice Tax :: %s %.2f\n", koreaInvoice.Currency, koreaInvoice.Tax)
	fmt.Printf("Korea Invoice Total :: %s %.2f\n", koreaInvoice.Currency, koreaInvoice.Total)
	fmt.Printf("Korea Corporate Tax :: KRW %.2f\n", koreaCorporateTax)
}
