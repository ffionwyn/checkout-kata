package main

import "fmt"

// item struct to represent each item.
type Item struct {
    Name         string
    UnitPrice    int
    SpecialPrice SpecialPrice
}

// special price struct to represent special pricing for an item.
type SpecialPrice struct {
    Quantity int
    Price    int
}

// checkout struct to represent the checkout system.
type Checkout struct {
    items map[string]Item
    cart  map[string]int
}

// start checkout initializes a new instance with provided items
func StartCheckout(items map[string]Item) *Checkout {
    return &Checkout{
        items: items,
        cart:  make(map[string]int),
    }
}

// scan adds an item to the empty cart
func (c *Checkout) Scan(itemName string) error {
  _, ok := c.items[itemName] 
    if !ok {
        return fmt.Errorf("item %s not found", itemName)
    }
    c.cart[itemName]++ 
    return nil
}

// get total price calculates the total price of the items in the cart
func (c *Checkout) GetTotalPrice() int {
    totalPrice := 0
    for itemName, quantity := range c.cart {
        item := c.items[itemName]
        specialPrice := item.SpecialPrice
        if specialPrice.Quantity > 0 && quantity >= specialPrice.Quantity {
            totalPrice += (quantity / specialPrice.Quantity) * specialPrice.Price
            quantity %= specialPrice.Quantity
        }
        totalPrice += quantity * item.UnitPrice
    }
    return totalPrice
}

// uses checkout system by scanning items and calculating the total price
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
    checkout.Scan("B")
    checkout.Scan("A")
    checkout.Scan("C")
    checkout.Scan("D")

    totalPrice := checkout.GetTotalPrice()
    fmt.Println("Total Price:", totalPrice) // output should be 175
}