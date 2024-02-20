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
	_, ok := c.items[itemName]
	if !ok {
		return fmt.Errorf("item %s not found", itemName)
	}
	c.cart[itemName]++
	return nil
}

// get total price calculates the total price of the items in the cart.
func (c *Checkout) GetTotalPrice() int {
	totalPrice := 0
	for itemName, quantity := range c.cart {
		item := c.items[itemName]
		specialPrice := item.SpecialPrice
		if specialPrice.Quantity > 0 && quantity >= specialPrice.Quantity {
			specialOfferCount := quantity / specialPrice.Quantity
			totalPrice += specialOfferCount * specialPrice.Price
			quantity %= specialPrice.Quantity
		}
		totalPrice += quantity * item.UnitPrice
	}
	return totalPrice
}
