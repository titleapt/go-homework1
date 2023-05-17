package main 

import (
	"fmt"
)

func calculateVat (price float64) float64 {
	return price * 0.07
}

func main() {
	productPrice := 100.0
	vatAmount := calculateVat(productPrice)
	totalPrice := productPrice + vatAmount

    fmt.Printf("Product Price: ฿%.2f\n", productPrice)
	fmt.Printf("VAT Amount (7%%): ฿%.2f\n", vatAmount)
	fmt.Printf("Total Price (including VAT): ฿%.2f\n", totalPrice)
}