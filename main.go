package main

import (
	"fmt"
)

// uses checkout system by scanning items and calculating the total price.
func main() {
	items := map[string]Item{
		"A": {
			Name:      "A",
			UnitPrice: 50,
			SpecialPrice: SpecialPrice{
				Quantity: 3, 
				Price:    130,
			},
		},
		"B": {
			Name:      "B",
			UnitPrice: 30,
			SpecialPrice: SpecialPrice{
				Quantity: 2, 
				Price:    45,
			},
		},
		"C": {
			Name:      "C",
			UnitPrice: 20,
		},
		"D": {
			Name:      "D",
			UnitPrice: 15,
		},
	}

	checkout := StartCheckout(items)
	checkout.Scan("A") 
	checkout.Scan("A")
	checkout.Scan("A")
	checkout.Scan("B") 
	checkout.Scan("B")
	checkout.Scan("C") 
	checkout.Scan("D") 

	totalPrice, err := checkout.GetTotalPrice()
	if err != nil {
		fmt.Println("Error calculating total price:", err)
		return
	}
	fmt.Println("Total Price:", totalPrice)

	itemQuantities := checkout.getItemCounts()
	fmt.Println("Item Quantities:")
	for itemName, quantity := range itemQuantities {
		fmt.Printf("%s: %d\n", itemName, quantity)
	}
}
