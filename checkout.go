package main

import "fmt"

// check out struct to represent the checkout system.
type Checkout struct {
	items map[string]Item
	cart  map[string]int
}

// start checkout initialises a new instance with provided items.
func StartCheckout(items map[string]Item) *Checkout {
	return &Checkout{
		items: items,
		cart:  make(map[string]int),
	}
}

// scan adds an item to the empty cart.
func (c *Checkout) Scan(itemName string) error {
    item, ok := c.items[itemName]
    if !ok {
        return fmt.Errorf("item %s not found", itemName)
    }
    if item.Name == "" {
        return fmt.Errorf("invalid item name: %s", itemName)
    }
    if item.SpecialPrice.Quantity <= 0 || item.SpecialPrice.Price <= 0 {
        return fmt.Errorf("invalid special price for item %s", itemName)
    }
    c.cart[itemName]++
    return nil
}

// get total price calculates the total price of the items in the cart.
func (c *Checkout) GetTotalPrice() (int, error) {
    totalPrice := 0
    for itemName, quantity := range c.cart {
        item := c.items[itemName]
        specialPrice := item.SpecialPrice
        if specialPrice.Quantity > 0 && quantity >= specialPrice.Quantity {
            if specialPrice.Quantity <= 0 || specialPrice.Price <= 0 {
                return 0, fmt.Errorf("invalid special price configuration for item %s", itemName)
            }
            specialOfferCount := quantity / specialPrice.Quantity
            totalPrice += specialOfferCount * specialPrice.Price
            quantity %= specialPrice.Quantity
        }
        if item.UnitPrice <= 0 {
            return 0, fmt.Errorf("invalid unit price for item %s", itemName)
        }
        totalPrice += quantity * item.UnitPrice
    }
    return totalPrice, nil
}
