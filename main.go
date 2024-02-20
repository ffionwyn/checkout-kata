package main

import "fmt"

// uses checkout system by scanning items and calculating the total price.
func main() {
    items := map[string]Item{
        "A": {
            Name:      "A",
            UnitPrice: 50,
            SpecialPrice: SpecialPrice{
                Quantity: 3, // special price: 3 for 130
                Price:    130,
            },
        },
        "B": {
            Name:      "B",
            UnitPrice: 30,
            SpecialPrice: SpecialPrice{
                Quantity: 2, // special price: 2 for 45
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
    checkout.Scan("A") // trigger special price for A
    checkout.Scan("A") 
	checkout.Scan("A") 
    checkout.Scan("B") // trigger special price for B
	checkout.Scan("B") 
    checkout.Scan("C") // normal price for C
    checkout.Scan("D") // normal price for D

	totalPrice, err := checkout.GetTotalPrice()
   		if err != nil {
        fmt.Println("Error calculating total price:", err)
        return
    }
    fmt.Println("Total Price:", totalPrice) 
}
