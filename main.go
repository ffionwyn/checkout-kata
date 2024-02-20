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


